// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: post.proto

package post_grpc

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _post_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Post) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	return nil
}

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on CreateReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CreateReq) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 20 {
		return CreateReqValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 2000 {
		return CreateReqValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 2000 runes, inclusive",
		}
	}

	if m.GetUserId() < 1 {
		return CreateReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateReqValidationError is the validation error returned by
// CreateReq.Validate if the designated constraints aren't met.
type CreateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReqValidationError) ErrorName() string { return "CreateReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReqValidationError{}

// Validate checks the field values on UpdateReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UpdateReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return UpdateReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 20 {
		return UpdateReqValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 2000 {
		return UpdateReqValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 2000 runes, inclusive",
		}
	}

	if m.GetUserId() < 1 {
		return UpdateReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// UpdateReqValidationError is the validation error returned by
// UpdateReq.Validate if the designated constraints aren't met.
type UpdateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReqValidationError) ErrorName() string { return "UpdateReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReqValidationError{}

// Validate checks the field values on ListReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ListReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDatetime() == nil {
		return ListReqValidationError{
			field:  "Datetime",
			reason: "value is required",
		}
	}

	if val := m.GetNum(); val < 1 || val > 30 {
		return ListReqValidationError{
			field:  "Num",
			reason: "value must be inside range [1, 30]",
		}
	}

	return nil
}

// ListReqValidationError is the validation error returned by ListReq.Validate
// if the designated constraints aren't met.
type ListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReqValidationError) ErrorName() string { return "ListReqValidationError" }

// Error satisfies the builtin error interface
func (e ListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReqValidationError{}

// Validate checks the field values on DeleteReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DeleteReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return DeleteReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return DeleteReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeleteReqValidationError is the validation error returned by
// DeleteReq.Validate if the designated constraints aren't met.
type DeleteReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReqValidationError) ErrorName() string { return "DeleteReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReqValidationError{}

// Validate checks the field values on ID with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ID) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return IDValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// IDValidationError is the validation error returned by ID.Validate if the
// designated constraints aren't met.
type IDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IDValidationError) ErrorName() string { return "IDValidationError" }

// Error satisfies the builtin error interface
func (e IDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IDValidationError{}

// Validate checks the field values on ListPost with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ListPost) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPostValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListPostValidationError is the validation error returned by
// ListPost.Validate if the designated constraints aren't met.
type ListPostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostValidationError) ErrorName() string { return "ListPostValidationError" }

// Error satisfies the builtin error interface
func (e ListPostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostValidationError{}

// Validate checks the field values on DeleteRes with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DeleteRes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Deleted

	return nil
}

// DeleteResValidationError is the validation error returned by
// DeleteRes.Validate if the designated constraints aren't met.
type DeleteResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResValidationError) ErrorName() string { return "DeleteResValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResValidationError{}
