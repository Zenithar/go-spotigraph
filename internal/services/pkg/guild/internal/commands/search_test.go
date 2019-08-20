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

func TestGuild_Search(t *testing.T) {
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
			name: "Invalid request",
			req: &guildv1.SearchRequest{
				GuildId: &types.StringValue{Value: "azerty"},
			},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &guildv1.SearchRequest{},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
			},
		},
		{
			name:    "Database error",
			req:     &guildv1.SearchRequest{},
			wantErr: true,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, db.ErrNoModification).Times(1)
			},
		},
		{
			name: "Filter by name",
			req: &guildv1.SearchRequest{
				Label: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
			},
		},
		{
			name: "Filter by GuildID",
			req: &guildv1.SearchRequest{
				GuildId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, guilds *mock.MockGuild) {
				guilds.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Guild{}, 0, nil).Times(1)
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
			guilds := mock.NewMockGuild(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, guilds)
			}

			// Prepare service
			underTest := commands.SearchHandler(guilds)

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
