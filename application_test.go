package jobrunner

import (
	"crypto/tls"
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func functionPointerEquals(expectFunc interface{}, actualFunc interface{}) bool {
	var expectValue = fmt.Sprintf("%v", reflect.ValueOf(expectFunc))
	var actualValue = fmt.Sprintf("%v", reflect.ValueOf(actualFunc))
	return expectValue == actualValue
}

func assertFunctionEquals(t *testing.T, expectFunc interface{}, actualFunc interface{}) {
	assert.True(t, functionPointerEquals(expectFunc, actualFunc))
}

func TestNewApplication_NilCustomization(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyVersion = "some version"
	var dummyInstances = rand.IntN(100)
	var dummySchedule Schedule
	var dummyOverlap = rand.IntN(100) > 50
	var dummyCustomization Customization
	var dummySessionID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummyCustomization).Returns(true).Once()
	m.Mock(uuid.New).Expects().Returns(dummySessionID).Once()

	// SUT
	var result = NewApplication(
		dummyName,
		dummyVersion,
		dummyInstances,
		dummySchedule,
		dummyOverlap,
		dummyCustomization,
	)

	// act
	var value, ok = result.(*application)

	// assert
	assert.True(t, ok)
	assert.NotNil(t, value)
	assert.Equal(t, dummyName, value.name)
	assert.Equal(t, dummyVersion, value.version)
	assert.Equal(t, dummyInstances, value.instances)
	assert.Equal(t, dummySchedule, value.schedule)
	assert.Equal(t, dummyOverlap, value.overlap)
	assert.NotNil(t, value.session)
	assert.Equal(t, dummySessionID, value.session.id)
	assert.Equal(t, 0, value.session.index)
	assert.Empty(t, value.session.attachment)
	assert.Equal(t, customizationDefault, value.session.customization)
	assert.Equal(t, customizationDefault, value.customization)
}

func TestNewApplication_HasCustomization(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyVersion = "some version"
	var dummyInstances = rand.IntN(100)
	var dummySchedule Schedule
	var dummyOverlap = rand.IntN(100) > 50
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySessionID = uuid.New()

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummyCustomization).Returns(false).Once()
	m.Mock(uuid.New).Expects().Returns(dummySessionID).Once()

	// SUT
	var result = NewApplication(
		dummyName,
		dummyVersion,
		dummyInstances,
		dummySchedule,
		dummyOverlap,
		dummyCustomization,
	)

	// act
	var value, ok = result.(*application)

	// assert
	assert.True(t, ok)
	assert.NotNil(t, value)
	assert.Equal(t, dummyName, value.name)
	assert.Equal(t, dummyVersion, value.version)
	assert.Equal(t, dummyInstances, value.instances)
	assert.Equal(t, dummySchedule, value.schedule)
	assert.Equal(t, dummyOverlap, value.overlap)
	assert.NotNil(t, value.session)
	assert.Equal(t, dummySessionID, value.session.id)
	assert.Equal(t, 0, value.session.index)
	assert.Empty(t, value.session.attachment)
	assert.Equal(t, dummyCustomization, value.session.customization)
	assert.Equal(t, dummyCustomization, value.customization)
}

func TestApplication_Start(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(startApplication).Expects(dummyApplication).Returns().Once()

	// SUT + act
	dummyApplication.Start()
}

func TestApplication_IsRunning(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: rand.IntN(100) > 50,
	}

	// SUT + act
	var result = dummyApplication.IsRunning()

	// assert
	assert.Equal(t, dummyApplication.started, result)
}

func TestApplication_LastErrors(t *testing.T) {
	// arrange
	var dummyLastErrors = []error{
		errors.New("some error 1"),
		errors.New("some error 2"),
		errors.New("some error 3"),
	}
	var dummyApplication = &application{
		name:       "some name",
		lastErrors: dummyLastErrors,
	}

	// SUT + act
	var result = dummyApplication.LastErrors()

	// assert
	assert.Equal(t, dummyApplication.lastErrors, result)
}

