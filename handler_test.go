package jobrunner

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInitiateSession(t *testing.T) {
	// arrange
	var dummyCustomization = &dummyCustomization{t: t}
	var dummyApplication = &application{
		customization: dummyCustomization,
	}
	var dummyIndex = rand.Intn(65536)
	var dummyReruns = rand.Intn(65536)
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// expect
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return dummySessionID
	}

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

	// verify
	verifyAll(t)
}

type dummyCustomizationRecoverPanic struct {
	dummyCustomization
	recoverPanic func(session Session, recoverResult interface{}) error
}

func (dummyCustomizationRecoverPanic *dummyCustomizationRecoverPanic) RecoverPanic(session Session, recoverResult interface{}) error {
	if dummyCustomizationRecoverPanic.recoverPanic != nil {
		return dummyCustomizationRecoverPanic.recoverPanic(session, recoverResult)
	}
	assert.Fail(dummyCustomizationRecoverPanic.t, "Unexpected call to RecoverPanic")
	return nil
}

func TestFinalizeSession_NoErrorResult(t *testing.T) {
	// arrange
	var dummyCustomizationRecoverPanic = &dummyCustomizationRecoverPanic{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummySession = &session{
		customization: dummyCustomizationRecoverPanic,
	}
	var dummyErrorResult error
	var dummyRecoverResult = errors.New("some recover result")
	var dummyRecoverError = errors.New("some recover error")
	var customizationRecoverPanicExpected int
	var customizationRecoverPanicCalled int

	// mock
	createMock(t)

	// expect
	customizationRecoverPanicExpected = 1
	dummyCustomizationRecoverPanic.recoverPanic = func(session Session, recoverResult interface{}) error {
		customizationRecoverPanicCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyRecoverResult, recoverResult)
		return dummyRecoverError
	}

	// SUT + act
	var err = finalizeSession(
		dummySession,
		dummyErrorResult,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyRecoverError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationRecoverPanicExpected, customizationRecoverPanicCalled, "Unexpected number of calls to method customization.RecoverPanic")
}

func TestFinalizeSession_WithErrorResult(t *testing.T) {
	// arrange
	var dummyCustomizationRecoverPanic = &dummyCustomizationRecoverPanic{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummySession = &session{
		customization: dummyCustomizationRecoverPanic,
	}
	var dummyErrorResult = errors.New("some error result")
	var dummyRecoverResult = errors.New("some recover result")
	var dummyRecoverError = errors.New("some recover error")
	var dummyError = errors.New("some error")
	var customizationRecoverPanicExpected int
	var customizationRecoverPanicCalled int

	// mock
	createMock(t)

	// expect
	customizationRecoverPanicExpected = 1
	dummyCustomizationRecoverPanic.recoverPanic = func(session Session, recoverResult interface{}) error {
		customizationRecoverPanicCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyRecoverResult, recoverResult)
		return dummyRecoverError
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, "Original Error: %w\nRecover Error: %v", format)
		assert.Equal(t, 2, len(a))
		assert.Equal(t, dummyErrorResult, a[0])
		assert.Equal(t, dummyRecoverError, a[1])
		return dummyError
	}

	// SUT + act
	var err = finalizeSession(
		dummySession,
		dummyErrorResult,
		dummyRecoverResult,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationRecoverPanicExpected, customizationRecoverPanicCalled, "Unexpected number of calls to method customization.RecoverPanic")
}

type dummyCustomizationProcessSession struct {
	dummyCustomization
	preAction  func(Session) error
	actionFunc func(Session) error
	postAction func(Session) error
}

func (dummyCustomizationProcessSession *dummyCustomizationProcessSession) PreAction(session Session) error {
	if dummyCustomizationProcessSession.preAction != nil {
		return dummyCustomizationProcessSession.preAction(session)
	}
	assert.Fail(dummyCustomizationProcessSession.t, "Unexpected call to PreAction")
	return nil
}

func (dummyCustomizationProcessSession *dummyCustomizationProcessSession) ActionFunc(session Session) error {
	if dummyCustomizationProcessSession.actionFunc != nil {
		return dummyCustomizationProcessSession.actionFunc(session)
	}
	assert.Fail(dummyCustomizationProcessSession.t, "Unexpected call to ActionFunc")
	return nil
}

