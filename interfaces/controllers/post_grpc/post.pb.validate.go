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

	// no validation rules for FishingSpotTypeId

	// no validation rules for PrefectureId

	// no validation rules for MeetingPlaceId

	if v, ok := interface{}(m.GetMeetingAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "MeetingAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for MaxApply

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

// Validate checks the field values on ApplyPost with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ApplyPost) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for PostId

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplyPostValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplyPostValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ApplyPostValidationError is the validation error returned by
// ApplyPost.Validate if the designated constraints aren't met.
type ApplyPostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplyPostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplyPostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplyPostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplyPostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplyPostValidationError) ErrorName() string { return "ApplyPostValidationError" }

// Error satisfies the builtin error interface
func (e ApplyPostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplyPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplyPostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplyPostValidationError{}

// Validate checks the field values on GetPostReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetPostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return GetPostReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetPostReqValidationError is the validation error returned by
// GetPostReq.Validate if the designated constraints aren't met.
type GetPostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostReqValidationError) ErrorName() string { return "GetPostReqValidationError" }

// Error satisfies the builtin error interface
func (e GetPostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostReqValidationError{}

// Validate checks the field values on ListPostsReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListPostsReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFilter() == nil {
		return ListPostsReqValidationError{
			field:  "Filter",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPostsReqValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetPageSize() > 30 {
		return ListPostsReqValidationError{
			field:  "PageSize",
			reason: "value must be less than or equal to 30",
		}
	}

	// no validation rules for PageToken

	return nil
}

// ListPostsReqValidationError is the validation error returned by
// ListPostsReq.Validate if the designated constraints aren't met.
type ListPostsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostsReqValidationError) ErrorName() string { return "ListPostsReqValidationError" }

// Error satisfies the builtin error interface
func (e ListPostsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostsReqValidationError{}

// Validate checks the field values on ListPostsRes with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListPostsRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPostsResValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	return nil
}

// ListPostsResValidationError is the validation error returned by
// ListPostsRes.Validate if the designated constraints aren't met.
type ListPostsResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostsResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostsResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostsResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostsResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostsResValidationError) ErrorName() string { return "ListPostsResValidationError" }

