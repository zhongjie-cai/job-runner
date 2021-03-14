package jobrunner

import (
	"time"
)

type Schedule interface {
	// WaitForNextExecution should internally pause the running goroutine and wait till a next schedule is anticipated
	//                      return true when a next execution should be conducted, or false to indicate a termination
	WaitForNextExecution() bool
}

type CronSchedule struct {
	seconds  []bool
	minutes  []bool
	hours    []bool
	weekdays []bool
	days     []bool
	months   []bool
	years    map[int]bool
	from     *time.Time
	till     *time.Time
	last     *time.Time
}

// NewCronSchedule creates an empty CronSchedule for consumer to manually configure a preferred schedule
func NewCronSchedule() *CronSchedule {
	return &CronSchedule{
		[]bool{},
		[]bool{},
		[]bool{},
		[]bool{},
		[]bool{},
		[]bool{},
		map[int]bool{},
		nil,
		nil,
		nil,
	}
}

func isValueMatch(value int, values []bool) bool {
	if len(values) == 0 {
		return true
	}
	return values[value]
}

func isTimeForNextRun(timeNext time.Time, cronSchedule *CronSchedule) bool {
	if !isValueMatchFunc(timeNext.Second(), cronSchedule.seconds) {
		return false
	}
	if !isValueMatchFunc(timeNext.Minute(), cronSchedule.minutes) {
		return false
	}
	if !isValueMatchFunc(timeNext.Hour(), cronSchedule.hours) {
		return false
	}
	if !isValueMatchFunc(int(timeNext.Weekday()), cronSchedule.weekdays) {
		return false
	}
	if !isValueMatchFunc(timeNext.Day(), cronSchedule.days) {
		return false
	}
	if !isValueMatchFunc(int(timeNext.Month()), cronSchedule.months) {
		return false
	}
	if len(cronSchedule.years) > 0 &&
		!cronSchedule.years[timeNext.Year()] {
		return false
	}
	return true
}

func calculateNextWaitDuration(cronSchedule *CronSchedule) (time.Time, time.Duration) {
	var timeLast = timeNow()
	if cronSchedule.last != nil {
		timeLast = *cronSchedule.last
	}
	if cronSchedule.from != nil &&
		cronSchedule.from.After(timeLast) {
		timeLast = *cronSchedule.from
	}
	var timeNext = timeLast.Truncate(time.Second)
	// this can be costly, especially if you set it up that the next run would be actually in years
	//   gladly, the trade-off here is acceptable as you anyway would wait for years for a next run
	for {
		timeNext = timeNext.Add(1 * time.Second)
		if isTimeForNextRunFunc(
			timeNext,
			cronSchedule,
		) {
			return timeNext, timeNext.Sub(timeNow())
		}
	}
}

func (cronSchedule *CronSchedule) WaitForNextExecution() bool {
	var currentLocalTime = timeNow()
	if cronSchedule.till != nil &&
		cronSchedule.till.Before(currentLocalTime) {
		// causes the schedule to terminate
		return false
	}
	// pause the running goroutine for the calculated duration
	var timeNext, waitDuration = calculateNextWaitDurationFunc(
		cronSchedule,
	)
	<-timeAfter(
		waitDuration,
	)
	cronSchedule.last = &timeNext
	return true
}

func generateFlagsData(data []bool, total int, values ...int) []bool {
	if len(data) == 0 {
		data = make([]bool, total)
	}
	if len(values) == 0 {
		for value := 0; value < total; value++ {
			data[value] = true
		}
	} else {
		for _, value := range values {
			if value >= 0 && value < total {
				data[value] = true
			}
		}
	}
	return data
}

// OnSeconds sets up the CronSchedule on second's level; if not called or called with no parameters, then every second is considered to be set up
func (cronSchedule *CronSchedule) OnSeconds(seconds ...int) *CronSchedule {
	cronSchedule.seconds = generateFlagsDataFunc(
		cronSchedule.seconds,
		60,
		seconds...,
	)
	return cronSchedule
}

// OnMinutes sets up the CronSchedule on minute's level; if not called or called with no parameters, then every minute is considered to be set up
func (cronSchedule *CronSchedule) OnMinutes(minutes ...int) *CronSchedule {
	cronSchedule.minutes = generateFlagsDataFunc(
		cronSchedule.minutes,
		60,
		minutes...,
	)
	return cronSchedule
}

// AtHours sets up the CronSchedule on hour's level; if not called or called with no parameters, then every hour is considered to be set up
func (cronSchedule *CronSchedule) AtHours(hours ...int) *CronSchedule {
	cronSchedule.hours = generateFlagsDataFunc(
		cronSchedule.hours,
		24,
		hours...,
	)
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on weekday's level; if not called or called with no parameters, then every weekday is considered to be set up
func (cronSchedule *CronSchedule) OnWeekdays(weekdays ...time.Weekday) *CronSchedule {
	var weekdaysInInt = []int{}
	for _, weekday := range weekdays {
		weekdaysInInt = append(weekdaysInInt, int(weekday))
	}
	cronSchedule.weekdays = generateFlagsDataFunc(
		cronSchedule.weekdays,
		7,
		weekdaysInInt...,
	)
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on day's level; if not called or called with no parameters, then every day is considered to be set up
func (cronSchedule *CronSchedule) OnDays(days ...int) *CronSchedule {
	cronSchedule.days = generateFlagsDataFunc(
		cronSchedule.days,
		31,
		days...,
	)
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on month's level; if not called or called with no parameters, then every month is considered to be set up
func (cronSchedule *CronSchedule) InMonths(months ...time.Month) *CronSchedule {
	var monthsInInt = []int{}
	for _, month := range months {
		monthsInInt = append(monthsInInt, int(month))
	}
	cronSchedule.months = generateFlagsDataFunc(
		cronSchedule.months,
		13, // due to the fact that month value start from 1 instead of 0...
		monthsInInt...,
	)
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on year's level; if not called or called with no parameters, then every year is considered to be set up
func (cronSchedule *CronSchedule) InYears(years ...int) *CronSchedule {
	if len(years) > 0 {
		for _, year := range years {
			cronSchedule.years[year] = true
		}
	}
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on second's level; if not called or called with no parameters, then every second is considered to be set up
func (cronSchedule *CronSchedule) From(start time.Time) *CronSchedule {
	cronSchedule.from = &start
	return cronSchedule
}

// OnSeconds sets up the CronSchedule on second's level; if not called or called with no parameters, then every second is considered to be set up
func (cronSchedule *CronSchedule) Till(end time.Time) *CronSchedule {
	cronSchedule.till = &end
	return cronSchedule
}
