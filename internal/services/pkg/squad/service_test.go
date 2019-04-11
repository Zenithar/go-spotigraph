package squad_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func Test_Squad_Creation(t *testing.T) {
	// Testcases
	testCases := []struct {
		name      string
		req       *spotigraph.SquadCreateReq
		publicErr *spotigraph.Error
		wantErr   bool
		prepare   func(ctx context.Context, squads *mock.MockSquad)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.SquadCreateReq{},
			wantErr: true,
		},
		{
			name: "Empty name",
			req: &spotigraph.SquadCreateReq{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &spotigraph.SquadCreateReq{
				Name: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing squad",
			req: &spotigraph.SquadCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				t1 := models.NewSquad("Foo")
				squads.EXPECT().FindByName(ctx, "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing squad",
			req: &spotigraph.SquadCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Create(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing squad with database error",
			req: &spotigraph.SquadCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Create(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, squads)
			}

			// Prepare service
			underTest := squad.New(squads)

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

func Test_Squad_Get(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.SquadGetReq
		wantErr bool
		prepare func(ctx context.Context, squads *mock.MockSquad)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.SquadGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.SquadGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.SquadGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Existing entity",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Non-Existing entity",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, squads)
			}

			// Prepare service
			underTest := squad.New(squads)

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

func Test_Squad_Update(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.SquadUpdateReq
		wantErr bool
		prepare func(ctx context.Context, squads *mock.MockSquad)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.SquadUpdateReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.SquadUpdateReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.SquadUpdateReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.SquadUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &spotigraph.SquadUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with principal update",
			req: &spotigraph.SquadUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("toto@foo.org")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				squads.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict name",
			req: &spotigraph.SquadUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				squads.EXPECT().FindByName(ctx, "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &spotigraph.SquadUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				squads.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Update(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, squads)
			}

			// Prepare service
			underTest := squad.New(squads)

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

func Test_Squad_Delete(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.SquadGetReq
		wantErr bool
		prepare func(ctx context.Context, squads *mock.MockSquad)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.SquadGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.SquadGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.SquadGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				squads.EXPECT().Delete(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with database error",
			req: &spotigraph.SquadGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				squads.EXPECT().Delete(ctx, gomock.Any()).Return(db.ErrNoResult).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, squads)
			}

			// Prepare service
			underTest := squad.New(squads)

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

func Test_Squad_Search(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.SquadSearchReq
		wantErr bool
		prepare func(ctx context.Context, squads *mock.MockSquad)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		}, {
			name:    "Empty request",
			req:     &spotigraph.SquadSearchReq{},
			wantErr: false,
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Squad{}, 0, nil).Times(1)
			},
		}, {
			name:    "Database error",
			req:     &spotigraph.SquadSearchReq{},
			wantErr: true,
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Squad{}, 0, db.ErrNoModification).Times(1)
			},
		}, {
			name: "Filter by name",
			req: &spotigraph.SquadSearchReq{
				Name: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Squad{}, 0, nil).Times(1)
			},
		}, {
			name: "Filter by SquadID",
			req: &spotigraph.SquadSearchReq{
				SquadId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Squad{}, 0, nil).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, squads)
			}

			// Prepare service
			underTest := squad.New(squads)

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
