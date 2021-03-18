package jobrunner

import (
	"errors"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewScheduleMaker(t *testing.T) {
	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestGenerateFlagsData_EmptyValues(t *testing.T) {
	// arrange
	var dummyData = []bool{
		rand.Intn(100) > 50,
		rand.Intn(100) > 50,
		rand.Intn(100) > 50,
	}
	var dummyTotal = rand.Intn(5) + 5

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestGenerateFlagsData_ValidValues(t *testing.T) {
	// arrange
	var dummyData = []bool{
		rand.Intn(100) > 50,
		rand.Intn(100) > 50,
		rand.Intn(100) > 50,
	}
	var dummyTotal = 5
	var dummyValue1 = 0
	var dummyValue2 = 2
	var dummyValue3 = 4
	var dummyValue4 = 6

	// mock
	createMock(t)

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
	assert.Equal(t, 5, len(result))
	assert.True(t, result[0])
	assert.True(t, result[1])
	assert.True(t, result[2])
	assert.False(t, result[3])
	assert.True(t, result[4])

	// verify
	verifyAll(t)
}

func TestScheduleMaker_OnSeconds(t *testing.T) {
	// arrange
	var dummyOldSeconds = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		seconds: dummyOldSeconds,
	}
	var dummySecond1 = rand.Intn(60)
	var dummySecond2 = rand.Intn(60)
	var dummySecond3 = rand.Intn(60)
	var dummyNewSeconds = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldSeconds, data)
		assert.Equal(t, 60, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, dummySecond1, values[0])
		assert.Equal(t, dummySecond2, values[1])
		assert.Equal(t, dummySecond3, values[2])
		return dummyNewSeconds
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_OnMinutes(t *testing.T) {
	// arrange
	var dummyOldMinutes = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		minutes: dummyOldMinutes,
	}
	var dummyMinute1 = rand.Intn(60)
	var dummyMinute2 = rand.Intn(60)
	var dummyMinute3 = rand.Intn(60)
	var dummyNewMinutes = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldMinutes, data)
		assert.Equal(t, 60, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, dummyMinute1, values[0])
		assert.Equal(t, dummyMinute2, values[1])
		assert.Equal(t, dummyMinute3, values[2])
		return dummyNewMinutes
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_AtHours(t *testing.T) {
	// arrange
	var dummyOldHours = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		hours: dummyOldHours,
	}
	var dummyHour1 = rand.Intn(24)
	var dummyHour2 = rand.Intn(24)
	var dummyHour3 = rand.Intn(24)
	var dummyNewHours = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldHours, data)
		assert.Equal(t, 24, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, dummyHour1, values[0])
		assert.Equal(t, dummyHour2, values[1])
		assert.Equal(t, dummyHour3, values[2])
		return dummyNewHours
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_OnWeekdays(t *testing.T) {
	// arrange
	var dummyOldWeekdays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		weekdays: dummyOldWeekdays,
	}
	var dummyWeekday1 = time.Weekday(rand.Intn(7))
	var dummyWeekday2 = time.Weekday(rand.Intn(7))
	var dummyWeekday3 = time.Weekday(rand.Intn(7))
	var dummyNewWeekdays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldWeekdays, data)
		assert.Equal(t, 7, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, int(dummyWeekday1), values[0])
		assert.Equal(t, int(dummyWeekday2), values[1])
		assert.Equal(t, int(dummyWeekday3), values[2])
		return dummyNewWeekdays
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_OnDays(t *testing.T) {
	// arrange
	var dummyOldDays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		days: dummyOldDays,
	}
	var dummyDay1 = 1 + rand.Intn(31)
	var dummyDay2 = 1 + rand.Intn(31)
	var dummyDay3 = 1 + rand.Intn(31)
	var dummyNewDays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldDays, data)
		assert.Equal(t, 31, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, dummyDay1-1, values[0])
		assert.Equal(t, dummyDay2-1, values[1])
		assert.Equal(t, dummyDay3-1, values[2])
		return dummyNewDays
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_InMonths(t *testing.T) {
	// arrange
	var dummyOldMonths = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyScheduleMaker = &scheduleMaker{
		months: dummyOldMonths,
	}
	var dummyMonth1 = time.Month(rand.Intn(12))
	var dummyMonth2 = time.Month(rand.Intn(12))
	var dummyMonth3 = time.Month(rand.Intn(12))
	var dummyNewMonths = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}

	// mock
	createMock(t)

	// expect
	generateFlagsDataFuncExpected = 1
	generateFlagsDataFunc = func(data []bool, total int, values ...int) []bool {
		generateFlagsDataFuncCalled++
		assert.Equal(t, dummyOldMonths, data)
		assert.Equal(t, 12, total)
		assert.Equal(t, 3, len(values))
		assert.Equal(t, int(dummyMonth1)-1, values[0])
		assert.Equal(t, int(dummyMonth2)-1, values[1])
		assert.Equal(t, int(dummyMonth3)-1, values[2])
		return dummyNewMonths
	}

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_InYears_EmptyList(t *testing.T) {
	// arrange
	var dummyOldYears = map[int]bool{
		rand.Intn(100): rand.Intn(100) > 50,
		rand.Intn(100): rand.Intn(100) > 50,
		rand.Intn(100): rand.Intn(100) > 50,
	}
	var dummyScheduleMaker = &scheduleMaker{
		years: dummyOldYears,
	}

	// mock
	createMock(t)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.InYears()

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyOldYears, dummyScheduleMaker.years)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_InYears_ValidList(t *testing.T) {
	// arrange
	var dummyYear1 = rand.Intn(100)
	var dummyYear2 = rand.Intn(100)
	var dummyYear3 = rand.Intn(100)
	var dummyOldYears = map[int]bool{
		dummyYear1: rand.Intn(100) > 50,
		dummyYear2: rand.Intn(100) > 50,
		dummyYear3: rand.Intn(100) > 50,
	}
	var dummyScheduleMaker = &scheduleMaker{
		years: dummyOldYears,
	}
	var dummyYear4 = 100 + rand.Intn(100)
	var dummyYear5 = 100 + rand.Intn(100)
	var dummyYear6 = 100 + rand.Intn(100)

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestScheduleMaker_From(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyStart = time.Now()

	// mock
	createMock(t)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.From(
		dummyStart,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyStart, *dummyScheduleMaker.from)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_Till(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyEnd = time.Now()

	// mock
	createMock(t)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Till(
		dummyEnd,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyEnd, *dummyScheduleMaker.till)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_Timezone_NilTimezone(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyTimezone *time.Location

	// mock
	createMock(t)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Timezone(
		dummyTimezone,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, time.Local, dummyScheduleMaker.timezone)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_Timezone_ValidTimezone(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{}
	var dummyTimezone, _ = time.LoadLocation("Asia/Shanghai")

	// mock
	createMock(t)

	// SUT
	var sut = dummyScheduleMaker

	// act
	var result = sut.Timezone(
		dummyTimezone,
	)

	// assert
	assert.Equal(t, dummyScheduleMaker, result)
	assert.Equal(t, dummyTimezone, dummyScheduleMaker.timezone)

	// verify
	verifyAll(t)
}

func TestConstructValueSlice_EmptyValues(t *testing.T) {
	// arrange
	var dummyValues = []bool{}
	var dummyTotal = rand.Intn(10) + 5

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
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

	// mock
	createMock(t)

	// SUT + act
	var result = constructValueSlice(
		dummyValues,
		dummyTotal,
	)

	// assert
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 4, result[2])

	// verify
	verifyAll(t)
}

func TestConstructWeekdayMap_EmptyWeekdays(t *testing.T) {
	// arrange
	var dummyWeekdays = []bool{}

	// mock
	createMock(t)

	// SUT + act
	var result = constructWeekdayMap(
		dummyWeekdays,
	)

	// assert
	assert.Equal(t, 7, len(result))
	assert.True(t, result[time.Monday])
	assert.True(t, result[time.Tuesday])
	assert.True(t, result[time.Wednesday])
	assert.True(t, result[time.Thursday])
	assert.True(t, result[time.Friday])
	assert.True(t, result[time.Saturday])
	assert.True(t, result[time.Sunday])

	// verify
	verifyAll(t)
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

	// mock
	createMock(t)

	// SUT + act
	var result = constructWeekdayMap(
		dummyWeekdays,
	)

	// assert
	assert.Equal(t, 3, len(result))
	assert.True(t, result[time.Monday])
	assert.True(t, result[time.Wednesday])
	assert.True(t, result[time.Friday])

	// verify
	verifyAll(t)
}

func TestConstructYearSlice_EmptyYears(t *testing.T) {
	// arrange
	var dummyYears = map[int]bool{}
	var dummyTime = time.Now()

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyTime
	}

	// SUT + act
	var result = constructYearSlice(
		dummyYears,
	)

	// assert
	assert.Equal(t, 100, len(result))
	for year := 0; year < 100; year++ {
		assert.Equal(t, dummyTime.Year()+year, result[year])
	}

	// verify
	verifyAll(t)
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

	// mock
	createMock(t)

	// expect
	sortIntsExpected = 1
	sortInts = func(a []int) {
		sortIntsCalled++
		sort.Ints(a)
	}

	// SUT + act
	var result = constructYearSlice(
		dummyYears,
	)

	// assert
	assert.Equal(t, 3, len(result))
	assert.Equal(t, 2020, result[0])
	assert.Equal(t, 2022, result[1])
	assert.Equal(t, 2024, result[2])

	// verify
	verifyAll(t)
}

func TestConstructScheduleTemplate(t *testing.T) {
	// arrange
	var dummyMakerSeconds = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerMinutes = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerHours = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerWeekdays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerDays = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerMonths = []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50}
	var dummyMakerYears = map[int]bool{
		rand.Intn(100): rand.Intn(100) > 50,
		rand.Intn(100): rand.Intn(100) > 50,
		rand.Intn(100): rand.Intn(100) > 50,
	}
	var dummyMakerTill = time.Now()
	var dummyScheduleMaker = &scheduleMaker{
		seconds:  dummyMakerSeconds,
		minutes:  dummyMakerMinutes,
		hours:    dummyMakerHours,
		weekdays: dummyMakerWeekdays,
		days:     dummyMakerDays,
		months:   dummyMakerMonths,
		years:    dummyMakerYears,
		till:     &dummyMakerTill,
	}
	var dummyScheduleSeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleHours = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleDays = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleMonths = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleYears = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyScheduleWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
	}

	// mock
	createMock(t)

	// expect
	constructValueSliceFuncExpected = 5
	constructValueSliceFunc = func(values []bool, total int) []int {
		constructValueSliceFuncCalled++
		if constructValueSliceFuncCalled == 1 {
			assert.Equal(t, dummyMakerSeconds, values)
			assert.Equal(t, 60, total)
			return dummyScheduleSeconds
		} else if constructValueSliceFuncCalled == 2 {
			assert.Equal(t, dummyMakerMinutes, values)
			assert.Equal(t, 60, total)
			return dummyScheduleMinutes
		} else if constructValueSliceFuncCalled == 3 {
			assert.Equal(t, dummyMakerHours, values)
			assert.Equal(t, 24, total)
			return dummyScheduleHours
		} else if constructValueSliceFuncCalled == 4 {
			assert.Equal(t, dummyMakerDays, values)
			assert.Equal(t, 31, total)
			return dummyScheduleDays
		} else if constructValueSliceFuncCalled == 5 {
			assert.Equal(t, dummyMakerMonths, values)
			assert.Equal(t, 12, total)
			return dummyScheduleMonths
		}
		return nil
	}
	constructYearSliceFuncExpected = 1
	constructYearSliceFunc = func(years map[int]bool) []int {
		constructYearSliceFuncCalled++
		assert.Equal(t, dummyMakerYears, years)
		return dummyScheduleYears
	}
	constructWeekdayMapFuncExpected = 1
	constructWeekdayMapFunc = func(weekdays []bool) map[time.Weekday]bool {
		constructWeekdayMapFuncCalled++
		assert.Equal(t, dummyMakerWeekdays, weekdays)
		return dummyScheduleWeekdays
	}

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

	// verify
	verifyAll(t)
}

