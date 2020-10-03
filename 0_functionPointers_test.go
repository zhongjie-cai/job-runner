package jobrunner

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	isInterfaceValueNilFuncExpected        int
	isInterfaceValueNilFuncCalled          int
	uuidNewExpected                        int
	uuidNewCalled                          int
	startApplicationFuncExpected           int
	startApplicationFuncCalled             int
	preBootstrapingFuncExpected            int
	preBootstrapingFuncCalled              int
	bootstrapFuncExpected                  int
	bootstrapFuncCalled                    int
	postBootstrapingFuncExpected           int
	postBootstrapingFuncCalled             int
	endApplicationFuncExpected             int
	endApplicationFuncCalled               int
	beginApplicationFuncExpected           int
	beginApplicationFuncCalled             int
	logAppRootFuncExpected                 int
	logAppRootFuncCalled                   int
	handleSessionFuncExpected              int
	handleSessionFuncCalled                int
	runInstancesFuncExpected               int
	runInstancesFuncCalled                 int
	repeatExecutionFuncExpected            int
	repeatExecutionFuncCalled              int
	timeAfterExpected                      int
	timeAfterCalled                        int
	runApplicationFuncExpected             int
	runApplicationFuncCalled               int
	initializeHTTPClientsFuncExpected      int
	initializeHTTPClientsFuncCalled        int
	fmtPrintfExpected                      int
	fmtPrintfCalled                        int
	fmtSprintfExpected                     int
	fmtSprintfCalled                       int
	marshalIgnoreErrorFuncExpected         int
	marshalIgnoreErrorFuncCalled           int
	debugStackExpected                     int
	debugStackCalled                       int
	stringsSplitExpected                   int
	stringsSplitCalled                     int
	strconvAtoiExpected                    int
	strconvAtoiCalled                      int
	initiateSessionFuncExpected            int
	initiateSessionFuncCalled              int
	getTimeNowUTCFuncExpected              int
	getTimeNowUTCFuncCalled                int
	finalizeSessionFuncExpected            int
	finalizeSessionFuncCalled              int
	timeSinceExpected                      int
	timeSinceCalled                        int
	jsonNewEncoderExpected                 int
	jsonNewEncoderCalled                   int
	stringsTrimRightExpected               int
	stringsTrimRightCalled                 int
	jsonUnmarshalExpected                  int
	jsonUnmarshalCalled                    int
	fmtErrorfExpected                      int
	fmtErrorfCalled                        int
	reflectTypeOfExpected                  int
	reflectTypeOfCalled                    int
	stringsToLowerExpected                 int
	stringsToLowerCalled                   int
	strconvParseBoolExpected               int
	strconvParseBoolCalled                 int
	strconvParseIntExpected                int
	strconvParseIntCalled                  int
	strconvParseFloatExpected              int
	strconvParseFloatCalled                int
	strconvParseUintExpected               int
	strconvParseUintCalled                 int
	tryUnmarshalPrimitiveTypesFuncExpected int
	tryUnmarshalPrimitiveTypesFuncCalled   int
	prepareLoggingFuncExpected             int
	prepareLoggingFuncCalled               int
	sortStringsExpected                    int
	sortStringsCalled                      int
	stringsJoinExpected                    int
	stringsJoinCalled                      int
	regexpMatchStringExpected              int
	regexpMatchStringCalled                int
	reflectValueOfExpected                 int
	reflectValueOfCalled                   int
	ioutilReadAllExpected                  int
	ioutilReadAllCalled                    int
	ioutilNopCloserExpected                int
	ioutilNopCloserCalled                  int
	bytesNewBufferExpected                 int
	bytesNewBufferCalled                   int
	constructResponseFuncExpected          int
	constructResponseFuncCalled            int
	logProcessEnterFuncExpected            int
	logProcessEnterFuncCalled              int
	logProcessExitFuncExpected             int
	logProcessExitFuncCalled               int
	logProcessResponseFuncExpected         int
	logProcessResponseFuncCalled           int
	logProcessRequestFuncExpected          int
	logProcessRequestFuncCalled            int
	processSessionFuncExpected             int
	processSessionFuncCalled               int
	httpStatusTextExpected                 int
	httpStatusTextCalled                   int
	strconvItoaExpected                    int
	strconvItoaCalled                      int
	tryUnmarshalFuncExpected               int
	tryUnmarshalFuncCalled                 int
	jsonMarshalExpected                    int
	jsonMarshalCalled                      int
	runtimeCallerExpected                  int
	runtimeCallerCalled                    int
	runtimeFuncForPCExpected               int
	runtimeFuncForPCCalled                 int
	getMethodNameFuncExpected              int
	getMethodNameFuncCalled                int
	logMethodEnterFuncExpected             int
	logMethodEnterFuncCalled               int
	logMethodParameterFuncExpected         int
	logMethodParameterFuncCalled           int
	logMethodLogicFuncExpected             int
	logMethodLogicFuncCalled               int
	logMethodReturnFuncExpected            int
	logMethodReturnFuncCalled              int
	logMethodExitFuncExpected              int
	logMethodExitFuncCalled                int
	timeNowExpected                        int
	timeNowCalled                          int
	clientDoFuncExpected                   int
	clientDoFuncCalled                     int
	timeSleepExpected                      int
	timeSleepCalled                        int
	getHTTPTransportFuncExpected           int
	getHTTPTransportFuncCalled             int
	stringsNewReaderExpected               int
	stringsNewReaderCalled                 int
	httpNewRequestExpected                 int
	httpNewRequestCalled                   int
	logWebcallStartFuncExpected            int
	logWebcallStartFuncCalled              int
	logWebcallRequestFuncExpected          int
	logWebcallRequestFuncCalled            int
	logWebcallResponseFuncExpected         int
	logWebcallResponseFuncCalled           int
	logWebcallFinishFuncExpected           int
	logWebcallFinishFuncCalled             int
	createHTTPRequestFuncExpected          int
	createHTTPRequestFuncCalled            int
	getClientForRequestFuncExpected        int
	getClientForRequestFuncCalled          int
	clientDoWithRetryFuncExpected          int
	clientDoWithRetryFuncCalled            int
	logErrorResponseFuncExpected           int
	logErrorResponseFuncCalled             int
	logSuccessResponseFuncExpected         int
	logSuccessResponseFuncCalled           int
	doRequestProcessingFuncExpected        int
	doRequestProcessingFuncCalled          int
	parseResponseFuncExpected              int
	parseResponseFuncCalled                int
)

