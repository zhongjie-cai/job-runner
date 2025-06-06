# job-runner
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

![Test](https://github.com/zhongjie-cai/job-runner/actions/workflows/ci.yaml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhongjie-cai/job-runner)](https://goreportcard.com/report/github.com/zhongjie-cai/job-runner)
[![Go Reference](https://pkg.go.dev/badge/github.com/zhongjie-cai/job-runner.svg)](https://pkg.go.dev/github.com/zhongjie-cai/job-runner)

This library is provided as a wrapper utility for quickly create and run your jobs in an application.

Original source: https://github.com/zhongjie-cai/job-runner

Library dependencies (must be present in vendor folder or in Go path):
* [UUID](https://github.com/google/uuid): `go get -u github.com/google/uuid`
* [Testify](https://github.com/stretchr/testify): `go get -u github.com/stretchr/testify` (For tests only)
* [gomocker](https://github.com/zhongjie-cai/gomocker): `go get -u github.com/zhongjie-cai/gomocker` (For tests only)

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
	var schedule, scheduleError = jobrunner.NewScheduleMaker().OnSeconds(
		0, 30,
	).OnMinutes(
		0, 15, 30, 45,
	).AtHours(
		0, 6, 12, 18,
	).OnWeekdays(
		time.Sunday,
	).Schedule()
	if scheduleError != nil {
		panic(scheduleError)
	}
	var application = jobrunner.NewApplication(
		"some job runner",
		"1.2.3",
		3,        // this instructs the application to start 3 instances for each round of job execution, each assigned with a dedicated session and sequential index
		schedule, // this instructs the application to repeat the job execution rounds for every given schedule
		false,    // this instructs the application to not start a new execution if a previous one did not finish
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
session.LogMethodParameter(parameters ...any)
session.LogMethodLogic(logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...any)
session.LogMethodReturn(returns ...any)
session.LogMethodExit()
```

The `Enter`, `Parameter`, `Return` and `Exit` are limited to the scope of method boundary area loggings. 
The `Logic` is the normal logging that can be used in any place at any level in the codebase to enforce the user's customized logging entries.

# Session Attachment

The registered session contains an attachment dictionary, which allows the user to attach any object into the given session associated to a session ID.

```golang
var myAttachmentName = "my attachment name"
var myAttachmentObject = anyStruct {
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
var retrievedAttachment anyStruct
var success = session.GetAttachment(myAttachmentName, &retrievedAttachment)
if !success {
	// failed to retrieve an attachment: add your customized logic here if needed
} else {
	// succeeded to retrieve an attachment: add your customized logic here if needed
}

// alternatively, use the sugar-function to retrieve attachment like below
var retrievedAttachment, success = jobrunner.GetAttachmentFromSession[anyStruct](session, myAttachmentName)
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

You can reuse a same struct for multiple HTTP status codes, as long as the structures in JSON format are compatible.
If there is no receiver entry registered for a particular HTTP status code, the corresponding response body is ignored for deserialization when that HTTP status code is received.

```golang
...

var webcallRequest = session.CreateWebcallRequest(
	HTTP.POST,                       // Method
	"https://www.example.com/tests", // URL
	"{\"foo\":\"bar\"}",             // Payload
	true,                            // SendClientCert
)
var responseOn200 responseOn200Struct
var responseOn400 responseOn400Struct
var responseOn500 responseOn500Struct
webcallRequest.AddHeader(
	"Content-Type",
	"application/json",
).AddHeader(
	"Accept",
	"application/json",
).Anticipate(
	http.StatusOK,
	http.StatusBadRequest,
	&responseOn200,
).Anticipate(
	http.StatusBadRequest,
	http.StatusInternalServerError,
	&responseOn400,
).Anticipate(
	http.StatusInternalServerError,
	999,
	&responseOn500,
)
var statusCode, responseHeader, responseError = webcallRequest.Process()

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
