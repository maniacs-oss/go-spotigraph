// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpc

import (
	"context"
	"crypto/tls"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"go.zenithar.org/spotigraph/cli/spotigraph/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Injectors from wire.go:

func setupLocalPostgreSQL(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	server, err := grpcServer(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

func grpcServer(ctx context.Context, cfg *config.Configuration) (*grpc.Server, error) {
	sopts := []grpc.ServerOption{}
	grpc_zap.ReplaceGrpcLogger(zap.L())

	sopts = append(sopts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_zap.StreamServerInterceptor(zap.L()), grpc_recovery.StreamServerInterceptor())), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_recovery.UnaryServerInterceptor(), grpc_zap.UnaryServerInterceptor(zap.L()))), grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)

	if cfg.Server.GRPC.UseTLS {

		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.GRPC.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.GRPC.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.GRPC.TLS.CertificatePath,
			CAFile:     cfg.Server.GRPC.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", zap.Error(err))
			return nil, err
		}

		sopts = append(sopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	} else {
		log.For(ctx).Info("No transport authentication enabled")
	}

	server := grpc.NewServer(sopts...)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	reflection.Register(server)

	err := view.Register(ochttp.ServerRequestCountView, ochttp.ServerRequestBytesView, ochttp.ServerResponseBytesView, ochttp.ServerLatencyView, ochttp.ServerRequestCountByMethod, ochttp.ServerResponseCountByStatusCode)
	if err != nil {
		log.For(ctx).Fatal("Unable to register HTTP stat views", zap.Error(err))
	}

	err = view.Register(ocgrpc.DefaultServerViews...)
	if err != nil {
		log.For(ctx).Fatal("Unable to register gRPC stat views", zap.Error(err))
	}

	return server, nil
}
