package models_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.zenithar.org/spotigraph/internal/models"
)

func TestMetadataCreation(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := &models.Metadata{}
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.Len()).To(Equal(0), "Metadata should be empty")
}

func TestMetadataModification(t *testing.T) {
	g := NewGomegaWithT(t)

	obj := &models.Metadata{}
	g.Expect(obj).ToNot(BeNil(), "Entity should not be nil")
	g.Expect(obj.Len()).To(Equal(0), "Metadata should be empty")

	obj.Set("foo.test")
	g.Expect(obj.Len()).To(Equal(0), "Metadata should be empty")

	obj.Set("foo.test", "123456", "456789")
	g.Expect(obj.Len()).To(Equal(1), "Metadata should not be empty")
	g.Expect(obj.Has("foo.test")).To(BeTrue(), "Metadata should contain given key")
	g.Expect(obj.Get("foo.test")).ToNot(BeEmpty(), "Metadata key should return non empty list")
	g.Expect(len(obj.Get("foo.test"))).To(Equal(2), "Metadata key should return 2 element list")

	obj.Append("foo.test")
	g.Expect(obj.Len()).To(Equal(1), "Metadata should not be empty")
	g.Expect(obj.Has("foo.test")).To(BeTrue(), "Metadata should contain given key")
	g.Expect(obj.Get("foo.test")).ToNot(BeEmpty(), "Metadata key should return non empty list")
	g.Expect(len(obj.Get("foo.test"))).To(Equal(2), "Metadata key should return 2 element list")

	obj.Append("foo.test", "000000")
	g.Expect(obj.Len()).To(Equal(1), "Metadata should not be empty")
	g.Expect(obj.Has("foo.test")).To(BeTrue(), "Metadata should contain given key")
	g.Expect(obj.Get("foo.test")).ToNot(BeEmpty(), "Metadata key should return non empty list")
	g.Expect(len(obj.Get("foo.test"))).To(Equal(3), "Metadata key should return 3 element list")

}
