package jobrunner

import (
	"encoding/json"
	"errors"
	"math/rand"
	"runtime"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSessionGetID_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetID()

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestSessionGetID_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// act
	var result = dummySession.GetID()

	// assert
	assert.Equal(t, dummySessionID, result)

	// verify
	verifyAll(t)
}

func TestSessionGetIndex_NilSessionObject(t *testing.T) {
	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetIndex()

	// assert
	assert.Zero(t, result)

	// verify
	verifyAll(t)
}

func TestSessionGetIndex_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyIndex = rand.Int()

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		index: dummyIndex,
	}

	// act
	var result = dummySession.GetIndex()

	// assert
	assert.Equal(t, dummyIndex, result)

	// verify
	verifyAll(t)
}

type dummyAttachment struct {
	ID   uuid.UUID
	Foo  string
	Test int
}

func TestSessionAttach_NilSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result = dummySession.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestSessionAttach_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		attachment: nil,
	}

	// act
	var result = dummySession.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummySession.attachment[dummyName])

	// verify
	verifyAll(t)
}

func TestSessionAttach_WithAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.Intn(100),
	}

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// act
	var result = dummySession.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummySession.attachment[dummyName])

	// verify
	verifyAll(t)
}

func TestSessionDetach_NilSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result = dummySession.Detach(
		dummyName,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestSessionDetach_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		attachment: nil,
	}

	// act
	var result = dummySession.Detach(
		dummyName,
	)

	// assert
	assert.True(t, result)
	var _, found = dummySession.attachment[dummyName]
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestSessionDetach_WithAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: "some value",
		},
	}

	// act
	var result = dummySession.Detach(
		dummyName,
	)

	// assert
	assert.True(t, result)
	var _, found = dummySession.attachment[dummyName]
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestSessionGetRawAttachment_NoSession(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result, found = dummySession.GetRawAttachment(
		dummyName,
	)

	// assert
	assert.Nil(t, result)
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestSessionGetRawAttachment_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{}

	// act
	var result, found = dummySession.GetRawAttachment(
		dummyName,
	)

	// assert
	assert.Nil(t, result)
	assert.False(t, found)

	// verify
	verifyAll(t)
}

func TestSessionGetRawAttachment_Success(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result, found = dummySession.GetRawAttachment(
		dummyName,
	)

	// assert
	assert.Equal(t, dummyValue, result)
	assert.True(t, found)

	// verify
	verifyAll(t)
}

func TestSessionGetAttachment_NoSession(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestSessionGetAttachment_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{}

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestSessionGetAttachment_MarshalError(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return nil, errors.New("some marshal error")
	}

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestSessionGetAttachment_UnmarshalError(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate int

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return json.Marshal(v)
	}
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return json.Unmarshal(data, v)
	}

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestSessionGetAttachment_Success(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.Intn(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

	// mock
	createMock(t)

	// expect
	jsonMarshalExpected = 1
	jsonMarshal = func(v interface{}) ([]byte, error) {
		jsonMarshalCalled++
		assert.Equal(t, dummyValue, v)
		return json.Marshal(v)
	}
	jsonUnmarshalExpected = 1
	jsonUnmarshal = func(data []byte, v interface{}) error {
		jsonUnmarshalCalled++
		return json.Unmarshal(data, v)
	}

	// SUT
	var dummySession = &session{
		attachment: map[string]interface{}{
			dummyName: dummyValue,
		},
	}

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)

	// verify
	verifyAll(t)
}

func TestSessionGetMethodName_UnknownCaller(t *testing.T) {
	// arrange
	var dummyPC = uintptr(rand.Int())
	var dummyFile = "some file"
	var dummyLine = rand.Int()
	var dummyOK = false

	// mock
	createMock(t)

	// expect
	runtimeCallerExpected = 1
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		assert.Equal(t, 3, skip)
		return dummyPC, dummyFile, dummyLine, dummyOK
	}

	// SUT + act
	var result = getMethodName()

	// assert
	assert.Equal(t, "?", result)

	// verify
	verifyAll(t)
}

