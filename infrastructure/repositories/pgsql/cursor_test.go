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
	"testing"

	"zntr.io/spotigraph/pkg/types"
)

func Test_encodeCursor(t *testing.T) {
	type args struct {
		object any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "nil",
			wantErr: false,
			want:    "-cSIyAhl4fJwi0mxhNCnPyA",
		},
		{
			name: "int",
			args: args{
				object: &personListCursor{
					ID: "123456789",
				},
			},
			wantErr: false,
			want:    "OyFcmg9HUxsluyog28n9sldGjBn_KA",
		},
		{
			name: "string",
			args: args{
				object: &personListCursor{
					Principal: types.StringRef("test"),
				},
			},
			wantErr: false,
			want:    "a_Xx2bGWS_e8-otKKgdVyFRc7zZXTvo",
		},
		{
			name: "full",
			args: args{
				object: &personListCursor{
					ID:        "123456789",
					Principal: types.StringRef("test"),
				},
			},
			wantErr: false,
			want:    "B_qBfHFtXrrIkRg8b0ZquFRGjBn_KOpHRs7V",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeCursor(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeCursor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encodeCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}
