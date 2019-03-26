package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

// StringArray describes string array type
type StringArray []string

// -----------------------------------------------------------------------------

// Value exports metadata as json for sql driver
func (s StringArray) Value() (driver.Value, error) {
	j, err := json.Marshal(s)
	return j, err
}

// Scan deserialize metadata from sql driver
func (s *StringArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*s, ok = i.([]string)
	if !ok {
		return errors.New("type assertion .([]string) failed.")
	}

	return nil
}

// -----------------------------------------------------------------------------

// Contains checks if item is in collection
func (s StringArray) Contains(item string) bool {
	for _, v := range s {
		if strings.ToLower(item) == strings.ToLower(v) {
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
		if strings.ToLower(item) == strings.ToLower(v) {
			idx = i
			break
		}
	}
	if idx < 0 {
		return
	}
	*s = append((*s)[:idx], (*s)[idx+1:]...)
}