func TestFindValueMatch_EmptyValues(t *testing.T) {
	// arrange
	var dummyValue = rand.Intn(10)
	var dummyValues = []int{}

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestFindValueMatch_ValidValues_NoMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.Intn(10)
	var dummyValues = []int{
		rand.Intn(10),
		rand.Intn(10) + 10,
		rand.Intn(10) + 20,
	}

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestFindValueMatch_ValidValues_ExactMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.Intn(10)
	var dummyValues = []int{
		rand.Intn(10),
		rand.Intn(10) + 10,
		rand.Intn(10) + 20,
		dummyValue,
		rand.Intn(10) + 40,
		rand.Intn(10) + 50,
	}

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestFindValueMatch_ValidValues_SimMatch(t *testing.T) {
	// arrange
	var dummyValue = 30 + rand.Intn(10)
	var dummyValues = []int{
		rand.Intn(10),
		rand.Intn(10) + 10,
		rand.Intn(10) + 20,
		rand.Intn(10) + 40,
		rand.Intn(10) + 50,
	}

	// mock
	createMock(t)

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

	// verify
	verifyAll(t)
}

func TestIsWeekdayMatch_EmptyWeekdays(t *testing.T) {
	// arrange
	var dummyYear = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyWeekdays = map[time.Weekday]bool{}

	// mock
	createMock(t)

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestIsWeekdayMatch_ValidWeekdays_NotFound(t *testing.T) {
	// arrange
	var dummyYear = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday + 1: true,
	}

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return dummyDate
	}

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsWeekdayMatch_ValidWeekdays_NotValid(t *testing.T) {
	// arrange
	var dummyYear = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday:     false,
		dummyWeekday + 1: true,
	}

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return dummyDate
	}

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.False(t, result)

	// verify
	verifyAll(t)
}

