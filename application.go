package jobrunner

import (
	"sync"
	"sync/atomic"
)

// Application is the interface for job runner application
type Application interface {
	// Start starts the job runner according to given specifications for number of instances (in parallel) and schedule frequency defined in application
	Start()
	// IsRunning returns true if the job has been successfully started and is currently running
	IsRunning() bool
	// LastErrors returns the list of errors occurred during the execution of job instances up until now
	LastErrors() []error
	// Stop interrupts the job runner hosting, causing the job runner to forcefully shutdown
	Stop()
}

type application struct {
	name          string
	version       string
	instances     int
	reruns        []int32
	schedule      Schedule
	overlap       bool
	session       *session
	customization Customization
	shutdown      chan bool
	started       bool
	lastErrors    []error
	waits         sync.WaitGroup
}

// NewApplication creates a new application for job runner hosting
//   instances marks how many action functions to be executed in parallel at once for a single scheduled execution
//   schedule is a CRON schedule managing when the action functions should be executed until stop signal is given
//   overlap marks a new execution should be executed or not when a previous execution has not yet completed
func NewApplication(
	name string,
	version string,
	instances int,
	schedule Schedule,
	overlap bool,
	customization Customization,
) Application {
	if isInterfaceValueNilFunc(customization) {
		customization = customizationDefault
	}
	var application = &application{
		name:      name,
		version:   version,
		instances: instances,
		reruns:    make([]int32, instances),
		schedule:  schedule,
		overlap:   overlap,
		session: &session{
			id:            uuidNew(),
			index:         0,
			attachment:    map[string]interface{}{},
			customization: customization,
		},
		customization: customization,
		shutdown:      make(chan bool),
		started:       false,
		lastErrors:    []error{},
		waits:         sync.WaitGroup{},
	}
	return application
}

func (app *application) Start() {
	startApplicationFunc(
		app,
	)
}

func (app *application) IsRunning() bool {
	return app.started
}

func (app *application) LastErrors() []error {
	return app.lastErrors
}

func (app *application) Stop() {
	if !app.started {
		return
	}
	app.shutdown <- true
}

func startApplication(app *application) {
	if app.started {
		return
	}
	if !preBootstrapingFunc(app) {
		return
	}
	bootstrapFunc(app)
	if !postBootstrapingFunc(app) {
		return
	}
	defer endApplicationFunc(app)
	beginApplicationFunc(app)
}

func preBootstraping(app *application) bool {
	var preBootstrapError = app.customization.PreBootstrap()
	if preBootstrapError != nil {
		logAppRootFunc(
			app.session,
			"application",
			"preBootstraping",
			"Failed to execute customization.PreBootstrap. Error: %+v",
			preBootstrapError,
		)
		app.lastErrors = append(
			app.lastErrors,
			preBootstrapError,
		)
		return false
	}
	logAppRootFunc(
		app.session,
		"application",
		"preBootstraping",
		"customization.PreBootstrap executed successfully",
	)
	return true
}

func bootstrap(app *application) {
	initializeHTTPClientsFunc(
		app.customization.DefaultTimeout(),
		app.customization.SkipServerCertVerification(),
		app.customization.ClientCert(),
		app.customization.RoundTripper,
	)
	logAppRootFunc(
		app.session,
		"application",
		"bootstrap",
		"Application bootstrapped successfully",
	)
}

func postBootstraping(app *application) bool {
	var postBootstrapError = app.customization.PostBootstrap()
	if postBootstrapError != nil {
		logAppRootFunc(
			app.session,
			"application",
			"postBootstraping",
			"Failed to execute customization.PostBootstrap. Error: %+v",
			postBootstrapError,
		)
		app.lastErrors = append(
			app.lastErrors,
			postBootstrapError,
		)
		return false
	}
	logAppRootFunc(
		app.session,
		"application",
		"postBootstraping",
		"customization.PostBootstrap executed successfully",
	)
	return true
}

func waitForNextRun(app *application) {
	var timeNext = app.schedule.NextSchedule()
	if timeNext == nil {
		logAppRootFunc(
			app.session,
			"application",
			"waitForNextRun",
			"No next schedule available, terminating execution",
		)
		app.started = false
		return
	}
	var waitDuration = timeNext.Sub(
		timeNow(),
	)
	logAppRootFunc(
		app.session,
		"application",
		"waitForNextRun",
		"Next run at [%v]: waiting for [%v]",
		*timeNext,
		waitDuration,
	)
	<-timeAfter(
		waitDuration,
	)
}

func runInstances(app *application) {
	var waitGroup sync.WaitGroup
	for id := 0; id < app.instances; id++ {
		waitGroup.Add(1)
		atomic.AddInt32(&app.reruns[id], 1)
		go func(index int, reruns int) {
			var sessionError = handleSessionFunc(
				app,
				index,
				reruns,
			)
			if sessionError != nil {
				app.lastErrors = append(
					app.lastErrors,
					sessionError,
				)
			}
			waitGroup.Done()
		}(id, int(app.reruns[id]))
	}
	waitGroup.Wait()
	if app.overlap {
		app.waits.Done()
	}
}

func scheduleExecution(app *application) {
	for {
		waitForNextRunFunc(
			app,
		)
		if !app.started {
			break
		}
		if app.overlap {
			app.waits.Add(1)
			go runInstancesFunc(
				app,
			)
		} else {
			runInstancesFunc(
				app,
			)
		}
	}
	if app.overlap {
		app.waits.Wait()
	}
}

func runApplication(app *application) {
	if isInterfaceValueNilFunc(app.schedule) {
		runInstancesFunc(
			app,
		)
	} else {
		scheduleExecutionFunc(
			app,
		)
	}
	app.shutdown <- true
}

func beginApplication(app *application) {
	logAppRootFunc(
		app.session,
		"application",
		"beginApplication",
		"Trying to start runner [%v] (v-%v)",
		app.name,
		app.version,
	)
	app.started = true
	go runApplicationFunc(app)
	<-app.shutdown
	app.started = false
	logAppRootFunc(
		app.session,
		"application",
		"beginApplication",
		"Runner terminated",
	)
}

func endApplication(app *application) {
	var appClosingError = app.customization.AppClosing()
	if appClosingError != nil {
		logAppRootFunc(
			app.session,
			"application",
			"endApplication",
			"Failed to execute customization.AppClosing. Error: %+v",
			appClosingError,
		)
		app.lastErrors = append(
			app.lastErrors,
			appClosingError,
		)
	} else {
		logAppRootFunc(
			app.session,
			"application",
			"endApplication",
			"customization.AppClosing executed successfully",
		)
	}
}
