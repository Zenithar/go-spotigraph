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

package opaque

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// Yes, this is hardcoded, no big deal.
// We do not really care about Confidentiality - this is merely for obfuscation
// Was generated with `openssl rand -hex 32`
var opaqueCursorEncryptionKey = [32]byte{0xdc, 0x41, 0xbb, 0x0a, 0xb6, 0x13, 0x24, 0x1b, 0xa3, 0x59, 0x7b, 0x2b, 0xa1, 0xda, 0x3e, 0x8c, 0xda, 0xfb, 0xc8, 0xa1, 0x01, 0xd8, 0xc2, 0x21, 0x01, 0x03, 0xd6, 0x84, 0x05, 0x7d, 0x71, 0x67}

// We are using a fixed nonce by design here.
// Using a fixed nonce is FATAL cryptography security flaw in normal cases
// But in this case we mostly care of obscuring / making opaque the key
var fixedNonce = [24]byte{0xdd, 0x67, 0x09, 0x6b, 0xb5, 0x79, 0x3c, 0xc5, 0xbd, 0x10, 0x4e, 0x58, 0xe0, 0x9d, 0xba, 0x49, 0x6f, 0x5c, 0x55, 0xd3, 0x9b, 0xdd, 0x33, 0x09}

// Encode obfuscates (encrypts) internal keys to be used as
// pagination cursors sent to frontend
func Encode(internalKey []byte) string {
	ciphertext := secretbox.Seal(nil, internalKey, &fixedNonce, &opaqueCursorEncryptionKey)
	return base64.RawURLEncoding.EncodeToString(ciphertext)
}

// EncodeString obfuscates (encrypts) internal keys to be used as
// pagination cursors sent to frontend
func EncodeString(internalKey string) string {
	return Encode([]byte(internalKey))
}

// Decode de-obfuscates (decrypts) internal keys to be used as
// pagination cursors sent to frontend
func Decode(opaqueKey string) ([]byte, error) {
	ciphertext, err := base64.RawURLEncoding.DecodeString(opaqueKey)
	if err != nil {
		return nil, err
	}

	plaintext, ok := secretbox.Open(nil, ciphertext, &fixedNonce, &opaqueCursorEncryptionKey)
	if !ok {
		return nil, errors.New("decryption failed")
	}

	return plaintext, nil
}

// DecodeToString de-obfuscates (decrypts) internal keys to be used as
// pagination cursors sent to frontend
func DecodeToString(opaqueKey string) (string, error) {
	out, err := Decode(opaqueKey)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