func createMock(t *testing.T) {
	isInterfaceValueNilFuncExpected = 0
	isInterfaceValueNilFuncCalled = 0
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		return false
	}
	uuidNewExpected = 0
	uuidNewCalled = 0
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return uuid.Nil
	}
	startApplicationFuncExpected = 0
	startApplicationFuncCalled = 0
	startApplicationFunc = func(app *application) {
		startApplicationFuncCalled++
	}
	preBootstrapingFuncExpected = 0
	preBootstrapingFuncCalled = 0
	preBootstrapingFunc = func(app *application) bool {
		preBootstrapingFuncCalled++
		return false
	}
	bootstrapFuncExpected = 0
	bootstrapFuncCalled = 0
	bootstrapFunc = func(app *application) {
		bootstrapFuncCalled++
	}
	postBootstrapingFuncExpected = 0
	postBootstrapingFuncCalled = 0
	postBootstrapingFunc = func(app *application) bool {
		postBootstrapingFuncCalled++
		return false
	}
	endApplicationFuncExpected = 0
	endApplicationFuncCalled = 0
	endApplicationFunc = func(app *application) {
		endApplicationFuncCalled++
	}
	beginApplicationFuncExpected = 0
	beginApplicationFuncCalled = 0
	beginApplicationFunc = func(app *application) {
		beginApplicationFuncCalled++
	}
	logAppRootFuncExpected = 0
	logAppRootFuncCalled = 0
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
	}
	handleSessionFuncExpected = 0
	handleSessionFuncCalled = 0
	handleSessionFunc = func(app *application, index int) error {
		handleSessionFuncCalled++
		return nil
	}
	runInstancesFuncExpected = 0
	runInstancesFuncCalled = 0
	runInstancesFunc = func(app *application) {
		runInstancesFuncCalled++
	}
	repeatExecutionFuncExpected = 0
	repeatExecutionFuncCalled = 0
	repeatExecutionFunc = func(app *application) {
		repeatExecutionFuncCalled++
	}
	timeAfterExpected = 0
	timeAfterCalled = 0
	timeAfter = func(d time.Duration) <-chan time.Time {
		timeAfterCalled++
		return nil
	}
	runApplicationFuncExpected = 0
	runApplicationFuncCalled = 0
	runApplicationFunc = func(app *application) {
		runApplicationFuncCalled++
	}
	initializeHTTPClientsFuncExpected = 0
	initializeHTTPClientsFuncCalled = 0
	initializeHTTPClientsFunc = func(webcallTimeout time.Duration, skipServerCertVerification bool, clientCertificate *tls.Certificate, roundTripperWrapper func(originalTransport http.RoundTripper) http.RoundTripper) {
		initializeHTTPClientsFuncCalled++
	}
	fmtPrintfExpected = 0
	fmtPrintfCalled = 0
	fmtPrintf = func(format string, a ...interface{}) (n int, err error) {
		fmtPrintfCalled++
		return 0, nil
	}
	fmtSprintfExpected = 0
	fmtSprintfCalled = 0
	fmtSprintf = func(format string, a ...interface{}) string {
		fmtSprintfCalled++
		return ""
	}
	marshalIgnoreErrorFuncExpected = 0
	marshalIgnoreErrorFuncCalled = 0
	marshalIgnoreErrorFunc = func(v interface{}) string {
		marshalIgnoreErrorFuncCalled++
		return ""
	}
	debugStackExpected = 0
	debugStackCalled = 0
	debugStack = func() []byte {
		debugStackCalled++
		return nil
	}
	stringsSplitExpected = 0
	stringsSplitCalled = 0
	stringsSplit = func(s, sep string) []string {
		stringsSplitCalled++
		return nil
	}
	strconvAtoiExpected = 0
	strconvAtoiCalled = 0
	strconvAtoi = func(s string) (int, error) {
		strconvAtoiCalled++
		return 0, nil
	}
	initiateSessionFuncExpected = 0
	initiateSessionFuncCalled = 0
	initiateSessionFunc = func(app *application, index int) *session {
		initiateSessionFuncCalled++
		return nil
	}
	getTimeNowUTCFuncExpected = 0
	getTimeNowUTCFuncCalled = 0
	getTimeNowUTCFunc = func() time.Time {
		getTimeNowUTCFuncCalled++
		return time.Time{}
	}
	finalizeSessionFuncExpected = 0
	finalizeSessionFuncCalled = 0
	finalizeSessionFunc = func(session *session, resultError error) error {
		finalizeSessionFuncCalled++
		return nil
	}
	timeSinceExpected = 0
	timeSinceCalled = 0
	timeSince = func(t time.Time) time.Duration {
		timeSinceCalled++
		return 0
	}
	jsonNewEncoderExpected = 0
	jsonNewEncoderCalled = 0
	jsonNewEncoder = func(w io.Writer) *json.Encoder {
		jsonNewEncoderCalled++
		return nil
	}
	stringsTrimRightExpected = 0
	stringsTrimRightCalled = 0
	stringsTrimRight = func(s string, cutset string) string {
		stringsTrimRightCalled++
		return ""
	}
	jsonUnmarshalExpected = 0
	jsonUnmarshalCalled = 0
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return nil
	}
	fmtErrorfExpected = 0
	fmtErrorfCalled = 0
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		return nil
	}
	reflectTypeOfExpected = 0
	reflectTypeOfCalled = 0
	reflectTypeOf = func(i interface{}) reflect.Type {
		reflectTypeOfCalled++
		return nil
	}
	stringsToLowerExpected = 0
	stringsToLowerCalled = 0
	stringsToLower = func(s string) string {
		stringsToLowerCalled++
		return ""
	}
	strconvParseBoolExpected = 0
	strconvParseBoolCalled = 0
	strconvParseBool = func(str string) (bool, error) {
		strconvParseBoolCalled++
		return false, nil
	}
	strconvParseIntExpected = 0
	strconvParseIntCalled = 0
	strconvParseInt = func(s string, base int, bitSize int) (int64, error) {
		strconvParseIntCalled++
		return 0, nil
	}
	strconvParseFloatExpected = 0
	strconvParseFloatCalled = 0
	strconvParseFloat = func(s string, bitSize int) (float64, error) {
		strconvParseFloatCalled++
		return 0, nil
	}
	strconvParseUintExpected = 0
	strconvParseUintCalled = 0
	strconvParseUint = func(s string, base int, bitSize int) (uint64, error) {
		strconvParseUintCalled++
		return 0, nil
	}
	tryUnmarshalPrimitiveTypesFuncExpected = 0
	tryUnmarshalPrimitiveTypesFuncCalled = 0
	tryUnmarshalPrimitiveTypesFunc = func(value string, dataTemplate interface{}) bool {
		tryUnmarshalPrimitiveTypesFuncCalled++
		return false
	}
	prepareLoggingFuncExpected = 0
	prepareLoggingFuncCalled = 0
	prepareLoggingFunc = func(session *session, logType LogType, logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		prepareLoggingFuncCalled++
	}
	sortStringsExpected = 0
	sortStringsCalled = 0
	sortStrings = func(a []string) {
		sortStringsCalled++
	}
	stringsJoinExpected = 0
	stringsJoinCalled = 0
	stringsJoin = func(a []string, sep string) string {
		stringsJoinCalled++
		return ""
	}
	regexpMatchStringExpected = 0
	regexpMatchStringCalled = 0
	regexpMatchString = func(pattern string, s string) (bool, error) {
		regexpMatchStringCalled++
		return false, nil
	}
	reflectValueOfExpected = 0
	reflectValueOfCalled = 0
	reflectValueOf = func(i interface{}) reflect.Value {
		reflectValueOfCalled++
		return reflect.Value{}
	}
	ioutilReadAllExpected = 0
	ioutilReadAllCalled = 0
	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		ioutilReadAllCalled++
		return nil, nil
	}
	ioutilNopCloserExpected = 0
	ioutilNopCloserCalled = 0
	ioutilNopCloser = func(r io.Reader) io.ReadCloser {
		ioutilNopCloserCalled++
		return nil
	}
	bytesNewBufferExpected = 0
	bytesNewBufferCalled = 0
	bytesNewBuffer = func(buf []byte) *bytes.Buffer {
		bytesNewBufferCalled++
		return nil
	}
	logProcessEnterFuncExpected = 0
	logProcessEnterFuncCalled = 0
	logProcessEnterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessEnterFuncCalled++
	}
	logProcessExitFuncExpected = 0
	logProcessExitFuncCalled = 0
	logProcessExitFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessExitFuncCalled++
	}
	logProcessResponseFuncExpected = 0
	logProcessResponseFuncCalled = 0
	logProcessResponseFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessResponseFuncCalled++
	}
	logProcessRequestFuncExpected = 0
	logProcessRequestFuncCalled = 0
	logProcessRequestFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessRequestFuncCalled++
	}
	processSessionFuncExpected = 0
	processSessionFuncCalled = 0
	processSessionFunc = func(session Session, customization Customization) error {
		processSessionFuncCalled++
		return nil
	}
	httpStatusTextExpected = 0
	httpStatusTextCalled = 0
	httpStatusText = func(code int) string {
		httpStatusTextCalled++
		return ""
	}
	strconvItoaExpected = 0
	strconvItoaCalled = 0
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return ""
	}
	tryUnmarshalFuncExpected = 0
	tryUnmarshalFuncCalled = 0
	tryUnmarshalFunc = func(value string, dataTemplate interface{}) error {
		tryUnmarshalFuncCalled++
		return nil
	}
	jsonMarshalExpected = 0
	jsonMarshalCalled = 0
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		return nil, nil
	}
	runtimeCallerExpected = 0
	runtimeCallerCalled = 0
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		return 0, "", 0, false
	}
	runtimeFuncForPCExpected = 0
	runtimeFuncForPCCalled = 0
	runtimeFuncForPC = func(pc uintptr) *runtime.Func {
		runtimeFuncForPCCalled++
		return nil
	}
	getMethodNameFuncExpected = 0
	getMethodNameFuncCalled = 0
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return ""
	}
	logMethodEnterFuncExpected = 0
	logMethodEnterFuncCalled = 0
	logMethodEnterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodEnterFuncCalled++
	}
	logMethodParameterFuncExpected = 0
	logMethodParameterFuncCalled = 0
	logMethodParameterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodParameterFuncCalled++
	}
	logMethodLogicFuncExpected = 0
	logMethodLogicFuncCalled = 0
	logMethodLogicFunc = func(session *session, logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodLogicFuncCalled++
	}
	logMethodReturnFuncExpected = 0
	logMethodReturnFuncCalled = 0
	logMethodReturnFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodReturnFuncCalled++
	}
	logMethodExitFuncExpected = 0
	logMethodExitFuncCalled = 0
	logMethodExitFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodExitFuncCalled++
	}
	timeNowExpected = 0
	timeNowCalled = 0
	timeNow = func() time.Time {
		timeNowCalled++
		return time.Time{}
	}
	clientDoFuncExpected = 0
	clientDoFuncCalled = 0
	clientDoFunc = func(httpClient *http.Client, httpRequest *http.Request) (*http.Response, error) {
		clientDoFuncCalled++
		return nil, nil
	}
	timeSleepExpected = 0
	timeSleepCalled = 0
	timeSleep = func(time.Duration) {
		timeSleepCalled++
	}
	getHTTPTransportFuncExpected = 0
	getHTTPTransportFuncCalled = 0
	getHTTPTransportFunc = func(skipServerCertVerification bool, clientCertificate *tls.Certificate, roundTripperWrapper func(originalTransport http.RoundTripper) http.RoundTripper) http.RoundTripper {
		getHTTPTransportFuncCalled++
		return nil
	}
	stringsNewReaderExpected = 0
	stringsNewReaderCalled = 0
	stringsNewReader = func(s string) *strings.Reader {
		stringsNewReaderCalled++
		return nil
	}
	httpNewRequestExpected = 0
	httpNewRequestCalled = 0
	httpNewRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		httpNewRequestCalled++
		return nil, nil
	}
	logWebcallStartFuncExpected = 0
	logWebcallStartFuncCalled = 0
	logWebcallStartFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logWebcallStartFuncCalled++
	}
	logWebcallRequestFuncExpected = 0
	logWebcallRequestFuncCalled = 0
	logWebcallRequestFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logWebcallRequestFuncCalled++
	}
	logWebcallResponseFuncExpected = 0
	logWebcallResponseFuncCalled = 0
	logWebcallResponseFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logWebcallResponseFuncCalled++
	}
	logWebcallFinishFuncExpected = 0
	logWebcallFinishFuncCalled = 0
	logWebcallFinishFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logWebcallFinishFuncCalled++
	}
	createHTTPRequestFuncExpected = 0
	createHTTPRequestFuncCalled = 0
	createHTTPRequestFunc = func(webRequest *webRequest) (*http.Request, error) {
		createHTTPRequestFuncCalled++
		return nil, nil
	}
	getClientForRequestFuncExpected = 0
	getClientForRequestFuncCalled = 0
	getClientForRequestFunc = func(sendClientCert bool) *http.Client {
		getClientForRequestFuncCalled++
		return nil
	}
	clientDoWithRetryFuncExpected = 0
	clientDoWithRetryFuncCalled = 0
	clientDoWithRetryFunc = func(httpClient *http.Client, httpRequest *http.Request, connectivityRetryCount int, httpStatusRetryCount map[int]int, retryDelay time.Duration) (*http.Response, error) {
		clientDoWithRetryFuncCalled++
		return nil, nil
	}
	logErrorResponseFuncExpected = 0
	logErrorResponseFuncCalled = 0
	logErrorResponseFunc = func(session *session, responseError error, startTime time.Time) {
		logErrorResponseFuncCalled++
	}
	logSuccessResponseFuncExpected = 0
	logSuccessResponseFuncCalled = 0
	logSuccessResponseFunc = func(session *session, response *http.Response, startTime time.Time) {
		logSuccessResponseFuncCalled++
	}
	doRequestProcessingFuncExpected = 0
	doRequestProcessingFuncCalled = 0
	doRequestProcessingFunc = func(webRequest *webRequest) (*http.Response, error) {
		doRequestProcessingFuncCalled++
		return nil, nil
	}
	parseResponseFuncExpected = 0
	parseResponseFuncCalled = 0
	parseResponseFunc = func(session *session, body io.ReadCloser, dataTemplate interface{}) error {
		parseResponseFuncCalled++
		return nil
	}
}

