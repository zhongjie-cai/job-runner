package jobrunner

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func initiateSession(
	app *application,
	index int,
	reruns int,
) *session {
	return &session{
		id:            uuid.New(),
		index:         index,
		reruns:        reruns,
		attachment:    map[string]interface{}{},
		customization: app.customization,
	}
}

func finalizeSession(
	session *session,
	errorResult error,
	recoverResult interface{},
) error {
	var recoverError = session.customization.RecoverPanic(
		session,
		recoverResult,
	)
	if errorResult == nil {
		return recoverError
	}
	return fmt.Errorf(
		"Original Error: %w\nRecover Error: %v",
		errorResult,
		recoverError,
	)
}

func processSession(
	session Session,
	customization Customization,
) error {
	var preActionError = customization.PreAction(
		session,
	)
	if preActionError != nil {
		return preActionError
	}
	var actionError = customization.ActionFunc(
		session,
	)
	if actionError != nil {
		return actionError
	}
	var postActionError = customization.PostAction(
		session,
	)
	if postActionError != nil {
		return postActionError
	}
	return nil
}

// handleSession wraps the HTTP handler with session related operations
func handleSession(
	app *application,
	index int,
	reruns int,
) (err error) {
	var session = initiateSession(
		app,
		index,
		reruns,
	)
	logProcessEnter(
		session,
		app.name,
		"",
		"",
	)
	logProcessRequest(
		session,
		app.name,
		"InstanceIndex",
		"%v",
		index,
	)
	defer func(startTime time.Time) {
		err = finalizeSession(
			session,
			err,
			recover(),
		)
		logProcessResponse(
			session,
			app.name,
			"",
			"%v",
			err,
		)
		logProcessExit(
			session,
			app.name,
			"Duration",
			"%s",
			time.Since(startTime),
		)
	}(
		time.Now().UTC(),
	)
	return processSession(
		session,
		app.customization,
	)
}