func TestIsWeekdayMatch_ValidWeekdays_FoundValid(t *testing.T) {
	// arrange
	var dummyYear = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDate = time.Now()
	var dummyWeekday = dummyDate.Weekday()
	var dummyWeekdays = map[time.Weekday]bool{
		dummyWeekday - 1: true,
		dummyWeekday:     true,
		dummyWeekday + 1: true,
	}

	// mock
	createMock(t)

	// expect
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, time.Local, loc)
		return dummyDate
	}

	// SUT + act
	var result = isWeekdayMatch(
		dummyYear,
		dummyMonth,
		dummyDay,
		dummyWeekdays,
	)

	// assert
	assert.True(t, result)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_YearOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummySchedule = &schedule{
		seconds: dummySeconds,
		minutes: dummyMinutes,
		hours:   dummyHours,
		days:    dummyDays,
		months:  dummyMonths,
		years:   dummyYears,
	}
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyFormat = "Invalid schedule configuration: no valid next execution time available"
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 1
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		assert.Equal(t, dummyStart.Year(), value)
		assert.Equal(t, dummyYears, values)
		return dummyYear, dummyYearIndex, false, true
	}
	fmtErrorfExpected = 1
	fmtErrorf = func(format string, a ...interface{}) error {
		fmtErrorfCalled++
		assert.Equal(t, dummyFormat, format)
		assert.Empty(t, a)
		return dummyError
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyStart, start)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_YearIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 1
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		assert.Equal(t, dummyStart.Year(), value)
		assert.Equal(t, dummyYears, values)
		return dummyYear, dummyYearIndex, true, false
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.January, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_MonthOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 2
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		}
		assert.Equal(t, int(dummyStart.Month())-1, value)
		assert.Equal(t, dummyMonths, values)
		return dummyMonth, dummyMonthIndex, false, true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year()+1, year)
		assert.Equal(t, time.January, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_MonthIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 2
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		}
		assert.Equal(t, int(dummyStart.Month())-1, value)
		assert.Equal(t, dummyMonths, values)
		return dummyMonth, dummyMonthIndex, true, false
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_DayOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 3
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		}
		assert.Equal(t, dummyStart.Day()-1, value)
		assert.Equal(t, dummyDays, values)
		return dummyDay, dummyDayIndex, false, true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year(), year)
		assert.Equal(t, dummyStart.Month()+1, month)
		assert.Equal(t, 1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_DayIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 3
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		}
		assert.Equal(t, dummyStart.Day()-1, value)
		assert.Equal(t, dummyDays, values)
		return dummyDay, dummyDayIndex, true, false
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_WeekdayMismatch(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 3
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		}
		assert.Equal(t, dummyStart.Day()-1, value)
		assert.Equal(t, dummyDays, values)
		return dummyDay, dummyDayIndex, false, false
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return false
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year(), year)
		assert.Equal(t, dummyStart.Month(), month)
		assert.Equal(t, dummyStart.Day()+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_HourOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 4
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		}
		assert.Equal(t, dummyStart.Hour(), value)
		assert.Equal(t, dummyHours, values)
		return dummyHour, dummyHourIndex, false, true
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year(), year)
		assert.Equal(t, dummyStart.Month(), month)
		assert.Equal(t, dummyStart.Day()+1, day)
		assert.Zero(t, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_HourIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 4
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		}
		assert.Equal(t, dummyStart.Hour(), value)
		assert.Equal(t, dummyHours, values)
		return dummyHour, dummyHourIndex, true, false
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Equal(t, dummyHour, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_MinuteOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyMinute = rand.Intn(60)
	var dummyMinuteIndex = rand.Intn(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 5
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		} else if findValueMatchFuncCalled == 4 {
			assert.Equal(t, dummyStart.Hour(), value)
			assert.Equal(t, dummyHours, values)
			return dummyHour, dummyHourIndex, false, false
		}
		assert.Equal(t, dummyStart.Minute(), value)
		assert.Equal(t, dummyMinutes, values)
		return dummyMinute, dummyMinuteIndex, false, true
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year(), year)
		assert.Equal(t, dummyStart.Month(), month)
		assert.Equal(t, dummyStart.Day(), day)
		assert.Equal(t, dummyStart.Hour()+1, hour)
		assert.Zero(t, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_MinuteIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyMinute = rand.Intn(60)
	var dummyMinuteIndex = rand.Intn(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 5
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		} else if findValueMatchFuncCalled == 4 {
			assert.Equal(t, dummyStart.Hour(), value)
			assert.Equal(t, dummyHours, values)
			return dummyHour, dummyHourIndex, false, false
		}
		assert.Equal(t, dummyStart.Minute(), value)
		assert.Equal(t, dummyMinutes, values)
		return dummyMinute, dummyMinuteIndex, true, false
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, time.Month(dummyMonth+1), month)
		assert.Equal(t, dummyDay+1, day)
		assert.Equal(t, dummyHour, hour)
		assert.Equal(t, dummyMinute, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_SecondOverflow(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyMinute = rand.Intn(60)
	var dummyMinuteIndex = rand.Intn(60)
	var dummySecond = rand.Intn(60)
	var dummySecondIndex = rand.Intn(60)
	var dummyTime = time.Now().Add(10 * time.Second)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 6
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		} else if findValueMatchFuncCalled == 4 {
			assert.Equal(t, dummyStart.Hour(), value)
			assert.Equal(t, dummyHours, values)
			return dummyHour, dummyHourIndex, false, false
		} else if findValueMatchFuncCalled == 5 {
			assert.Equal(t, dummyStart.Minute(), value)
			assert.Equal(t, dummyMinutes, values)
			return dummyMinute, dummyMinuteIndex, false, false
		}
		assert.Equal(t, dummyStart.Second(), value)
		assert.Equal(t, dummySeconds, values)
		return dummySecond, dummySecondIndex, false, true
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}
	timeDateExpected = 1
	timeDate = func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
		timeDateCalled++
		assert.Equal(t, dummyStart.Year(), year)
		assert.Equal(t, dummyStart.Month(), month)
		assert.Equal(t, dummyStart.Day(), day)
		assert.Equal(t, dummyStart.Hour(), hour)
		assert.Equal(t, dummyStart.Minute()+1, min)
		assert.Zero(t, sec)
		assert.Zero(t, nsec)
		assert.Equal(t, dummyLocation, loc)
		return dummyTime
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.False(t, completed)
	assert.Equal(t, dummyTime, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestDetermineScheduleIndex_NoOverflow_NoIncrement(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySeconds = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyMinutes = []int{rand.Intn(60), rand.Intn(60), rand.Intn(60)}
	var dummyHours = []int{rand.Intn(24), rand.Intn(24), rand.Intn(24)}
	var dummyDays = []int{rand.Intn(31), rand.Intn(31), rand.Intn(31)}
	var dummyMonths = []int{rand.Intn(12), rand.Intn(12), rand.Intn(12)}
	var dummyYears = []int{rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	var dummyLocation, _ = time.LoadLocation("Asia/Shanghai")
	var dummyWeekdays = map[time.Weekday]bool{
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
		time.Weekday(rand.Intn(7)): rand.Intn(100) > 50,
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
	var dummyYear = rand.Intn(100)
	var dummyYearIndex = rand.Intn(100)
	var dummyMonth = rand.Intn(12)
	var dummyMonthIndex = rand.Intn(12)
	var dummyDay = rand.Intn(31)
	var dummyDayIndex = rand.Intn(31)
	var dummyHour = rand.Intn(24)
	var dummyHourIndex = rand.Intn(24)
	var dummyMinute = rand.Intn(60)
	var dummyMinuteIndex = rand.Intn(60)
	var dummySecond = rand.Intn(60)
	var dummySecondIndex = rand.Intn(60)

	// mock
	createMock(t)

	// expect
	findValueMatchFuncExpected = 6
	findValueMatchFunc = func(value int, values []int) (int, int, bool, bool) {
		findValueMatchFuncCalled++
		if findValueMatchFuncCalled == 1 {
			assert.Equal(t, dummyStart.Year(), value)
			assert.Equal(t, dummyYears, values)
			return dummyYear, dummyYearIndex, false, false
		} else if findValueMatchFuncCalled == 2 {
			assert.Equal(t, int(dummyStart.Month())-1, value)
			assert.Equal(t, dummyMonths, values)
			return dummyMonth, dummyMonthIndex, false, false
		} else if findValueMatchFuncCalled == 3 {
			assert.Equal(t, dummyStart.Day()-1, value)
			assert.Equal(t, dummyDays, values)
			return dummyDay, dummyDayIndex, false, false
		} else if findValueMatchFuncCalled == 4 {
			assert.Equal(t, dummyStart.Hour(), value)
			assert.Equal(t, dummyHours, values)
			return dummyHour, dummyHourIndex, false, false
		} else if findValueMatchFuncCalled == 5 {
			assert.Equal(t, dummyStart.Minute(), value)
			assert.Equal(t, dummyMinutes, values)
			return dummyMinute, dummyMinuteIndex, false, false
		}
		assert.Equal(t, dummyStart.Second(), value)
		assert.Equal(t, dummySeconds, values)
		return dummySecond, dummySecondIndex, false, false
	}
	isWeekdayMatchFuncExpected = 1
	isWeekdayMatchFunc = func(year, month, day int, weekdays map[time.Weekday]bool) bool {
		isWeekdayMatchFuncCalled++
		assert.Equal(t, dummyYear, year)
		assert.Equal(t, dummyMonth, month)
		assert.Equal(t, dummyDay, day)
		assert.Equal(t, dummyWeekdays, weekdays)
		return true
	}

	// SUT + act
	var completed, start, err = determineScheduleIndex(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.True(t, completed)
	assert.Equal(t, dummyStart, start)
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestInitialiseSchedule_Error(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySchedule = &schedule{second: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	determineScheduleIndexFuncExpected = 1
	determineScheduleIndexFunc = func(start time.Time, schedule *schedule) (bool, time.Time, error) {
		determineScheduleIndexFuncCalled++
		assert.Equal(t, dummyStart, start)
		assert.Equal(t, dummySchedule, schedule)
		return false, dummyStart, dummyError
	}

	// SUT + act
	var err = initialiseSchedule(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestInitialiseSchedule_Success(t *testing.T) {
	// arrange
	var dummyStart = time.Now()
	var dummySchedule = &schedule{second: rand.Int()}

	// mock
	createMock(t)

	// expect
	determineScheduleIndexFuncExpected = 2
	determineScheduleIndexFunc = func(start time.Time, schedule *schedule) (bool, time.Time, error) {
		determineScheduleIndexFuncCalled++
		assert.Equal(t, dummyStart, start)
		assert.Equal(t, dummySchedule, schedule)
		return determineScheduleIndexFuncCalled == determineScheduleIndexFuncExpected, dummyStart, nil
	}

	// SUT + act
	var err = initialiseSchedule(
		dummyStart,
		dummySchedule,
	)

	// assert
	assert.NoError(t, err)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_Schedule_WithoutFrom(t *testing.T) {
	// arrange
	var dummyScheduleMaker = &scheduleMaker{
		seconds: []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50},
	}
	var dummyTimeNow = time.Now()
	var dummySchedule = &schedule{year: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyTimeNow
	}
	constructScheduleTemplateFuncExpected = 1
	constructScheduleTemplateFunc = func(scheduleMaker *scheduleMaker) *schedule {
		constructScheduleTemplateFuncCalled++
		assert.Equal(t, dummyScheduleMaker, scheduleMaker)
		return dummySchedule
	}
	initialiseScheduleFuncExpected = 1
	initialiseScheduleFunc = func(start time.Time, schedule *schedule) error {
		initialiseScheduleFuncCalled++
		assert.Equal(t, dummyTimeNow, start)
		assert.Equal(t, dummySchedule, schedule)
		return dummyError
	}

	// SUT
	var sut, err = dummyScheduleMaker.Schedule()

	// act
	var result, ok = sut.(*schedule)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummySchedule, result)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
}

func TestScheduleMaker_Schedule_WithFrom(t *testing.T) {
	// arrange
	var dummyFrom = time.Now().Add(100 * time.Second)
	var dummyScheduleMaker = &scheduleMaker{
		seconds: []bool{rand.Intn(100) > 50, rand.Intn(100) > 50, rand.Intn(100) > 50},
		from:    &dummyFrom,
	}
	var dummyTimeNow = time.Now()
	var dummySchedule = &schedule{year: rand.Int()}
	var dummyError = errors.New("some error")

	// mock
	createMock(t)

	// expect
	timeNowExpected = 1
	timeNow = func() time.Time {
		timeNowCalled++
		return dummyTimeNow
	}
	constructScheduleTemplateFuncExpected = 1
	constructScheduleTemplateFunc = func(scheduleMaker *scheduleMaker) *schedule {
		constructScheduleTemplateFuncCalled++
		assert.Equal(t, dummyScheduleMaker, scheduleMaker)
		return dummySchedule
	}
	initialiseScheduleFuncExpected = 1
	initialiseScheduleFunc = func(start time.Time, schedule *schedule) error {
		initialiseScheduleFuncCalled++
		assert.Equal(t, dummyFrom, start)
		assert.Equal(t, dummySchedule, schedule)
		return dummyError
	}

	// SUT
	var sut, err = dummyScheduleMaker.Schedule()

	// act
	var result, ok = sut.(*schedule)

	// assert
	assert.True(t, ok)
	assert.Equal(t, dummySchedule, result)
	assert.Equal(t, dummyError, err)

	// verify
	verifyAll(t)
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

	for given, expect := range testData {
		// arrange
		var timeStart, _ = time.Parse(layout, given)
		var timeExpect, _ = time.Parse(layout, expect)

		// stub
		timeNow = func() time.Time {
			return timeStart
		}

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
