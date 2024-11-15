package jobrunner

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestSessionGetID_NilSessionObject(t *testing.T) {
	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetID()

	// assert
	assert.Zero(t, result)
}

func TestSessionGetID_ValidSessionObject(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// act
	var result = dummySession.GetID()

	// assert
	assert.Equal(t, dummySessionID, result)
}

func TestSessionGetIndex_NilSessionObject(t *testing.T) {
	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetIndex()

	// assert
	assert.Zero(t, result)
}

func TestSessionGetIndex_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyIndex = rand.Int()

	// SUT
	var dummySession = &session{
		index: dummyIndex,
	}

	// act
	var result = dummySession.GetIndex()

	// assert
	assert.Equal(t, dummyIndex, result)
}

func TestSessionGetReruns_NilSessionObject(t *testing.T) {
	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetReruns()

	// assert
	assert.Zero(t, result)
}

func TestSessionGetReruns_ValidSessionObject(t *testing.T) {
	// arrange
	var dummyIndex = rand.Int()

	// SUT
	var dummySession = &session{
		reruns: dummyIndex,
	}

	// act
	var result = dummySession.GetReruns()

	// assert
	assert.Equal(t, dummyIndex, result)
}

func TestSessionAttach_NilSessionObject(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.IntN(100),
	}

	// SUT
	var dummySession *session

	// act
	var result = dummySession.Attach(
		dummyName,
		dummyValue,
	)

	// assert
	assert.False(t, result)
}

func TestSessionAttach_NoAttachment(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.IntN(100),
	}

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
}

func TestSessionAttach_WithAttachment(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.IntN(100),
	}

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
}

func TestSessionDetach_NilSessionObject(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// SUT
	var dummySession *session

	// act
	var result = dummySession.Detach(
		dummyName,
	)

	// assert
	assert.False(t, result)
}

func TestSessionDetach_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

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
}

func TestSessionDetach_WithAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

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
}

func TestSessionGetRawAttachment_NoSession(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// SUT
	var dummySession *session

	// act
	var result, found = dummySession.GetRawAttachment(
		dummyName,
	)

	// assert
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestSessionGetRawAttachment_NoAttachment(t *testing.T) {
	// arrange
	var dummyName = "some name"

	// SUT
	var dummySession = &session{}

	// act
	var result, found = dummySession.GetRawAttachment(
		dummyName,
	)

	// assert
	assert.Nil(t, result)
	assert.False(t, found)
}

func TestSessionGetRawAttachment_Success(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		ID:   uuid.New(),
		Foo:  "bar",
		Test: rand.IntN(100),
	}

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
}

func TestSessionGetAttachment_NoSession(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyDataTemplate map[string]interface{}

	// SUT
	var dummySession *session

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, result)
}

func TestSessionGetAttachment_NoAttachment(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyDataTemplate dummyAttachment

	// SUT
	var dummySession = &session{}

	// act
	var result = dummySession.GetAttachment(
		dummyName,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, result)
}

func TestSessionGetAttachment_MarshalError(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.IntN(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(json.Marshal).Expects(dummyValue).Returns(nil, errors.New("some marshal error")).Once()

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
	assert.Zero(t, result)
}

func TestSessionGetAttachment_UnmarshalError(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.IntN(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate int

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
	assert.Zero(t, result)
}

func TestSessionGetAttachment_Success(t *testing.T) {
	// arrange
	type dummyAttachment struct {
		ID   uuid.UUID
		Foo  string
		Test int
	}
	var dummyName = "some name"
	var dummyValue = dummyAttachment{
		Foo:  "bar",
		Test: rand.IntN(100),
		ID:   uuid.New(),
	}
	var dummyDataTemplate dummyAttachment

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
}

func TestSessionGetMethodName_UnknownCaller(t *testing.T) {
	// arrange
	var dummyPC = uintptr(rand.Int())
	var dummyFile = "some file"
	var dummyLine = rand.Int()
	var dummyOK = false

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(runtime.Caller).Expects(3).Returns(dummyPC, dummyFile, dummyLine, dummyOK).Once()

	// SUT + act
	var result = getMethodName()

	// assert
	assert.Equal(t, "?", result)
}

func TestSessionGetMethodName_HappyPath(t *testing.T) {
	// SUT + act
	var result = getMethodName()

	// assert
	assert.Contains(t, result, "runtime.goexit")
}

func TestSessionLogMethodEnter(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	var m = gomocker.NewMocker(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	m.Mock(getMethodName).Expects().Returns(dummyMethodName).Once()
	m.Mock(logMethodEnter).Expects(dummySession, dummyMethodName, "", "").Returns().Once()

	// act
	dummySession.LogMethodEnter()
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
	var m = gomocker.NewMocker(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	m.Mock(getMethodName).Expects().Returns(dummyMethodName).Once()
	m.Mock(logMethodParameter).Expects(dummySession, dummyMethodName, "0", "%v", dummyParameters[0]).Returns().Once()
	m.Mock(logMethodParameter).Expects(dummySession, dummyMethodName, "1", "%v", dummyParameters[1]).Returns().Once()
	m.Mock(logMethodParameter).Expects(dummySession, dummyMethodName, "2", "%v", dummyParameters[2]).Returns().Once()

	// act
	dummySession.LogMethodParameter(
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
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
	var m = gomocker.NewMocker(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	m.Mock(logMethodLogic).Expects(
		dummySession,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

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
	var m = gomocker.NewMocker(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	m.Mock(getMethodName).Expects().Returns(dummyMethodName).Once()
	m.Mock(logMethodReturn).Expects(dummySession, dummyMethodName, "0", "%v", dummyReturns[0]).Returns().Once()
	m.Mock(logMethodReturn).Expects(dummySession, dummyMethodName, "1", "%v", dummyReturns[1]).Returns().Once()
	m.Mock(logMethodReturn).Expects(dummySession, dummyMethodName, "2", "%v", dummyReturns[2]).Returns().Once()

	// act
	dummySession.LogMethodReturn(
		dummyReturn1,
		dummyReturn2,
		dummyReturn3,
	)
}

func TestSessionLogMethodExit(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethodName = "some method name"

	// mock
	var m = gomocker.NewMocker(t)

	// SUT
	var dummySession = &session{
		id: dummySessionID,
	}

	// expect
	m.Mock(getMethodName).Expects().Returns(dummyMethodName).Once()
	m.Mock(logMethodExit).Expects(dummySession, dummyMethodName, "", "").Returns().Once()

	// act
	dummySession.LogMethodExit()
}

func TestSessionCreateWebcallRequest(t *testing.T) {
	// arrange
	var dummySessionID = uuid.New()
	var dummyMethod = "some method"
	var dummyURL = "some URL"
	var dummyPayload = "some payload"
	var dummySendClientCert = rand.IntN(100) < 50

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
}
