// Code generated by protoc-gen-validate
// source: cilium/api/network_filter.proto
// DO NOT EDIT!!!

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
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e NetworkFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = NetworkFilterValidationError{}