func verifyAll(t *testing.T) {
	isInterfaceValueNilFunc = isInterfaceValueNil
	assert.Equal(t, isInterfaceValueNilFuncExpected, isInterfaceValueNilFuncCalled, "Unexpected number of calls to method isInterfaceValueNilFunc")
	uuidNew = uuid.New
	assert.Equal(t, uuidNewExpected, uuidNewCalled, "Unexpected number of calls to method uuidNew")
	startApplicationFunc = startApplication
	assert.Equal(t, startApplicationFuncExpected, startApplicationFuncCalled, "Unexpected number of calls to method startApplicationFunc")
	preBootstrapingFunc = preBootstraping
	assert.Equal(t, preBootstrapingFuncExpected, preBootstrapingFuncCalled, "Unexpected number of calls to method preBootstrapingFunc")
	bootstrapFunc = bootstrap
	assert.Equal(t, bootstrapFuncExpected, bootstrapFuncCalled, "Unexpected number of calls to method bootstrapFunc")
	postBootstrapingFunc = postBootstraping
	assert.Equal(t, postBootstrapingFuncExpected, postBootstrapingFuncCalled, "Unexpected number of calls to method postBootstrapingFunc")
	endApplicationFunc = endApplication
	assert.Equal(t, endApplicationFuncExpected, endApplicationFuncCalled, "Unexpected number of calls to method endApplicationFunc")
	beginApplicationFunc = beginApplication
	assert.Equal(t, beginApplicationFuncExpected, beginApplicationFuncCalled, "Unexpected number of calls to method beginApplicationFunc")
	logAppRootFunc = logAppRoot
	assert.Equal(t, logAppRootFuncExpected, logAppRootFuncCalled, "Unexpected number of calls to method logAppRootFunc")
	handleSessionFunc = handleSession
	assert.Equal(t, handleSessionFuncExpected, handleSessionFuncCalled, "Unexpected number of calls to method handleSessionFunc")
	runInstancesFunc = runInstances
	assert.Equal(t, runInstancesFuncExpected, runInstancesFuncCalled, "Unexpected number of calls to method runInstancesFunc")
	repeatExecutionFunc = repeatExecution
	assert.Equal(t, repeatExecutionFuncExpected, repeatExecutionFuncCalled, "Unexpected number of calls to method repeatExecutionFunc")
	timeAfter = time.After
	assert.Equal(t, timeAfterExpected, timeAfterCalled, "Unexpected number of calls to method timeAfter")
	runApplicationFunc = runApplication
	assert.Equal(t, runApplicationFuncExpected, runApplicationFuncCalled, "Unexpected number of calls to method runApplicationFunc")
	initializeHTTPClientsFunc = initializeHTTPClients
	assert.Equal(t, initializeHTTPClientsFuncExpected, initializeHTTPClientsFuncCalled, "Unexpected number of calls to method initializeHTTPClientsFunc")
	fmtPrintf = fmt.Printf
	assert.Equal(t, fmtPrintfExpected, fmtPrintfCalled, "Unexpected number of calls to method fmtPrintf")
	fmtSprintf = fmt.Sprintf
	assert.Equal(t, fmtSprintfExpected, fmtSprintfCalled, "Unexpected number of calls to method fmtSprintf")
	marshalIgnoreErrorFunc = marshalIgnoreError
	assert.Equal(t, marshalIgnoreErrorFuncExpected, marshalIgnoreErrorFuncCalled, "Unexpected number of calls to method marshalIgnoreErrorFunc")
	debugStack = debug.Stack
	assert.Equal(t, debugStackExpected, debugStackCalled, "Unexpected number of calls to debugStack")
	stringsSplit = strings.Split
	assert.Equal(t, stringsSplitExpected, stringsSplitCalled, "Unexpected number of calls to method stringsSplit")
	strconvAtoi = strconv.Atoi
	assert.Equal(t, strconvAtoiExpected, strconvAtoiCalled, "Unexpected number of calls to method strconvAtoi")
	initiateSessionFunc = initiateSession
	assert.Equal(t, initiateSessionFuncExpected, initiateSessionFuncCalled, "Unexpected number of calls to method initiateSessionFunc")
	getTimeNowUTCFunc = getTimeNowUTC
	assert.Equal(t, getTimeNowUTCFuncExpected, getTimeNowUTCFuncCalled, "Unexpected number of calls to method getTimeNowUTCFunc")
	finalizeSessionFunc = finalizeSession
	assert.Equal(t, finalizeSessionFuncExpected, finalizeSessionFuncCalled, "Unexpected number of calls to method finalizeSessionFunc")
	timeSince = time.Since
	assert.Equal(t, timeSinceExpected, timeSinceCalled, "Unexpected number of calls to method timeSince")
	jsonNewEncoder = json.NewEncoder
	assert.Equal(t, jsonNewEncoderExpected, jsonNewEncoderCalled, "Unexpected number of calls to jsonNewEncoder")
	stringsTrimRight = strings.TrimRight
	assert.Equal(t, stringsTrimRightExpected, stringsTrimRightCalled, "Unexpected number of calls to stringsTrimRight")
	jsonUnmarshal = json.Unmarshal
	assert.Equal(t, jsonUnmarshalExpected, jsonUnmarshalCalled, "Unexpected number of calls to jsonUnmarshal")
	fmtErrorf = fmt.Errorf
	assert.Equal(t, fmtErrorfExpected, fmtErrorfCalled, "Unexpected number of calls to fmtErrorf")
	reflectTypeOf = reflect.TypeOf
	assert.Equal(t, reflectTypeOfExpected, reflectTypeOfCalled, "Unexpected number of calls to reflectTypeOf")
	stringsToLower = strings.ToLower
	assert.Equal(t, stringsToLowerExpected, stringsToLowerCalled, "Unexpected number of calls to stringsToLower")
	strconvParseBool = strconv.ParseBool
	assert.Equal(t, strconvParseBoolExpected, strconvParseBoolCalled, "Unexpected number of calls to strconvParseBool")
	strconvParseInt = strconv.ParseInt
	assert.Equal(t, strconvParseIntExpected, strconvParseIntCalled, "Unexpected number of calls to strconvParseInt")
	strconvParseFloat = strconv.ParseFloat
	assert.Equal(t, strconvParseFloatExpected, strconvParseFloatCalled, "Unexpected number of calls to strconvParseFloat")
	strconvParseUint = strconv.ParseUint
	assert.Equal(t, strconvParseUintExpected, strconvParseUintCalled, "Unexpected number of calls to strconvParseUint")
	tryUnmarshalPrimitiveTypesFunc = tryUnmarshalPrimitiveTypes
	assert.Equal(t, tryUnmarshalPrimitiveTypesFuncExpected, tryUnmarshalPrimitiveTypesFuncCalled, "Unexpected number of calls to tryUnmarshalPrimitiveTypesFunc")
	prepareLoggingFunc = prepareLogging
	assert.Equal(t, prepareLoggingFuncExpected, prepareLoggingFuncCalled, "Unexpected number of calls to prepareLoggingFunc")
	sortStrings = sort.Strings
	assert.Equal(t, sortStringsExpected, sortStringsCalled, "Unexpected number of calls to sortStrings")
	stringsJoin = strings.Join
	assert.Equal(t, stringsJoinExpected, stringsJoinCalled, "Unexpected number of calls to stringsJoin")
	regexpMatchString = regexp.MatchString
	assert.Equal(t, regexpMatchStringExpected, regexpMatchStringCalled, "Unexpected number of calls to regexpMatchString")
	reflectValueOf = reflect.ValueOf
	assert.Equal(t, reflectValueOfExpected, reflectValueOfCalled, "Unexpected number of calls to reflectValueOf")
	ioutilReadAll = ioutil.ReadAll
	assert.Equal(t, ioutilReadAllExpected, ioutilReadAllCalled, "Unexpected number of calls to ioutilReadAll")
	ioutilNopCloser = ioutil.NopCloser
	assert.Equal(t, ioutilNopCloserExpected, ioutilNopCloserCalled, "Unexpected number of calls to ioutilNopCloser")
	bytesNewBuffer = bytes.NewBuffer
	assert.Equal(t, bytesNewBufferExpected, bytesNewBufferCalled, "Unexpected number of calls to bytesNewBuffer")
	logProcessEnterFunc = logProcessEnter
	assert.Equal(t, logProcessEnterFuncExpected, logProcessEnterFuncCalled, "Unexpected number of calls to method logProcessEnterFunc")
	logProcessExitFunc = logProcessExit
	assert.Equal(t, logProcessExitFuncExpected, logProcessExitFuncCalled, "Unexpected number of calls to method logProcessExitFunc")
	logProcessResponseFunc = logProcessResponse
	assert.Equal(t, logProcessResponseFuncExpected, logProcessResponseFuncCalled, "Unexpected number of calls to method logProcessResponseFunc")
	logProcessRequestFunc = logProcessRequest
	assert.Equal(t, logProcessRequestFuncExpected, logProcessRequestFuncCalled, "Unexpected number of calls to method logProcessRequestFunc")
	processSessionFunc = processSession
	assert.Equal(t, processSessionFuncExpected, processSessionFuncCalled, "Unexpected number of calls to method processSessionFunc")
	httpStatusText = http.StatusText
	assert.Equal(t, httpStatusTextExpected, httpStatusTextCalled, "Unexpected number of calls to method httpStatusText")
	strconvItoa = strconv.Itoa
	assert.Equal(t, strconvItoaExpected, strconvItoaCalled, "Unexpected number of calls to method strconvItoa")
	tryUnmarshalFunc = tryUnmarshal
	assert.Equal(t, tryUnmarshalFuncExpected, tryUnmarshalFuncCalled, "Unexpected number of calls to method tryUnmarshalFunc")
	jsonMarshal = json.Marshal
	assert.Equal(t, jsonMarshalExpected, jsonMarshalCalled, "Unexpected number of calls to method jsonMarshal")
	runtimeCaller = runtime.Caller
	assert.Equal(t, runtimeCallerExpected, runtimeCallerCalled, "Unexpected number of calls to method runtimeCaller")
	runtimeFuncForPC = runtime.FuncForPC
	assert.Equal(t, runtimeFuncForPCExpected, runtimeFuncForPCCalled, "Unexpected number of calls to method runtimeFuncForPC")
	getMethodNameFunc = getMethodName
	assert.Equal(t, getMethodNameFuncExpected, getMethodNameFuncCalled, "Unexpected number of calls to method getMethodNameFunc")
	logMethodEnterFunc = logMethodEnter
	assert.Equal(t, logMethodEnterFuncExpected, logMethodEnterFuncCalled, "Unexpected number of calls to method logMethodEnterFunc")
	logMethodParameterFunc = logMethodParameter
	assert.Equal(t, logMethodParameterFuncExpected, logMethodParameterFuncCalled, "Unexpected number of calls to method logMethodParameterFunc")
	logMethodLogicFunc = logMethodLogic
	assert.Equal(t, logMethodLogicFuncExpected, logMethodLogicFuncCalled, "Unexpected number of calls to method logMethodLogicFunc")
	logMethodReturnFunc = logMethodReturn
	assert.Equal(t, logMethodReturnFuncExpected, logMethodReturnFuncCalled, "Unexpected number of calls to method logMethodReturnFunc")
	logMethodExitFunc = logMethodExit
	assert.Equal(t, logMethodExitFuncExpected, logMethodExitFuncCalled, "Unexpected number of calls to method logMethodExitFunc")
	timeNow = time.Now
	assert.Equal(t, timeNowExpected, timeNowCalled, "Unexpected number of calls to timeNow")
	clientDoFunc = clientDo
	assert.Equal(t, clientDoFuncExpected, clientDoFuncCalled, "Unexpected number of calls to method clientDoFunc")
	timeSleep = time.Sleep
	assert.Equal(t, timeSleepExpected, timeSleepCalled, "Unexpected number of calls to method timeSleep")
	getHTTPTransportFunc = getHTTPTransport
	assert.Equal(t, getHTTPTransportFuncExpected, getHTTPTransportFuncCalled, "Unexpected number of calls to method getHTTPTransportFunc")
	stringsNewReader = strings.NewReader
	assert.Equal(t, stringsNewReaderExpected, stringsNewReaderCalled, "Unexpected number of calls to method stringsNewReader")
	httpNewRequest = http.NewRequest
	assert.Equal(t, httpNewRequestExpected, httpNewRequestCalled, "Unexpected number of calls to method httpNewRequest")
	logWebcallStartFunc = logWebcallStart
	assert.Equal(t, logWebcallStartFuncExpected, logWebcallStartFuncCalled, "Unexpected number of calls to method logWebcallStartFunc")
	logWebcallRequestFunc = logWebcallRequest
	assert.Equal(t, logWebcallRequestFuncExpected, logWebcallRequestFuncCalled, "Unexpected number of calls to method logWebcallRequestFunc")
	logWebcallResponseFunc = logWebcallResponse
	assert.Equal(t, logWebcallResponseFuncExpected, logWebcallResponseFuncCalled, "Unexpected number of calls to method logWebcallResponseFunc")
	logWebcallFinishFunc = logWebcallFinish
	assert.Equal(t, logWebcallFinishFuncExpected, logWebcallFinishFuncCalled, "Unexpected number of calls to method logWebcallFinishFunc")
	createHTTPRequestFunc = createHTTPRequest
	assert.Equal(t, createHTTPRequestFuncExpected, createHTTPRequestFuncCalled, "Unexpected number of calls to method createHTTPRequestFunc")
	getClientForRequestFunc = getClientForRequest
	assert.Equal(t, getClientForRequestFuncExpected, getClientForRequestFuncCalled, "Unexpected number of calls to method getClientForRequestFunc")
	clientDoWithRetryFunc = clientDoWithRetry
	assert.Equal(t, clientDoWithRetryFuncExpected, clientDoWithRetryFuncCalled, "Unexpected number of calls to method clientDoWithRetryFunc")
	logErrorResponseFunc = logErrorResponse
	assert.Equal(t, logErrorResponseFuncExpected, logErrorResponseFuncCalled, "Unexpected number of calls to method logErrorResponseFunc")
	logSuccessResponseFunc = logSuccessResponse
	assert.Equal(t, logSuccessResponseFuncExpected, logSuccessResponseFuncCalled, "Unexpected number of calls to method logSuccessResponseFunc")
	doRequestProcessingFunc = doRequestProcessing
	assert.Equal(t, doRequestProcessingFuncExpected, doRequestProcessingFuncCalled, "Unexpected number of calls to method doRequestProcessingFunc")
	parseResponseFunc = parseResponse
	assert.Equal(t, parseResponseFuncExpected, parseResponseFuncCalled, "Unexpected number of calls to method parseResponseFunc")
}

