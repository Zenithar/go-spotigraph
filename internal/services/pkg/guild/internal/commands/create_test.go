package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/commands"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
)

func TestGuild_Create(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild)
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
			req:     &guildv1.CreateRequest{},
			wantErr: true,
		},
		{
			name: "Empty label",
			req: &guildv1.CreateRequest{
				Label: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &guildv1.CreateRequest{
				Label: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing guild",
			req: &guildv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				t1 := models.NewGuild("Foo")
				guilds.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing guild",
			req: &guildv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing guild with database error",
			req: &guildv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, guilds)
			}

			// Prepare handler
			underTest := commands.CreateHandler(guilds)

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
		})
	}
}
