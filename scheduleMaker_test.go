package jobrunner

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhongjie-cai/gomocker/v2"
)

func TestNewScheduleMaker(t *testing.T) {
	// SUT
	var sut = NewScheduleMaker()

	// act
	var result, ok = sut.(*scheduleMaker)

	// assert
	assert.True(t, ok)
	assert.NotNil(t, result)
	assert.Empty(t, result.seconds)
	assert.Empty(t, result.minutes)
	assert.Empty(t, result.hours)
	assert.Empty(t, result.weekdays)
	assert.Empty(t, result.days)
	assert.Empty(t, result.months)
	assert.Empty(t, result.years)
	assert.Nil(t, result.from)
	assert.Nil(t, result.till)
}

func TestGenerateFlagsData_EmptyValues(t *testing.T) {
	// arrange
	var dummyData = []bool{
		rand.IntN(100) > 50,
		rand.IntN(100) > 50,
		rand.IntN(100) > 50,
	}
	var dummyTotal = rand.IntN(5) + 5

	// SUT + act
	var result = generateFlagsData(
		dummyData,
		dummyTotal,
	)

	// assert
	assert.Equal(t, dummyTotal, len(result))
	for _, value := range result {
		assert.True(t, value)
	}
}

func TestGenerateFlagsData_ValidValues(t *testing.T) {
	// arrange
	var dummyData = []bool{
		rand.IntN(100) > 50,
		rand.IntN(100) > 50,
		rand.IntN(100) > 50,
	}
	var dummyTotal = 5
	var dummyValue1 = 0
	var dummyValue2 = 2
	var dummyValue3 = 4
	var dummyValue4 = 6

	// SUT + act
	var result = generateFlagsData(
		dummyData,
		dummyTotal,
		dummyValue1,
		dummyValue2,
		dummyValue3,
		dummyValue4,
	)

	// assert
	assert.Len(t, result, 5)
	assert.True(t, result[0])
	assert.True(t, result[1])
	assert.True(t, result[2])
	assert.False(t, result[3])
	assert.True(t, result[4])
}

func TestScheduleMaker_OnSeconds(t *testing.T) {
	// arrange
	var dummyOldSeconds = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		seconds: dummyOldSeconds,
	}
	var dummySecond1 = rand.IntN(60)
	var dummySecond2 = rand.IntN(60)
	var dummySecond3 = rand.IntN(60)
	var dummyNewSeconds = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldSeconds, 60, dummySecond1, dummySecond2, dummySecond3).Returns(dummyNewSeconds).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.OnSeconds(
		dummySecond1,
		dummySecond2,
		dummySecond3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewSeconds, dummyScheduleMaker.seconds)
}

func TestScheduleMaker_OnMinutes(t *testing.T) {
	// arrange
	var dummyOldMinutes = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		minutes: dummyOldMinutes,
	}
	var dummyMinute1 = rand.IntN(60)
	var dummyMinute2 = rand.IntN(60)
	var dummyMinute3 = rand.IntN(60)
	var dummyNewMinutes = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldMinutes, 60, dummyMinute1, dummyMinute2, dummyMinute3).Returns(dummyNewMinutes).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.OnMinutes(
		dummyMinute1,
		dummyMinute2,
		dummyMinute3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewMinutes, dummyScheduleMaker.minutes)
}

func TestScheduleMaker_AtHours(t *testing.T) {
	// arrange
	var dummyOldHours = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		hours: dummyOldHours,
	}
	var dummyHour1 = rand.IntN(24)
	var dummyHour2 = rand.IntN(24)
	var dummyHour3 = rand.IntN(24)
	var dummyNewHours = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldHours, 24, dummyHour1, dummyHour2, dummyHour3).Returns(dummyNewHours).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.AtHours(
		dummyHour1,
		dummyHour2,
		dummyHour3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewHours, dummyScheduleMaker.hours)
}

