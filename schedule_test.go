package jobrunner

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestMoveValueIndex_IndexOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.IntN(10) + 5
	var dummyValues = []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}
	var dummyMaxValue = rand.Int()

	// SUT + act
	var value, index, reset = moveValueIndex(
		dummyOldIndex,
		dummyValues,
		dummyMaxValue,
	)

	// assert
	assert.Equal(t, dummyValues[0], value)
	assert.Zero(t, index)
	assert.True(t, reset)
}

func TestMoveValueIndex_ValueOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.IntN(5)
	var dummyValues = []int{
		rand.IntN(100) + 10,
		rand.IntN(100) + 10,
		rand.IntN(100) + 10,
		rand.IntN(100) + 10,
		rand.IntN(100) + 10,
	}
	var dummyMaxValue = rand.IntN(10)

	// SUT + act
	var value, index, reset = moveValueIndex(
		dummyOldIndex,
		dummyValues,
		dummyMaxValue,
	)

	// assert
	assert.Equal(t, dummyValues[0], value)
	assert.Zero(t, index)
	assert.True(t, reset)
}

func TestMoveValueIndex_NoOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.IntN(5)
	var dummyValues = []int{
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
		rand.IntN(100),
	}
	var dummyMaxValue = rand.IntN(100) + 100

	// SUT + act
	var value, index, reset = moveValueIndex(
		dummyOldIndex,
		dummyValues,
		dummyMaxValue,
	)

	// assert
	assert.Equal(t, dummyValues[dummyOldIndex+1], value)
	assert.Equal(t, dummyOldIndex+1, index)
	assert.False(t, reset)
}

func TestGetDaysOfMonth_Mocked(t *testing.T) {
	// arrange
	var dummyYear = 2000 + rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyTime = time.Now()

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+2),
		0, 0, 0, 0, 0, time.Local).Returns(dummyTime).Once()

	// SUT + act
	var result = getDaysOfMonth(
		dummyYear,
		dummyMonth,
	)

	// assert
	assert.Equal(t, dummyTime.Day(), result)
}

func TestGetDaysOfMonth_Integration(t *testing.T) {
	// arrange
	var dummyYear = 2020
	var dummyMonth = 1

	// SUT + act
	var result = getDaysOfMonth(
		dummyYear,
		dummyMonth,
	)

	// assert
	assert.Equal(t, 29, result)
}

func TestConstructTimeBySchedule(t *testing.T) {
	// arrange
	var dummyLocation, _ = time.LoadLocation("China/Beijing")
	var dummySchedule = &schedule{
		year:     rand.IntN(100) + 2000,
		month:    rand.IntN(12),
		day:      rand.IntN(31),
		hour:     rand.IntN(24),
		minute:   rand.IntN(60),
		second:   rand.IntN(60),
		timezone: dummyLocation,
	}
	var dummyTime = time.Now()

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Date).Expects(dummySchedule.year, time.Month(dummySchedule.month+1),
		dummySchedule.day+1, dummySchedule.hour, dummySchedule.minute,
		dummySchedule.second, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var result = constructTimeBySchedule(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyTime, result)
}

func TestUpdateScheduleIndex_NoReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).
		Returns(dummyNewSecond, dummyNewSecondIndex, false).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyOldMinute, dummySchedule.minute)
	assert.Equal(t, dummyOldMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyOldHour, dummySchedule.hour)
	assert.Equal(t, dummyOldHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyOldDay, dummySchedule.day)
	assert.Equal(t, dummyOldDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyOldMonth, dummySchedule.month)
	assert.Equal(t, dummyOldMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyOldYear, dummySchedule.year)
	assert.Equal(t, dummyOldYearIndex, dummySchedule.yearIndex)
}

