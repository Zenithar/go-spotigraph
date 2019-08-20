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

func TestPerson_Search(t *testing.T) {
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
			name: "Invalid request",
			req: &personv1.SearchRequest{
				PersonId: &types.StringValue{Value: ""},
			},
			wantErr: true,
		},
		{
			name:    "Empty request",
			req:     &personv1.SearchRequest{},
			wantErr: false,
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Person{}, 0, nil).Times(1)
			},
		},
		{
			name:    "Database error",
			req:     &personv1.SearchRequest{},
			wantErr: true,
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Person{}, 0, db.ErrNoModification).Times(1)
			},
		},
		{
			name: "Filter by name",
			req: &personv1.SearchRequest{
				Principal: &types.StringValue{Value: "Foo"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Person{}, 0, nil).Times(1)
			},
		},
		{
			name: "Filter by PersonID",
			req: &personv1.SearchRequest{
				PersonId: &types.StringValue{Value: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e"},
			},
			wantErr: false,
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.Person{}, 0, nil).Times(1)
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
			persons := mock.NewMockPerson(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, persons)
			}

			// Prepare service
			underTest := commands.SearchHandler(persons)

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
