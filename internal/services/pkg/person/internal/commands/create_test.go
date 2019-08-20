package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories/test/mock"
	"go.zenithar.org/spotigraph/internal/services/pkg/person/internal/commands"
	personv1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/person/v1"
)

func TestPerson_Create(t *testing.T) {
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
			req:     &personv1.CreateRequest{},
			wantErr: true,
		},
		{
			name: "Empty principal",
			req: &personv1.CreateRequest{
				Principal: "",
			},
			wantErr: true,
		},
		{
			name: "Existing person",
			req: &personv1.CreateRequest{
				Principal: "Foo",
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				t1 := models.NewPerson("Foo")
				persons.EXPECT().FindByPrincipal(gomock.Any(), "tNHSUKxnBogs3bFk26hMUJxC0yTdQyvDHo8nliLfxlXM8QOlwX+6ABo59SvTQteVlRzhkB011EB3AJB4aTusRA").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing person",
			req: &personv1.CreateRequest{
				Principal: "Foo",
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().FindByPrincipal(gomock.Any(), "tNHSUKxnBogs3bFk26hMUJxC0yTdQyvDHo8nliLfxlXM8QOlwX+6ABo59SvTQteVlRzhkB011EB3AJB4aTusRA").Return(nil, db.ErrNoResult).Times(1)
				persons.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing person with database error",
			req: &personv1.CreateRequest{
				Principal: "Foo",
			},
			prepare: func(ctx context.Context, persons *mock.MockPerson) {
				persons.EXPECT().FindByPrincipal(gomock.Any(), "tNHSUKxnBogs3bFk26hMUJxC0yTdQyvDHo8nliLfxlXM8QOlwX+6ABo59SvTQteVlRzhkB011EB3AJB4aTusRA").Return(nil, db.ErrNoResult).Times(1)
				persons.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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

			// Prepare handler
			underTest := commands.CreateHandler(persons)

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