func TestUpdateScheduleIndex_SecondReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)
	var dummyNewMinute = rand.IntN(60)
	var dummyNewMinuteIndex = rand.IntN(60)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).Returns(dummyNewSecond, dummyNewSecondIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.minuteIndex, dummySchedule.minutes, 60).Returns(dummyNewMinute, dummyNewMinuteIndex, false).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyNewMinute, dummySchedule.minute)
	assert.Equal(t, dummyNewMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyOldHour, dummySchedule.hour)
	assert.Equal(t, dummyOldHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyOldDay, dummySchedule.day)
	assert.Equal(t, dummyOldDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyOldMonth, dummySchedule.month)
	assert.Equal(t, dummyOldMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyOldYear, dummySchedule.year)
	assert.Equal(t, dummyOldYearIndex, dummySchedule.yearIndex)
}

func TestUpdateScheduleIndex_MinuteReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)
	var dummyNewMinute = rand.IntN(60)
	var dummyNewMinuteIndex = rand.IntN(60)
	var dummyNewHour = rand.IntN(24)
	var dummyNewHourIndex = rand.IntN(24)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).Returns(dummyNewSecond, dummyNewSecondIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.minuteIndex, dummySchedule.minutes, 60).Returns(dummyNewMinute, dummyNewMinuteIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.hourIndex, dummySchedule.hours, 24).Returns(dummyNewHour, dummyNewHourIndex, false).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyNewMinute, dummySchedule.minute)
	assert.Equal(t, dummyNewMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyNewHour, dummySchedule.hour)
	assert.Equal(t, dummyNewHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyOldDay, dummySchedule.day)
	assert.Equal(t, dummyOldDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyOldMonth, dummySchedule.month)
	assert.Equal(t, dummyOldMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyOldYear, dummySchedule.year)
	assert.Equal(t, dummyOldYearIndex, dummySchedule.yearIndex)
}

func TestUpdateScheduleIndex_HourReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)
	var dummyNewMinute = rand.IntN(60)
	var dummyNewMinuteIndex = rand.IntN(60)
	var dummyNewHour = rand.IntN(24)
	var dummyNewHourIndex = rand.IntN(24)
	var dummyNewDay = rand.IntN(31)
	var dummyNewDayIndex = rand.IntN(31)
	var dummyDaysOfMonth = rand.IntN(31)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).Returns(dummyNewSecond, dummyNewSecondIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.minuteIndex, dummySchedule.minutes, 60).Returns(dummyNewMinute, dummyNewMinuteIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.hourIndex, dummySchedule.hours, 24).Returns(dummyNewHour, dummyNewHourIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.dayIndex, dummySchedule.days, dummyDaysOfMonth).Returns(dummyNewDay, dummyNewDayIndex, false).Once()
	m.Mock(getDaysOfMonth).Expects(dummyOldYear, dummyOldMonth).Returns(dummyDaysOfMonth).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyNewMinute, dummySchedule.minute)
	assert.Equal(t, dummyNewMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyNewHour, dummySchedule.hour)
	assert.Equal(t, dummyNewHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyNewDay, dummySchedule.day)
	assert.Equal(t, dummyNewDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyOldMonth, dummySchedule.month)
	assert.Equal(t, dummyOldMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyOldYear, dummySchedule.year)
	assert.Equal(t, dummyOldYearIndex, dummySchedule.yearIndex)
}

func TestUpdateScheduleIndex_DayReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)
	var dummyNewMinute = rand.IntN(60)
	var dummyNewMinuteIndex = rand.IntN(60)
	var dummyNewHour = rand.IntN(24)
	var dummyNewHourIndex = rand.IntN(24)
	var dummyNewDay = rand.IntN(31)
	var dummyNewDayIndex = rand.IntN(31)
	var dummyDaysOfMonth = rand.IntN(31)
	var dummyNewMonth = rand.IntN(31)
	var dummyNewMonthIndex = rand.IntN(31)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).Returns(dummyNewSecond, dummyNewSecondIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.minuteIndex, dummySchedule.minutes, 60).Returns(dummyNewMinute, dummyNewMinuteIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.hourIndex, dummySchedule.hours, 24).Returns(dummyNewHour, dummyNewHourIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.dayIndex, dummySchedule.days, dummyDaysOfMonth).Returns(dummyNewDay, dummyNewDayIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.monthIndex, dummySchedule.months, 12).Returns(dummyNewMonth, dummyNewMonthIndex, false).Once()
	m.Mock(getDaysOfMonth).Expects(dummyOldYear, dummyOldMonth).Returns(dummyDaysOfMonth).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyNewMinute, dummySchedule.minute)
	assert.Equal(t, dummyNewMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyNewHour, dummySchedule.hour)
	assert.Equal(t, dummyNewHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyNewDay, dummySchedule.day)
	assert.Equal(t, dummyNewDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyNewMonth, dummySchedule.month)
	assert.Equal(t, dummyNewMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyOldYear, dummySchedule.year)
	assert.Equal(t, dummyOldYearIndex, dummySchedule.yearIndex)
}

