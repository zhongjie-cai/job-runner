package jobrunner

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/google/uuid"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestPrepareLoggingFunc_NilSession(t *testing.T) {
	// arrange
	var dummySession *session
	var dummyLogType = LogType(rand.IntN(100))
	var dummyLogLevel = LogLevel(rand.IntN(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// SUT + act
	prepareLogging(
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestPrepareLoggingFunc_HappyPath(t *testing.T) {
	// arrange
	type customization struct {
		Customization
	}
	var dummyCustomization = &customization{}
	var dummySession = &session{
		customization: dummyCustomization,
	}
	var dummyLogType = LogType(rand.IntN(100))
	var dummyLogLevel = LogLevel(rand.IntN(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format %v %v %v"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")
	var dummyDescription = fmt.Sprintf(dummyMessageFormat, dummyParameter1, dummyParameter2, dummyParameter3)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock((*customization).Log).Expects(
		dummyCustomization,
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyDescription,
	).Returns().Once()

	// SUT + act
	prepareLogging(
		dummySession,
		dummyLogType,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)

}

func TestLogAppRoot(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeAppRoot,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logAppRoot(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogProcessEnter(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeProcessEnter,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logProcessEnter(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogProcessRequest(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeProcessRequest,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logProcessRequest(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogMethodEnter(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeMethodEnter,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logMethodEnter(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogMethodParameter(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeMethodParameter,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logMethodParameter(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogMethodLogic(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyLogLevel = LogLevel(rand.IntN(100))
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeMethodLogic,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logMethodLogic(
		dummySession,
		dummyLogLevel,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogWebcallStart(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeWebcallStart,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logWebcallStart(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogWebcallRequest(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeWebcallRequest,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logWebcallRequest(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogWebcallResponse(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeWebcallResponse,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logWebcallResponse(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogWebcallFinish(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeWebcallFinish,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logWebcallFinish(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogMethodReturn(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeMethodReturn,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logMethodReturn(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogMethodExit(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeMethodExit,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logMethodExit(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogProcessResponse(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeProcessResponse,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logProcessResponse(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}

func TestLogProcessExit(t *testing.T) {
	// arrange
	var dummySession = &session{
		id: uuid.New(),
	}
	var dummyCategory = "some category"
	var dummySubcategory = "some subcategory"
	var dummyMessageFormat = "some message format"
	var dummyParameter1 = "some parameter 1"
	var dummyParameter2 = rand.Int()
	var dummyParameter3 = errors.New("some parameter 3")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(prepareLogging).Expects(
		dummySession,
		LogTypeProcessExit,
		LogLevelInfo,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	).Returns().Once()

	// SUT + act
	logProcessExit(
		dummySession,
		dummyCategory,
		dummySubcategory,
		dummyMessageFormat,
		dummyParameter1,
		dummyParameter2,
		dummyParameter3,
	)
}
