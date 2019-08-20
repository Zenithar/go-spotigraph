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
	"go.zenithar.org/spotigraph/internal/services/pkg/guild/internal/commands"
	guildv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/guild/v1"
)

func TestGuild_Leave(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		want    interface{}
		wantErr bool
		prepare func(ctx context.Context, guilds *mock.MockGuild, persons *mock.MockPerson, memberships *mock.MockMembership)
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
			req:     &guildv1.LeaveRequest{},
			wantErr: true,
		},
		{
			name: "Empty Guild ID",
			req: &guildv1.LeaveRequest{
				GuildId: "",
			},
			wantErr: true,
		},
		{
			name: "Empty Person ID",
			req: &guildv1.LeaveRequest{
				PersonId: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid Guild ID",
			req: &guildv1.LeaveRequest{
				GuildId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Invalid Guild ID",
			req: &guildv1.LeaveRequest{
				PersonId: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Guild not exists",
			req: &guildv1.LeaveRequest{
				GuildId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild, persons *mock.MockPerson, memberships *mock.MockMembership) {
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Person not exists",
			req: &guildv1.LeaveRequest{
				GuildId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Valid request",
			req: &guildv1.LeaveRequest{
				GuildId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Leave(gomock.Any(), u1, c1).Return(nil).Times(1)
			},
			wantErr: false,
			want:    &guildv1.LeaveResponse{},
		},
		{
			name: "Database error",
			req: &guildv1.LeaveRequest{
				GuildId:  "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				PersonId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f",
			},
			prepare: func(ctx context.Context, guilds *mock.MockGuild, persons *mock.MockPerson, memberships *mock.MockMembership) {
				c1 := models.NewGuild("Foo")
				guilds.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(c1, nil).Times(1)
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55f").Return(u1, nil).Times(1)
				memberships.EXPECT().Leave(gomock.Any(), u1, c1).Return(db.ErrNoModification).Times(1)
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
			persons := mock.NewMockPerson(ctrl)
			memberships := mock.NewMockMembership(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, guilds, persons, memberships)
			}

			// Prepare service
			underTest := commands.LeaveHandler(guilds, persons, memberships)

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