func TestUpdateScheduleIndex_MonthReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.IntN(60)
	var dummyOldSecondIndex = rand.IntN(60)
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldMinute = rand.IntN(60)
	var dummyOldMinuteIndex = rand.IntN(60)
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyOldHour = rand.IntN(24)
	var dummyOldHourIndex = rand.IntN(24)
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyOldDay = rand.IntN(31)
	var dummyOldDayIndex = rand.IntN(31)
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyOldMonth = rand.IntN(12)
	var dummyOldMonthIndex = rand.IntN(12)
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyOldYear = rand.IntN(100)
	var dummyOldYearIndex = rand.IntN(100)
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		second:      dummyOldSecond,
		secondIndex: dummyOldSecondIndex,
		seconds:     dummySeconds,
		minute:      dummyOldMinute,
		minuteIndex: dummyOldMinuteIndex,
		minutes:     dummyMinutes,
		hour:        dummyOldHour,
		hourIndex:   dummyOldHourIndex,
		hours:       dummyHours,
		day:         dummyOldDay,
		dayIndex:    dummyOldDayIndex,
		days:        dummyDays,
		month:       dummyOldMonth,
		monthIndex:  dummyOldMonthIndex,
		months:      dummyMonths,
		year:        dummyOldYear,
		yearIndex:   dummyOldYearIndex,
		years:       dummyYears,
	}
	var dummyNewSecond = rand.IntN(60)
	var dummyNewSecondIndex = rand.IntN(60)
	var dummyNewMinute = rand.IntN(60)
	var dummyNewMinuteIndex = rand.IntN(60)
	var dummyNewHour = rand.IntN(24)
	var dummyNewHourIndex = rand.IntN(24)
	var dummyNewDay = rand.IntN(31)
	var dummyNewDayIndex = rand.IntN(31)
	var dummyDaysOfMonth = rand.IntN(31)
	var dummyNewMonth = rand.IntN(31)
	var dummyNewMonthIndex = rand.IntN(31)
	var dummyNewYear = rand.IntN(100)
	var dummyNewYearIndex = rand.IntN(100)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(moveValueIndex).Expects(dummySchedule.secondIndex, dummySchedule.seconds, 60).Returns(dummyNewSecond, dummyNewSecondIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.minuteIndex, dummySchedule.minutes, 60).Returns(dummyNewMinute, dummyNewMinuteIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.hourIndex, dummySchedule.hours, 24).Returns(dummyNewHour, dummyNewHourIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.dayIndex, dummySchedule.days, dummyDaysOfMonth).Returns(dummyNewDay, dummyNewDayIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.monthIndex, dummySchedule.months, 12).Returns(dummyNewMonth, dummyNewMonthIndex, true).Once()
	m.Mock(moveValueIndex).Expects(dummySchedule.yearIndex, dummySchedule.years, 9999).Returns(dummyNewYear, dummyNewYearIndex, rand.IntN(100) > 50).Once()
	m.Mock(getDaysOfMonth).Expects(dummyOldYear, dummyOldMonth).Returns(dummyDaysOfMonth).Once()

	// SUT + act
	updateScheduleIndex(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyNewSecond, dummySchedule.second)
	assert.Equal(t, dummyNewSecondIndex, dummySchedule.secondIndex)
	assert.Equal(t, dummyNewMinute, dummySchedule.minute)
	assert.Equal(t, dummyNewMinuteIndex, dummySchedule.minuteIndex)
	assert.Equal(t, dummyNewHour, dummySchedule.hour)
	assert.Equal(t, dummyNewHourIndex, dummySchedule.hourIndex)
	assert.Equal(t, dummyNewDay, dummySchedule.day)
	assert.Equal(t, dummyNewDayIndex, dummySchedule.dayIndex)
	assert.Equal(t, dummyNewMonth, dummySchedule.month)
	assert.Equal(t, dummyNewMonthIndex, dummySchedule.monthIndex)
	assert.Equal(t, dummyNewYear, dummySchedule.year)
	assert.Equal(t, dummyNewYearIndex, dummySchedule.yearIndex)
}

