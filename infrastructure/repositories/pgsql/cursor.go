// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pgsql

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"

	"zntr.io/spotigraph/pkg/security/opaque"
)

func decodeCursor(token string, result any) error {
	// Deobfuscate cursor
	cursorBytes, err := opaque.Decode(token)
	if err != nil {
		return fmt.Errorf("invalid cursor: %w", err)
	}

	// No error
	return cbor.Unmarshal(cursorBytes, result)
}

func encodeCursor(object any) (string, error) {
	// No error
	cursorBytes, err := cbor.Marshal(object)
	if err != nil {
		return "", fmt.Errorf("unable to encode cursor content: %w", err)
	}

	// Encode as opaque token
	token := opaque.Encode(cursorBytes)

	// No error
	return token, nil
}

func mustEncodeCursor(object any) string {
	out, err := encodeCursor(object)
	if err != nil {
		panic(err)
	}

	return out
}
