package jobrunner

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestDefaultCustomization_PreBootstrap(t *testing.T) {
	// SUT + act
	var err = customizationDefault.PreBootstrap()

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_PostBootstrap(t *testing.T) {
	// SUT + act
	var err = customizationDefault.PostBootstrap()

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_AppClosing(t *testing.T) {
	// SUT + act
	var err = customizationDefault.AppClosing()

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_Log_HappyPath(t *testing.T) {
	// arrange
	var dummySession = &session{}
	var dummyTimeNow = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	var dummyTimeString = "some time string"
	var dummyLogType = LogType(rand.IntN(100))
	var dummyLogLevel = LogLevel(rand.IntN(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyDescription = "some description"
	var dummySessionID = uuid.New()
	var dummySessionIndex = rand.Int()
	var dummyFormat = "[%v] <%v|%v> (%v|%v) [%v|%v] %v\n"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Now).Expects().Returns(dummyTimeNow).Once()
	m.Mock(formatDateTime).Expects(dummyTimeNow).Returns(dummyTimeString).Once()
	m.Mock((*session).GetID).Expects(dummySession).Returns(dummySessionID).Once()
	m.Mock((*session).GetIndex).Expects(dummySession).Returns(dummySessionIndex).Once()
	m.Mock(fmt.Printf).Expects(
		dummyFormat,
		dummyTimeString,
		dummySessionID,
		dummySessionIndex,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyDescription,
	).Returns(rand.Int(), errors.New("some error")).Once()

	// SUT + act
	customizationDefault.Log(
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyDescription,
	)
}

func TestDefaultCustomization_PreAction(t *testing.T) {
	// arrange
	var dummySession Session

	// SUT + act
	var err = customizationDefault.PreAction(dummySession)

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_PostAction(t *testing.T) {
	// arrange
	var dummySession Session

	// SUT + act
	var err = customizationDefault.PostAction(dummySession)

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_ActionFunc(t *testing.T) {
	// arrange
	var dummySession Session

	// SUT + act
	var err = customizationDefault.ActionFunc(dummySession)

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_RecoverPanic_NilRecoverResult(t *testing.T) {
	// arrange
	var dummySession Session
	var dummyRecoverResult *int

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummyRecoverResult).Returns(true).Once()

	// SUT + act
	var err = customizationDefault.RecoverPanic(
		dummySession,
		dummyRecoverResult,
	)

	// assert
	assert.NoError(t, err)
}

func TestDefaultCustomization_RecoverPanic_RecoverResultAsError(t *testing.T) {
	// arrange
	var dummySession Session
	var dummyRecoverResult = errors.New("some recover result")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummyRecoverResult).Returns(false).Once()

	// SUT + act
	var err = customizationDefault.RecoverPanic(
		dummySession,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyRecoverResult, err)
}

func TestDefaultCustomization_RecoverPanic_RecoverResultAsNonError(t *testing.T) {
	// arrange
	var dummySession Session
	var dummyRecoverResult = "some recover result"
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummyRecoverResult).Returns(false).Once()
	m.Mock(fmt.Errorf).Expects("%v", dummyRecoverResult).Returns(dummyError).Once()

	// SUT + act
	var err = customizationDefault.RecoverPanic(
		dummySession,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyError, err)
}

func TestDefaultCustomization_ClientCert(t *testing.T) {
	// SUT + act
	var result = customizationDefault.ClientCert()

	// assert
	assert.Nil(t, result)
}

func TestDefaultCustomization_DefaultTimeout(t *testing.T) {
	// SUT + act
	var result = customizationDefault.DefaultTimeout()

	// assert
	assert.Equal(t, 3*time.Minute, result)
}

func TestDefaultCustomization_SkipServerCertVerification(t *testing.T) {
	// SUT + act
	var result = customizationDefault.SkipServerCertVerification()

	// assert
	assert.False(t, result)
}

func TestDefaultCustomization_RoundTripper(t *testing.T) {
	// arrange
	var dummyTransport = &http.Transport{}

	// SUT + act
	var result = customizationDefault.RoundTripper(dummyTransport)

	// assert
	assert.Equal(t, dummyTransport, result)
}

func TestDefaultCustomization_WrapRequest(t *testing.T) {
	// arrange
	var dummySession Session
	var dummyRequest = &http.Request{Host: "some host"}

	// SUT + act
	var result = customizationDefault.WrapRequest(dummySession, dummyRequest)

	// assert
	assert.Equal(t, dummyRequest, result)
}
