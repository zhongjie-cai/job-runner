package jobrunner

import (
	"sync"
	"time"
)

// Application is the interface for job runner application
type Application interface {
	// Start starts the job runner according to given specifications for number of instances (in parallel) and repeat frequency defined in application
	Start()
	// Stop interrupts the job runner hosting, causing the job runner to forcefully shutdown
	Stop()
}

type application struct {
	name          string
	version       string
	instances     int
	repeat        *time.Duration
	session       *session
	customization Customization
	shutdown      chan bool
	started       bool
}

// NewApplication creates a new application for job runner hosting
func NewApplication(
	name string,
	version string,
	instances int,
	repeat *time.Duration,
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
		&session{
			uuidNew(),
			0,
			map[string]interface{}{},
			customization,
		},
		customization,
		make(chan bool),
		false,
	}
	return application
}

func (app *application) Start() {
	startApplicationFunc(
		app,
	)
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
			handleSessionFunc(
				app,
				index,
			)
			waitGroup.Done()
		}(id)
	}
	waitGroup.Wait()
}

func repeatExecution(app *application) {
	for {
		go runInstancesFunc(
			app,
		)
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
	go runApplicationFunc(app)
	<-app.shutdown
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
	} else {
		logAppRootFunc(
			app.session,
			"application",
			"endApplication",
			"customization.AppClosing executed successfully",
		)
	}
}
