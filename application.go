package jobrunner

import (
	"sync"
	"time"
)

// Application is the interface for job runner application
type Application interface {
	// Start starts the job runner according to given specifications for number of instances (in parallel) and repeat frequency defined in application
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
	repeat        *time.Duration
	overlap       bool
	session       *session
	customization Customization
	shutdown      chan bool
	started       bool
	lastErrors    []error
}

// NewApplication creates a new application for job runner hosting
//   instances marks how many action functions to be executed in parallel at once
//   repeat marks how often the action functions should be repeated until stop signal is given
//   overlap marks whether allow overlapping when a previous execution is taking longer than repeat duration
func NewApplication(
	name string,
	version string,
	instances int,
	repeat *time.Duration,
	overlap bool,
	customization Customization,
) Application {
	if isInterfaceValueNilFunc(customization) {
		customization = customizationDefault
	}
	var application = &application{
		name,
		version,
		instances,
		repeat,
		overlap,
		&session{
			uuidNew(),
			0,
			map[string]interface{}{},
			customization,
		},
		customization,
		make(chan bool),
		false,
		[]error{},
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

func runInstances(app *application) {
	var waitGroup sync.WaitGroup
	for id := 0; id < app.instances; id++ {
		waitGroup.Add(1)
		go func(index int) {
			var sessionError = handleSessionFunc(
				app,
				index,
			)
			if sessionError != nil {
				app.lastErrors = append(
					app.lastErrors,
					sessionError,
				)
			}
			waitGroup.Done()
		}(id)
	}
	waitGroup.Wait()
}

func repeatExecution(app *application) {
	for {
		if app.overlap {
			go runInstancesFunc(
				app,
			)
		} else {
			runInstancesFunc(
				app,
			)
		}
		<-timeAfter(
			*app.repeat,
		)
		if !app.started {
			break
		}
	}
}

func runApplication(app *application) {
	if app.repeat == nil {
		runInstancesFunc(
			app,
		)
	} else {
		repeatExecutionFunc(
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
