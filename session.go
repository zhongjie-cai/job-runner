package jobrunner

import (
	"encoding/json"
	"runtime"
	"strconv"

	"github.com/google/uuid"
)

// Session is the storage for the current HTTP request session, containing information needed for logging, monitoring, etc.
type Session interface {
	SessionMeta
	SessionAttachment
	SessionLogging
	SessionWebcall
}

// SessionMeta is a subset of Session interface, containing only meta data related methods
type SessionMeta interface {
	// GetID returns the ID of this registered session object
	GetID() uuid.UUID

	// GetIndex returns the instance index registered to session object for given session ID
	GetIndex() int

	// GetReruns returns the rerun count for the same instance since first scheduled
	GetReruns() int
}

// SessionAttachment is a subset of Session interface, containing only attachment related methods
type SessionAttachment interface {
	// Attach attaches any value object into the given session associated to the session ID
	Attach(name string, value interface{}) bool

	// Detach detaches any value object from the given session associated to the session ID
	Detach(name string) bool

	// GetRawAttachment retrieves any value object from the given session associated to the session ID and returns the raw interface (consumer needs to manually cast, but works for struct with private fields)
	GetRawAttachment(name string) (interface{}, bool)

	// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template (only works for structs with exported fields)
	GetAttachment(name string, dataTemplate interface{}) bool
}

// SessionLogging is a subset of Session interface, containing only logging related methods
type SessionLogging interface {
	// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
	LogMethodEnter()

	// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
	LogMethodParameter(parameters ...interface{})

	// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
	LogMethodLogic(logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{})

	// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
	LogMethodReturn(returns ...interface{})

	// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
	LogMethodExit()
}

// SessionWebcall is a subset of Session interface, containing only webcall related methods
type SessionWebcall interface {
	// CreateWebcallRequest generates a webcall request object to the targeted external web service for the given session associated to the session ID
	CreateWebcallRequest(method string, url string, payload string, sendClientCert bool) WebRequest
}

type session struct {
	id            uuid.UUID
	index         int
	reruns        int
	attachment    map[string]interface{}
	customization Customization
}

// GetID returns the ID of this registered session object
func (session *session) GetID() uuid.UUID {
	if session == nil {
		return uuid.Nil
	}
	return session.id
}

// GetIndex returns the instance index registered to session object for given session ID
func (session *session) GetIndex() int {
	if session == nil {
		return 0
	}
	return session.index
}

// GetReruns returns the rerun count for the same instance since first scheduled
func (session *session) GetReruns() int {
	if session == nil {
		return 0
	}
	return session.reruns
}

// Attach attaches any value object into the given session associated to the session ID
func (session *session) Attach(name string, value interface{}) bool {
	if session == nil {
		return false
	}
	if session.attachment == nil {
		session.attachment = map[string]interface{}{}
	}
	session.attachment[name] = value
	return true
}

// Detach detaches any value object from the given session associated to the session ID
func (session *session) Detach(name string) bool {
	if session == nil {
		return false
	}
	if session.attachment != nil {
		delete(session.attachment, name)
	}
	return true
}

// GetRawAttachment retrieves any value object from the given session associated to the session ID and returns the raw interface (consumer needs to manually cast, but works for struct with private fields)
func (session *session) GetRawAttachment(name string) (interface{}, bool) {
	if session == nil {
		return nil, false
	}
	var attachment, found = session.attachment[name]
	if !found {
		return nil, false
	}
	return attachment, true
}

// GetAttachment retrieves any value object from the given session associated to the session ID and unmarshals the content to given data template
func (session *session) GetAttachment(name string, dataTemplate interface{}) bool {
	if session == nil {
		return false
	}
	var attachment, found = session.GetRawAttachment(name)
	if !found {
		return false
	}
	var bytes, marshalError = json.Marshal(attachment)
	if marshalError != nil {
		return false
	}
	var unmarshalError = json.Unmarshal(
		bytes,
		dataTemplate,
	)
	return unmarshalError == nil
}

func getMethodName() string {
	var pc, _, _, ok = runtime.Caller(3)
	if !ok {
		return "?"
	}
	var fn = runtime.FuncForPC(pc)
	return fn.Name()
}

// LogMethodEnter sends a logging entry of MethodEnter log type for the given session associated to the session ID
func (session *session) LogMethodEnter() {
	var methodName = getMethodName()
	logMethodEnter(
		session,
		methodName,
		"",
		"",
	)
}

// LogMethodParameter sends a logging entry of MethodParameter log type for the given session associated to the session ID
func (session *session) LogMethodParameter(parameters ...interface{}) {
	var methodName = getMethodName()
	for index, parameter := range parameters {
		logMethodParameter(
			session,
			methodName,
			strconv.Itoa(index),
			"%v",
			parameter,
		)
	}
}

// LogMethodLogic sends a logging entry of MethodLogic log type for the given session associated to the session ID
func (session *session) LogMethodLogic(logLevel LogLevel, category string, subcategory string, messageFormat string, parameters ...interface{}) {
	logMethodLogic(
		session,
		logLevel,
		category,
		subcategory,
		messageFormat,
		parameters...,
	)
}

// LogMethodReturn sends a logging entry of MethodReturn log type for the given session associated to the session ID
func (session *session) LogMethodReturn(returns ...interface{}) {
	var methodName = getMethodName()
	for index, returnValue := range returns {
		logMethodReturn(
			session,
			methodName,
			strconv.Itoa(index),
			"%v",
			returnValue,
		)
	}
}

// LogMethodExit sends a logging entry of MethodExit log type for the given session associated to the session ID
func (session *session) LogMethodExit() {
	var methodName = getMethodName()
	logMethodExit(
		session,
		methodName,
		"",
		"",
	)
}

// CreateWebcallRequest generates a webcall request object to the targeted external web service for the given session associated to the session ID
func (session *session) CreateWebcallRequest(
	method string,
	url string,
	payload string,
	sendClientCert bool,
) WebRequest {
	return &webRequest{
		session,
		method,
		url,
		payload,
		map[string][]string{},
		map[string][]string{},
		0,
		nil,
		sendClientCert,
		0,
		[]dataReceiver{},
	}
}
