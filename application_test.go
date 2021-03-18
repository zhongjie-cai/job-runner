package jobrunner

import (
	"crypto/tls"
	"errors"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication_NilCustomization(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyVersion = "some version"
	var dummyInstances = rand.Int()
	var dummySchedule Schedule
	var dummyOverlap = rand.Intn(100) > 50
	var dummyCustomization Customization
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummyCustomization, i)
		return true
	}
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return dummySessionID
	}

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

	// verify
	verifyAll(t)
}

func TestNewApplication_HasCustomization(t *testing.T) {
	// arrange
	var dummyName = "some name"
	var dummyVersion = "some version"
	var dummyInstances = rand.Int()
	var dummySchedule Schedule
	var dummyOverlap = rand.Intn(100) > 50
	var dummyCustomization = &dummyCustomization{t: t}
	var dummySessionID = uuid.New()

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummyCustomization, i)
		return false
	}
	uuidNewExpected = 1
	uuidNew = func() uuid.UUID {
		uuidNewCalled++
		return dummySessionID
	}

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

	// verify
	verifyAll(t)
}

func TestApplication_Start(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	createMock(t)

	// expect
	startApplicationFuncExpected = 1
	startApplicationFunc = func(app *application) {
		startApplicationFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}

	// SUT + act
	dummyApplication.Start()

	// verify
	verifyAll(t)
}

func TestApplication_IsRunning(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: rand.Intn(100) > 50,
	}

	// mock
	createMock(t)

	// SUT + act
	var result = dummyApplication.IsRunning()

	// assert
	assert.Equal(t, dummyApplication.started, result)

	// verify
	verifyAll(t)
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

	// mock
	createMock(t)

	// SUT + act
	var result = dummyApplication.LastErrors()

	// assert
	assert.Equal(t, dummyApplication.lastErrors, result)

	// verify
	verifyAll(t)
}

func TestApplication_Stop_NotStarted(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: false,
	}

	// mock
	createMock(t)

	// SUT + act
	dummyApplication.Stop()

	// verify
	verifyAll(t)
}

func TestApplication_Stop_HasStarted(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		started:  true,
	}

	// mock
	createMock(t)

	// SUT + act
	go dummyApplication.Stop()

	// assert
	assert.True(t, <-dummyShutdown)

	// verify
	verifyAll(t)
}

func TestStartApplication_AlreadyStarted(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
	}

	// mock
	createMock(t)

	// SUT + act
	startApplication(dummyApplication)

	// verify
	verifyAll(t)
}

func TestStartApplication_PreBootstrapingFailure(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	createMock(t)

	// expect
	preBootstrapingFuncExpected = 1
	preBootstrapingFunc = func(app *application) bool {
		preBootstrapingFuncCalled++
		assert.Equal(t, dummyApplication, app)
		return false
	}

	// SUT + act
	startApplication(dummyApplication)

	// verify
	verifyAll(t)
}

func TestStartApplication_PostBootstrapingFailure(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	createMock(t)

	// expect
	preBootstrapingFuncExpected = 1
	preBootstrapingFunc = func(app *application) bool {
		preBootstrapingFuncCalled++
		assert.Equal(t, dummyApplication, app)
		return true
	}
	bootstrapFuncExpected = 1
	bootstrapFunc = func(app *application) {
		bootstrapFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}
	postBootstrapingFuncExpected = 1
	postBootstrapingFunc = func(app *application) bool {
		postBootstrapingFuncCalled++
		assert.Equal(t, dummyApplication, app)
		return false
	}

	// SUT + act
	startApplication(dummyApplication)

	// verify
	verifyAll(t)
}

