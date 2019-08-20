// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: spotigraph/chapter/v1/chapter_api.proto

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
var _chapter_api_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetLabel()); l < 2 || l > 50 {
		return CreateRequestValidationError{
			field:  "Label",
			reason: "value length must be between 2 and 50 bytes, inclusive",
		}
	}

	if !_CreateRequest_Label_Pattern.MatchString(m.GetLabel()) {
		return CreateRequestValidationError{
			field:  "Label",
			reason: "value does not match regex pattern \"^[A-Za-z]+([ 0-9A-Za-z]+)*$\"",
		}
	}

	return nil
}

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

var _CreateRequest_Label_Pattern = regexp.MustCompile("^[A-Za-z]+([ 0-9A-Za-z]+)*$")

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 32 {
		return GetRequestValidationError{
			field:  "Id",
			reason: "value length must be 32 runes",
		}
	}

	if !_GetRequest_Id_Pattern.MatchString(m.GetId()) {
		return GetRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	return nil
}

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

var _GetRequest_Id_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 32 {
		return UpdateRequestValidationError{
			field:  "Id",
			reason: "value length must be 32 runes",
		}
	}

	if !_UpdateRequest_Id_Pattern.MatchString(m.GetId()) {
		return UpdateRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	{
		tmp := m.GetLabel()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateRequestValidationError{
					field:  "Label",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

var _UpdateRequest_Id_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) != 32 {
		return DeleteRequestValidationError{
			field:  "Id",
			reason: "value length must be 32 runes",
		}
	}

	if !_DeleteRequest_Id_Pattern.MatchString(m.GetId()) {
		return DeleteRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	return nil
}

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

var _DeleteRequest_Id_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Page

	// no validation rules for PerPage

	{
		tmp := m.GetCursor()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchRequestValidationError{
					field:  "Cursor",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetChapterId(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) != 32 {
			return SearchRequestValidationError{
				field:  "ChapterId",
				reason: "value length must be 32 runes",
			}
		}

		if !_SearchRequest_ChapterId_Pattern.MatchString(wrapper.GetValue()) {
			return SearchRequestValidationError{
				field:  "ChapterId",
				reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
			}
		}

	}

	{
		tmp := m.GetLabel()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchRequestValidationError{
					field:  "Label",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

var _SearchRequest_ChapterId_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on CreateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetEntity()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateResponseValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetEntity()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetResponseValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on UpdateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetEntity()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateResponseValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// UpdateResponseValidationError is the validation error returned by
// UpdateResponse.Validate if the designated constraints aren't met.
type UpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResponseValidationError) ErrorName() string { return "UpdateResponseValidationError" }

// Error satisfies the builtin error interface
func (e UpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResponseValidationError{}

// Validate checks the field values on DeleteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeleteResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// DeleteResponseValidationError is the validation error returned by
// DeleteResponse.Validate if the designated constraints aren't met.
type DeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResponseValidationError) ErrorName() string { return "DeleteResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResponseValidationError{}

// Validate checks the field values on SearchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Total

	// no validation rules for PerPage

	// no validation rules for Count

	// no validation rules for CurrentPage

	{
		tmp := m.GetNextCursor()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  "NextCursor",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetMembers() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SearchResponseValidationError{
						field:  fmt.Sprintf("Members[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// SearchResponseValidationError is the validation error returned by
// SearchResponse.Validate if the designated constraints aren't met.
type SearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchResponseValidationError) ErrorName() string { return "SearchResponseValidationError" }

// Error satisfies the builtin error interface
func (e SearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchResponseValidationError{}

// Validate checks the field values on JoinRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JoinRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPersonId()) != 32 {
		return JoinRequestValidationError{
			field:  "PersonId",
			reason: "value length must be 32 runes",
		}
	}

	if !_JoinRequest_PersonId_Pattern.MatchString(m.GetPersonId()) {
		return JoinRequestValidationError{
			field:  "PersonId",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	if utf8.RuneCountInString(m.GetChapterId()) != 32 {
		return JoinRequestValidationError{
			field:  "ChapterId",
			reason: "value length must be 32 runes",
		}
	}

	if !_JoinRequest_ChapterId_Pattern.MatchString(m.GetChapterId()) {
		return JoinRequestValidationError{
			field:  "ChapterId",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	return nil
}

// JoinRequestValidationError is the validation error returned by
// JoinRequest.Validate if the designated constraints aren't met.
type JoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinRequestValidationError) ErrorName() string { return "JoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e JoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinRequestValidationError{}

var _JoinRequest_PersonId_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

var _JoinRequest_ChapterId_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on JoinResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JoinResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return JoinResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// JoinResponseValidationError is the validation error returned by
// JoinResponse.Validate if the designated constraints aren't met.
type JoinResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinResponseValidationError) ErrorName() string { return "JoinResponseValidationError" }

// Error satisfies the builtin error interface
func (e JoinResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinResponseValidationError{}

// Validate checks the field values on LeaveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LeaveRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPersonId()) != 32 {
		return LeaveRequestValidationError{
			field:  "PersonId",
			reason: "value length must be 32 runes",
		}
	}

	if !_LeaveRequest_PersonId_Pattern.MatchString(m.GetPersonId()) {
		return LeaveRequestValidationError{
			field:  "PersonId",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	if utf8.RuneCountInString(m.GetChapterId()) != 32 {
		return LeaveRequestValidationError{
			field:  "ChapterId",
			reason: "value length must be 32 runes",
		}
	}

	if !_LeaveRequest_ChapterId_Pattern.MatchString(m.GetChapterId()) {
		return LeaveRequestValidationError{
			field:  "ChapterId",
			reason: "value does not match regex pattern \"^[0-9A-Za-z]+$\"",
		}
	}

	return nil
}

// LeaveRequestValidationError is the validation error returned by
// LeaveRequest.Validate if the designated constraints aren't met.
type LeaveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LeaveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LeaveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LeaveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LeaveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LeaveRequestValidationError) ErrorName() string { return "LeaveRequestValidationError" }

// Error satisfies the builtin error interface
func (e LeaveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLeaveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LeaveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LeaveRequestValidationError{}

var _LeaveRequest_PersonId_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

var _LeaveRequest_ChapterId_Pattern = regexp.MustCompile("^[0-9A-Za-z]+$")

// Validate checks the field values on LeaveResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LeaveResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetError()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LeaveResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// LeaveResponseValidationError is the validation error returned by
// LeaveResponse.Validate if the designated constraints aren't met.
type LeaveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LeaveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LeaveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LeaveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LeaveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LeaveResponseValidationError) ErrorName() string { return "LeaveResponseValidationError" }

// Error satisfies the builtin error interface
func (e LeaveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLeaveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LeaveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LeaveResponseValidationError{}
