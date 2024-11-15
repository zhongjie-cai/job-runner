package jobrunner

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// Customization holds all customization methods
type Customization interface {
	// BootstrapCustomization holds customization methods related to bootstrapping
	BootstrapCustomization
	// HandlerCustomization holds customization methods related to handlers
	HandlerCustomization
	// LoggingCustomization holds customization methods related to logging
	LoggingCustomization
	// WebRequestCustomization holds customization methods related to web requests
	WebRequestCustomization
}

// BootstrapCustomization holds customization methods related to bootstrapping
type BootstrapCustomization interface {
	// PreBootstrap is to customize the pre-processing logic before bootstrapping
	PreBootstrap() error

	// PostBootstrap is to customize the post-processing logic after bootstrapping
	PostBootstrap() error

	// AppClosing is to customize the application closing logic after runner shutdown
	AppClosing() error
}

// HandlerCustomization holds customization methods related to handlers
type HandlerCustomization interface {
	// PreAction is to customize the pre-action used before each job action takes place, e.g. authorization, etc.
	PreAction(session Session) error

	// PostAction is to customize the post-action used after each job action takes place successfully, e.g. finalization, etc.
	PostAction(session Session) error

	// ActionFunc is to customize the action function to be executed when the application starts
	ActionFunc(session Session) error

	// RecoverPanic is to customize the recovery of panic into a valid response and error in case it happens (for recoverable panic only)
	RecoverPanic(session Session, recoverResult interface{}) error
}

// LoggingCustomization holds customization methods related to logging
type LoggingCustomization interface {
	// Log is to customize the logging backend for the whole application
	Log(session Session, logType LogType, logLevel LogLevel, category, subcategory, description string)
}

// WebRequestCustomization holds customization methods related to web requests
type WebRequestCustomization interface {
	// ClientCert is to customize the client certificate for external requests; if not set or nil, no client certificate is sent to external web services
	ClientCert() *tls.Certificate

	// DefaultTimeout is to customize the default timeout for any webcall communications through HTTP/HTTPS by session
	DefaultTimeout() time.Duration

	// SkipServerCertVerification is to customize the skip of server certificate verification for any webcall communications through HTTP/HTTPS by session
	SkipServerCertVerification() bool

	// RoundTripper is to customize the creation of the HTTP transport for any webcall communications through HTTP/HTTPS by session
	RoundTripper(originalTransport http.RoundTripper) http.RoundTripper

	// WrapRequest is to customize the creation of the HTTP request for any webcall communications through HTTP/HTTPS by session; utilize this method if needed for new relic wrapping, etc.
	WrapRequest(session Session, httpRequest *http.Request) *http.Request
}

var (
	customizationDefault = &DefaultCustomization{}
)

// DefaultCustomization can be used for easier customization override
type DefaultCustomization struct{}

// PreBootstrap is to customize the pre-processing logic before bootstrapping
func (customization *DefaultCustomization) PreBootstrap() error {
	return nil
}

// PostBootstrap is to customize the post-processing logic after bootstrapping
func (customization *DefaultCustomization) PostBootstrap() error {
	return nil
}

// AppClosing is to customize the application closing logic after runner shutdown
func (customization *DefaultCustomization) AppClosing() error {
	return nil
}

// PreAction is to customize the pre-action used before each job action takes place, e.g. authorization, etc.
func (customization *DefaultCustomization) PreAction(session Session) error {
	return nil
}

// PostAction is to customize the post-action used after each job action takes place successfully, e.g. finalization, etc.
func (customization *DefaultCustomization) PostAction(session Session) error {
	return nil
}

// ActionFunc is to customize the action function to be executed when the application starts
func (customization *DefaultCustomization) ActionFunc(session Session) error {
	return nil
}

// RecoverPanic is to customize the recovery of panic into a valid response and error in case it happens (for recoverable panic only)
func (customization *DefaultCustomization) RecoverPanic(session Session, recoverResult interface{}) error {
	if isInterfaceValueNil(recoverResult) {
		return nil
	}
	var recoverError, ok = recoverResult.(error)
	if !ok {
		recoverError = fmt.Errorf("%v", recoverResult)
	}
	return recoverError
}

// Log is to customize the logging backend for the whole application
func (customization *DefaultCustomization) Log(session Session, logType LogType, logLevel LogLevel, category, subcategory, description string) {
	fmt.Printf(
		"[%v] <%v|%v> (%v|%v) [%v|%v] %v\n",
		formatDateTime(time.Now()),
		session.GetID(),
		session.GetIndex(),
		logType,
		logLevel,
		category,
		subcategory,
		description,
	)
}

// ClientCert is to customize the client certificate for external requests; if not set or nil, no client certificate is sent to external web services
func (customization *DefaultCustomization) ClientCert() *tls.Certificate {
	return nil
}

// DefaultTimeout is to customize the default timeout for any webcall communications through HTTP/HTTPS by session
func (customization *DefaultCustomization) DefaultTimeout() time.Duration {
	return 3 * time.Minute
}

// SkipServerCertVerification is to customize the skip of server certificate verification for any webcall communications through HTTP/HTTPS by session
func (customization *DefaultCustomization) SkipServerCertVerification() bool {
	return false
}

// RoundTripper is to customize the creation of the HTTP transport for any webcall communications through HTTP/HTTPS by session
func (customization *DefaultCustomization) RoundTripper(originalTransport http.RoundTripper) http.RoundTripper {
	return originalTransport
}

// WrapRequest is to customize the creation of the HTTP request for any webcall communications through HTTP/HTTPS by session; utilize this method if needed for new relic wrapping, etc.
func (customization *DefaultCustomization) WrapRequest(session Session, httpRequest *http.Request) *http.Request {
	return httpRequest
}