func TestScheduleMaker_OnWeekdays(t *testing.T) {
	// arrange
	var dummyOldWeekdays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		weekdays: dummyOldWeekdays,
	}
	var dummyWeekday1 = time.Weekday(rand.IntN(7))
	var dummyWeekday2 = time.Weekday(rand.IntN(7))
	var dummyWeekday3 = time.Weekday(rand.IntN(7))
	var dummyNewWeekdays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldWeekdays, 7, int(dummyWeekday1), int(dummyWeekday2), int(dummyWeekday3)).Returns(dummyNewWeekdays).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.OnWeekdays(
		dummyWeekday1,
		dummyWeekday2,
		dummyWeekday3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewWeekdays, dummyScheduleMaker.weekdays)
}

func TestScheduleMaker_OnDays(t *testing.T) {
	// arrange
	var dummyOldDays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		days: dummyOldDays,
	}
	var dummyDay1 = 1 + rand.IntN(31)
	var dummyDay2 = 1 + rand.IntN(31)
	var dummyDay3 = 1 + rand.IntN(31)
	var dummyNewDays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldDays, 31, dummyDay1-1, dummyDay2-1, dummyDay3-1).Returns(dummyNewDays).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.OnDays(
		dummyDay1,
		dummyDay2,
		dummyDay3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewDays, dummyScheduleMaker.days)
}

func TestScheduleMaker_InMonths(t *testing.T) {
	// arrange
	var dummyOldMonths = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		months: dummyOldMonths,
	}
	var dummyMonth1 = time.Month(rand.IntN(12))
	var dummyMonth2 = time.Month(rand.IntN(12))
	var dummyMonth3 = time.Month(rand.IntN(12))
	var dummyNewMonths = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(generateFlagsData).Expects(dummyOldMonths, 12, int(dummyMonth1)-1, int(dummyMonth2)-1, int(dummyMonth3)-1).Returns(dummyNewMonths).Once()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.InMonths(
		dummyMonth1,
		dummyMonth2,
		dummyMonth3,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyNewMonths, dummyScheduleMaker.months)
}

func TestScheduleMaker_InYears_EmptyList(t *testing.T) {
	// arrange
	var dummyOldYears = map[int]bool{
		rand.IntN(100): rand.IntN(100) > 50,
		rand.IntN(100): rand.IntN(100) > 50,
		rand.IntN(100): rand.IntN(100) > 50,
	}
	var dummyScheduleMaker = &scheduleMaker{
		years: dummyOldYears,
	}

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.InYears()

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyOldYears, dummyScheduleMaker.years)
}

func TestScheduleMaker_InYears_ValidList(t *testing.T) {
	// arrange
	var dummyYear1 = rand.IntN(100)
	var dummyYear2 = rand.IntN(100)
	var dummyYear3 = rand.IntN(100)
	var dummyOldYears = map[int]bool{
		dummyYear1: rand.IntN(100) > 50,
		dummyYear2: rand.IntN(100) > 50,
		dummyYear3: rand.IntN(100) > 50,
	}
	var dummyScheduleMaker = &scheduleMaker{
		years: dummyOldYears,
	}
	var dummyYear4 = 100 + rand.IntN(100)
	var dummyYear5 = 100 + rand.IntN(100)
	var dummyYear6 = 100 + rand.IntN(100)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.InYears(
		dummyYear4,
		dummyYear5,
		dummyYear6,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyOldYears[dummyYear1], dummyScheduleMaker.years[dummyYear1])
	assert.Equal(t, dummyOldYears[dummyYear2], dummyScheduleMaker.years[dummyYear2])
	assert.Equal(t, dummyOldYears[dummyYear3], dummyScheduleMaker.years[dummyYear3])
	assert.True(t, dummyScheduleMaker.years[dummyYear4])
	assert.True(t, dummyScheduleMaker.years[dummyYear5])
	assert.True(t, dummyScheduleMaker.years[dummyYear6])
}

func TestScheduleMaker_From(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyStart = time.Now()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.From(
		dummyStart,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyStart, *dummyScheduleMaker.from)
}

