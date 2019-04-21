package helpers

import (
	"github.com/dchest/uniuri"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// IDGeneratedLength defines the length of the id string
const IDGeneratedLength = 32

// IDGeneratorFunc returns a randomly generated string useable as identifier
var IDGeneratorFunc = func() string {
	return uniuri.NewLen(IDGeneratedLength)
}

// IDValidationRules describes identifier contract for syntaxic validation
var IDValidationRules = []validation.Rule{
	validation.Length(IDGeneratedLength, IDGeneratedLength),
	is.Alphanumeric,
}
