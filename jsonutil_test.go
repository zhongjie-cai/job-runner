package jobrunner

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestMarshalIgnoreError_Empty(t *testing.T) {
	// arrange
	var dummyObject *struct {
		Foo  string
		Test int
	}
	var expectedResult = "null"

	// SUT + act
	var result = marshalIgnoreError(
		dummyObject,
	)

	// assert
	assert.Equal(t, expectedResult, result)
}

func TestMarshalIgnoreError_Success(t *testing.T) {
	// arrange
	var dummyObject = struct {
		Foo  string
		Test int
	}{
		"<bar />",
		123,
	}
	var expectedResult = "{\"Foo\":\"<bar />\",\"Test\":123}"

	// SUT + act
	var result = marshalIgnoreError(
		dummyObject,
	)

	// assert
	assert.Equal(t, expectedResult, result)
}

func TestTryUnmarshalPrimitiveTypes_Empty(t *testing.T) {
	// arrange
	var dummyValue string
	var dummyDataTemplate string

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_String(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate string

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Bool_Error(t *testing.T) {
	// arrange
	var dummyValue = "some invalid bool"
	var dummyDataTemplate bool

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Bool_NoError(t *testing.T) {
	// arrange
	var dummyValue = rand.IntN(100) < 50
	var dummyValueString = fmt.Sprintf("%v", dummyValue)
	var dummyDataTemplate bool

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Integer_Error(t *testing.T) {
	// arrange
	var dummyValue = "some invalid integer"
	var dummyDataTemplate int

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Integer_NoError(t *testing.T) {
	// arrange
	var dummyValue = rand.IntN(math.MaxInt32)
	var dummyValueString = fmt.Sprintf("%v", dummyValue)
	var dummyDataTemplate int

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Int64_Error(t *testing.T) {
	// arrange
	var dummyValue = "some invalid int64"
	var dummyDataTemplate int64

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Int64_NoError(t *testing.T) {
	// arrange
	var dummyValue = rand.Int64()
	var dummyValueString = fmt.Sprintf("%v", dummyValue)
	var dummyDataTemplate int64

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Float64_Error(t *testing.T) {
	// arrange
	var dummyValue = "some invalid float64"
	var dummyDataTemplate float64

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Float64_NoError(t *testing.T) {
	// arrange
	var dummyValue = rand.Float64()
	var dummyValueString = fmt.Sprintf("%v", dummyValue)
	var dummyDataTemplate float64

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Byte_Error(t *testing.T) {
	// arrange
	var dummyValue = "some invalid byte"
	var dummyDataTemplate byte

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_Byte_NoError(t *testing.T) {
	// arrange
	var dummyValue = byte(rand.IntN(math.MaxUint8))
	var dummyValueString = fmt.Sprintf("%v", dummyValue)
	var dummyDataTemplate byte

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.True(t, result)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshalPrimitiveTypes_OtherTypes(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate error

	// SUT + act
	var result = tryUnmarshalPrimitiveTypes(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.False(t, result)
	assert.Zero(t, dummyDataTemplate)
}

func TestTryUnmarshal_NilDataTemplate(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate any

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(true).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Nil(t, dummyDataTemplate)
}

func TestTryUnmarshal_PrimitiveType(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate string
	var dummyData = "some data"

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(false).Once()
	m.Mock(tryUnmarshalPrimitiveTypes).Expects(dummyValue, gomocker.Anything()).Returns(true).SideEffects(
		gomocker.ParamSideEffect(1, 2, func(value *string) { *value = dummyData })).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyData, dummyDataTemplate)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Primitive(t *testing.T) {
	// arrange
	var dummyValue = rand.Int()
	var dummyValueString = strconv.Itoa(dummyValue)
	var dummyDataTemplate int

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(false).Once()
	m.Mock(tryUnmarshalPrimitiveTypes).Expects(dummyValueString, gomocker.Anything()).Returns(false).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshal_NoQuoteJSONSuccess_Struct(t *testing.T) {
	// arrange
	var dummyValueString = "{\"foo\":\"bar\",\"test\":123}"
	var dummyDataTemplate struct {
		Foo  string `json:"foo"`
		Test int    `json:"test"`
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(false).Once()
	m.Mock(tryUnmarshalPrimitiveTypes).Expects(dummyValueString, gomocker.Anything()).Returns(false).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValueString,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "bar", dummyDataTemplate.Foo)
	assert.Equal(t, 123, dummyDataTemplate.Test)
}

func TestTryUnmarshal_WithQuoteJSONSuccess(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate string

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(false).Once()
	m.Mock(tryUnmarshalPrimitiveTypes).Expects(dummyValue, gomocker.Anything()).Returns(false).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, dummyValue, dummyDataTemplate)
}

func TestTryUnmarshal_Failure(t *testing.T) {
	// arrange
	var dummyValue = "some value"
	var dummyDataTemplate uuid.UUID
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(isInterfaceValueNil).Expects(&dummyDataTemplate).Returns(false).Once()
	m.Mock(tryUnmarshalPrimitiveTypes).Expects(dummyValue, gomocker.Anything()).Returns(false).Once()
	m.Mock(fmt.Errorf).Expects("Unable to unmarshal value [%v] into data template", dummyValue).Returns(dummyError).Once()

	// SUT + act
	var err = tryUnmarshal(
		dummyValue,
		&dummyDataTemplate,
	)

	// assert
	assert.Equal(t, dummyError, err)
	assert.Zero(t, dummyDataTemplate)
}
