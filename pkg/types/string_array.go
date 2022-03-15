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

package types

import "strings"

// StringArray describes string array type
type StringArray []string

// -----------------------------------------------------------------------------

// Contains checks if item is in collection
func (s StringArray) Contains(item string) bool {
	for _, v := range s {
		if strings.EqualFold(item, v) {
			return true
		}
	}

	return false
}

// AddIfNotContains add items if not already in collection
func (s *StringArray) AddIfNotContains(items ...string) {
	for _, it := range items {
		if s.Contains(it) {
			return
		}
		*s = append(*s, it)
	}
}

// Remove item from collection
func (s *StringArray) Remove(item string) {
	idx := -1
	for i, v := range *s {
		if strings.EqualFold(item, v) {
			idx = i
			break
		}
	}
	if idx < 0 {
		return
	}
	*s = append((*s)[:idx], (*s)[idx+1:]...)
}

// HasOneOf returns true when one of provided items is found in array.
func (s StringArray) HasOneOf(items ...string) bool {
	for _, item := range items {
		if s.Contains(item) {
			return true
		}
	}

	return false
}

// HasAll returns true when all of provided items is found in array.
func (s StringArray) HasAll(items ...string) bool {
	found := false

	for _, item := range items {
		found = s.Contains(item)
		if !found {
			break
		}
	}

	return found
}

// Apply string array with given function.
func (s StringArray) Apply(predicate func(string) *string) StringArray {
	res := []string{}

	for _, item := range s {
		n := predicate(item)
		if n != nil {
			res = append(res, *n)
		}
	}

	return StringArray(res)
}

// Unique returns a new StringArry without duplicate items.
func (s StringArray) Unique() StringArray {
	res := StringArray([]string{})

	for _, item := range s {
		if res.Contains(item) {
			continue
		}

		res = append(res, item)
	}

	return res
}