func TestScheduleMaker_Till(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyEnd = time.Now()

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Till(
		dummyEnd,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyEnd, *dummyScheduleMaker.till)
}

func TestScheduleMaker_SkipOverdue(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.SkipOverdue()

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.True(t, dummyScheduleMaker.skipOverdue)
}

func TestScheduleMaker_Timezone_NilTimezone(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyTimezone *time.Location

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Timezone(
		dummyTimezone,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, time.Local, dummyScheduleMaker.timezone)
}

func TestScheduleMaker_Timezone_ValidTimezone(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyTimezone, _ = time.LoadLocation("Asia/Shanghai")

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Timezone(
		dummyTimezone,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyTimezone, dummyScheduleMaker.timezone)
}

func TestConstructValueSlice_EmptyValues(t *testing.T) {
	// arrange
	var dummyValues = []bool{}
	var dummyTotal = rand.IntN(10) + 5

	// SUT + act
	var result = constructValueSlice(
		dummyValues,
		dummyTotal,
	)

	// assert
	assert.Equal(t, dummyTotal, len(result))
	for i := 0; i < len(result); i++ {
		assert.Equal(t, i, result[i])
	}
}

func TestConstructValueSlice_ValidValues(t *testing.T) {
	// arrange
	var dummyValues = []bool{
		true,
		false,
		true,
		false,
		true,
		false,
	}
	var dummyTotal = len(dummyValues)

	// SUT + act
	var result = constructValueSlice(
		dummyValues,
		dummyTotal,
	)

	// assert
	assert.Len(t, result, 3)
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 4, result[2])
}

func TestConstructWeekdayMap_EmptyWeekdays(t *testing.T) {
	// arrange
	var dummyWeekdays = []bool{}

	// SUT + act
	var result = constructWeekdayMap(
		dummyWeekdays,
	)

	// assert
	assert.Len(t, result, 7)
	assert.True(t, result[time.Monday])
	assert.True(t, result[time.Tuesday])
	assert.True(t, result[time.Wednesday])
	assert.True(t, result[time.Thursday])
	assert.True(t, result[time.Friday])
	assert.True(t, result[time.Saturday])
	assert.True(t, result[time.Sunday])
}

func TestConstructWeekdayMap_ValidWeekdays(t *testing.T) {
	// arrange
	var dummyWeekdays = []bool{
		false,
		true,
		false,
		true,
		false,
		true,
		false,
	}

	// SUT + act
	var result = constructWeekdayMap(
		dummyWeekdays,
	)

	// assert
	assert.Len(t, result, 3)
	assert.True(t, result[time.Monday])
	assert.True(t, result[time.Wednesday])
	assert.True(t, result[time.Friday])
}

func TestConstructYearSlice_EmptyYears(t *testing.T) {
	// arrange
	var dummyYears = map[int]bool{}
	var dummyTime = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Now).Expects().Returns(dummyTime).Once()

	// SUT + act
	var result = constructYearSlice(
		dummyYears,
	)

	// assert
	assert.Len(t, result, 100)
	for year := 0; year < 100; year++ {
		assert.Equal(t, dummyTime.Year()+year, result[year])
	}
}

func TestConstructYearSlice_ValidYears(t *testing.T) {
	// arrange
	var dummyYears = map[int]bool{
		2020: true,
		2021: false,
		2022: true,
		2023: false,
		2024: true,
	}

	// SUT + act
	var result = constructYearSlice(
		dummyYears,
	)

	// assert
	assert.Len(t, result, 3)
	assert.Equal(t, 2020, result[0])
	assert.Equal(t, 2022, result[1])
	assert.Equal(t, 2024, result[2])
}

