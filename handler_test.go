package jobrunner

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestInitiateSession(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		customization: dummyCustomization,
	}
	var dummyIndex = rand.IntN(65536)
	var dummyReruns = rand.IntN(65536)
	var dummySessionID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(uuid.New).Expects().Returns(dummySessionID).Once()

	// SUT + act
	var session = initiateSession(
		dummyApplication,
		dummyIndex,
		dummyReruns,
	)

	// assert
	assert.NotNil(t, session)
	assert.Equal(t, dummySessionID, session.id)
	assert.Equal(t, dummyIndex, session.index)
	assert.Equal(t, dummyReruns, session.reruns)
	assert.Empty(t, session.attachment)
	assert.Equal(t, dummyCustomization, session.customization)
}

func TestFinalizeSession_NoErrorResult(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyErrorResult error
	var dummyRecoverResult = errors.New("some recover result")
	var dummyRecoverError = errors.New("some recover error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).RecoverPanic).Expects(dummyCustomization, dummySession, dummyRecoverResult).Returns(dummyRecoverError).Once()

	// SUT + act
	var err = finalizeSession(
		dummySession,
		dummyErrorResult,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyRecoverError, err)

}

func TestFinalizeSession_WithErrorResult(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyErrorResult = errors.New("some error result")
	var dummyRecoverResult = errors.New("some recover result")
	var dummyRecoverError = errors.New("some recover error")
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).RecoverPanic).Expects(dummyCustomization, dummySession, dummyRecoverResult).Returns(dummyRecoverError).Once()
	m.Mock(fmt.Errorf).Expects("Original Error: %w\nRecover Error: %v", dummyErrorResult, dummyRecoverError).Returns(dummyError).Once()

	// SUT + act
	var err = finalizeSession(
		dummySession,
		dummyErrorResult,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyError, err)

}

func TestProcessSession_PreActionError(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreAction).Expects(dummyCustomization, dummySession).Returns(dummyError).Once()

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomization,
	)

	// assert
	assert.Equal(t, dummyError, err)

}

func TestProcessSession_ActionFuncError(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreAction).Expects(dummyCustomization, dummySession).Returns(nil).Once()
	m.Mock((*customization).ActionFunc).Expects(dummyCustomization, dummySession).Returns(dummyError).Once()

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomization,
	)

	// assert
	assert.Equal(t, dummyError, err)

}

func TestProcessSession_PostActionError(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreAction).Expects(dummyCustomization, dummySession).Returns(nil).Once()
	m.Mock((*customization).ActionFunc).Expects(dummyCustomization, dummySession).Returns(nil).Once()
	m.Mock((*customization).PostAction).Expects(dummyCustomization, dummySession).Returns(dummyError).Once()

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomization,
	)

	// assert
	assert.Equal(t, dummyError, err)

}

func TestProcessSession_Success(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreAction).Expects(dummyCustomization, dummySession).Returns(nil).Once()
	m.Mock((*customization).ActionFunc).Expects(dummyCustomization, dummySession).Returns(nil).Once()
	m.Mock((*customization).PostAction).Expects(dummyCustomization, dummySession).Returns(nil).Once()

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomization,
	)

	// assert
	assert.NoError(t, err)

}

func TestHandleSession_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		name:          dummyName,
		customization: dummyCustomization,
	}
	var dummyIndex = rand.Int()
	var dummyReruns = rand.Int()
	var dummySession = &session{id: uuid.New()}
	var dummyTimeNow = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	var dummyDuration = time.Duration(rand.IntN(100))
	var dummyProcessError = errors.New("some process error")
	var dummyFinalError = errors.New("some final error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(initiateSession).Expects(dummyApplication, dummyIndex, dummyReruns).Returns(dummySession).Once()
	m.Mock(logProcessEnter).Expects(dummySession, dummyName, "", "").Returns().Once()
	m.Mock(logProcessRequest).Expects(dummySession, dummyName, "InstanceIndex", "%v", dummyIndex).Returns().Once()
	m.Mock(time.Now).Expects().Returns(dummyTimeNow).Once()
	m.Mock(processSession).Expects(dummySession, dummyCustomization).Returns(dummyProcessError).Once()
	m.Mock(finalizeSession).Expects(dummySession, dummyProcessError, recover()).Returns(dummyFinalError).Once()
	m.Mock(time.Since).Expects(dummyTimeNow).Returns(dummyDuration).Once()
	m.Mock(logProcessResponse).Expects(dummySession, dummyName, "", "%v", dummyFinalError).Returns().Once()
	m.Mock(logProcessExit).Expects(dummySession, dummyName, "Duration", "%s", dummyDuration).Returns().Once()

	// SUT + act
	var err = handleSession(
		dummyApplication,
		dummyIndex,
		dummyReruns,
	)

	// assert
	assert.Equal(t, dummyFinalError, err)
}
