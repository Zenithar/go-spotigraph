// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
	bmock "go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/internal/publisher/mock"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/reactor/pkg/chapter/handlers/commands"
	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/repositories/test/mock"

	chapterv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/chapter/v1"
)

func TestChapter_Create(t *testing.T) {
	// Testcases
	testCases := []struct {
		name    string
		req     interface{}
		wantErr bool
		prepare func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, broker *bmock.MockPublisher)
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
			req:     &chapterv1.CreateRequest{},
			wantErr: true,
		},
		{
			name: "Empty label",
			req: &chapterv1.CreateRequest{
				Label:    "",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			wantErr: true,
		},
		{
			name: "Invalid name",
			req: &chapterv1.CreateRequest{
				Label:    "&é=",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			wantErr: true,
		},
		{
			name: "Existing chapter",
			req: &chapterv1.CreateRequest{
				Label:    "Foo",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, broker *bmock.MockPublisher) {
				t1 := models.NewChapter("Foo")
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(t1, nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Leader not found",
			req: &chapterv1.CreateRequest{
				Label:    "Foo",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, broker *bmock.MockPublisher) {
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(nil, db.ErrNoResult).Times(1)
			},
			wantErr: true,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter",
			req: &chapterv1.CreateRequest{
				Label:    "Foo",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, broker *bmock.MockPublisher) {
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				u1 := models.NewPerson("toto@foo.org")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				broker.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1)
			},
			wantErr: false,
		},
		// ---------------------------------------------------------------------
		{
			name: "Non-Existing chapter with database error",
			req: &chapterv1.CreateRequest{
				Label:    "Foo",
				LeaderId: "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e",
			},
			prepare: func(ctx context.Context, chapters *mock.MockChapter, persons *mock.MockPerson, broker *bmock.MockPublisher) {
				chapters.EXPECT().FindByLabel(gomock.Any(), "Foo").Return(nil, db.ErrNoResult).Times(1)
				u1 := models.NewPerson("toto@foo.org")
				persons.EXPECT().Get(gomock.Any(), "0NeNLNeGwxRtS4YPzM2QV4suGMs6Q55e").Return(u1, nil).Times(1)
				chapters.EXPECT().Create(gomock.Any(), gomock.Any()).Return(db.ErrNoModification).Times(1)
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
			chapters := mock.NewMockChapter(ctrl)
			persons := mock.NewMockPerson(ctrl)
			broker := bmock.NewMockPublisher(ctrl)

			// Prepare the mocks:
			if testCase.prepare != nil {
				testCase.prepare(ctx, chapters, persons, broker)
			}

			// Prepare handler
			underTest := commands.CreateHandler(chapters, persons, broker)

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