func TestConstructScheduleTemplate(t *testing.T) {
	// arrange
	var dummyMakerSeconds = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerMinutes = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerHours = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerWeekdays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerDays = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerMonths = []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50}
	var dummyMakerYears = map[int]bool{
		rand.IntN(100): rand.IntN(100) > 50,
		rand.IntN(100): rand.IntN(100) > 50,
		rand.IntN(100): rand.IntN(100) > 50,
	}
	var dummyMakerTill = time.Now()
	var dummySkipOverdue = rand.IntN(100) > 50
	var dummyScheduleMaker = &scheduleMaker{
		seconds:     dummyMakerSeconds,
		minutes:     dummyMakerMinutes,
		hours:       dummyMakerHours,
		weekdays:    dummyMakerWeekdays,
		days:        dummyMakerDays,
		months:      dummyMakerMonths,
		years:       dummyMakerYears,
		till:        &dummyMakerTill,
		skipOverdue: dummySkipOverdue,
	}
	var dummyScheduleSeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleHours = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleDays = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleMonths = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleYears = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyScheduleWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(constructValueSlice).Expects(dummyMakerSeconds, 60).Returns(dummyScheduleSeconds).Once()
	m.Mock(constructValueSlice).Expects(dummyMakerMinutes, 60).Returns(dummyScheduleMinutes).Once()
	m.Mock(constructValueSlice).Expects(dummyMakerHours, 24).Returns(dummyScheduleHours).Once()
	m.Mock(constructValueSlice).Expects(dummyMakerDays, 31).Returns(dummyScheduleDays).Once()
	m.Mock(constructValueSlice).Expects(dummyMakerMonths, 12).Returns(dummyScheduleMonths).Once()
	m.Mock(constructYearSlice).Expects(dummyMakerYears).Returns(dummyScheduleYears).Once()
	m.Mock(constructWeekdayMap).Expects(dummyMakerWeekdays).Returns(dummyScheduleWeekdays).Once()

	// SUT + act
	var result = constructScheduleTemplate(
		dummyScheduleMaker,
	)

	// assert
	assert.Equal(t, 0, result.secondIndex)
	assert.Equal(t, dummyScheduleSeconds[0], result.second)
	assert.Equal(t, dummyScheduleSeconds, result.seconds)
	assert.Equal(t, 0, result.minuteIndex)
	assert.Equal(t, dummyScheduleMinutes[0], result.minute)
	assert.Equal(t, dummyScheduleMinutes, result.minutes)
	assert.Equal(t, 0, result.hourIndex)
	assert.Equal(t, dummyScheduleHours[0], result.hour)
	assert.Equal(t, dummyScheduleHours, result.hours)
	assert.Equal(t, 0, result.dayIndex)
	assert.Equal(t, dummyScheduleDays[0], result.day)
	assert.Equal(t, dummyScheduleDays, result.days)
	assert.Equal(t, 0, result.monthIndex)
	assert.Equal(t, dummyScheduleMonths[0], result.month)
	assert.Equal(t, dummyScheduleMonths, result.months)
	assert.Equal(t, 0, result.yearIndex)
	assert.Equal(t, dummyScheduleYears[0], result.year)
	assert.Equal(t, dummyScheduleYears, result.years)
	assert.Equal(t, dummyScheduleWeekdays, result.weekdays)
	assert.Equal(t, &dummyMakerTill, result.till)
	assert.Equal(t, dummySkipOverdue, result.skipOverdue)
}

func TestFindValueMatch_EmptyValues(t *testing.T) {
	// arrange
	var dummyValue = rand.IntN(10)
	var dummyValues = []int{}

	// SUT + act
	var value, index, reset, overflow = findValueMatch(
		dummyValue,
		dummyValues,
	)

	// assert
	assert.Zero(t, value)
	assert.Zero(t, index)
	assert.False(t, reset)
	assert.False(t, overflow)
}

func TestFindValueMatch_ValidValues_NoMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.IntN(10)
	var dummyValues = []int{
		rand.IntN(10),
		rand.IntN(10) + 10,
		rand.IntN(10) + 20,
	}

	// SUT + act
	var value, index, reset, overflow = findValueMatch(
		dummyValue,
		dummyValues,
	)

	// assert
	assert.Equal(t, dummyValues[0], value)
	assert.Zero(t, index)
	assert.True(t, reset)
	assert.True(t, overflow)
}

