package jobrunner

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestIsInterfaceValueNil_NilInterface(t *testing.T) {
	// arrange
	var dummyInterface http.ResponseWriter

	// SUT + act
	var result = isInterfaceValueNil(
		dummyInterface,
	)

	// assert
	assert.True(t, result)
}

func TestIsInterfaceValueNil_NilValue(t *testing.T) {
	// arrange
	var dummyInterface *DefaultCustomization

	// SUT + act
	var result = isInterfaceValueNil(
		dummyInterface,
	)

	// assert
	assert.True(t, result)
}

func TestIsInterfaceValueNil_EmptyValue(t *testing.T) {
	// arrange
	var dummyInterface = 0

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(reflect.Value.IsValid).Expects(gomocker.Anything()).Returns(false).Once()

	// SUT + act
	var result = isInterfaceValueNil(
		dummyInterface,
	)

	// assert
	assert.True(t, result)
}

func TestIsInterfaceValueNil_ValidValue(t *testing.T) {
	// arrange
	var dummyInterface = 0

	// SUT + act
	var result = isInterfaceValueNil(
		dummyInterface,
	)

	// assert
	assert.False(t, result)
}