func TestStartApplication_HappyPath(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name: "some name",
	}

	// mock
	createMock(t)

	// expect
	preBootstrapingFuncExpected = 1
	preBootstrapingFunc = func(app *application) bool {
		preBootstrapingFuncCalled++
		assert.Equal(t, dummyApplication, app)
		return true
	}
	bootstrapFuncExpected = 1
	bootstrapFunc = func(app *application) {
		bootstrapFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}
	postBootstrapingFuncExpected = 1
	postBootstrapingFunc = func(app *application) bool {
		postBootstrapingFuncCalled++
		assert.Equal(t, dummyApplication, app)
		return true
	}
	beginApplicationFuncExpected = 1
	beginApplicationFunc = func(app *application) {
		beginApplicationFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}
	endApplicationFuncExpected = 1
	endApplicationFunc = func(app *application) {
		endApplicationFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}

	// SUT + act
	startApplication(dummyApplication)

	// verify
	verifyAll(t)
}

type dummyCustomizationPreBootstrapping struct {
	dummyCustomization
	preBootstrap func() error
}

func (customization *dummyCustomizationPreBootstrapping) PreBootstrap() error {
	if customization.preBootstrap != nil {
		return customization.preBootstrap()
	}
	assert.Fail(customization.t, "Unexpected call to PreBootstrap")
	return nil
}

func TestPreBootstraping_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationPreBootstrapping{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationPreBootstrapExpected int
	var customizationPreBootstrapCalled int
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.PreBootstrap. Error: %+v"

	// mock
	createMock(t)

	// expect
	customizationPreBootstrapExpected = 1
	dummyCustomization.preBootstrap = func() error {
		customizationPreBootstrapCalled++
		return dummyError
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "preBootstraping", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	var result = preBootstraping(
		dummyApplication,
	)

	// assert
	assert.False(t, result)
	assert.Equal(t, 1, len(dummyApplication.lastErrors))
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreBootstrapExpected, customizationPreBootstrapCalled, "Unexpected number of calls to method customization.PreBootstrap")
}

func TestPreBootstraping_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationPreBootstrapping{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationPreBootstrapExpected int
	var customizationPreBootstrapCalled int
	var dummyMessageFormat = "customization.PreBootstrap executed successfully"

	// mock
	createMock(t)

	// expect
	customizationPreBootstrapExpected = 1
	dummyCustomization.preBootstrap = func() error {
		customizationPreBootstrapCalled++
		return nil
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "preBootstraping", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	var result = preBootstraping(
		dummyApplication,
	)

	// assert
	assert.True(t, result)
	assert.Empty(t, dummyApplication.lastErrors)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPreBootstrapExpected, customizationPreBootstrapCalled, "Unexpected number of calls to method customization.PreBootstrap")
}

type dummyCustomizationBootstrap struct {
	dummyCustomization
	clientCert                 func() *tls.Certificate
	defaultTimeout             func() time.Duration
	skipServerCertVerification func() bool
	roundTripper               func(http.RoundTripper) http.RoundTripper
}

func (customization *dummyCustomizationBootstrap) ClientCert() *tls.Certificate {
	if customization.clientCert != nil {
		return customization.clientCert()
	}
	assert.Fail(customization.t, "Unexpected call to ClientCert")
	return nil
}

func (customization *dummyCustomizationBootstrap) DefaultTimeout() time.Duration {
	if customization.defaultTimeout != nil {
		return customization.defaultTimeout()
	}
	assert.Fail(customization.t, "Unexpected call to DefaultTimeout")
	return 0
}

func (customization *dummyCustomizationBootstrap) SkipServerCertVerification() bool {
	if customization.skipServerCertVerification != nil {
		return customization.skipServerCertVerification()
	}
	assert.Fail(customization.t, "Unexpected call to SkipServerCertVerification")
	return false
}

func (customization *dummyCustomizationBootstrap) RoundTripper(originalTransport http.RoundTripper) http.RoundTripper {
	if customization.roundTripper != nil {
		return customization.roundTripper(originalTransport)
	}
	assert.Fail(customization.t, "Unexpected call to RoundTripper")
	return nil
}

