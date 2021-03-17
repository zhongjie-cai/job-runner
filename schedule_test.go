package jobrunner

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMoveValueIndex_IndexOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.Intn(10) + 5
	var dummyValues = []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}
	var dummyMaxValue = rand.Int()

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestMoveValueIndex_ValueOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.Intn(5)
	var dummyValues = []int{
		rand.Intn(100) + 10,
		rand.Intn(100) + 10,
		rand.Intn(100) + 10,
		rand.Intn(100) + 10,
		rand.Intn(100) + 10,
	}
	var dummyMaxValue = rand.Intn(10)

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestMoveValueIndex_NoOverflow(t *testing.T) {
	// arrange
	var dummyOldIndex = rand.Intn(5)
	var dummyValues = []int{
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
	}
	var dummyMaxValue = rand.Intn(100) + 100

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestGetDaysOfMonth_Mocked(t *testing.T) {
	// arrange
	var dummyYear = 2000 + rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyTime = time.Now()

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+2), month)
		assert.Zero(t, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return dummyTime
	}

	// SUT + act
	var result = getDaysOfMonth(
		dummyYear,
		dummyMonth,
	)

	// assert
	assert.Equal(t, dummyTime.Day(), result)

	// verify
	verifyAll(t)
}

func TestGetDaysOfMonth_Integration(t *testing.T) {
	// arrange
	var dummyYear = 2020
	var dummyMonth = 1

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+2), month)
		assert.Zero(t, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return time.Date(year, month, day, hour, min, sec, nsec, loc)
	}

	// SUT + act
	var result = getDaysOfMonth(
		dummyYear,
		dummyMonth,
	)

	// assert
	assert.Equal(t, 29, result)

	// verify
	verifyAll(t)
}

