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

package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/cmd/spotigraph/internal/models"
)

func TestSquadCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewSquad("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Label).To(Equal("foo"), "Entity should have the matching label")
	g.Expect(obj.GetGroupType()).To(Equal("squad"), "Entity should have a valid group type")
	g.Expect(obj.GetGroupID()).To(Equal(obj.ID), "Entity should have a valid group id")

	leader := models.NewPerson("toto")
	obj.SetProductOwner(leader)
	g.Expect(obj.ProductOwnerID).To(Equal(leader.ID), "Entity should have a valid leader id")
}

func TestSquadValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, f := range []struct {
		label     string
		expectErr bool
	}{
		{"a", true},
		{"aa", true},
		{"foo", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
	} {
		obj := models.NewSquad(f.label)
		g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")

		if err := obj.Validate(); err != nil {
			if !f.expectErr {
				t.Errorf("Validation error should not be raised, %v raised", err)
			}
		} else {
			if f.expectErr {
				t.Error("Validation error should be raised")
			}
		}
	}
}
