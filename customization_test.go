package jobrunner

import (
	"errors"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDefaultCustomization_PreBootstrap(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var err = customizationDefault.PreBootstrap()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_PostBootstrap(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var err = customizationDefault.PostBootstrap()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_AppClosing(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var err = customizationDefault.AppClosing()

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

type dummySessionLog struct {
	dummySession
	getID    func() uuid.UUID
	getIndex func() int
}

func (session *dummySessionLog) GetID() uuid.UUID {
	if session.getID != nil {
		return session.getID()
	}
	assert.Fail(session.t, "Unexpected call to GetID")
	return uuid.Nil
}

func (session *dummySessionLog) GetIndex() int {
	if session.getIndex != nil {
		return session.getIndex()
	}
	assert.Fail(session.t, "Unexpected call to GetIndex")
	return 0
}

func TestDefaultCustomization_Log_NilSession(t *testing.T) {
	// arrange
	var dummySession *session
	var dummyLogType = LogType(rand.Intn(100))
	var dummyLogLevel = LogLevel(rand.Intn(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyDescription = "some description"

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummySession, i)
		return true
	}

	// SUT + act
	customizationDefault.Log(
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_Log_HappyPath(t *testing.T) {
	// arrange
	var dummySession = &dummySessionLog{dummySession: dummySession{t: t}}
	var dummyLogType = LogType(rand.Intn(100))
	var dummyLogLevel = LogLevel(rand.Intn(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyDescription = "some description"
	var sessionGetIDExpected int
	var sessionGetIDCalled int
	var sessionGetIndexExpected int
	var sessionGetIndexCalled int
	var dummySessionID = uuid.New()
	var dummySessionIndex = rand.Int()
	var dummyFormat = "<%v|%v> (%v|%v) [%v|%v] %v\n"

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummySession, i)
		return false
	}
	sessionGetIDExpected = 1
	dummySession.getID = func() uuid.UUID {
		sessionGetIDCalled++
		return dummySessionID
	}
	sessionGetIndexExpected = 1
	dummySession.getIndex = func() int {
		sessionGetIndexCalled++
		return dummySessionIndex
	}
	fmtPrintfExpected = 1
	fmtPrintf = func(format string, a ...interface{}) (n int, err error) {
		fmtPrintfCalled++
		assert.Equal(t, dummyFormat, format)
		assert.Equal(t, 7, len(a))
		assert.Equal(t, dummySessionID, a[0])
		assert.Equal(t, dummySessionIndex, a[1])
		assert.Equal(t, dummyLogType, a[2])
		assert.Equal(t, dummyLogLevel, a[3])
		assert.Equal(t, dummyCategory, a[4])
		assert.Equal(t, dummySubcategory, a[5])
		assert.Equal(t, dummyDescription, a[6])
		return rand.Int(), errors.New("some error")
	}

	// SUT + act
	customizationDefault.Log(
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyDescription,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, sessionGetIDExpected, sessionGetIDCalled, "Unexpected number of calls to method sessionGetID")
	assert.Equal(t, sessionGetIndexExpected, sessionGetIndexCalled, "Unexpected number of calls to method sessionGetIndex")
}

func TestDefaultCustomization_PreAction(t *testing.T) {
	// arrange
	var dummySession Session

	// mock
	createMock(t)

	// SUT + act
	var err = customizationDefault.PreAction(dummySession)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_PostAction(t *testing.T) {
	// arrange
	var dummySession Session

	// mock
	createMock(t)

	// SUT + act
	var err = customizationDefault.PostAction(dummySession)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_ClientCert(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var result = customizationDefault.ClientCert()

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_DefaultTimeout(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var result = customizationDefault.DefaultTimeout()

	// assert
	assert.Equal(t, 3*time.Minute, result)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_SkipServerCertVerification(t *testing.T) {
	// mock
	createMock(t)

	// SUT + act
	var result = customizationDefault.SkipServerCertVerification()

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_RoundTripper(t *testing.T) {
	// arrange
	var dummyTransport = &dummyTransport{t: t}

	// mock
	createMock(t)

	// SUT + act
	var result = customizationDefault.RoundTripper(dummyTransport)

	// assert
	assert.Equal(t, dummyTransport, result)

	// verify
	verifyAll(t)
}

func TestDefaultCustomization_WrapRequest(t *testing.T) {
	// arrange
	var dummySession Session
	var dummyRequest = &http.Request{Host: "some host"}

	// mock
	createMock(t)

	// SUT + act
	var result = customizationDefault.WrapRequest(dummySession, dummyRequest)

	// assert
	assert.Equal(t, dummyRequest, result)

	// verify
	verifyAll(t)
}
