package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/squad/internal/commands"
	squadv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/squad/v1"
)

func TestSquad_Create(t *testing.T) {
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
			req:     &squadv1.CreateRequest{},
			wantErr: true,
		},
		{
			name: "Empty label",
			req: &squadv1.CreateRequest{
				Label: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &squadv1.CreateRequest{
				Label: "&é=",
			},
			wantErr: true,
		},
		{
			name: "Existing squad",
			req: &squadv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				t1 := models.NewSquad("Foo")
				squads.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing squad",
			req: &squadv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing squad with database error",
			req: &squadv1.CreateRequest{
				Label: "Foo",
			},
			prepare: func(ctx context.Context, squads *mock.MockSquad) {
				squads.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				squads.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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

			// Prepare handler
			underTest := commands.CreateHandler(squads)

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