func TestNextSchedule_ExceededTillTime(t *testing.T) {
	// arrange
	var dummyTill = time.Now().Add(-10 * time.Second)

	// SUT
	var sut = &schedule{
		till: &dummyTill,
	}

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Nil(t, result)
}

func TestNextSchedule_AlreadyCompleted(t *testing.T) {
	// SUT
	var sut = &schedule{
		completed: true,
	}

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Nil(t, result)
}

func TestNextSchedule_NextInFuture(t *testing.T) {
	// arrange
	var dummySchedule = &schedule{}
	var dummyTimeNext = time.Now().Add(10 * time.Second)
	var dummyCompleted = rand.IntN(100) > 50

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(constructTimeBySchedule).Expects(dummySchedule).Returns(dummyTimeNext).Once()
	m.Mock(updateScheduleIndex).Expects(dummySchedule).Returns(dummyCompleted).Once()

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Equal(t, dummyTimeNext, *result)
	assert.Equal(t, dummyCompleted, dummySchedule.completed)
}

func TestNextSchedule_NextInPast_NoSkipOverdue(t *testing.T) {
	// arrange
	var dummySchedule = &schedule{
		skipOverdue: false,
	}
	var dummyTimeNext = time.Now().Add(-10 * time.Second)
	var dummyCompleted = rand.IntN(100) > 50

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(constructTimeBySchedule).Expects(dummySchedule).Returns(dummyTimeNext).Once()
	m.Mock(updateScheduleIndex).Expects(dummySchedule).Returns(dummyCompleted).Once()

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Equal(t, dummyTimeNext, *result)
	assert.Equal(t, dummyCompleted, dummySchedule.completed)
}

func TestNextSchedule_NextInPast_SkipOverdue_Completed(t *testing.T) {
	// arrange
	var dummySchedule = &schedule{
		skipOverdue: true,
	}
	var dummyTimeNext = time.Now().Add(-10 * time.Second)
	var dummyCompleted = true

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(constructTimeBySchedule).Expects(dummySchedule).Returns(dummyTimeNext).Once()
	m.Mock(updateScheduleIndex).Expects(dummySchedule).Returns(dummyCompleted).Once()

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Nil(t, result)
	assert.True(t, dummySchedule.completed)
}

func TestNextSchedule_NextInPast_SkipOverdue_NotCompleted(t *testing.T) {
	// arrange
	var dummySchedule = &schedule{
		skipOverdue: true,
	}
	var dummyTimeNextInPast = time.Now().Add(-10 * time.Second)
	var dummyTimeNextInFuture = time.Now().Add(10 * time.Second)
	var dummyCompleted = false

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(constructTimeBySchedule).Expects(dummySchedule).Returns(dummyTimeNextInPast).Once()
	m.Mock(constructTimeBySchedule).Expects(dummySchedule).Returns(dummyTimeNextInFuture).Once()
	m.Mock(updateScheduleIndex).Expects(dummySchedule).Returns(dummyCompleted).Twice()

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Equal(t, dummyTimeNextInFuture, *result)
	assert.False(t, dummySchedule.completed)
}