func TestBootstrap_HappyPath(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationBootstrap{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationDefaultTimeoutExpected int
	var customizationDefaultTimeoutCalled int
	var customizationSkipServerCertVerificationExpected int
	var customizationSkipServerCertVerificationCalled int
	var customizationClientCertExpected int
	var customizationClientCertCalled int
	var customizationRoundTripperExpected int
	var customizationRoundTripperCalled int
	var dummyWebcallTimeout = time.Duration(rand.Intn(100))
	var dummySkipCertVerification = rand.Intn(100) > 50
	var dummyClientCertificate = &tls.Certificate{Certificate: [][]byte{{0}}}
	var dummyOriginalTransport = &dummyTransport{t: t}
	var dummyMessageFormat = "Application bootstrapped successfully"

	// mock
	createMock(t)

	// expect
	initializeHTTPClientsFuncExpected = 1
	initializeHTTPClientsFunc = func(webcallTimeout time.Duration, skipServerCertVerification bool, clientCertificate *tls.Certificate, roundTripperWrapper func(originalTransport http.RoundTripper) http.RoundTripper) {
		initializeHTTPClientsFuncCalled++
		assert.Equal(t, dummyWebcallTimeout, webcallTimeout)
		assert.Equal(t, dummySkipCertVerification, skipServerCertVerification)
		assert.Equal(t, dummyClientCertificate, clientCertificate)
		roundTripperWrapper(dummyOriginalTransport)
	}
	customizationDefaultTimeoutExpected = 1
	dummyCustomization.defaultTimeout = func() time.Duration {
		customizationDefaultTimeoutCalled++
		return dummyWebcallTimeout
	}
	customizationSkipServerCertVerificationExpected = 1
	dummyCustomization.skipServerCertVerification = func() bool {
		customizationSkipServerCertVerificationCalled++
		return dummySkipCertVerification
	}
	customizationClientCertExpected = 1
	dummyCustomization.clientCert = func() *tls.Certificate {
		customizationClientCertCalled++
		return dummyClientCertificate
	}
	customizationRoundTripperExpected = 1
	dummyCustomization.roundTripper = func(originalTransport http.RoundTripper) http.RoundTripper {
		customizationRoundTripperCalled++
		assert.Equal(t, dummyOriginalTransport, originalTransport)
		return originalTransport
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "bootstrap", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	bootstrap(
		dummyApplication,
	)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationDefaultTimeoutExpected, customizationDefaultTimeoutCalled, "Unexpected number of calls to method customization.DefaultTimeout")
	assert.Equal(t, customizationSkipServerCertVerificationExpected, customizationSkipServerCertVerificationCalled, "Unexpected number of calls to method customization.SkipServerCertVerification")
	assert.Equal(t, customizationClientCertExpected, customizationClientCertCalled, "Unexpected number of calls to method customization.ClientCert")
	assert.Equal(t, customizationRoundTripperExpected, customizationRoundTripperCalled, "Unexpected number of calls to method customization.RoundTripper")
}

type dummyCustomizationPostBootstrapping struct {
	dummyCustomization
	postBootstrap func() error
}

func (customization *dummyCustomizationPostBootstrapping) PostBootstrap() error {
	if customization.postBootstrap != nil {
		return customization.postBootstrap()
	}
	assert.Fail(customization.t, "Unexpected call to PostBootstrap")
	return nil
}

func TestPostBootstraping_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationPostBootstrapping{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationPostBootstrapExpected int
	var customizationPostBootstrapCalled int
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.PostBootstrap. Error: %+v"

	// mock
	createMock(t)

	// expect
	customizationPostBootstrapExpected = 1
	dummyCustomization.postBootstrap = func() error {
		customizationPostBootstrapCalled++
		return dummyError
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "postBootstraping", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	var result = postBootstraping(
		dummyApplication,
	)

	// assert
	assert.False(t, result)
	assert.Equal(t, 1, len(dummyApplication.lastErrors))
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPostBootstrapExpected, customizationPostBootstrapCalled, "Unexpected number of calls to method customization.PostBootstrap")
}

func TestPostBootstraping_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationPostBootstrapping{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationPostBootstrapExpected int
	var customizationPostBootstrapCalled int
	var dummyMessageFormat = "customization.PostBootstrap executed successfully"

	// mock
	createMock(t)

	// expect
	customizationPostBootstrapExpected = 1
	dummyCustomization.postBootstrap = func() error {
		customizationPostBootstrapCalled++
		return nil
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "postBootstraping", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	var result = postBootstraping(
		dummyApplication,
	)

	// assert
	assert.True(t, result)
	assert.Empty(t, dummyApplication.lastErrors)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationPostBootstrapExpected, customizationPostBootstrapCalled, "Unexpected number of calls to method customization.PostBootstrap")
}

func TestWaitForNextRun_NilNextSchedule(t *testing.T) {
	// arrange
	var dummySchedule = &dummySchedule{t: t}
	var dummySession = &session{id: uuid.New()}
	var dummyApplication = &application{
		name:     "some name",
		schedule: dummySchedule,
		session:  dummySession,
		started:  true,
	}
	var dummyTimeNext *time.Time
	var dummyMessageFormat = "No next schedule available, terminating execution"
	var scheduleNextScheduleExpected int
	var scheduleNextScheduleCalled int

	// mock
	createMock(t)

	// expect
	scheduleNextScheduleExpected = 1
	dummySchedule.nextSchedule = func() *time.Time {
		scheduleNextScheduleCalled++
		return dummyTimeNext
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category, subcategory, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "waitForNextRun", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	waitForNextRun(
		dummyApplication,
	)

	// assert
	assert.False(t, dummyApplication.started)

	// verify
	verifyAll(t)
	assert.Equal(t, scheduleNextScheduleExpected, scheduleNextScheduleCalled, "Unexpected number of calls to schedule.Wait")
}

func TestWaitForNextRun_ValidNextSchedule(t *testing.T) {
	// arrange
	var dummySchedule = &dummySchedule{t: t}
	var dummySession = &session{id: uuid.New()}
	var dummyApplication = &application{
		name:     "some name",
		schedule: dummySchedule,
		session:  dummySession,
	}
	var dummyTimeNow = time.Now()
	var dummyDuration = time.Duration(rand.Intn(1000)) + 10*time.Second
	var dummyTimeNext = dummyTimeNow.Add(dummyDuration)
	var dummyMessageFormat = "Next run at [%v]: waiting for [%v]"
	var dummyControlChannel = make(chan time.Time)
	var scheduleNextScheduleExpected int
	var scheduleNextScheduleCalled int

	// mock
	createMock(t)

	// expect
	scheduleNextScheduleExpected = 1
	dummySchedule.nextSchedule = func() *time.Time {
		scheduleNextScheduleCalled++
		return &dummyTimeNext
	}
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyTimeNow
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category, subcategory, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "waitForNextRun", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 2, len(parameters))
		assert.Equal(t, dummyTimeNext, parameters[0])
		assert.Equal(t, dummyDuration, parameters[1])
	}
	timeAfterExpected = 1
	timeAfter = func(d time.Duration) <-chan time.Time {
		timeAfterCalled++
		assert.Equal(t, dummyDuration, d)
		return dummyControlChannel
	}

	// SUT + act
	go waitForNextRun(
		dummyApplication,
	)

	// push
	dummyControlChannel <- dummyTimeNext

	// verify
	verifyAll(t)
	assert.Equal(t, scheduleNextScheduleExpected, scheduleNextScheduleCalled, "Unexpected number of calls to schedule.Wait")
}

