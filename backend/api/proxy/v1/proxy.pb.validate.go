// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proxy/v1/proxy.proto

package proxyv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RequestProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RequestProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestProxyRequestMultiError, or nil if none found.
func (m *RequestProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetService()) < 1 {
		err := RequestProxyRequestValidationError{
			field:  "Service",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHttpMethod()) < 1 {
		err := RequestProxyRequestValidationError{
			field:  "HttpMethod",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := RequestProxyRequestValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestProxyRequestValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestProxyRequestValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestProxyRequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RequestProxyRequestMultiError(errors)
	}

	return nil
}

// RequestProxyRequestMultiError is an error wrapping multiple validation
// errors returned by RequestProxyRequest.ValidateAll() if the designated
// constraints aren't met.
type RequestProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestProxyRequestMultiError) AllErrors() []error { return m }

// RequestProxyRequestValidationError is the validation error returned by
// RequestProxyRequest.Validate if the designated constraints aren't met.
type RequestProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestProxyRequestValidationError) ErrorName() string {
	return "RequestProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RequestProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestProxyRequestValidationError{}

// Validate checks the field values on RequestProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RequestProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestProxyResponseMultiError, or nil if none found.
func (m *RequestProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HttpStatus

	{
		sorted_keys := make([]string, len(m.GetHeaders()))
		i := 0
		for key := range m.GetHeaders() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetHeaders()[key]
			_ = val

			// no validation rules for Headers[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, RequestProxyResponseValidationError{
							field:  fmt.Sprintf("Headers[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, RequestProxyResponseValidationError{
							field:  fmt.Sprintf("Headers[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return RequestProxyResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestProxyResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestProxyResponseValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestProxyResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RequestProxyResponseMultiError(errors)
	}

	return nil
}

// RequestProxyResponseMultiError is an error wrapping multiple validation
// errors returned by RequestProxyResponse.ValidateAll() if the designated
// constraints aren't met.
type RequestProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestProxyResponseMultiError) AllErrors() []error { return m }

// RequestProxyResponseValidationError is the validation error returned by
// RequestProxyResponse.Validate if the designated constraints aren't met.
type RequestProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestProxyResponseValidationError) ErrorName() string {
	return "RequestProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RequestProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestProxyResponseValidationError{}