func functionPointerEquals(t *testing.T, expectFunc interface{}, actualFunc interface{}) {
	var expectValue = fmt.Sprintf("%v", reflect.ValueOf(expectFunc))
	var actualValue = fmt.Sprintf("%v", reflect.ValueOf(actualFunc))
	assert.Equal(t, expectValue, actualValue)
}

// mock structs
type dummyApplication struct {
	t *testing.T
}

func (application *dummyApplication) Start(numberOfInstances int, repeatPeriod *time.Duration) {
	assert.Fail(application.t, "Unexpected call to Start")
}

func (application *dummyApplication) Stop() {
	assert.Fail(application.t, "Unexpected call to Stop")
}

type dummyCustomization struct {
	t *testing.T
}

func (customization *dummyCustomization) PreBootstrap() error {
	assert.Fail(customization.t, "Unexpected call to PreBootstrap")
	return nil
}

func (customization *dummyCustomization) PostBootstrap() error {
	assert.Fail(customization.t, "Unexpected call to PostBootstrap")
	return nil
}

func (customization *dummyCustomization) AppClosing() error {
	assert.Fail(customization.t, "Unexpected call to AppClosing")
	return nil
}

func (customization *dummyCustomization) Log(session Session, logType LogType, logLevel LogLevel, category, subcategory, description string) {
	assert.Fail(customization.t, "Unexpected call to Log")
}