func TestApplication_Stop_NotStarted(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: false,
	}

	// SUT + act
	dummyApplication.Stop()
}

func TestApplication_Stop_HasStarted(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		started:  true,
	}

	// SUT + act
	go dummyApplication.Stop()

	// assert
	assert.True(t, <-dummyShutdown)
}

func TestStartApplication_AlreadyStarted(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
	}

	// SUT + act
	startApplication(dummyApplication)
}

func TestStartApplication_PreBootstrapingFailure(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(preBootstraping).Expects(dummyApplication).Returns(false).Once()

	// SUT + act
	startApplication(dummyApplication)
}

func TestStartApplication_PostBootstrapingFailure(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(preBootstraping).Expects(dummyApplication).Returns(true).Once()
	m.Mock(bootstrap).Expects(dummyApplication).Returns().Once()
	m.Mock(postBootstraping).Expects(dummyApplication).Returns(false).Once()

	// SUT + act
	startApplication(dummyApplication)
}

func TestStartApplication_HappyPath(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(preBootstraping).Expects(dummyApplication).Returns(true).Once()
	m.Mock(bootstrap).Expects(dummyApplication).Returns().Once()
	m.Mock(postBootstraping).Expects(dummyApplication).Returns(true).Once()
	m.Mock(beginApplication).Expects(dummyApplication).Returns().Once()
	m.Mock(endApplication).Expects(dummyApplication).Returns().Once()

	// SUT + act
	startApplication(dummyApplication)
}

func TestPreBootstraping_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.PreBootstrap. Error: %+v"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreBootstrap).Expects(dummyCustomization).Returns(dummyError).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "preBootstraping",
		dummyMessageFormat, dummyError).Returns().Once()

	// SUT + act
	var result = preBootstraping(
		dummyApplication,
	)

	// assert
	assert.False(t, result)
	assert.Len(t, dummyApplication.lastErrors, 1)
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])
}

func TestPreBootstraping_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyMessageFormat = "customization.PreBootstrap executed successfully"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PreBootstrap).Expects(dummyCustomization).Returns(nil).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application",
		"preBootstraping", dummyMessageFormat).Returns().Once()

	// SUT + act
	var result = preBootstraping(
		dummyApplication,
	)

	// assert
	assert.True(t, result)
	assert.Empty(t, dummyApplication.lastErrors)
}

func TestBootstrap_HappyPath(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyWebcallTimeout = time.Duration(rand.IntN(100))
	var dummySkipCertVerification = rand.IntN(100) > 50
	var dummyClientCertificate = &tls.Certificate{Certificate: [][]byte{{0}}}
	var dummyMessageFormat = "Application bootstrapped successfully"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(initializeHTTPClients).Expects(dummyWebcallTimeout, dummySkipCertVerification,
		dummyClientCertificate, gomocker.Anything()).Returns().Once()
	m.Mock((*customization).DefaultTimeout).Expects(dummyCustomization).Returns(dummyWebcallTimeout).Once()
	m.Mock((*customization).SkipServerCertVerification).Expects(dummyCustomization).Returns(dummySkipCertVerification).Once()
	m.Mock((*customization).ClientCert).Expects(dummyCustomization).Returns(dummyClientCertificate).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "bootstrap", dummyMessageFormat).Returns().Once()

	// SUT + act
	bootstrap(
		dummyApplication,
	)
}

func TestPostBootstraping_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.PostBootstrap. Error: %+v"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PostBootstrap).Expects(dummyCustomization).Returns(dummyError).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "postBootstraping",
		dummyMessageFormat, dummyError).Returns().Once()

	// SUT + act
	var result = postBootstraping(
		dummyApplication,
	)

	// assert
	assert.False(t, result)
	assert.Len(t, dummyApplication.lastErrors, 1)
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])
}

func TestPostBootstraping_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyMessageFormat = "customization.PostBootstrap executed successfully"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).PostBootstrap).Expects(dummyCustomization).Returns(nil).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application",
		"postBootstraping", dummyMessageFormat).Returns().Once()

	// SUT + act
	var result = postBootstraping(
		dummyApplication,
	)

	// assert
	assert.True(t, result)
	assert.Empty(t, dummyApplication.lastErrors)
}

