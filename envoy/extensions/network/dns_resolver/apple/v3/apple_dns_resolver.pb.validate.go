// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/network/dns_resolver/apple/v3/apple_dns_resolver.proto

package envoy_extensions_network_dns_resolver_apple_v3

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
)

// Validate checks the field values on AppleDnsResolverConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AppleDnsResolverConfig) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AppleDnsResolverConfigValidationError is the validation error returned by
// AppleDnsResolverConfig.Validate if the designated constraints aren't met.
type AppleDnsResolverConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppleDnsResolverConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppleDnsResolverConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppleDnsResolverConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppleDnsResolverConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppleDnsResolverConfigValidationError) ErrorName() string {
	return "AppleDnsResolverConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AppleDnsResolverConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppleDnsResolverConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppleDnsResolverConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppleDnsResolverConfigValidationError{}
