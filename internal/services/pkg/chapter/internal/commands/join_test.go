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

func TestChapter_Join(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		want    interface{}
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, memberships *mock.MockMembership)
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
			req:     &chapterv1.JoinRequest{},
			wantErr: true,
		},
		{
			name: "Empty Chapter ID",
			req: &chapterv1.JoinRequest{
				ChapterId: "",
			},
			wantErr: true,
		},
		{
			name: "Empty Person ID",
			req: &chapterv1.JoinRequest{
				PersonId: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid Chapter ID",
			req: &chapterv1.JoinRequest{
				ChapterId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Invalid Chapter ID",
			req: &chapterv1.JoinRequest{
				PersonId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Chapter not exists",
			req: &chapterv1.JoinRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, memberships *mock.MockMembership) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Person not exists",
			req: &chapterv1.JoinRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Valid request",
			req: &chapterv1.JoinRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Join(gomock.Any(), u1, c1).Return(nil).Times(1)
			},
			wantErr: false,
			want:    &chapterv1.JoinResponse{},
		},
		{
			name: "Database error",
			req: &chapterv1.JoinRequest{
				ChapterId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Join(gomock.Any(), u1, c1).Return(db.ErrNoModification).Times(1)
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
			persons := mock.NewMockPerson(ctrl)
			memberships := mock.NewMockMembership(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, chapters, persons, memberships)
			}

			// Prepare service
			underTest := commands.JoinHandler(chapters, persons, memberships)

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
