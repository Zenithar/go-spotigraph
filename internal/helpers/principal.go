package helpers

import (
	"encoding/base64"

	"golang.org/x/crypto/blake2b"
)

var principalHashKey = []byte(`7EcP%Sm5=Wgoce5Sb"%[E.<&xG8t5soYU$CzdIMTgK@^4i(Zo|)LoDB'!g"R2]8$`)

// PrincipalHashFunc return the principal hashed using Blake2b keyed algorithm
var PrincipalHashFunc = func(principal string) string {
	// Prepare hasher
	hasher, err := blake2b.New512(principalHashKey)
	if err != nil {
		panic(err)
	}

	// Append principal
	_, err = hasher.Write([]byte(principal))
	if err != nil {
		panic(err)
	}

	// Return base64 hash value of the principal hash
	return base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
}