func TestWaitForNextRun_NilNextSchedule(t *testing.T) {
	// arrange
	type schedule struct {
		Schedule
	}
	var dummySchedule = &schedule{}
	var dummySession = &session{id: uuid.New()}
	var dummyApplication = &application{
		name:     "some name",
		schedule: dummySchedule,
		session:  dummySession,
		started:  true,
	}
	var dummyTimeNext *time.Time
	var dummyMessageFormat = "No next schedule available, terminating execution"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*schedule).NextSchedule).Expects(dummySchedule).Returns(dummyTimeNext).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application",
		"waitForNextRun", dummyMessageFormat).Returns().Once()

	// SUT + act
	waitForNextRun(
		dummyApplication,
	)

	// assert
	assert.False(t, dummyApplication.started)
}

func TestWaitForNextRun_ValidNextSchedule(t *testing.T) {
	// arrange
	type schedule struct {
		Schedule
	}
	var dummySchedule = &schedule{}
	var dummySession = &session{id: uuid.New()}
	var dummyApplication = &application{
		name:     "some name",
		schedule: dummySchedule,
		session:  dummySession,
	}
	var dummyTimeNow = time.Now()
	var dummyDuration = time.Duration(rand.IntN(1000)) + 10*time.Second
	var dummyTimeNext = dummyTimeNow.Add(dummyDuration)
	var dummyMessageFormat = "Next run at [%v]: waiting for [%v]"
	var dummyControlChannel = make(chan time.Time)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*schedule).NextSchedule).Expects(dummySchedule).Returns(&dummyTimeNext).Once()
	m.Mock(time.Now).Expects().Returns(dummyTimeNow).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "waitForNextRun",
		dummyMessageFormat, dummyTimeNext, dummyDuration).Returns().Once()
	m.Mock(time.After).Expects(dummyDuration).Returns(dummyControlChannel).Once()

	// SUT + act
	go waitForNextRun(
		dummyApplication,
	)

	// push
	dummyControlChannel <- dummyTimeNext
}

func TestRunInstances_ZeroInstance(t *testing.T) {
	// arrange
	var dummyApplication = &application{}

	// SUT + act
	runInstances(
		dummyApplication,
	)
}

func TestRunInstances_SingleInstance(t *testing.T) {
	// arrange
	var dummyReruns = rand.Int32N(65535)
	var dummyApplication = &application{
		instances: 1,
		reruns:    []int32{dummyReruns},
	}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(handleSession).Expects(dummyApplication, 0, int(dummyReruns)+1).Returns(dummyError).Once()

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// assert
	assert.Len(t, dummyApplication.lastErrors, 1)
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])
}

func TestRunInstances_MultipleInstances(t *testing.T) {
	// arrange
	var dummyErrors = []error{
		errors.New("some error 1"),
		errors.New("some error 2"),
		errors.New("some error 3"),
	}
	var dummyApplication = &application{
		instances: 3,
		reruns:    make([]int32, 3),
	}
	var calls = map[int]bool{}
	var lock = sync.RWMutex{}

	// stub
	var callChecker = func(value interface{}) bool {
		lock.Lock()
		defer lock.Unlock()
		var called, found = calls[value.(int)]
		if found || called {
			return false
		}
		calls[value.(int)] = true
		return true
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(handleSession).Expects(dummyApplication, gomocker.Matches(callChecker), 1).Returns(dummyErrors[0]).Once()
	m.Mock(handleSession).Expects(dummyApplication, gomocker.Matches(callChecker), 1).Returns(dummyErrors[1]).Once()
	m.Mock(handleSession).Expects(dummyApplication, gomocker.Matches(callChecker), 1).Returns(dummyErrors[2]).Once()

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// assert
	assert.Len(t, dummyApplication.lastErrors, 3)
	assert.ElementsMatch(t, dummyErrors, dummyApplication.lastErrors)
}

func TestRunInstances_Overlap(t *testing.T) {
	// arrange
	var dummyReruns = rand.Int32N(65535)
	var dummyApplication = &application{
		instances: 1,
		reruns:    []int32{dummyReruns},
		overlap:   true,
	}
	var dummyError = errors.New("some error")

	// stub
	dummyApplication.waits.Add(1)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(handleSession).Expects(dummyApplication, 0, int(dummyReruns)+1).Returns(dummyError).Once()

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// assert
	assert.Len(t, dummyApplication.lastErrors, 1)
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])
}

