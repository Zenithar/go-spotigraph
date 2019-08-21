// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: spotigraph/chapter/v1/chapter.proto

package chapterv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _chapter_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Chapter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Chapter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Label

	// no validation rules for LeaderId

	// no validation rules for Urn

	return nil
}

// ChapterValidationError is the validation error returned by Chapter.Validate
// if the designated constraints aren't met.
type ChapterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChapterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChapterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChapterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChapterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChapterValidationError) ErrorName() string { return "ChapterValidationError" }

// Error satisfies the builtin error interface
func (e ChapterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChapter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChapterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChapterValidationError{}