func (customization *dummyCustomization) PreAction(session Session) error {
	assert.Fail(customization.t, "Unexpected call to PreAction")
	return nil
}

func (customization *dummyCustomization) PostAction(session Session) error {
	assert.Fail(customization.t, "Unexpected call to PostAction")
	return nil
}

func (customization *dummyCustomization) ActionFunc(session Session) error {
	assert.Fail(customization.t, "Unexpected call to ActionFunc")
	return nil
}

func (customization *dummyCustomization) RecoverPanic(session Session, recoverResult interface{}) error {
	assert.Fail(customization.t, "Unexpected call to RecoverPanic")
	return nil
}

func (customization *dummyCustomization) ClientCert() *tls.Certificate {
	assert.Fail(customization.t, "Unexpected call to ClientCert")
	return nil
}

func (customization *dummyCustomization) DefaultTimeout() time.Duration {
	assert.Fail(customization.t, "Unexpected call to DefaultTimeout")
	return 0
}

func (customization *dummyCustomization) SkipServerCertVerification() bool {
	assert.Fail(customization.t, "Unexpected call to SkipServerCertVerification")
	return false
}

func (customization *dummyCustomization) RoundTripper(originalTransport http.RoundTripper) http.RoundTripper {
	assert.Fail(customization.t, "Unexpected call to RoundTripper")
	return nil
}

