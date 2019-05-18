package models

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

// AddIfNotContains add item if not already in collection
func (s *StringArray) AddIfNotContains(item string) {
	if s.Contains(item) {
		return
	}
	*s = append(*s, item)
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
