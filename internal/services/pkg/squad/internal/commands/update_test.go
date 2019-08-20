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
	"go.zenithar.org/spotigraph/internal/services/pkg/squad/internal/commands"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

func TestSquad_Update(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, squads *mock.MockSquad)
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
			req:     &squadv1.UpdateRequest{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &squadv1.UpdateRequest{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &squadv1.UpdateRequest{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &squadv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity without update",
			req: &squadv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with label update",
			req: &squadv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("toto@foo.org")
				squads.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				squads.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		}, {
			name: "Existent entity with conflict label",
			req: &squadv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				squads.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(u1, nil).Times(1)
			},
			wantErr: true,
		}, {
			name: "Existent entity with error during update",
			req: &squadv1.UpdateRequest{
				Id:    "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				Label: &types.StringValue{Value: "Fuu"},
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				u1 := models.NewSquad("Foo")
				squads.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				squads.EXPECT().FindByLabel(gomock.Any(), "Fuu").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Update(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			squads := mock.NewMockSquad(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, squads)
			}

			// Prepare service
			underTest := commands.UpdateHandler(squads)

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
