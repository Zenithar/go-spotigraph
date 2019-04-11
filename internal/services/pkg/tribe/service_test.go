package tribe_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/tribe"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func Test_Tribe_Creation(t *testing.T) {
	// Testcases
	testCases := []struct {
		name      string
		req       *spotigraph.TribeCreateReq
		publicErr *spotigraph.Error
		wantErr   bool
		prepare   func(ctx context.Context, tribes *mock.MockTribe)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.TribeCreateReq{},
			wantErr: true,
		},
		{
			name: "Empty name",
			req: &spotigraph.TribeCreateReq{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &spotigraph.TribeCreateReq{
				Name: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing tribe",
			req: &spotigraph.TribeCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				t1 := models.NewTribe("Foo")
				tribes.EXPECT().FindByName(ctx, "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing tribe",
			req: &spotigraph.TribeCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				tribes.EXPECT().Create(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing tribe with database error",
			req: &spotigraph.TribeCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				tribes.EXPECT().Create(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			tribes := mock.NewMockTribe(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, tribes)
			}

			// Prepare service
			underTest := tribe.New(tribes)

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

func Test_Tribe_Get(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.TribeGetReq
		wantErr bool
		prepare func(ctx context.Context, tribes *mock.MockTribe)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.TribeGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.TribeGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.TribeGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Existing entity",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Non-Existing entity",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
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
			tribes := mock.NewMockTribe(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, tribes)
			}

			// Prepare service
			underTest := tribe.New(tribes)

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

func Test_Tribe_Update(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.TribeUpdateReq
		wantErr bool
		prepare func(ctx context.Context, tribes *mock.MockTribe)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.TribeUpdateReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.TribeUpdateReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.TribeUpdateReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.TribeUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &spotigraph.TribeUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with principal update",
			req: &spotigraph.TribeUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("toto@foo.org")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				tribes.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				tribes.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict name",
			req: &spotigraph.TribeUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				tribes.EXPECT().FindByName(ctx, "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &spotigraph.TribeUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				tribes.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				tribes.EXPECT().Update(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			tribes := mock.NewMockTribe(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, tribes)
			}

			// Prepare service
			underTest := tribe.New(tribes)

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

func Test_Tribe_Delete(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.TribeGetReq
		wantErr bool
		prepare func(ctx context.Context, tribes *mock.MockTribe)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.TribeGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.TribeGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.TribeGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				tribes.EXPECT().Delete(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with database error",
			req: &spotigraph.TribeGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				u1 := models.NewTribe("Foo")
				tribes.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				tribes.EXPECT().Delete(ctx, gomock.Any()).Return(db.ErrNoResult).Times(1)
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
			tribes := mock.NewMockTribe(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, tribes)
			}

			// Prepare service
			underTest := tribe.New(tribes)

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

func Test_Tribe_Search(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.TribeSearchReq
		wantErr bool
		prepare func(ctx context.Context, tribes *mock.MockTribe)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		}, {
			name:    "Empty request",
			req:     &spotigraph.TribeSearchReq{},
			wantErr: false,
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Tribe{}, 0, nil).Times(1)
			},
		}, {
			name:    "Database error",
			req:     &spotigraph.TribeSearchReq{},
			wantErr: true,
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Tribe{}, 0, db.ErrNoModification).Times(1)
			},
		}, {
			name: "Filter by name",
			req: &spotigraph.TribeSearchReq{
				Name: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Tribe{}, 0, nil).Times(1)
			},
		}, {
			name: "Filter by TribeID",
			req: &spotigraph.TribeSearchReq{
				TribeId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, tribes *mock.MockTribe) {
				tribes.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Tribe{}, 0, nil).Times(1)
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
			tribes := mock.NewMockTribe(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, tribes)
			}

			// Prepare service
			underTest := tribe.New(tribes)

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