func TestRunInstances_ZeroInstance(t *testing.T) {
	// arrange
	var dummyApplication = &application{}

	// mock
	createMock(t)

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// verify
	verifyAll(t)
}

func TestRunInstances_SingleInstance(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		instances: 1,
	}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	handleSessionFuncExpected = 1
	handleSessionFunc = func(app *application, index int) error {
		handleSessionFuncCalled++
		assert.Equal(t, dummyApplication, app)
		assert.Equal(t, 0, index)
		return dummyError
	}

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// assert
	assert.Equal(t, 1, len(dummyApplication.lastErrors))
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])

	// verify
	verifyAll(t)
}

func TestRunInstances_MultipleInstances(t *testing.T) {
	// arrange
	var dummyErrors = []error{
		errors.New("some error 1"),
		errors.New("some error 2"),
		errors.New("some error 3"),
	}
	var dummyApplication = &application{
		instances: len(dummyErrors),
	}
	var expectedIndex = map[int]bool{}
	var writeLock sync.Mutex

	// mock
	createMock(t)

	// expect
	handleSessionFuncExpected = dummyApplication.instances
	handleSessionFunc = func(app *application, index int) error {
		handleSessionFuncCalled++
		assert.Equal(t, dummyApplication, app)
		writeLock.Lock()
		expectedIndex[index] = true
		writeLock.Unlock()
		return dummyErrors[index]
	}

	// SUT + act
	runInstances(
		dummyApplication,
	)

	// assert
	assert.Equal(t, dummyApplication.instances, len(expectedIndex))
	assert.True(t, expectedIndex[0])
	assert.True(t, expectedIndex[1])
	assert.True(t, expectedIndex[2])
	assert.Equal(t, dummyApplication.instances, len(dummyApplication.lastErrors))
	assert.ElementsMatch(t, dummyErrors, dummyApplication.lastErrors)

	// verify
	verifyAll(t)
}