// Error satisfies the builtin error interface
func (e ListPostsResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostsRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostsResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostsResValidationError{}

// Validate checks the field values on CreatePostReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreatePostReq) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 20 {
		return CreatePostReqValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 2000 {
		return CreatePostReqValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 2000 runes, inclusive",
		}
	}

	if val := m.GetFishingSpotTypeId(); val < 1 || val > 4 {
		return CreatePostReqValidationError{
			field:  "FishingSpotTypeId",
			reason: "value must be inside range [1, 4]",
		}
	}

	if len(m.GetFishTypeIds()) < 1 {
		return CreatePostReqValidationError{
			field:  "FishTypeIds",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_CreatePostReq_FishTypeIds_Unique := make(map[int64]struct{}, len(m.GetFishTypeIds()))

	for idx, item := range m.GetFishTypeIds() {
		_, _ = idx, item

		if _, exists := _CreatePostReq_FishTypeIds_Unique[item]; exists {
			return CreatePostReqValidationError{
				field:  fmt.Sprintf("FishTypeIds[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_CreatePostReq_FishTypeIds_Unique[item] = struct{}{}
		}

		if val := item; val < 1 || val > 95 {
			return CreatePostReqValidationError{
				field:  fmt.Sprintf("FishTypeIds[%v]", idx),
				reason: "value must be inside range [1, 95]",
			}
		}

	}

	if val := m.GetPrefectureId(); val < 1 || val > 47 {
		return CreatePostReqValidationError{
			field:  "PrefectureId",
			reason: "value must be inside range [1, 47]",
		}
	}

	if l := utf8.RuneCountInString(m.GetMeetingPlaceId()); l < 27 || l > 255 {
		return CreatePostReqValidationError{
			field:  "MeetingPlaceId",
			reason: "value length must be between 27 and 255 runes, inclusive",
		}
	}

	if t := m.GetMeetingAt(); t != nil {
		ts, err := ptypes.Timestamp(t)
		if err != nil {
			return CreatePostReqValidationError{
				field:  "MeetingAt",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
		}

		now := time.Now()

		if ts.Sub(now) <= 0 {
			return CreatePostReqValidationError{
				field:  "MeetingAt",
				reason: "value must be greater than now",
			}
		}

	}

	if m.GetMaxApply() < 1 {
		return CreatePostReqValidationError{
			field:  "MaxApply",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return CreatePostReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreatePostReqValidationError is the validation error returned by
// CreatePostReq.Validate if the designated constraints aren't met.
type CreatePostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostReqValidationError) ErrorName() string { return "CreatePostReqValidationError" }

// Error satisfies the builtin error interface
func (e CreatePostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostReqValidationError{}

// Validate checks the field values on UpdatePostReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdatePostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return UpdatePostReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 20 {
		return UpdatePostReqValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 2000 {
		return UpdatePostReqValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 2000 runes, inclusive",
		}
	}

	if val := m.GetFishingSpotTypeId(); val < 1 || val > 4 {
		return UpdatePostReqValidationError{
			field:  "FishingSpotTypeId",
			reason: "value must be inside range [1, 4]",
		}
	}

	if len(m.GetFishTypeIds()) < 1 {
		return UpdatePostReqValidationError{
			field:  "FishTypeIds",
			reason: "value must contain at least 1 item(s)",
		}
	}

	_UpdatePostReq_FishTypeIds_Unique := make(map[int64]struct{}, len(m.GetFishTypeIds()))

	for idx, item := range m.GetFishTypeIds() {
		_, _ = idx, item

		if _, exists := _UpdatePostReq_FishTypeIds_Unique[item]; exists {
			return UpdatePostReqValidationError{
				field:  fmt.Sprintf("FishTypeIds[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_UpdatePostReq_FishTypeIds_Unique[item] = struct{}{}
		}

		if val := item; val < 1 || val > 95 {
			return UpdatePostReqValidationError{
				field:  fmt.Sprintf("FishTypeIds[%v]", idx),
				reason: "value must be inside range [1, 95]",
			}
		}

	}

	if val := m.GetPrefectureId(); val < 1 || val > 47 {
		return UpdatePostReqValidationError{
			field:  "PrefectureId",
			reason: "value must be inside range [1, 47]",
		}
	}

	if l := utf8.RuneCountInString(m.GetMeetingPlaceId()); l < 27 || l > 255 {
		return UpdatePostReqValidationError{
			field:  "MeetingPlaceId",
			reason: "value length must be between 27 and 255 runes, inclusive",
		}
	}

	if t := m.GetMeetingAt(); t != nil {
		ts, err := ptypes.Timestamp(t)
		if err != nil {
			return UpdatePostReqValidationError{
				field:  "MeetingAt",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
		}

		now := time.Now()

		if ts.Sub(now) <= 0 {
			return UpdatePostReqValidationError{
				field:  "MeetingAt",
				reason: "value must be greater than now",
			}
		}

	}

	if m.GetMaxApply() < 1 {
		return UpdatePostReqValidationError{
			field:  "MaxApply",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// UpdatePostReqValidationError is the validation error returned by
// UpdatePostReq.Validate if the designated constraints aren't met.
type UpdatePostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostReqValidationError) ErrorName() string { return "UpdatePostReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdatePostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostReqValidationError{}

// Validate checks the field values on DeletePostReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeletePostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() < 1 {
		return DeletePostReqValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeletePostReqValidationError is the validation error returned by
// DeletePostReq.Validate if the designated constraints aren't met.
type DeletePostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostReqValidationError) ErrorName() string { return "DeletePostReqValidationError" }

// Error satisfies the builtin error interface
func (e DeletePostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostReqValidationError{}

// Validate checks the field values on DeletePostRes with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeletePostRes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// DeletePostResValidationError is the validation error returned by
// DeletePostRes.Validate if the designated constraints aren't met.
type DeletePostResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostResValidationError) ErrorName() string { return "DeletePostResValidationError" }

// Error satisfies the builtin error interface
func (e DeletePostResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostResValidationError{}

// Validate checks the field values on GetApplyPostReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetApplyPostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPostId() < 1 {
		return GetApplyPostReqValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return GetApplyPostReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// GetApplyPostReqValidationError is the validation error returned by
// GetApplyPostReq.Validate if the designated constraints aren't met.
type GetApplyPostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetApplyPostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetApplyPostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetApplyPostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetApplyPostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetApplyPostReqValidationError) ErrorName() string { return "GetApplyPostReqValidationError" }

// Error satisfies the builtin error interface
func (e GetApplyPostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetApplyPostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetApplyPostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetApplyPostReqValidationError{}

// Validate checks the field values on ListApplyPostsReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListApplyPostsReq) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListApplyPostsReqValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListApplyPostsReqValidationError is the validation error returned by
// ListApplyPostsReq.Validate if the designated constraints aren't met.
type ListApplyPostsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplyPostsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplyPostsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplyPostsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplyPostsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplyPostsReqValidationError) ErrorName() string {
	return "ListApplyPostsReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplyPostsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplyPostsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplyPostsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplyPostsReqValidationError{}

// Validate checks the field values on ListApplyPostsRes with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListApplyPostsRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetApplyPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListApplyPostsResValidationError{
					field:  fmt.Sprintf("ApplyPosts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListApplyPostsResValidationError is the validation error returned by
// ListApplyPostsRes.Validate if the designated constraints aren't met.
type ListApplyPostsResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplyPostsResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplyPostsResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplyPostsResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplyPostsResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplyPostsResValidationError) ErrorName() string {
	return "ListApplyPostsResValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplyPostsResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplyPostsRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplyPostsResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplyPostsResValidationError{}

// Validate checks the field values on CreateApplyPostReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateApplyPostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPostId() < 1 {
		return CreateApplyPostReqValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return CreateApplyPostReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// CreateApplyPostReqValidationError is the validation error returned by
// CreateApplyPostReq.Validate if the designated constraints aren't met.
type CreateApplyPostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateApplyPostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateApplyPostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateApplyPostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateApplyPostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateApplyPostReqValidationError) ErrorName() string {
	return "CreateApplyPostReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateApplyPostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateApplyPostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateApplyPostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateApplyPostReqValidationError{}

// Validate checks the field values on DeleteApplyPostReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteApplyPostReq) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPostId() < 1 {
		return DeleteApplyPostReqValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
	}

	if m.GetUserId() < 1 {
		return DeleteApplyPostReqValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DeleteApplyPostReqValidationError is the validation error returned by
// DeleteApplyPostReq.Validate if the designated constraints aren't met.
type DeleteApplyPostReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteApplyPostReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteApplyPostReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteApplyPostReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteApplyPostReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteApplyPostReqValidationError) ErrorName() string {
	return "DeleteApplyPostReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteApplyPostReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteApplyPostReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteApplyPostReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteApplyPostReqValidationError{}

// Validate checks the field values on ListPostsReq_Filter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPostsReq_Filter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPrefectureId() > 47 {
		return ListPostsReq_FilterValidationError{
			field:  "PrefectureId",
			reason: "value must be less than or equal to 47",
		}
	}

	if m.GetFishingSpotTypeId() > 4 {
		return ListPostsReq_FilterValidationError{
			field:  "FishingSpotTypeId",
			reason: "value must be less than or equal to 4",
		}
	}

	_ListPostsReq_Filter_FishTypeIds_Unique := make(map[int64]struct{}, len(m.GetFishTypeIds()))

	for idx, item := range m.GetFishTypeIds() {
		_, _ = idx, item

		if _, exists := _ListPostsReq_Filter_FishTypeIds_Unique[item]; exists {
			return ListPostsReq_FilterValidationError{
				field:  fmt.Sprintf("FishTypeIds[%v]", idx),
				reason: "repeated value must contain unique items",
			}
		} else {
			_ListPostsReq_Filter_FishTypeIds_Unique[item] = struct{}{}
		}

		// no validation rules for FishTypeIds[idx]
	}

	if v, ok := interface{}(m.GetMeetingAtFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPostsReq_FilterValidationError{
				field:  "MeetingAtFrom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMeetingAtTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPostsReq_FilterValidationError{
				field:  "MeetingAtTo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CanApply

	if _, ok := ListPostsReq_Filter_OrderBy_name[int32(m.GetOrderBy())]; !ok {
		return ListPostsReq_FilterValidationError{
			field:  "OrderBy",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := ListPostsReq_Filter_SortBy_name[int32(m.GetSortBy())]; !ok {
		return ListPostsReq_FilterValidationError{
			field:  "SortBy",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for UserId

	return nil
}

// ListPostsReq_FilterValidationError is the validation error returned by
// ListPostsReq_Filter.Validate if the designated constraints aren't met.
type ListPostsReq_FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostsReq_FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostsReq_FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostsReq_FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostsReq_FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostsReq_FilterValidationError) ErrorName() string {
	return "ListPostsReq_FilterValidationError"
}

// Error satisfies the builtin error interface
func (e ListPostsReq_FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostsReq_Filter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostsReq_FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostsReq_FilterValidationError{}

// Validate checks the field values on ListApplyPostsReq_Filter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListApplyPostsReq_Filter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for PostId

	return nil
}

// ListApplyPostsReq_FilterValidationError is the validation error returned by
// ListApplyPostsReq_Filter.Validate if the designated constraints aren't met.
type ListApplyPostsReq_FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListApplyPostsReq_FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListApplyPostsReq_FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListApplyPostsReq_FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListApplyPostsReq_FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListApplyPostsReq_FilterValidationError) ErrorName() string {
	return "ListApplyPostsReq_FilterValidationError"
}

// Error satisfies the builtin error interface
func (e ListApplyPostsReq_FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListApplyPostsReq_Filter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListApplyPostsReq_FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListApplyPostsReq_FilterValidationError{}