func TestSessionGetMethodName_HappyPath(t *testing.T) {
	// mock
	createMock(t)

	// expect
	runtimeCallerExpected = 1
	runtimeCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		runtimeCallerCalled++
		assert.Equal(t, 3, skip)
		return runtime.Caller(2)
	}
	runtimeFuncForPCExpected = 1
	runtimeFuncForPC = func(pc uintptr) *runtime.Func {
		runtimeFuncForPCCalled++
		assert.NotZero(t, pc)
		return runtime.FuncForPC(pc)
	}

	// SUT + act
	var result = getMethodName()

	// assert
	assert.Contains(t, result, "TestSessionGetMethodName_HappyPath")

	// verify
	verifyAll(t)
}

func TestSessionLogMethodEnter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	logMethodEnterFuncExpected = 1
	logMethodEnterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodEnterFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Zero(t, subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// act
	dummySession.LogMethodEnter()

	// verify
	verifyAll(t)
}

func TestSessionLogMethodParameter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyParameter1 = "foo"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("test")
	var dummyParameters = []interface{}{
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	}
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	strconvItoaExpected = 3
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	logMethodParameterFuncExpected = 3
	logMethodParameterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodParameterFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Equal(t, strconv.Itoa(logMethodParameterFuncCalled-1), subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyParameters[logMethodParameterFuncCalled-1], parameters[0])
	}

	// act
	dummySession.LogMethodParameter(
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// verify
	verifyAll(t)
}

func TestSessionLogMethodLogic(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyLogLevel = LogLevel(rand.Int())
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "foo"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("test")

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	logMethodLogicFuncExpected = 1
	logMethodLogicFunc = func(session *session, logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodLogicFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyLogLevel, logLevel)
		assert.Equal(t, dummyCategory, category)
		assert.Equal(t, dummySubcategory, subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 3, len(parameters))
		assert.Equal(t, dummyParameter1, parameters[0])
		assert.Equal(t, dummyParameter2, parameters[1])
		assert.Equal(t, dummyParameter3, parameters[2])
	}

	// act
	dummySession.LogMethodLogic(
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

	// verify
	verifyAll(t)
}

func TestSessionLogMethodReturn(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyReturn1 = "foo"
	var dummyReturn2 = rand.Int()
	var dummyReturn3 = errors.New("test")
	var dummyReturns = []interface{}{
		dummyReturn1,
		dummyReturn2,
		dummyReturn3,
	}
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	strconvItoaExpected = 3
	strconvItoa = func(i int) string {
		strconvItoaCalled++
		return strconv.Itoa(i)
	}
	logMethodReturnFuncExpected = 3
	logMethodReturnFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodReturnFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Equal(t, strconv.Itoa(logMethodReturnFuncCalled-1), subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyReturns[logMethodReturnFuncCalled-1], parameters[0])
	}

	// act
	dummySession.LogMethodReturn(
		dummyReturn1,
		dummyReturn2,
		dummyReturn3,
	)

	// verify
	verifyAll(t)
}

func TestSessionLogMethodExit(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	getMethodNameFuncExpected = 1
	getMethodNameFunc = func() string {
		getMethodNameFuncCalled++
		return dummyMethodName
	}
	logMethodExitFuncExpected = 1
	logMethodExitFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logMethodExitFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyMethodName, category)
		assert.Zero(t, subcategory)
		assert.Zero(t, messageFormat)
		assert.Empty(t, parameters)
	}

	// act
	dummySession.LogMethodExit()

	// verify
	verifyAll(t)
}

func TestSessionCreateWebcallRequest(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummySendClientCert = rand.Intn(100) < 50

	// mock
	createMock(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// act
	var result = dummySession.CreateWebcallRequest(
		dummyMethod,
		dummyURL,
		dummyPayload,
		dummySendClientCert,
	)
	var webrequest, ok = result.(*webRequest)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummySession, webrequest.session)
	assert.Equal(t, dummyMethod, webrequest.method)
	assert.Equal(t, dummyURL, webrequest.url)
	assert.Equal(t, dummyPayload, webrequest.payload)
	assert.NotNil(t, webrequest.query)
	assert.Empty(t, webrequest.query)
	assert.NotNil(t, webrequest.header)
	assert.Empty(t, webrequest.header)
	assert.Zero(t, webrequest.connRetry)
	assert.Nil(t, webrequest.httpRetry)
	assert.Equal(t, dummySendClientCert, webrequest.sendClientCert)
	assert.Zero(t, webrequest.retryDelay)
	assert.Empty(t, webrequest.dataReceivers)

	// verify
	verifyAll(t)
}