func (dummyCustomizationProcessSession *dummyCustomizationProcessSession) PostAction(session Session) error {
	if dummyCustomizationProcessSession.postAction != nil {
		return dummyCustomizationProcessSession.postAction(session)
	}
	assert.Fail(dummyCustomizationProcessSession.t, "Unexpected call to PostAction")
	return nil
}

func TestProcessSession_PreActionError(t *testing.T) {
	// arrange
	var dummySession = &dummySession{t: t}
	var dummyCustomizationProcessSession = &dummyCustomizationProcessSession{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyError = errors.New("some error")
	var customizationPreActionExpected int
	var customizationPreActionCalled int
	var customizationActionFuncExpected int
	var customizationActionFuncCalled int
	var customizationPostActionExpected int
	var customizationPostActionCalled int

	// mock
	createMock(t)

	// expect
	customizationPreActionExpected = 1
	dummyCustomizationProcessSession.preAction = func(session Session) error {
		customizationPreActionCalled++
		assert.Equal(t, dummySession, session)
		return dummyError
	}

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomizationProcessSession,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreActionExpected, customizationPreActionCalled, "Unexpected number of calls to method customization.PreAction")
	assert.Equal(t, customizationActionFuncExpected, customizationActionFuncCalled, "Unexpected number of calls to method customization.ActionFunc")
	assert.Equal(t, customizationPostActionExpected, customizationPostActionCalled, "Unexpected number of calls to method customization.PostAction")
}

func TestProcessSession_ActionFuncError(t *testing.T) {
	// arrange
	var dummySession = &dummySession{t: t}
	var dummyCustomizationProcessSession = &dummyCustomizationProcessSession{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyError = errors.New("some error")
	var customizationPreActionExpected int
	var customizationPreActionCalled int
	var customizationActionFuncExpected int
	var customizationActionFuncCalled int
	var customizationPostActionExpected int
	var customizationPostActionCalled int

	// mock
	createMock(t)

	// expect
	customizationPreActionExpected = 1
	dummyCustomizationProcessSession.preAction = func(session Session) error {
		customizationPreActionCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}
	customizationActionFuncExpected = 1
	dummyCustomizationProcessSession.actionFunc = func(session Session) error {
		customizationActionFuncCalled++
		assert.Equal(t, dummySession, session)
		return dummyError
	}

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomizationProcessSession,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreActionExpected, customizationPreActionCalled, "Unexpected number of calls to method customization.PreAction")
	assert.Equal(t, customizationActionFuncExpected, customizationActionFuncCalled, "Unexpected number of calls to method customization.ActionFunc")
	assert.Equal(t, customizationPostActionExpected, customizationPostActionCalled, "Unexpected number of calls to method customization.PostAction")
}

func TestProcessSession_PostActionError(t *testing.T) {
	// arrange
	var dummySession = &dummySession{t: t}
	var dummyCustomizationProcessSession = &dummyCustomizationProcessSession{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyError = errors.New("some error")
	var customizationPreActionExpected int
	var customizationPreActionCalled int
	var customizationActionFuncExpected int
	var customizationActionFuncCalled int
	var customizationPostActionExpected int
	var customizationPostActionCalled int

	// mock
	createMock(t)

	// expect
	customizationPreActionExpected = 1
	dummyCustomizationProcessSession.preAction = func(session Session) error {
		customizationPreActionCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}
	customizationActionFuncExpected = 1
	dummyCustomizationProcessSession.actionFunc = func(session Session) error {
		customizationActionFuncCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}
	customizationPostActionExpected = 1
	dummyCustomizationProcessSession.postAction = func(session Session) error {
		customizationPostActionCalled++
		assert.Equal(t, dummySession, session)
		return dummyError
	}

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomizationProcessSession,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreActionExpected, customizationPreActionCalled, "Unexpected number of calls to method customization.PreAction")
	assert.Equal(t, customizationActionFuncExpected, customizationActionFuncCalled, "Unexpected number of calls to method customization.ActionFunc")
	assert.Equal(t, customizationPostActionExpected, customizationPostActionCalled, "Unexpected number of calls to method customization.PostAction")
}

