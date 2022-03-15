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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToOpaque(t *testing.T) {
	tests := []struct {
		key    string
		opaque string
		encode func(in string) string
		decode func(in string) (string, error)
	}{
		{"super:secret:internalkey", "OuZVo68Zmyi4rBzr5PLPP6Up-ydAB_1WQM_EWDW5jf3W5fMHPqNDFA", EncodeString, DecodeToString},
		{"blah:super:secret", "BWMUz3lmUv4kyaVMxD_dq7Qw6ioITvtDRs-bX2qzkezH", EncodeString, DecodeToString},
		{"", "GZaO6ng6jtBjrutPdWf-6A", EncodeString, DecodeToString},
	}

	for _, test := range tests {
		outres := test.encode(test.key)
		assert.Equal(t, test.opaque, outres)

		inres, err := test.decode(test.opaque)
		assert.NoError(t, err)
		assert.Equal(t, test.key, inres)
	}
}
