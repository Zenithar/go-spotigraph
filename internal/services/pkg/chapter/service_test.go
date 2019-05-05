package chapter_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/chapter"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func Test_Chapter_Creation(t *testing.T) {
	// Testcases
	testCases := []struct {
		name      string
		req       *spotigraph.ChapterCreateReq
		publicErr *spotigraph.Error
		wantErr   bool
		prepare   func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.ChapterCreateReq{},
			wantErr: true,
		},
		{
			name: "Empty name",
			req: &spotigraph.ChapterCreateReq{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &spotigraph.ChapterCreateReq{
				Name: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing chapter",
			req: &spotigraph.ChapterCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				t1 := models.NewChapter("Foo")
				chapters.EXPECT().FindByName(gomock.Any(), "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter",
			req: &spotigraph.ChapterCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().FindByName(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter with database error",
			req: &spotigraph.ChapterCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().FindByName(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := chapter.NewWithDecorators(chapters,
				chapter.WithLogger(log.Default()),
				chapter.WithTracer(),
				chapter.WithCache(cache.BigCache(), 5*time.Minute),
			)

			// Do the query
			got, err := underTest.Create(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).ToNot(BeNil(), "Public error should be set")
				if tt.publicErr != nil {
					g.Expect(got.Error).To(Equal(tt.publicErr), "Public error should be as expected")
				}
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})
	}
}

func Test_Chapter_Get(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.ChapterGetReq
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.ChapterGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.ChapterGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.ChapterGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Existing entity",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Non-Existing entity",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := chapter.NewWithDecorators(chapters, chapter.WithLogger(log.Default()), chapter.WithTracer())

			// Do the query
			got, err := underTest.Get(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).ToNot(BeNil(), "Public error should be set")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})
	}
}

func Test_Chapter_Update(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.ChapterUpdateReq
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.ChapterUpdateReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.ChapterUpdateReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.ChapterUpdateReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.ChapterUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &spotigraph.ChapterUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with principal update",
			req: &spotigraph.ChapterUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("toto@foo.org")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().FindByName(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict name",
			req: &spotigraph.ChapterUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().FindByName(gomock.Any(), "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &spotigraph.ChapterUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().FindByName(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				chapters.EXPECT().Update(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := chapter.NewWithDecorators(chapters,
				chapter.WithLogger(log.Default()),
				chapter.WithTracer(),
				chapter.WithMetric(),
			)

			// Do the query
			got, err := underTest.Update(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).ToNot(BeNil(), "Public error should be set")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})
	}
}

func Test_Chapter_Delete(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.ChapterGetReq
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.ChapterGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.ChapterGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.ChapterGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with database error",
			req: &spotigraph.ChapterGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				u1 := models.NewChapter("Foo")
				chapters.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := chapter.NewWithDecorators(chapters, chapter.WithLogger(log.Default()), chapter.WithTracer())

			// Do the query
			got, err := underTest.Delete(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).ToNot(BeNil(), "Public error should be set")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
			}
		})
	}
}

func Test_Chapter_Search(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.ChapterSearchReq
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		}, {
			name:    "Empty request",
			req:     &spotigraph.ChapterSearchReq{},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, nil).Times(1)
			},
		}, {
			name:    "Database error",
			req:     &spotigraph.ChapterSearchReq{},
			wantErr: true,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, db.ErrNoModification).Times(1)
			},
		}, {
			name: "Filter by name",
			req: &spotigraph.ChapterSearchReq{
				Name: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, nil).Times(1)
			},
		}, {
			name: "Filter by ChapterID",
			req: &spotigraph.ChapterSearchReq{
				ChapterId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, chapters *mock.MockChapter) {
				chapters.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Chapter{}, 0, nil).Times(1)
			},
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			chapters := mock.NewMockChapter(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, chapters)
			}

			// Prepare service
			underTest := chapter.NewWithDecorators(chapters, chapter.WithLogger(log.Default()), chapter.WithTracer())

			// Do the query
			got, err := underTest.Search(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).ToNot(BeNil(), "Public error should be set")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
			}
		})
	}
}
