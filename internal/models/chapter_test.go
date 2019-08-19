package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/models"
)

func TestChapterCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewChapter("foo")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Label).To(Equal("foo"), "Entity should have the matching label")
	g.Expect(obj.URN()).ToNot(BeEmpty(), "Entity should have the expected urn")
	g.Expect(obj.GetGroupType()).To(Equal("chapter"), "Entity should have a valid group type")
	g.Expect(obj.GetGroupID()).To(Equal(obj.ID), "Entity should have a valid group id")

	leader := models.NewUser("toto")
	obj.SetLeader(leader)
	g.Expect(obj.LeaderID).To(Equal(leader.ID), "Entity should have a valid leader id")
}

func TestChapterValidation(t *testing.T) {
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
		obj := models.NewChapter(f.label)
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
