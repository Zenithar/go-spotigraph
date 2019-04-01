package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/models"
)

func TestUserCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := models.NewUser("toto@foo.com")
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.ID).ToNot(BeEmpty(), "Entity ID should not be blank")
	g.Expect(obj.Principal).To(Equal("ivMCUbISUB91+FQcltrLoT2Unp+j3cnAf6vkEYUEzbM9iqyrzlStfAYr1vDbTcUmDwxpxHbDKkkj9M5zkU9MgQ"), "Entity should have the matching principal")
	g.Expect(obj.URN()).ToNot(BeEmpty(), "Entity should have the expected urn")
}

func TestUserValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, f := range []struct {
		name      string
		expectErr bool
	}{
		{"toto@foo.com", false},
	} {
		obj := models.NewUser(f.name)
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