func TestConstructTimeBySchedule(t *testing.T) {
	// arrange
	var dummySchedule = &schedule{
		year:   rand.Intn(100) + 2000,
		month:  rand.Intn(12),
		day:    rand.Intn(31),
		hour:   rand.Intn(24),
		minute: rand.Intn(60),
		second: rand.Intn(60),
	}
	var dummyTime = time.Now()

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummySchedule.year, year)
		assert.Equal(t, time.Month(dummySchedule.month+1), month)
		assert.Equal(t, dummySchedule.day+1, day)
		assert.Equal(t, dummySchedule.hour, hour)
		assert.Equal(t, dummySchedule.minute, min)
		assert.Equal(t, dummySchedule.second, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return dummyTime
	}

	// SUT + act
	var result = constructTimeBySchedule(
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyTime, result)

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_NoReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 1
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		assert.Equal(t, dummySchedule.secondIndex, oldIndex)
		assert.Equal(t, dummySchedule.seconds, values)
		assert.Equal(t, 60, maxValue)
		return dummyNewSecond, dummyNewSecondIndex, false
	}

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

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_SecondReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)
	var dummyNewMinute = rand.Intn(60)
	var dummyNewMinuteIndex = rand.Intn(60)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 2
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		if moveValueIndexFuncCalled == 1 {
			assert.Equal(t, dummySchedule.secondIndex, oldIndex)
			assert.Equal(t, dummySchedule.seconds, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewSecond, dummyNewSecondIndex, true
		}
		assert.Equal(t, dummySchedule.minuteIndex, oldIndex)
		assert.Equal(t, dummySchedule.minutes, values)
		assert.Equal(t, 60, maxValue)
		return dummyNewMinute, dummyNewMinuteIndex, false
	}

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

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_MinuteReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)
	var dummyNewMinute = rand.Intn(60)
	var dummyNewMinuteIndex = rand.Intn(60)
	var dummyNewHour = rand.Intn(24)
	var dummyNewHourIndex = rand.Intn(24)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 3
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		if moveValueIndexFuncCalled == 1 {
			assert.Equal(t, dummySchedule.secondIndex, oldIndex)
			assert.Equal(t, dummySchedule.seconds, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewSecond, dummyNewSecondIndex, true
		} else if moveValueIndexFuncCalled == 2 {
			assert.Equal(t, dummySchedule.minuteIndex, oldIndex)
			assert.Equal(t, dummySchedule.minutes, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewMinute, dummyNewMinuteIndex, true
		}
		assert.Equal(t, dummySchedule.hourIndex, oldIndex)
		assert.Equal(t, dummySchedule.hours, values)
		assert.Equal(t, 24, maxValue)
		return dummyNewHour, dummyNewHourIndex, false
	}

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

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_HourReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)
	var dummyNewMinute = rand.Intn(60)
	var dummyNewMinuteIndex = rand.Intn(60)
	var dummyNewHour = rand.Intn(24)
	var dummyNewHourIndex = rand.Intn(24)
	var dummyNewDay = rand.Intn(31)
	var dummyNewDayIndex = rand.Intn(31)
	var dummyDaysOfMonth = rand.Intn(31)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 4
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		if moveValueIndexFuncCalled == 1 {
			assert.Equal(t, dummySchedule.secondIndex, oldIndex)
			assert.Equal(t, dummySchedule.seconds, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewSecond, dummyNewSecondIndex, true
		} else if moveValueIndexFuncCalled == 2 {
			assert.Equal(t, dummySchedule.minuteIndex, oldIndex)
			assert.Equal(t, dummySchedule.minutes, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewMinute, dummyNewMinuteIndex, true
		} else if moveValueIndexFuncCalled == 3 {
			assert.Equal(t, dummySchedule.hourIndex, oldIndex)
			assert.Equal(t, dummySchedule.hours, values)
			assert.Equal(t, 24, maxValue)
			return dummyNewHour, dummyNewHourIndex, true
		}
		assert.Equal(t, dummySchedule.dayIndex, oldIndex)
		assert.Equal(t, dummySchedule.days, values)
		assert.Equal(t, dummyDaysOfMonth, maxValue)
		return dummyNewDay, dummyNewDayIndex, false
	}
	getDaysOfMonthFuncExpected = 1
	getDaysOfMonthFunc = func(year, month int) int {
		getDaysOfMonthFuncCalled++
		assert.Equal(t, dummyOldYear, year)
		assert.Equal(t, dummyOldMonth, month)
		return dummyDaysOfMonth
	}

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

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_DayReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)
	var dummyNewMinute = rand.Intn(60)
	var dummyNewMinuteIndex = rand.Intn(60)
	var dummyNewHour = rand.Intn(24)
	var dummyNewHourIndex = rand.Intn(24)
	var dummyNewDay = rand.Intn(31)
	var dummyNewDayIndex = rand.Intn(31)
	var dummyDaysOfMonth = rand.Intn(31)
	var dummyNewMonth = rand.Intn(31)
	var dummyNewMonthIndex = rand.Intn(31)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 5
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		if moveValueIndexFuncCalled == 1 {
			assert.Equal(t, dummySchedule.secondIndex, oldIndex)
			assert.Equal(t, dummySchedule.seconds, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewSecond, dummyNewSecondIndex, true
		} else if moveValueIndexFuncCalled == 2 {
			assert.Equal(t, dummySchedule.minuteIndex, oldIndex)
			assert.Equal(t, dummySchedule.minutes, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewMinute, dummyNewMinuteIndex, true
		} else if moveValueIndexFuncCalled == 3 {
			assert.Equal(t, dummySchedule.hourIndex, oldIndex)
			assert.Equal(t, dummySchedule.hours, values)
			assert.Equal(t, 24, maxValue)
			return dummyNewHour, dummyNewHourIndex, true
		} else if moveValueIndexFuncCalled == 4 {
			assert.Equal(t, dummySchedule.dayIndex, oldIndex)
			assert.Equal(t, dummySchedule.days, values)
			assert.Equal(t, dummyDaysOfMonth, maxValue)
			return dummyNewDay, dummyNewDayIndex, true
		}
		assert.Equal(t, dummySchedule.monthIndex, oldIndex)
		assert.Equal(t, dummySchedule.months, values)
		assert.Equal(t, 12, maxValue)
		return dummyNewMonth, dummyNewMonthIndex, false
	}
	getDaysOfMonthFuncExpected = 1
	getDaysOfMonthFunc = func(year, month int) int {
		getDaysOfMonthFuncCalled++
		assert.Equal(t, dummyOldYear, year)
		assert.Equal(t, dummyOldMonth, month)
		return dummyDaysOfMonth
	}

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

	// verify
	verifyAll(t)
}

