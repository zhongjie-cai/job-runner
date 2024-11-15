package jobrunner

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_AppRoot(t *testing.T) {
	// arrange
	var appRootValue = 0

	// SUT
	var sut = LogType(appRootValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, LogTypeAppRoot, sut)
	assert.Equal(t, appRootLogTypeName, result)
}

func TestString_NonSupportedLogTypes(t *testing.T) {
	// arrange
	var unsupportedValue = 1 << 31

	// SUT
	var sut = LogType(unsupportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Zero(t, result)
}

func TestString_SingleSupportedLogType(t *testing.T) {
	// SUT
	var sut = LogTypeMethodLogic

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, methodLogicLogTypeName, result)
}

func TestString_MultipleSupportedLogTypes(t *testing.T) {
	// arrange
	var supportedValue = LogTypeProcessEnter | LogTypeProcessRequest | LogTypeMethodLogic | LogTypeProcessResponse | LogTypeProcessExit

	// SUT
	var sut = LogType(supportedValue)

	// act
	var result = sut.String()

	// assert
	assert.Equal(t, LogTypeGeneralLogging, sut)
	assert.True(t, strings.Contains(result, apiEnterLogTypeName))
	assert.True(t, strings.Contains(result, apiRequestLogTypeName))
	assert.True(t, strings.Contains(result, methodLogicLogTypeName))
	assert.True(t, strings.Contains(result, apiResponseLogTypeName))
	assert.True(t, strings.Contains(result, apiExitLogTypeName))
}

func TestHasFlag_FlagMatch_AppRoot(t *testing.T) {
	// arrange
	var flag = LogTypeAppRoot

	// SUT
	var sut = LogTypeAppRoot

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)
}

func TestHasFlag_FlagNoMatch_AppRoot(t *testing.T) {
	// arrange
	var flag = LogTypeAppRoot

	// SUT
	var sut = LogTypeProcessEnter | LogTypeProcessExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)
}

func TestHasFlag_FlagMatch_NotAppRoot(t *testing.T) {
	// arrange
	var flag = LogTypeMethodLogic

	// SUT
	var sut = LogTypeProcessEnter | LogTypeMethodLogic | LogTypeProcessExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.True(t, result)
}

func TestHasFlag_FlagNoMatch_NotAppRoot(t *testing.T) {
	// arrange
	var flag = LogTypeMethodLogic

	// SUT
	var sut = LogTypeProcessEnter | LogTypeProcessExit

	// act
	var result = sut.HasFlag(flag)

	// assert
	assert.False(t, result)
}

func TestNewLogType_NoMatchFound(t *testing.T) {
	// arrange
	var dummyValue = "some value"

	// SUT + act
	var result = NewLogType(dummyValue)

	// assert
	assert.Equal(t, LogTypeAppRoot, result)
}

func TestNewLogType_AppRoot(t *testing.T) {
	// arrange
	var dummyValue = appRootLogTypeName

	// SUT + act
	var result = NewLogType(dummyValue)

	// assert
	assert.Equal(t, LogTypeAppRoot, result)
}

func TestNewLogType_HappyPath(t *testing.T) {
	for key, value := range logTypeNameMapping {
		// SUT + act
		var result = NewLogType(key)

		// assert
		assert.Equal(t, value, result)
	}
}
