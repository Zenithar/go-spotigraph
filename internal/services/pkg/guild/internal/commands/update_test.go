package commands_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/commands"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
)

func TestGuild_Update(t *testing.T) {
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
			req:     &guildv1.UpdateRequest{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &guildv1.UpdateRequest{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &guildv1.UpdateRequest{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &guildv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &guildv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with label update",
			req: &guildv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("toto@foo.org")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict label",
			req: &guildv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &guildv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				u1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				guilds.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				guilds.EXPECT().Update(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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

			// Prepare service
			underTest := commands.UpdateHandler(guilds)

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
