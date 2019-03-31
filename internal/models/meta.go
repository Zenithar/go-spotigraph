package models

import "strings"

// Metadata represents object metadata
type Metadata map[string][]string

// Len returns the number of items in md.
func (md Metadata) Len() int {
	return len(md)
}

// Has returns true if metadata has given key
func (md Metadata) Has(k string) bool {
	k = strings.ToLower(k)
	_, ok := md[k]
	return ok
}

// Get obtains the values for a given key.
func (md Metadata) Get(k string) []string {
	k = strings.ToLower(k)
	return md[k]
}

// Set sets the value of a given key with a slice of values.
func (md Metadata) Set(k string, vals ...string) {
	if len(vals) == 0 {
		return
	}
	k = strings.ToLower(k)
	md[k] = vals
}

// Append adds the values to key k, not overwriting what was already stored at that key.
func (md Metadata) Append(k string, vals ...string) {
	if len(vals) == 0 {
		return
	}
	k = strings.ToLower(k)
	md[k] = append(md[k], vals...)
}