func TestScheduleExecution_WithOverlap(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
		overlap: true,
	}
	var controlFlag = make(chan bool)

	// mock
	createMock(t)

	// expect
	waitForNextRunFuncExpected = 2
	waitForNextRunFunc = func(app *application) {
		waitForNextRunFuncCalled++
		assert.Equal(t, dummyApplication, app)
		app.started = (waitForNextRunFuncCalled < waitForNextRunFuncExpected)
		<-controlFlag
	}
	runInstancesFuncExpected = 1
	runInstancesFunc = func(app *application) {
		runInstancesFuncCalled++
		assert.Equal(t, dummyApplication, app)
		controlFlag <- true
	}

	// SUT + act
	go scheduleExecution(
		dummyApplication,
	)

	// push
	controlFlag <- true

	// verify
	verifyAll(t)
}

func TestScheduleExecution_NoOverlap(t *testing.T) {
	// arrange
	var dummyApplication = &application{
		name:    "some name",
		started: true,
		overlap: false,
	}

	// mock
	createMock(t)

	// expect
	waitForNextRunFuncExpected = 2
	waitForNextRunFunc = func(app *application) {
		waitForNextRunFuncCalled++
		assert.Equal(t, dummyApplication, app)
		app.started = (waitForNextRunFuncCalled < waitForNextRunFuncExpected)
	}
	runInstancesFuncExpected = 1
	runInstancesFunc = func(app *application) {
		runInstancesFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}

	// SUT + act
	scheduleExecution(
		dummyApplication,
	)

	// verify
	verifyAll(t)
}

func TestRunApplication_NoSchedule(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	var dummySchedule = &dummySchedule{t: t}
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		schedule: dummySchedule,
	}

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummySchedule, i)
		return true
	}
	runInstancesFuncExpected = 1
	runInstancesFunc = func(app *application) {
		runInstancesFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}

	// SUT + act
	go runApplication(
		dummyApplication,
	)

	// assert
	assert.True(t, <-dummyShutdown)

	// verify
	verifyAll(t)
}

