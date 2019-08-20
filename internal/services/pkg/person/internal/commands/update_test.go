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
	"go.zenithar.org/spotigraph/internal/services/pkg/person/internal/commands"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
)

func TestPerson_Update(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, persons *mock.MockPerson)
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
			req:     &personv1.UpdateRequest{},
			wantErr: true,
		},
		{
			name: "Empty ID",
			req: &personv1.UpdateRequest{
				Id: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid ID",
			req: &personv1.UpdateRequest{
				Id: "123456789",
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existent entity",
			req: &personv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Existent entity without update",
			req: &personv1.UpdateRequest{
				Id: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Existent entity with FirstName update",
			req: &personv1.UpdateRequest{
				Id:        "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				FirstName: &types.StringValue{Value: "toto"},
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				persons.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Existent entity with LastName update",
			req: &personv1.UpdateRequest{
				Id:       "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				LastName: &types.StringValue{Value: "toto"},
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				persons.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Existent entity with update error",
			req: &personv1.UpdateRequest{
				Id:       "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
				LastName: &types.StringValue{Value: "toto"},
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				u1 := models.NewPerson("Foo")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				persons.EXPECT().Update(gomock.Any(), gomock.Any()).Return(db.ErrNoResult).Times(1)
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
			persons := mock.NewMockPerson(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, persons)
			}

			// Prepare service
			underTest := commands.UpdateHandler(persons)

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
