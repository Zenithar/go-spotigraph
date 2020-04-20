package queries_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/pkg/chapter/handlers/queries"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories/test/mock"

	chapterv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/chapter/v1"
)

func TestChapter_Search(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
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
			name: "Invalid request",
			req: &chapterv1.SearchRequest{
				ChapterId: &wrappers.StringValue{Value: "azerty"},
			},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &chapterv1.SearchRequest{},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				c1 := models.NewChapter("C1")
				c2 := models.NewChapter("C2")
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{c1, c2}, 0, nil).Times(1)
			},
		},
		{
			name:    "Database error",
			req:     &chapterv1.SearchRequest{},
			wantErr: true,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, db.ErrNoModification).Times(1)
			},
		},
		{
			name: "Filter by name",
			req: &chapterv1.SearchRequest{
				Label: &wrappers.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, nil).Times(1)
			},
		},
		{
			name: "Filter by ChapterID",
			req: &chapterv1.SearchRequest{
				ChapterId: &wrappers.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, nil).Times(1)
			},
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

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := queries.SearchHandler(chapters)

			// Do the query
			got, err := underTest.Handle(ctx, testCase.req)

			// assert results expectations
			if testCase.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
			}
		})
	}
}
