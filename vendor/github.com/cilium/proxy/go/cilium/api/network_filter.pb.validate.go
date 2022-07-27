// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/network_filter.proto

package cilium

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
var _network_filter_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on NetworkFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NetworkFilter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Proxylib

	// no validation rules for ProxylibParams

	// no validation rules for L7Proto

	// no validation rules for PolicyName

	return nil
}

// NetworkFilterValidationError is the validation error returned by
// NetworkFilter.Validate if the designated constraints aren't met.
type NetworkFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NetworkFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NetworkFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NetworkFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NetworkFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NetworkFilterValidationError) ErrorName() string { return "NetworkFilterValidationError" }

// Error satisfies the builtin error interface
func (e NetworkFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NetworkFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NetworkFilterValidationError{}