func TestScheduleExecution_WithOverlap(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
		overlap: true,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(waitForNextRun).Expects(dummyApplication).Returns().Once()
	m.Mock(runInstances).Expects(dummyApplication).Returns().SideEffect(
		func(index int, params ...interface{}) {
			dummyApplication.waits.Done()
		}).Once()
	m.Mock(waitForNextRun).Expects(dummyApplication).Returns().SideEffect(
		func(index int, params ...interface{}) {
			dummyApplication.started = false
		}).Once()

	// SUT + act
	scheduleExecution(
		dummyApplication,
	)
}

func TestScheduleExecution_NoOverlap(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
		overlap: false,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(waitForNextRun).Expects(dummyApplication).Returns().Once()
	m.Mock(runInstances).Expects(dummyApplication).Returns().Once()
	m.Mock(waitForNextRun).Expects(dummyApplication).Returns().SideEffect(func(index int, params ...interface{}) {
		dummyApplication.started = false
	}).Once()

	// SUT + act
	scheduleExecution(
		dummyApplication,
	)
}

func TestRunApplication_NoSchedule(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	type schedule struct {
		Schedule
	}
	var dummySchedule = &schedule{}
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		schedule: dummySchedule,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummySchedule).Returns(true).Once()
	m.Mock(runInstances).Expects(dummyApplication).Returns().Once()

	// SUT + act
	go runApplication(
		dummyApplication,
	)

	// assert
	assert.True(t, <-dummyShutdown)
}

func TestRunApplication_WithSchedule(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	type schedule struct {
		Schedule
	}
	var dummySchedule = &schedule{}
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		schedule: dummySchedule,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(dummySchedule).Returns(false).Once()
	m.Mock(scheduleExecution).Expects(dummyApplication).Returns().Once()

	// SUT + act
	go runApplication(
		dummyApplication,
	)

	// assert
	assert.True(t, <-dummyShutdown)
}

func TestBeginApplication_HappyPath(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyVersion = "some version"
	var dummySession = &session{id: uuid.New()}
	var dummyShutdown = make(chan bool)
	var dummyApplication = &application{
		name:     dummyName,
		version:  dummyVersion,
		session:  dummySession,
		shutdown: dummyShutdown,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(runApplication).Expects(dummyApplication).Returns().SideEffect(func(index int, params ...interface{}) {
		dummyShutdown <- true
	}).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "beginApplication",
		"Trying to start runner [%v] (v-%v)", dummyName, dummyVersion).Returns().Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "beginApplication", "Runner terminated").Returns().Once()

	// SUT + act
	beginApplication(
		dummyApplication,
	)

	// assert
	assert.False(t, dummyApplication.started)
}

func TestEndApplication_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.AppClosing. Error: %+v"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).AppClosing).Expects(dummyCustomization).Returns(dummyError).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application", "endApplication",
		dummyMessageFormat, dummyError).Returns().Once()

	// SUT + act
	endApplication(
		dummyApplication,
	)

	// assert
	assert.Len(t, dummyApplication.lastErrors, 1)
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])
}

func TestEndApplication_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var dummyMessageFormat = "customization.AppClosing executed successfully"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).AppClosing).Expects(dummyCustomization).Returns(nil).Once()
	m.Mock(logAppRoot).Expects(dummySession, "application",
		"endApplication", dummyMessageFormat).Returns().Once()

	// SUT + act
	endApplication(
		dummyApplication,
	)

	// assert
	assert.Empty(t, dummyApplication.lastErrors)
}
