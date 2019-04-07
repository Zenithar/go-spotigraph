package guild_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func Test_Guild_Creation(t *testing.T) {
	// Testcases
	testCases := []struct {
		name      string
		req       *spotigraph.GuildCreateReq
		publicErr *spotigraph.Error
		wantErr   bool
		prepare   func(ctx context.Context, guilds *mock.MockGuild)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.GuildCreateReq{},
			wantErr: true,
		},
		{
			name: "Empty name",
			req: &spotigraph.GuildCreateReq{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &spotigraph.GuildCreateReq{
				Name: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing guild",
			req: &spotigraph.GuildCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				t1 := models.NewGuild("Foo")
				guilds.EXPECT().FindByName(ctx, "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing guild",
			req: &spotigraph.GuildCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Create(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing guild with database error",
			req: &spotigraph.GuildCreateReq{
				Name: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().FindByName(ctx, "Foo").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Create(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := guild.New(guilds)

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

func Test_Guild_Get(t *testing.T) {

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.GuildGetReq
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.GuildGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.GuildGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.GuildGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Existing entity",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Non-Existing entity",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := guild.New(guilds)

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

func Test_Guild_Update(t *testing.T) {

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.GuildUpdateReq
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.GuildUpdateReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.GuildUpdateReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.GuildUpdateReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.GuildUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &spotigraph.GuildUpdateReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with principal update",
			req: &spotigraph.GuildUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("toto@foo.org")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Update(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict name",
			req: &spotigraph.GuildUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByName(ctx, "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &spotigraph.GuildUpdateReq{
				Id:   "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
				Name: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByName(ctx, "Fuu").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Update(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := guild.New(guilds)

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

func Test_Guild_Delete(t *testing.T) {

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.GuildGetReq
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.GuildGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.GuildGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.GuildGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				guilds.EXPECT().Delete(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with database error",
			req: &spotigraph.GuildGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
				guilds.EXPECT().Delete(ctx, gomock.Any()).Return(db.ErrNoResult).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := guild.New(guilds)

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

func Test_Guild_Search(t *testing.T) {

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.GuildSearchReq
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		}, {
			name:    "Empty request",
			req:     &spotigraph.GuildSearchReq{},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
			},
		}, {
			name:    "Database error",
			req:     &spotigraph.GuildSearchReq{},
			wantErr: true,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, db.ErrNoModification).Times(1)
			},
		}, {
			name: "Filter by name",
			req: &spotigraph.GuildSearchReq{
				Name: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
			},
		}, {
			name: "Filter by GuildID",
			req: &spotigraph.GuildSearchReq{
				GuildId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := guild.New(guilds)

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
