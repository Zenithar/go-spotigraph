package helpers

import "github.com/dchest/uniuri"

// IDGeneratorFunc returns a randomly generated string useable as identifier
var IDGeneratorFunc = func() string {
	return uniuri.NewLen(64)
}
