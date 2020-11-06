package jobrunner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// func pointers for injection / testing: application.go
var (
	isInterfaceValueNilFunc   = isInterfaceValueNil
	uuidNew                   = uuid.New
	startApplicationFunc      = startApplication
	preBootstrapingFunc       = preBootstraping
	bootstrapFunc             = bootstrap
	postBootstrapingFunc      = postBootstraping
	endApplicationFunc        = endApplication
	beginApplicationFunc      = beginApplication
	logAppRootFunc            = logAppRoot
	handleSessionFunc         = handleSession
	runInstancesFunc          = runInstances
	repeatExecutionFunc       = repeatExecution
	timeAfter                 = time.After
	runApplicationFunc        = runApplication
	initializeHTTPClientsFunc = initializeHTTPClients
)

// func pointers for injection / testing: customization.go
var (
	fmtPrintf              = fmt.Printf
	fmtSprintf             = fmt.Sprintf
	marshalIgnoreErrorFunc = marshalIgnoreError
	debugStack             = debug.Stack
)

// func pointers for injection / testing: handler.go
var (
	stringsSplit           = strings.Split
	strconvAtoi            = strconv.Atoi
	initiateSessionFunc    = initiateSession
	getTimeNowUTCFunc      = getTimeNowUTC
	finalizeSessionFunc    = finalizeSession
	timeSince              = time.Since
	logProcessEnterFunc    = logProcessEnter
	logProcessRequestFunc  = logProcessRequest
	logProcessResponseFunc = logProcessResponse
	logProcessExitFunc     = logProcessExit
	processSessionFunc     = processSession
)

// func pointers for injection / testing: jsonutil.go
var (
	jsonNewEncoder                 = json.NewEncoder
	stringsTrimRight               = strings.TrimRight
	reflectTypeOf                  = reflect.TypeOf
	strconvParseBool               = strconv.ParseBool
	stringsToLower                 = strings.ToLower
	strconvParseInt                = strconv.ParseInt
	strconvParseFloat              = strconv.ParseFloat
	strconvParseUint               = strconv.ParseUint
	tryUnmarshalPrimitiveTypesFunc = tryUnmarshalPrimitiveTypes
	jsonUnmarshal                  = json.Unmarshal
	fmtErrorf                      = fmt.Errorf
)

// func pointers for injection / testing: logger.go
var (
	prepareLoggingFunc = prepareLogging
)

// func pointers for injection / testing: logType.go
var (
	sortStrings = sort.Strings
	stringsJoin = strings.Join
)

// func pointers for injection / testing: parameter.go
var (
	regexpMatchString = regexp.MatchString
)

// func pointers for injection / testing: pointerutil.go
var (
	reflectValueOf = reflect.ValueOf
)

// func pointers for injection / testing: session.go
var (
	strconvItoa            = strconv.Itoa
	tryUnmarshalFunc       = tryUnmarshal
	jsonMarshal            = json.Marshal
	runtimeCaller          = runtime.Caller
	runtimeFuncForPC       = runtime.FuncForPC
	getMethodNameFunc      = getMethodName
	logMethodEnterFunc     = logMethodEnter
	logMethodParameterFunc = logMethodParameter
	logMethodLogicFunc     = logMethodLogic
	logMethodReturnFunc    = logMethodReturn
	logMethodExitFunc      = logMethodExit
)

// func pointers for injection / testing: timeutil.go
var (
	timeNow = time.Now
)

// func pointers for injection / testing: webRequest.go
var (
	clientDoFunc            = clientDo
	timeSleep               = time.Sleep
	getHTTPTransportFunc    = getHTTPTransport
	urlQueryEscape          = url.QueryEscape
	createQueryStringFunc   = createQueryString
	generateRequestURLFunc  = generateRequestURL
	stringsNewReader        = strings.NewReader
	httpNewRequest          = http.NewRequest
	logWebcallStartFunc     = logWebcallStart
	logWebcallRequestFunc   = logWebcallRequest
	logWebcallResponseFunc  = logWebcallResponse
	logWebcallFinishFunc    = logWebcallFinish
	createHTTPRequestFunc   = createHTTPRequest
	getClientForRequestFunc = getClientForRequest
	clientDoWithRetryFunc   = clientDoWithRetry
	logErrorResponseFunc    = logErrorResponse
	logSuccessResponseFunc  = logSuccessResponse
	doRequestProcessingFunc = doRequestProcessing
	ioutilReadAll           = ioutil.ReadAll
	ioutilNopCloser         = ioutil.NopCloser
	bytesNewBuffer          = bytes.NewBuffer
	httpStatusText          = http.StatusText
	parseResponseFunc       = parseResponse
)
