# job-runner

[![Build Status](https://travis-ci.com/zhongjie-cai/job-runner.svg?branch=master)](https://travis-ci.com/zhongjie-cai/job-runner)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/zhongjie-cai/job-runner)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhongjie-cai/job-runner)](https://goreportcard.com/report/github.com/zhongjie-cai/job-runner)
[![Coverage](http://gocover.io/_badge/github.com/zhongjie-cai/job-runner)](http://gocover.io/github.com/zhongjie-cai/job-runner)

This library is provided as a wrapper utility for quickly create and run your jobs in an application.

Original source: https://github.com/zhongjie-cai/job-runner

Library dependencies (must be present in vendor folder or in Go path):
* [UUID](https://github.com/google/uuid): `go get -u github.com/google/uuid`
* [Testify](https://github.com/stretchr/testify): `go get -u github.com/stretchr/testify` (For tests only)

A sample application is shown below:

# main.go
```golang
package main

import (
	"fmt"
	"time"

	jobrunner "github.com/zhongjie-cai/job-runner"
)

// This is a sample of how to setup application for a job runner
func main() {
	var period = time.Second * 5
	var application = jobrunner.NewApplication(
		"some job runner",
		"1.2.3",
		3,       // this instructs the application to start 3 instances for each round of job execution, each assigned with a dedicated session and sequential index
		&period, // this instructs the application to repeat the job execution rounds for every given period; a new round would start even if the previous round has not finished
		&myCustomization{},
	)
	application.Start()
}

// myCustomization inherits from the default customization so you can skip setting up all customization methods
//   alternatively, you could bring in your own struct that instantiate the jobrunner.Customization interface to have a verbosed control over what to customize
type myCustomization struct {
	jobrunner.DefaultCustomization
}

func (customization *myCustomization) Log(session jobrunner.Session, logType jobrunner.LogType, logLevel jobrunner.LogLevel, category, subcategory, description string) {
	fmt.Printf("[%v|%v] <%v|%v> %v\n", logType, logLevel, category, subcategory, description)
}

func (customization *myCustomization) ActionFunc(session jobrunner.Session) error {
	session.LogMethodLogic(
		jobrunner.LogLevelInfo,
		"test category",
		"test subcategory",
		"I am a running job for %v",
		session.GetIndex(),
	)
	return nil
}
```

# Logging

The library allows the user to customize its logging function by customizing the `Log` method. 
The logging is split into two management areas: log type and log level. 

## Log Type

The log type definitions can be found under the `logType.go` file. 
Apart from all `Method`-prefixed log types, all remainig log types are managed by the library internally and should not be worried by the consumer. 

## Log Level

The log level definitions can be found under the `logLevel.go` file. 
Log level only affects all `Method`-prefixed log types; for all other log types, the log level is default to `Info`. 

## Session Logging

The registered session allows the user to add manual logging to its codebase, through several listed methods as
```golang
session.LogMethodEnter()
session.LogMethodParameter(parameters ...interface{})
session.LogMethodLogic(logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})
session.LogMethodReturn(returns ...interface{})
session.LogMethodExit()
```

The `Enter`, `Parameter`, `Return` and `Exit` are limited to the scope of method boundary area loggings. 
The `Logic` is the normal logging that can be used in any place at any level in the codebase to enforce the user's customized logging entries.

# Session Attachment

The registered session contains an attachment dictionary, which allows the user to attach any object which is JSON serializable into the given session associated to a session ID.

```golang
var myAttachmentName = "my attachment name"
var myAttachmentObject = anyJSONSerializableStruct {
	...
}
var success = session.Attach(myAttachmentName, myAttachmentObject)
if !success {
	// failed to attach an object: add your customized logic here if needed
} else {
	// succeeded to attach an object: add your customized logic here if needed
}
```

To retrieve a previously attached object from session, simply use the following sample logic.

```golang
var myAttachmentName = "my attachment name"
var retrievedAttachment anyJSONSerializableStruct
var success = session.GetAttachment(myAttachmentName, &retrievedAttachment)
if !success {
	// failed to retrieve an attachment: add your customized logic here if needed
} else {
	// succeeded to retrieve an attachment: add your customized logic here if needed
}
```

In some situations, it is good to detach a certain attachment, especially if it is a big object consuming large memory, which can be done as following.

```golang
var myAttachmentName = "my attachment name"
var success = session.Detach(myAttachmentName)
if !success {
	// failed to detach an attachment: add your customized logic here if needed
} else {
	// succeeded to detach an attachment: add your customized logic here if needed
}
```

# External Webcall Requests

The library provides a way to send out HTTP/HTTPS requests to external web services based on current session. 
Using this provided feature ensures the logging of the web service requests into corresponding log type for the given session. 

```golang
...

var webcallRequest = session.CreateWebcallRequest(
	HTTP.POST,                       // Method
	"https://www.example.com/tests", // URL
	"{\"foo\":\"bar\"}",             // Payload
	map[string]string{               // Headers
		"Content-Type": "application/json",
		"Accept": "application/json",
	},
	true,                            // SendClientCert
)
var testSample testSampleStruct
var statusCode, responseHeader, responseError = webcallRequest.Process(
	&testSample,
)

...
```

If the response from external web services is anticipated to be different according to HTTP status code, use the following declaration and processing logic to intelligently receive and deserialize the response body to anticipated data template structure:

```golang
...

var responseOn200 responseOn200Struct
var responseOn400 responseOn400Struct
var responseOn500 responseOn500Struct
var statusCode, responseHeader, responseError = webcallRequest.Process(
	map[int]interface{}{
		http.StatusOK:                  &responseOn200,
		http.StatusBadRequest:          &responseOn400,
		http.StatusForbidden:           &responseOn400,
		http.StatusInternalServerError: &responseOn500,
	},
)

...
```

You can reuse a same struct for multiple HTTP status codes, as long as the structures in JSON format are compatible. If there is no receiver entry defined in data template map for a particular HTTP status code, the corresponding response body is ignored for deserialization when that HTTP status code is received.

Webcall requests would send out client certificate for mTLS communications if the following customization is in place.

```golang
func (customization *myCustomization) ClientCert() *tls.Certificate {
    return ... // replace with however you would load the client certificate
}
```

Webcall requests could also be customized for：

## HTTP Client's HTTP Transport (http.RoundTripper)

This is to enable the 3rd party monitoring libraries, e.g. new relic, to wrap the HTTP transport for better handling of webcall communications. 

```golang
func (customization *myCustomization) RoundTripper(originalTransport http.RoundTripper) http.RoundTripper {
	return ... // replace with whatever round trip wrapper logic you would like to have
}
```

## HTTP Request (http.Request)

This is to enable the 3rd party monitoring libraries, e.g. new relic, to wrap individual HTTP request for better handling of web requests.

```golang
func (customization *myCustomization) WrapRequest(session Session, httpRequest *http.Request) *http.Request {
	return ... // replace with whatever HTTP request wrapper logic you would like to have
}
```

## Webcall Timeout

This is to provide the default HTTP request timeouts for HTTP Client over all webcall communications.

```golang
func (customization *myCustomization) DefaultTimeout() time.Duration {
	return 3 * time.Minute // replace with whatever timeout duration you would like to have
}
```
