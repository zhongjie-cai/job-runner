package jobrunner

import (
	"time"
)

func initiateSession(
	app *application,
	index int,
	reruns int,
) *session {
	return &session{
		id:            uuidNew(),
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
	return fmtErrorf(
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
	var session = initiateSessionFunc(
		app,
		index,
		reruns,
	)
	logProcessEnterFunc(
		session,
		app.name,
		"",
		"",
	)
	logProcessRequestFunc(
		session,
		app.name,
		"InstanceIndex",
		"%v",
		index,
	)
	defer func(startTime time.Time) {
		err = finalizeSessionFunc(
			session,
			err,
			recover(),
		)
		logProcessResponseFunc(
			session,
			app.name,
			"",
			"%v",
			err,
		)
		logProcessExitFunc(
			session,
			app.name,
			"Duration",
			"%s",
			timeSince(startTime),
		)
	}(
		getTimeNowUTCFunc(),
	)
	return processSessionFunc(
		session,
		app.customization,
	)
}
