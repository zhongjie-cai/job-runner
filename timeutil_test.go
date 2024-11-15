package jobrunner

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	// arrange
	var dummyTime = time.Date(2345, 6, 7, 8, 9, 10, 11, time.UTC)
	var expectedResult = "2345-06-07"

	// SUT + act
	var result = formatDate(dummyTime)

	// assert
	assert.Equal(t, expectedResult, result)
}

func TestFormatTime(t *testing.T) {
	// arrange
	var dummyTime = time.Date(2345, 6, 7, 8, 9, 10, 11, time.UTC)
	var expectedResult = "08:09:10"

	// SUT + act
	var result = formatTime(dummyTime)

	// assert
	assert.Equal(t, expectedResult, result)
}

func TestFormatDateTime(t *testing.T) {
	// arrange
	var dummyTime = time.Date(2345, 6, 7, 8, 9, 10, 11, time.UTC)
	var expectedResult = "2345-06-07T08:09:10"

	// SUT + act
	var result = formatDateTime(dummyTime)

	// assert
	assert.Equal(t, expectedResult, result)
}