func TestUpdateScheduleIndex_MonthReset(t *testing.T) {
	// arrange
	var dummyOldSecond = rand.Intn(60)
	var dummyOldSecondIndex = rand.Intn(60)
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldMinute = rand.Intn(60)
	var dummyOldMinuteIndex = rand.Intn(60)
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyOldHour = rand.Intn(24)
	var dummyOldHourIndex = rand.Intn(24)
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyOldDay = rand.Intn(31)
	var dummyOldDayIndex = rand.Intn(31)
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyOldMonth = rand.Intn(12)
	var dummyOldMonthIndex = rand.Intn(12)
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyOldYear = rand.Intn(100)
	var dummyOldYearIndex = rand.Intn(100)
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyNewSecond = rand.Intn(60)
	var dummyNewSecondIndex = rand.Intn(60)
	var dummyNewMinute = rand.Intn(60)
	var dummyNewMinuteIndex = rand.Intn(60)
	var dummyNewHour = rand.Intn(24)
	var dummyNewHourIndex = rand.Intn(24)
	var dummyNewDay = rand.Intn(31)
	var dummyNewDayIndex = rand.Intn(31)
	var dummyDaysOfMonth = rand.Intn(31)
	var dummyNewMonth = rand.Intn(31)
	var dummyNewMonthIndex = rand.Intn(31)
	var dummyNewYear = rand.Intn(100)
	var dummyNewYearIndex = rand.Intn(100)

	// mock
	createMock(t)

	// expect
	moveValueIndexFuncExpected = 6
	moveValueIndexFunc = func(oldIndex int, values []int, maxValue int) (int, int, bool) {
		moveValueIndexFuncCalled++
		if moveValueIndexFuncCalled == 1 {
			assert.Equal(t, dummySchedule.secondIndex, oldIndex)
			assert.Equal(t, dummySchedule.seconds, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewSecond, dummyNewSecondIndex, true
		} else if moveValueIndexFuncCalled == 2 {
			assert.Equal(t, dummySchedule.minuteIndex, oldIndex)
			assert.Equal(t, dummySchedule.minutes, values)
			assert.Equal(t, 60, maxValue)
			return dummyNewMinute, dummyNewMinuteIndex, true
		} else if moveValueIndexFuncCalled == 3 {
			assert.Equal(t, dummySchedule.hourIndex, oldIndex)
			assert.Equal(t, dummySchedule.hours, values)
			assert.Equal(t, 24, maxValue)
			return dummyNewHour, dummyNewHourIndex, true
		} else if moveValueIndexFuncCalled == 4 {
			assert.Equal(t, dummySchedule.dayIndex, oldIndex)
			assert.Equal(t, dummySchedule.days, values)
			assert.Equal(t, dummyDaysOfMonth, maxValue)
			return dummyNewDay, dummyNewDayIndex, true
		} else if moveValueIndexFuncCalled == 5 {
			assert.Equal(t, dummySchedule.monthIndex, oldIndex)
			assert.Equal(t, dummySchedule.months, values)
			assert.Equal(t, 12, maxValue)
			return dummyNewMonth, dummyNewMonthIndex, true
		}
		assert.Equal(t, dummySchedule.yearIndex, oldIndex)
		assert.Equal(t, dummySchedule.years, values)
		assert.Equal(t, 9999, maxValue)
		return dummyNewYear, dummyNewYearIndex, rand.Intn(100) > 50
	}
	getDaysOfMonthFuncExpected = 1
	getDaysOfMonthFunc = func(year, month int) int {
		getDaysOfMonthFuncCalled++
		assert.Equal(t, dummyOldYear, year)
		assert.Equal(t, dummyOldMonth, month)
		return dummyDaysOfMonth
	}

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

	// verify
	verifyAll(t)
}

func TestNextSchedule_ExceededTillTime(t *testing.T) {
	// arrange
	var dummyCurrentLocalTime = time.Now()
	var dummyTill = dummyCurrentLocalTime.Add(-1 * time.Second)

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyCurrentLocalTime
	}

	// SUT
	var sut = &schedule{
		till: &dummyTill,
	}

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestNextSchedule_NextTimeInPast(t *testing.T) {
	// arrange
	var dummyCurrentLocalTime = time.Now()
	var dummyTill = dummyCurrentLocalTime.Add(1 * time.Second)
	var dummySchedule = &schedule{
		till: &dummyTill,
	}
	var dummyTimeNext = dummyCurrentLocalTime.Add(-1 * time.Second)

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyCurrentLocalTime
	}
	constructTimeByScheduleFuncExpected = 1
	constructTimeByScheduleFunc = func(schedule *schedule) time.Time {
		constructTimeByScheduleFuncCalled++
		assert.Equal(t, dummySchedule, schedule)
		return dummyTimeNext
	}

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Nil(t, result)

	// verify
	verifyAll(t)
}

func TestNextSchedule_HappyPath(t *testing.T) {
	// arrange
	var dummyCurrentLocalTime = time.Now()
	var dummyTill = dummyCurrentLocalTime.Add(2 * time.Second)
	var dummySchedule = &schedule{
		till: &dummyTill,
	}
	var dummyTimeNext = dummyCurrentLocalTime.Add(1 * time.Second)

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyCurrentLocalTime
	}
	constructTimeByScheduleFuncExpected = 1
	constructTimeByScheduleFunc = func(schedule *schedule) time.Time {
		constructTimeByScheduleFuncCalled++
		assert.Equal(t, dummySchedule, schedule)
		return dummyTimeNext
	}
	updateScheduleIndexFuncExpected = 1
	updateScheduleIndexFunc = func(schedule *schedule) {
		updateScheduleIndexFuncCalled++
		assert.Equal(t, dummySchedule, schedule)
	}

	// SUT
	var sut = dummySchedule

	// act
	var result = sut.NextSchedule()

	// assert
	assert.Equal(t, dummyTimeNext, *result)

	// verify
	verifyAll(t)
}