func TestFindValueMatch_ValidValues_ExactMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.IntN(10)
	var dummyValues = []int{
		rand.IntN(10),
		rand.IntN(10) + 10,
		rand.IntN(10) + 20,
		dummyValue,
		rand.IntN(10) + 40,
		rand.IntN(10) + 50,
	}

	// SUT + act
	var value, index, reset, overflow = findValueMatch(
		dummyValue,
		dummyValues,
	)

	// assert
	assert.Equal(t, dummyValues[3], value)
	assert.Equal(t, 3, index)
	assert.False(t, reset)
	assert.False(t, overflow)
}

func TestFindValueMatch_ValidValues_SimMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.IntN(10)
	var dummyValues = []int{
		rand.IntN(10),
		rand.IntN(10) + 10,
		rand.IntN(10) + 20,
		rand.IntN(10) + 40,
		rand.IntN(10) + 50,
	}

	// SUT + act
	var value, index, reset, overflow = findValueMatch(
		dummyValue,
		dummyValues,
	)

	// assert
	assert.Equal(t, dummyValues[3], value)
	assert.Equal(t, 3, index)
	assert.True(t, reset)
	assert.False(t, overflow)
}

func TestIsWeekdayMatch_EmptyWeekdays(t *testing.T) {
	// arrange
	var dummyYear = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyWeekdays = map[time.Weekday]bool{}

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.True(t, result)
}

func TestIsWeekdayMatch_ValidWeekdays_NotFound(t *testing.T) {
	// arrange
	var dummyYear = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday + 1: true,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		0, 0, 0, 0, time.Local).Returns(dummyDate).Once()

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.False(t, result)
}

func TestIsWeekdayMatch_ValidWeekdays_NotValid(t *testing.T) {
	// arrange
	var dummyYear = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday:     false,
		dummyWeekday + 1: true,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		0, 0, 0, 0, time.Local).Returns(dummyDate).Once()

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.False(t, result)
}

func TestIsWeekdayMatch_ValidWeekdays_FoundValid(t *testing.T) {
	// arrange
	var dummyYear = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday:     true,
		dummyWeekday + 1: true,
	}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		0, 0, 0, 0, time.Local).Returns(dummyDate).Once()

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.True(t, result)
}

func TestDetermineScheduleIndex_YearOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummySchedule = &schedule{
		seconds: dummySeconds,
		minutes: dummyMinutes,
		hours:   dummyHours,
		days:    dummyDays,
		months:  dummyMonths,
		years:   dummyYears,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyFormat = "Invalid schedule configuration: no valid next execution time available"
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, true).Once()
	m.Mock(fmt.Errorf).Expects(dummyFormat).Returns(dummyError).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyStart, start)
	assert.Equal(t, dummyError, err)
}