func TestRunApplication_WithSchedule(t *testing.T) {
	// arrange
	var dummyShutdown = make(chan bool)
	var dummySchedule = &dummySchedule{t: t}
	var dummyApplication = &application{
		name:     "some name",
		shutdown: dummyShutdown,
		schedule: dummySchedule,
	}

	// mock
	createMock(t)

	// expect
	isInterfaceValueNilFuncExpected = 1
	isInterfaceValueNilFunc = func(i interface{}) bool {
		isInterfaceValueNilFuncCalled++
		assert.Equal(t, dummySchedule, i)
		return false
	}
	scheduleExecutionFuncExpected = 1
	scheduleExecutionFunc = func(app *application) {
		scheduleExecutionFuncCalled++
		assert.Equal(t, dummyApplication, app)
	}

	// SUT + act
	go runApplication(
		dummyApplication,
	)

	// assert
	assert.True(t, <-dummyShutdown)

	// verify
	verifyAll(t)
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
	createMock(t)

	// expect
	runApplicationFuncExpected = 1
	runApplicationFunc = func(app *application) {
		runApplicationFuncCalled++
		assert.True(t, dummyApplication.started)
		assert.Equal(t, dummyApplication, app)
		dummyShutdown <- true
	}
	logAppRootFuncExpected = 2
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "beginApplication", subcategory)
		if logAppRootFuncCalled == 1 {
			assert.Equal(t, "Trying to start runner [%v] (v-%v)", messageFormat)
			assert.Equal(t, 2, len(parameters))
			assert.Equal(t, dummyName, parameters[0])
			assert.Equal(t, dummyVersion, parameters[1])
		} else if logAppRootFuncCalled == 2 {
			assert.Equal(t, "Runner terminated", messageFormat)
			assert.Empty(t, parameters)
		}
	}

	// SUT + act
	beginApplication(
		dummyApplication,
	)

	// assert
	assert.False(t, dummyApplication.started)

	// verify
	verifyAll(t)
}

type dummyCustomizationEndApplication struct {
	dummyCustomization
	appClosing func() error
}

func (customization *dummyCustomizationEndApplication) AppClosing() error {
	if customization.appClosing != nil {
		return customization.appClosing()
	}
	assert.Fail(customization.t, "Unexpected call to AppClosing")
	return nil
}

func TestEndApplication_Error(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationEndApplication{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationAppClosingExpected int
	var customizationAppClosingCalled int
	var dummyError = errors.New("some error")
	var dummyMessageFormat = "Failed to execute customization.AppClosing. Error: %+v"

	// mock
	createMock(t)

	// expect
	customizationAppClosingExpected = 1
	dummyCustomization.appClosing = func() error {
		customizationAppClosingCalled++
		return dummyError
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "endApplication", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Equal(t, 1, len(parameters))
		assert.Equal(t, dummyError, parameters[0])
	}

	// SUT + act
	endApplication(
		dummyApplication,
	)

	// assert
	assert.Equal(t, 1, len(dummyApplication.lastErrors))
	assert.Equal(t, dummyError, dummyApplication.lastErrors[0])

	// verify
	verifyAll(t)
	assert.Equal(t, customizationAppClosingExpected, customizationAppClosingCalled, "Unexpected number of calls to method customization.AppClosing")
}

func TestEndApplication_Success(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCustomization = &dummyCustomizationEndApplication{
		dummyCustomization: dummyCustomization{t: t},
	}
	var dummyApplication = &application{
		session:       dummySession,
		customization: dummyCustomization,
	}
	var customizationAppClosingExpected int
	var customizationAppClosingCalled int
	var dummyMessageFormat = "customization.AppClosing executed successfully"

	// mock
	createMock(t)

	// expect
	customizationAppClosingExpected = 1
	dummyCustomization.appClosing = func() error {
		customizationAppClosingCalled++
		return nil
	}
	logAppRootFuncExpected = 1
	logAppRootFunc = func(session *session, category string, subcategory string, messageFormat string, parameters ...interface{}) {
		logAppRootFuncCalled++
		assert.Equal(t, dummySession, session)
		assert.Equal(t, "application", category)
		assert.Equal(t, "endApplication", subcategory)
		assert.Equal(t, dummyMessageFormat, messageFormat)
		assert.Empty(t, parameters)
	}

	// SUT + act
	endApplication(
		dummyApplication,
	)

	// assert
	assert.Empty(t, dummyApplication.lastErrors)

	// verify
	verifyAll(t)
	assert.Equal(t, customizationAppClosingExpected, customizationAppClosingCalled, "Unexpected number of calls to method customization.AppClosing")
}
