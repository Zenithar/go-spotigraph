package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/user"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

func Test_User_Creation(t *testing.T) {
	t.Parallel()

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.UserCreateReq
		wantErr bool
		prepare func(ctx context.Context, users *mock.MockUser)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.UserCreateReq{},
			wantErr: true,
		},
		{
			name: "Empty principal",
			req: &spotigraph.UserCreateReq{
				Principal: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid principal",
			req: &spotigraph.UserCreateReq{
				Principal: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Existing principal",
			req: &spotigraph.UserCreateReq{
				Principal: "toto@foo.org",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				u1 := models.NewUser("toto@foo.org")
				users.EXPECT().FindByPrincipal(ctx, "mAxU9/tpO5WgDeKqrtwSSabfoK5eQrPd7PZ9c7liWtfr5W6J0SQo72LnazVD45UsZOP7ESffo07DDQuoa2hoPw").Return(u1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing principal",
			req: &spotigraph.UserCreateReq{
				Principal: "toto@foo.org",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().FindByPrincipal(ctx, "mAxU9/tpO5WgDeKqrtwSSabfoK5eQrPd7PZ9c7liWtfr5W6J0SQo72LnazVD45UsZOP7ESffo07DDQuoa2hoPw").Return(nil, db.ErrNoResult).Times(1)
				users.EXPECT().Create(ctx, gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing principal with database error",
			req: &spotigraph.UserCreateReq{
				Principal: "toto@foo.org",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().FindByPrincipal(ctx, "mAxU9/tpO5WgDeKqrtwSSabfoK5eQrPd7PZ9c7liWtfr5W6J0SQo72LnazVD45UsZOP7ESffo07DDQuoa2hoPw").Return(nil, db.ErrNoResult).Times(1)
				users.EXPECT().Create(ctx, gomock.Any()).Return(db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			users := mock.NewMockUser(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, users)
			}

			// Prepare service
			underTest := user.New(users)

			// Do the query
			got, err := underTest.Create(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})

	}
}

func Test_User_Get(t *testing.T) {
	t.Parallel()

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.UserGetReq
		wantErr bool
		prepare func(ctx context.Context, users *mock.MockUser)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.UserGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.UserGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.UserGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Existing entity",
			req: &spotigraph.UserGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				u1 := models.NewUser("toto@foo.org")
				users.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Database error",
			req: &spotigraph.UserGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoModification).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Non-Existing entity",
			req: &spotigraph.UserGetReq{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al",
			},
			prepare: func(ctx context.Context, users *mock.MockUser) {
				users.EXPECT().Get(ctx, "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e9HublDYim7SpJNu6j8IP7d6erd2i36Al").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			users := mock.NewMockUser(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, users)
			}

			// Prepare service
			underTest := user.New(users)

			// Do the query
			got, err := underTest.Get(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})

	}
}

func Test_User_Update(t *testing.T) {
	t.Parallel()

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.UserUpdateReq
		wantErr bool
		prepare func(ctx context.Context, users *mock.MockUser)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.UserUpdateReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.UserUpdateReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.UserUpdateReq{
				Id: "123456789",
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			users := mock.NewMockUser(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, users)
			}

			// Prepare service
			underTest := user.New(users)

			// Do the query
			got, err := underTest.Update(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
				g.Expect(got.Entity).ToNot(BeNil(), "Entity should not be nil")
			}
		})

	}
}

func Test_User_Delete(t *testing.T) {
	t.Parallel()

	// Testcases
	testCases := []struct {
		name    string
		req     *spotigraph.UserGetReq
		wantErr bool
		prepare func(ctx context.Context, users *mock.MockUser)
	}{
		// ---------------------------------------------------------------------
		{
			name:    "Null request",
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &spotigraph.UserGetReq{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &spotigraph.UserGetReq{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &spotigraph.UserGetReq{
				Id: "123456789",
			},
			wantErr: true,
		},
	}

	// Subtests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Arm mocks
			ctx := context.Background()
			users := mock.NewMockUser(ctrl)

			// Prepare the mocks:
			if tt.prepare != nil {
				tt.prepare(ctx, users)
			}

			// Prepare service
			underTest := user.New(users)

			// Do the query
			got, err := underTest.Delete(ctx, tt.req)

			// assert results expectations
			if tt.wantErr {
				g.Expect(err).ToNot(BeNil(), "Error should be raised")
			} else {
				g.Expect(err).To(BeNil(), "Error should not be raised")
				g.Expect(got).ToNot(BeNil(), "Result should not be nil")
				g.Expect(got.Error).To(BeNil(), "Public error should be nil")
			}
		})

	}
}
