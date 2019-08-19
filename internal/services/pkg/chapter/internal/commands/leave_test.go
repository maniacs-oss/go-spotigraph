package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter/internal/commands"
	chapterv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/chapter/v1"
)

func TestChapter_Leave(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		want    interface{}
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter, users *mock.MockUser, memberships *mock.MockMembership)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Invalid request type",
			req:     &struct{}{},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &chapterv1.LeaveRequest{},
			wantErr: true,
		},
		{
			name: "Empty Chapter ID",
			req: &chapterv1.LeaveRequest{
				ChapterId: "",
			},
			wantErr: true,
		},
		{
			name: "Empty User ID",
			req: &chapterv1.LeaveRequest{
				UserId: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid Chapter ID",
			req: &chapterv1.LeaveRequest{
				ChapterId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Invalid Chapter ID",
			req: &chapterv1.LeaveRequest{
				UserId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Chapter not exists",
			req: &chapterv1.LeaveRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				UserId:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, users *mock.MockUser, memberships *mock.MockMembership) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		{
			name: "User not exists",
			req: &chapterv1.LeaveRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				UserId:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, users *mock.MockUser, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				users.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Valid request",
			req: &chapterv1.LeaveRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				UserId:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, users *mock.MockUser, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewUser("Foo")
				users.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Leave(gomock.Any(), u1, c1).Return(nil).Times(1)
			},
			wantErr: false,
			want:    &chapterv1.LeaveResponse{},
		},
		{
			name: "Database error",
			req: &chapterv1.LeaveRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				UserId:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, users *mock.MockUser, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewUser("Foo")
				users.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Leave(gomock.Any(), u1, c1).Return(db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		testCase := tt
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)
			users := mock.NewMockUser(ctrl)
			memberships := mock.NewMockMembership(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, chapters, users, memberships)
			}

			// Prepare service
			underTest := commands.LeaveHandler(chapters, users, memberships)

			// Do the query
			got, err := underTest.Handle(ctx, testCase.req)

			// assert results expectations
			if testCase.wantErr && err == nil {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				t.Fatal()
			}
			if !testCase.wantErr && err != nil {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				t.Fatal()
			}
			if !testCase.wantErr && !cmp.Equal(got, testCase.want) {
				t.Fatalf("got '%v', wanted '%v'", got, testCase.want)
			}
		})
	}
}
