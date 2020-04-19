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

// SetPrincipalHashKey used to set the key of hash function
func SetPrincipalHashKey(key []byte) {
	if len(key) != 64 {
		panic("Principal hash key length must be 64bytes long.")
	}

	principalHashKey = key
}
