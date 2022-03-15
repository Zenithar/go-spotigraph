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
package respond

import (
	"encoding/json"
	"net/http"
)

// JSON serialize the data with matching requested encoding
func WithJSON(w http.ResponseWriter, code int, data interface{}) {
	// Marshal response as json
	body, _ := json.Marshal(data)

	// Set content type header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write status
	w.WriteHeader(code)

	// Write response
	_, err := w.Write(body)
	if err != nil {
		http.Error(w, "unable to write request body", http.StatusInternalServerError)
		return
	}
}