func TestProcessSession_Success(t *testing.T) {
	// arrange
	var dummySession = &dummySession{t: t}
	var dummyCustomizationProcessSession = &dummyCustomizationProcessSession{
		dummyCustomization: dummyCustomization{t: t},
	}
	var customizationPreActionExpected int
	var customizationPreActionCalled int
	var customizationActionFuncExpected int
	var customizationActionFuncCalled int
	var customizationPostActionExpected int
	var customizationPostActionCalled int

	// mock
	createMock(t)

	// expect
	customizationPreActionExpected = 1
	dummyCustomizationProcessSession.preAction = func(session Session) error {
		customizationPreActionCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}
	customizationActionFuncExpected = 1
	dummyCustomizationProcessSession.actionFunc = func(session Session) error {
		customizationActionFuncCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}
	customizationPostActionExpected = 1
	dummyCustomizationProcessSession.postAction = func(session Session) error {
		customizationPostActionCalled++
		assert.Equal(t, dummySession, session)
		return nil
	}

	// SUT + act
	var err = processSession(
		dummySession,
		dummyCustomizationProcessSession,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreActionExpected, customizationPreActionCalled, "Unexpected number of calls to method customization.PreAction")
	assert.Equal(t, customizationActionFuncExpected, customizationActionFuncCalled, "Unexpected number of calls to method customization.ActionFunc")
	assert.Equal(t, customizationPostActionExpected, customizationPostActionCalled, "Unexpected number of calls to method customization.PostAction")
}

func TestHandleSession_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyCustomization = &dummyCustomization{t: t}
	var dummyApplication = &application{
		name:          dummyName,
		customization: dummyCustomization,
	}
	var dummyIndex = rand.Int()
	var dummyReruns = rand.Int()
	var dummySession = &session{id: uuid.New()}
	var dummyTimeNow = time.Now()
	var dummyDuration = time.Duration(rand.Intn(100))
	var dummyProcessError = errors.New("some process error")
	var dummyFinalError = errors.New("some final error")

	// mock
	createMock(t)

	// expect
	initiateSessionFuncExpected = 1
	initiateSessionFunc = func(app *application, index int, reruns int) *session {
		initiateSessionFuncCalled++
		assert.Equal(t, dummyApplication, app)
		assert.Equal(t, dummyIndex, index)
		assert.Equal(t, dummyReruns, reruns)
		return dummySession
	}
	logProcessEnterFuncExpected = 1
	logProcessEnterFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessEnterFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyName, category)
		assert.Equal(t, "", subcategory)
		assert.Equal(t, "", messageFormat)
		assert.Empty(t, parameters)
	}
	logProcessRequestFuncExpected = 1
	logProcessRequestFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessRequestFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyName, category)
		assert.Equal(t, "InstanceIndex", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyIndex, parameters[0])
	}
	getTimeNowUTCFuncExpected = 1
	getTimeNowUTCFunc = func() time.Time {
		getTimeNowUTCFuncCalled++
		return dummyTimeNow
	}
	processSessionFuncExpected = 1
	processSessionFunc = func(session Session, customization Customization) error {
		processSessionFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyCustomization, customization)
		return dummyProcessError
	}
	finalizeSessionFuncExpected = 1
	finalizeSessionFunc = func(session *session, resultError error, recoverResult interface{}) error {
		finalizeSessionFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyProcessError, resultError)
		assert.Equal(t, recover(), recoverResult)
		return dummyFinalError
	}
	timeSinceExpected = 1
	timeSince = func(tm time.Time) time.Duration {
		timeSinceCalled++
		assert.Equal(t, dummyTimeNow, tm)
		return dummyDuration
	}
	logProcessResponseFuncExpected = 1
	logProcessResponseFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessResponseFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyName, category)
		assert.Equal(t, "", subcategory)
		assert.Equal(t, "%v", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyFinalError, parameters[0])
	}
	logProcessExitFuncExpected = 1
	logProcessExitFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logProcessExitFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, dummyName, category)
		assert.Equal(t, "Duration", subcategory)
		assert.Equal(t, "%s", messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyDuration, parameters[0])
	}

	// SUT + act
	var err = handleSession(
		dummyApplication,
		dummyIndex,
		dummyReruns,
	)

	// assert
	assert.Equal(t, dummyFinalError, err)

	// verify
	verifyAll(t)
}