func TestDetermineScheduleIndex_YearIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, true, false).Once()
	m.Mock(time.Date).Expects(dummyYear, time.January, 1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_MonthOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, true).Once()
	m.Mock(time.Date).Expects(dummyStart.Year()+1, time.January, 1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_MonthIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, true, false).Once()
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), 1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_DayOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, true).Once()
	m.Mock(time.Date).Expects(dummyStart.Year(), dummyStart.Month()+1, 1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_DayIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, true, false).Once()
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_WeekdayMismatch(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(false).Once()
	m.Mock(time.Date).Expects(dummyStart.Year(), dummyStart.Month(), dummyStart.Day()+1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_HourOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, false, true).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()
	m.Mock(time.Date).Expects(dummyStart.Year(), dummyStart.Month(), dummyStart.Day()+1,
		0, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_HourIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, true, false).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		dummyHour, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_MinuteOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyMinute = rand.IntN(60)
	var dummyMinuteIndex = rand.IntN(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Minute(), dummyMinutes).Returns(dummyMinute, dummyMinuteIndex, false, true).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()
	m.Mock(time.Date).Expects(dummyStart.Year(), dummyStart.Month(), dummyStart.Day(),
		dummyStart.Hour()+1, 0, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_MinuteIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyMinute = rand.IntN(60)
	var dummyMinuteIndex = rand.IntN(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Minute(), dummyMinutes).Returns(dummyMinute, dummyMinuteIndex, true, false).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()
	m.Mock(time.Date).Expects(dummyYear, time.Month(dummyMonth+1), dummyDay+1,
		dummyHour, dummyMinute, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_SecondOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyMinute = rand.IntN(60)
	var dummyMinuteIndex = rand.IntN(60)
	var dummySecond = rand.IntN(60)
	var dummySecondIndex = rand.IntN(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Minute(), dummyMinutes).Returns(dummyMinute, dummyMinuteIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Second(), dummySeconds).Returns(dummySecond, dummySecondIndex, false, true).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()
	m.Mock(time.Date).Expects(dummyStart.Year(), dummyStart.Month(), dummyStart.Day(),
		dummyStart.Hour(), dummyStart.Minute()+1, 0, 0, dummyLocation).Returns(dummyTime).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)
}

func TestDetermineScheduleIndex_NoOverflow_NoIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyMinutes = []int{rand.IntN(60), rand.IntN(60), rand.IntN(60)}
	var dummyHours = []int{rand.IntN(24), rand.IntN(24), rand.IntN(24)}
	var dummyDays = []int{rand.IntN(31), rand.IntN(31), rand.IntN(31)}
	var dummyMonths = []int{rand.IntN(12), rand.IntN(12), rand.IntN(12)}
	var dummyYears = []int{rand.IntN(100), rand.IntN(100), rand.IntN(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
		time.Weekday(rand.IntN(7)): rand.IntN(100) > 50,
	}
	var dummySchedule = &schedule{
		seconds:  dummySeconds,
		minutes:  dummyMinutes,
		hours:    dummyHours,
		days:     dummyDays,
		months:   dummyMonths,
		years:    dummyYears,
		timezone: dummyLocation,
		weekdays: dummyWeekdays,
	}
	var dummyYear = rand.IntN(100)
	var dummyYearIndex = rand.IntN(100)
	var dummyMonth = rand.IntN(12)
	var dummyMonthIndex = rand.IntN(12)
	var dummyDay = rand.IntN(31)
	var dummyDayIndex = rand.IntN(31)
	var dummyHour = rand.IntN(24)
	var dummyHourIndex = rand.IntN(24)
	var dummyMinute = rand.IntN(60)
	var dummyMinuteIndex = rand.IntN(60)
	var dummySecond = rand.IntN(60)
	var dummySecondIndex = rand.IntN(60)

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(findValueMatch).Expects(dummyStart.Year(), dummyYears).Returns(dummyYear, dummyYearIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(int(dummyStart.Month())-1, dummyMonths).Returns(dummyMonth, dummyMonthIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Day()-1, dummyDays).Returns(dummyDay, dummyDayIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Hour(), dummyHours).Returns(dummyHour, dummyHourIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Minute(), dummyMinutes).Returns(dummyMinute, dummyMinuteIndex, false, false).Once()
	m.Mock(findValueMatch).Expects(dummyStart.Second(), dummySeconds).Returns(dummySecond, dummySecondIndex, false, false).Once()
	m.Mock(isWeekdayMatch).Expects(dummyYear, dummyMonth, dummyDay, dummyWeekdays).Returns(true).Once()

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.True(t, completed)
	assert.Equal(t, dummyStart, start)
	assert.NoError(t, err)
}

func TestInitialiseSchedule_Error(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySchedule = &schedule{second: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(determineScheduleIndex).Expects(dummyStart, dummySchedule).Returns(false, dummyStart, dummyError).Once()

	// SUT + act
	var err = initialiseSchedule(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyError, err)
}

func TestInitialiseSchedule_Success(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySchedule = &schedule{second: rand.Int()}

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(determineScheduleIndex).Expects(dummyStart, dummySchedule).Returns(true, dummyStart, nil).Once()

	// SUT + act
	var err = initialiseSchedule(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.NoError(t, err)
}

func TestScheduleMaker_Schedule_WithoutFrom(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{
		seconds: []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50},
	}
	var dummyTimeNow = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	var dummySchedule = &schedule{year: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Now).Expects().Returns(dummyTimeNow).Once()
	m.Mock(constructScheduleTemplate).Expects(dummyScheduleMaker).Returns(dummySchedule).Once()
	m.Mock(initialiseSchedule).Expects(dummyTimeNow, dummySchedule).Returns(dummyError).Once()

	// SUT
	var sut, err = dummyScheduleMaker.Schedule()

	// act
	var result, ok = sut.(*schedule)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummySchedule, result)
	assert.Equal(t, dummyError, err)
}

func TestScheduleMaker_Schedule_WithFrom(t *testing.T) {
	// arrange
	var dummyFrom = time.Now().Add(100 * time.Second)
	var dummyScheduleMaker = &scheduleMaker{
		seconds: []bool{rand.IntN(100) > 50, rand.IntN(100) > 50, rand.IntN(100) > 50},
		from:    &dummyFrom,
	}
	var dummyTimeNow = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	var dummySchedule = &schedule{year: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	var m = gomocker.NewMocker(t)

	// expect
	m.Mock(time.Now).Expects().Returns(dummyTimeNow).Once()
	m.Mock(constructScheduleTemplate).Expects(dummyScheduleMaker).Returns(dummySchedule).Once()
	m.Mock(initialiseSchedule).Expects(dummyFrom, dummySchedule).Returns(dummyError).Once()

	// SUT
	var sut, err = dummyScheduleMaker.Schedule()

	// act
	var result, ok = sut.(*schedule)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummySchedule, result)
	assert.Equal(t, dummyError, err)
}

func TestScheduleMaker_Integration(t *testing.T) {
	// setup
	const layout = "2006-01-02 15:04:05"
	var testData = map[string]string{
		"2001-02-03 04:05:06": "2001-04-01 00:00:00",
		"2001-04-15 03:06:09": "2001-04-15 04:00:00",
		"2001-04-15 06:30:55": "2001-04-15 06:35:00",
		"2001-10-15 22:55:35": "2002-01-01 00:00:00",
		"2001-10-16 22:55:35": "2002-01-01 00:00:00",
		"2001-04-16 22:55:35": "2001-07-01 00:00:00",
		"2001-12-03 04:05:06": "2002-01-01 00:00:00",
		"2001-10-15 22:55:31": "2002-01-01 00:00:00",
		"2001-10-15 22:55:29": "2001-10-15 22:55:30",
		"2001-10-15 23:55:30": "2002-01-01 00:00:00",
	}

	// mock
	var m = gomocker.NewMocker(t)

	for given, expect := range testData {
		// arrange
		var timeStart, _ = time.Parse(layout, given)
		var timeExpect, _ = time.Parse(layout, expect)

		// stub
		m.Stub(time.Now).Returns(timeStart).Once()

		// SUT
		var scheduleMaker, _ = NewScheduleMaker().OnSeconds(
			0, 30,
		).OnMinutes(
			0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55,
		).AtHours(
			0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		).OnDays(
			1, 15,
		).InMonths(
			time.January, time.April, time.July, time.October,
		).InYears(
			2000, 2001, 2002,
		).Schedule()

		// act
		var timeNext = scheduleMaker.NextSchedule()

		// assert
		assert.NotNil(t, timeNext)
		assert.Equal(t, timeExpect.Year(), timeNext.Year())
		assert.Equal(t, timeExpect.Month(), timeNext.Month())
		assert.Equal(t, timeExpect.Day(), timeNext.Day())
		assert.Equal(t, timeExpect.Hour(), timeNext.Hour())
		assert.Equal(t, timeExpect.Minute(), timeNext.Minute())
		assert.Equal(t, timeExpect.Second(), timeNext.Second())
	}
}
