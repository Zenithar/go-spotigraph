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
	g.Expect(obj.Name).To(Equal("foo"), "Entity should have the matching name")
	g.Expect(obj.URN()).ToNot(BeEmpty(), "Entity should have the expected urn")
}

func TestGuildValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, f := range []struct {
		name      string
		expectErr bool
	}{
		{"a", true},
		{"aa", true},
		{"foo", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", true},
	} {
		obj := models.NewGuild(f.name)
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
	g.Expect(obj.Members).To(BeEmpty(), "Members should be empty")

	u1 := models.NewUser("toto@foo.com")
	obj.AddMember(u1)

	g.Expect(obj.Members).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.Members)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.Members[0]).To(Equal(u1.ID), "First element shouold match user id")

	obj.AddMember(u1)

	g.Expect(obj.Members).ToNot(BeEmpty(), "Members should not be empty")
	g.Expect(len(obj.Members)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.Members.Contains(u1.ID)).To(BeTrue(), "Members should contains u1")

	u2 := models.NewUser("titi@foo.com")
	obj.AddMember(u2)
	g.Expect(len(obj.Members)).To(Equal(2), "Members length should be 2")
	g.Expect(obj.Members.Contains(u2.ID)).To(BeTrue(), "Members should contains u2")

	obj.RemoveMember(u2)
	g.Expect(len(obj.Members)).To(Equal(1), "Members length should be 1")
	g.Expect(obj.Members.Contains(u2.ID)).To(BeFalse(), "Members should not contains u2")

	obj.RemoveMember(u1)
	g.Expect(obj.Members.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.Members).To(BeEmpty(), "Members should be empty")

	obj.RemoveMember(u1)
	g.Expect(obj.Members.Contains(u1.ID)).To(BeFalse(), "Members should not contains u1")
	g.Expect(obj.Members).To(BeEmpty(), "Members should be empty")
}
