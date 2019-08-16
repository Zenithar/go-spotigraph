package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/models"
)

func TestGuildCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewGuild("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Label).To(Equal("foo"), "Entity should have the matching label")
	g.Expect(obj.URN()).ToNot(BeEmpty(), "Entity should have the expected urn")
	g.Expect(obj.GetGroupType()).To(Equal("guild"), "Entity should have a valid group type")
	g.Expect(obj.GetGroupID()).To(Equal(obj.ID), "Entity should have a valid group id")
}

func TestGuildValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, f := range []struct {
		label     string
		expectErr bool
	}{
		{"a", true},
		{"aa", false},
		{"foo", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
	} {
		obj := models.NewGuild(f.label)
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

func TestGuildMembership(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewGuild("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.MemberIDs).To(BeEmpty(), "Members should be empty")

	u1 := models.NewUser("toto@foo.com")
	obj.AddMember(u1)
	obj.SetLeader(u1)

	g.Expect(obj.LeaderID).ToNot(BeEmpty(), "LeaderID should not be empty")
	g.Expect(obj.LeaderID).To(Equal(u1.ID), "Leader should match user id")

	g.Expect(obj.MemberIDs).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.MemberIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.MemberIDs[0]).To(Equal(u1.ID), "First element shouold match user id")

	obj.AddMember(u1)

	g.Expect(obj.MemberIDs).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.MemberIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.MemberIDs.Contains(u1.ID)).To(BeTrue(), "Members should contains u1")

	u2 := models.NewUser("titi@foo.com")
	obj.AddMember(u2)
	g.Expect(len(obj.MemberIDs)).To(Equal(2), "Members length should be 2")
	g.Expect(obj.MemberIDs.Contains(u2.ID)).To(BeTrue(), "Members should contains u2")

	obj.RemoveMember(u2)
	g.Expect(len(obj.MemberIDs)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.MemberIDs.Contains(u2.ID)).To(BeFalse(), "Members should not contains u2")

	obj.RemoveMember(u1)
	g.Expect(obj.MemberIDs.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.MemberIDs).To(BeEmpty(), "Members should be empty")

	obj.RemoveMember(u1)
	g.Expect(obj.MemberIDs.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.MemberIDs).To(BeEmpty(), "Members should be empty")
}