func (customization *dummyCustomization) WrapRequest(session Session, httpRequest *http.Request) *http.Request {
	assert.Fail(customization.t, "Unexpected call to WrapRequest")
	return nil
}

type dummySession struct {
	t *testing.T
}

func (session *dummySession) GetID() uuid.UUID {
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySession) GetIndex() int {
	assert.Fail(session.t, "Unexpected call to GetIndex")
	return 0
}

func (session *dummySession) Attach(name string, value interface{}) bool {
	assert.Fail(session.t, "Unexpected call to Attach")
	return false
}

func (session *dummySession) Detach(name string) bool {
	assert.Fail(session.t, "Unexpected call to Detach")
	return false
}

func (session *dummySession) GetRawAttachment(name string) (interface{}, bool) {
	assert.Fail(session.t, "Unexpected call to GetRawAttachment")
	return nil, false
}

func (session *dummySession) GetAttachment(name string, dataTemplate interface{}) bool {
	assert.Fail(session.t, "Unexpected call to GetAttachment")
	return false
}

func (session *dummySession) LogMethodEnter() {
	assert.Fail(session.t, "Unexpected call to LogMethodEnter")
}

func (session *dummySession) LogMethodParameter(parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodParameter")
}

func (session *dummySession) LogMethodLogic(logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodLogic")
}

func (session *dummySession) LogMethodReturn(returns ...interface{}) {
	assert.Fail(session.t, "Unexpected call to LogMethodReturn")
}

func (session *dummySession) LogMethodExit() {
	assert.Fail(session.t, "Unexpected call to LogMethodExit")
}

func (session *dummySession) CreateWebcallRequest(method string, url string, payload string, header map[string]string, sendClientCert bool) WebRequest {
	assert.Fail(session.t, "Unexpected call to CreateWebcallRequest")
	return nil
}

type dummyTransport struct {
	t *testing.T
}

func (transport *dummyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	assert.Fail(transport.t, "Unexpected call to RoundTrip")
	return nil, nil
}
