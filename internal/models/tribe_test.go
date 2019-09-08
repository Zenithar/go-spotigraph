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

	"go.zenithar.org/spotigraph/internal/models"
)

func TestTribeCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewTribe("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Label).To(Equal("foo"), "Entity should have the matching label")
	g.Expect(obj.GetGroupType()).To(Equal("tribe"), "Entity should have a valid group type")
	g.Expect(obj.GetGroupID()).To(Equal(obj.ID), "Entity should have a valid group id")
}

func TestTribeValidation(t *testing.T) {
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
		obj := models.NewTribe(f.label)
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

func TestTribeMembership(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewTribe("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.SquadIDs).To(BeEmpty(), "Members should be empty")

	u1 := models.NewSquad("squad-1")
	obj.AddSquad(u1)

	l1 := models.NewPerson("toto@foo.com")
	obj.SetLeader(l1)

	g.Expect(obj.LeaderID).ToNot(BeEmpty(), "LeaderID should not be empty")
	g.Expect(obj.LeaderID).To(Equal(l1.ID), "Leader should match user id")

	g.Expect(obj.SquadIDs).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.SquadIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.SquadIDs[0]).To(Equal(u1.ID), "First element should match squad id")

	obj.AddSquad(u1)

	g.Expect(obj.SquadIDs).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.SquadIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.SquadIDs.Contains(u1.ID)).To(BeTrue(), "Members should contains u1")

	u2 := models.NewSquad("squad-2")
	obj.AddSquad(u2)
	g.Expect(len(obj.SquadIDs)).To(Equal(2), "Members length should be 2")
	g.Expect(obj.SquadIDs.Contains(u2.ID)).To(BeTrue(), "Members should contains u2")

	obj.RemoveSquad(u2)
	g.Expect(len(obj.SquadIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.SquadIDs.Contains(u2.ID)).To(BeFalse(), "Members should not contains u2")

	obj.RemoveSquad(u1)
	g.Expect(obj.SquadIDs.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.SquadIDs).To(BeEmpty(), "Members should be empty")

	obj.RemoveSquad(u1)
	g.Expect(obj.SquadIDs.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.SquadIDs).To(BeEmpty(), "Members should be empty")
}
