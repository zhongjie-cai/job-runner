package jobrunner

import (
	"fmt"
	"sort"
	"time"
)

type ScheduleMaker interface {
	// OnSeconds sets up the schedule maker on second's level; if not called or called with no parameters, then every second is considered to be set up
	OnSeconds(seconds ...int) ScheduleMaker
	// OnMinutes sets up the schedule maker on minute's level; if not called or called with no parameters, then every minute is considered to be set up
	OnMinutes(minutes ...int) ScheduleMaker
	// AtHours sets up the schedule maker on hour's level; if not called or called with no parameters, then every hour is considered to be set up
	AtHours(hours ...int) ScheduleMaker
	// OnWeekdays sets up the schedule maker on weekday's level; if not called or called with no parameters, then every weekday is considered to be set up
	OnWeekdays(weekdays ...time.Weekday) ScheduleMaker
	// OnDays sets up the schedule maker on day's level; if not called or called with no parameters, then every day is considered to be set up
	OnDays(days ...int) ScheduleMaker
	// InMonths sets up the schedule maker on month's level; if not called or called with no parameters, then every month is considered to be set up
	InMonths(months ...time.Month) ScheduleMaker
	// InYears sets up the schedule maker on year's level; if not called or called with no parameters, then every year is considered to be set up
	InYears(years ...int) ScheduleMaker
	// From sets up the schedule maker for its start datetime; if not called or called with no parameters, then the schedule will start immediately
	From(start time.Time) ScheduleMaker
	// Till sets up the schedule maker for its end datetime; if not called or called with no parameters, then the schedule won't stop until its data increment
	Till(end time.Time) ScheduleMaker
	// Timezone sets up the schedule maker for its timezone location; if not called or called with no parameters, then the schedule takes time.Local as default
	Timezone(timezone *time.Location) ScheduleMaker
	// SkipOverdue sets up the schedule maker to skip an overdue schedule or not; if not called, then the schedule defaults to not skip overdues
	SkipOverdue() ScheduleMaker
	// Done returns a compiled schedule based on all previously configured settings
	Schedule() (Schedule, error)
}

type scheduleMaker struct {
	seconds     []bool
	minutes     []bool
	hours       []bool
	weekdays    []bool
	days        []bool
	months      []bool
	years       map[int]bool
	from        *time.Time
	till        *time.Time
	timezone    *time.Location
	skipOverdue bool
}

// NewScheduleMaker creates an empty scheduleMaker for consumer to manually configure a preferred schedule
func NewScheduleMaker() ScheduleMaker {
	return &scheduleMaker{
		[]bool{},       // seconds
		[]bool{},       // minutes
		[]bool{},       // hours
		[]bool{},       // weekdays
		[]bool{},       // days
		[]bool{},       // months
		map[int]bool{}, // years
		nil,            // from
		nil,            // till
		time.Local,     // timezone
		false,          // skipOverdue
	}
}

func generateFlagsData(data []bool, total int, values ...int) []bool {
	if len(values) == 0 {
		data = make([]bool, total)
		for value := 0; value < total; value++ {
			data[value] = true
		}
	} else {
		if len(data) < total {
			for count := len(data); count < total; count++ {
				data = append(data, false)
			}
		}
		for _, value := range values {
			data[value%total] = true
		}
	}
	return data
}

// OnSeconds sets up the scheduleMaker on second's level; if not called or called with no parameters, then every second is considered to be set up
func (scheduleMaker *scheduleMaker) OnSeconds(seconds ...int) ScheduleMaker {
	scheduleMaker.seconds = generateFlagsData(
		scheduleMaker.seconds,
		60,
		seconds...,
	)
	return scheduleMaker
}

// OnMinutes sets up the scheduleMaker on minute's level; if not called or called with no parameters, then every minute is considered to be set up
func (scheduleMaker *scheduleMaker) OnMinutes(minutes ...int) ScheduleMaker {
	scheduleMaker.minutes = generateFlagsData(
		scheduleMaker.minutes,
		60,
		minutes...,
	)
	return scheduleMaker
}

// AtHours sets up the scheduleMaker on hour's level; if not called or called with no parameters, then every hour is considered to be set up
func (scheduleMaker *scheduleMaker) AtHours(hours ...int) ScheduleMaker {
	scheduleMaker.hours = generateFlagsData(
		scheduleMaker.hours,
		24,
		hours...,
	)
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on weekday's level; if not called or called with no parameters, then every weekday is considered to be set up
func (scheduleMaker *scheduleMaker) OnWeekdays(weekdays ...time.Weekday) ScheduleMaker {
	var weekdaysInInt = []int{}
	for _, weekday := range weekdays {
		weekdaysInInt = append(weekdaysInInt, int(weekday))
	}
	scheduleMaker.weekdays = generateFlagsData(
		scheduleMaker.weekdays,
		7,
		weekdaysInInt...,
	)
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on day's level; if not called or called with no parameters, then every day is considered to be set up
func (scheduleMaker *scheduleMaker) OnDays(days ...int) ScheduleMaker {
	var daysInInt = []int{}
	for _, day := range days {
		// due to the fact that day value start from 1 instead of 0...
		daysInInt = append(daysInInt, day-1)
	}
	scheduleMaker.days = generateFlagsData(
		scheduleMaker.days,
		31,
		daysInInt...,
	)
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on month's level; if not called or called with no parameters, then every month is considered to be set up
func (scheduleMaker *scheduleMaker) InMonths(months ...time.Month) ScheduleMaker {
	var monthsInInt = []int{}
	for _, month := range months {
		// due to the fact that month value start from 1 instead of 0...
		monthsInInt = append(monthsInInt, int(month-1))
	}
	scheduleMaker.months = generateFlagsData(
		scheduleMaker.months,
		12,
		monthsInInt...,
	)
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on year's level; if not called or called with no parameters, then every year is considered to be set up
func (scheduleMaker *scheduleMaker) InYears(years ...int) ScheduleMaker {
	if len(years) > 0 {
		for _, year := range years {
			scheduleMaker.years[year] = true
		}
	}
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on second's level; if not called or called with no parameters, then every second is considered to be set up
func (scheduleMaker *scheduleMaker) From(start time.Time) ScheduleMaker {
	scheduleMaker.from = &start
	return scheduleMaker
}

// OnSeconds sets up the scheduleMaker on second's level; if not called or called with no parameters, then every second is considered to be set up
func (scheduleMaker *scheduleMaker) Till(end time.Time) ScheduleMaker {
	scheduleMaker.till = &end
	return scheduleMaker
}

// Timezone sets up the schedule maker for its timezone location; if not called or called with no parameters, then the schedule takes time.Local as default
func (scheduleMaker *scheduleMaker) Timezone(timezone *time.Location) ScheduleMaker {
	if timezone != nil {
		scheduleMaker.timezone = timezone
	} else {
		scheduleMaker.timezone = time.Local
	}
	return scheduleMaker
}

// SkipOverdue sets up the schedule maker to skip an overdue schedule or not; if not called, then the schedule defaults to not skip overdues
func (scheduleMaker *scheduleMaker) SkipOverdue() ScheduleMaker {
	scheduleMaker.skipOverdue = true
	return scheduleMaker
}

func constructValueSlice(values []bool, total int) []int {
	var data = []int{}
	if len(values) == 0 {
		for index := 0; index < total; index++ {
			data = append(data, index)
		}
	} else {
		for index, valid := range values {
			if valid {
				data = append(data, index)
			}
		}
	}
	return data
}

func constructWeekdayMap(weekdays []bool) map[time.Weekday]bool {
	var data = map[time.Weekday]bool{}
	if len(weekdays) == 0 {
		for weekday := 0; weekday < 7; weekday++ {
			data[time.Weekday(weekday)] = true
		}
	} else {
		for weekday, valid := range weekdays {
			if valid {
				data[time.Weekday(weekday)] = true
			}
		}
	}
	return data
}

func constructYearSlice(years map[int]bool) []int {
	var data = []int{}
	if len(years) == 0 {
		// hard-code this to allow execution for 100 years if no year is specified
		var currentTime = time.Now()
		var currentYear = currentTime.Year()
		for year := currentYear; year < currentYear+100; year++ {
			data = append(data, year)
		}
	} else {
		for year, valid := range years {
			if valid {
				data = append(data, year)
			}
		}
		sort.Ints(data)
	}
	return data
}

func constructScheduleTemplate(scheduleMaker *scheduleMaker) *schedule {
	var schedule = &schedule{
		secondIndex: 0,
		seconds:     constructValueSlice(scheduleMaker.seconds, 60),
		minuteIndex: 0,
		minutes:     constructValueSlice(scheduleMaker.minutes, 60),
		hourIndex:   0,
		hours:       constructValueSlice(scheduleMaker.hours, 24),
		dayIndex:    0,
		days:        constructValueSlice(scheduleMaker.days, 31),
		monthIndex:  0,
		months:      constructValueSlice(scheduleMaker.months, 12),
		yearIndex:   0,
		years:       constructYearSlice(scheduleMaker.years),
		weekdays:    constructWeekdayMap(scheduleMaker.weekdays),
		till:        scheduleMaker.till,
		timezone:    scheduleMaker.timezone,
		skipOverdue: scheduleMaker.skipOverdue,
	}
	schedule.second = schedule.seconds[schedule.secondIndex]
	schedule.minute = schedule.minutes[schedule.minuteIndex]
	schedule.hour = schedule.hours[schedule.hourIndex]
	schedule.day = schedule.days[schedule.dayIndex]
	schedule.month = schedule.months[schedule.monthIndex]
	schedule.year = schedule.years[schedule.yearIndex]
	return schedule
}

func findValueMatch(value int, values []int) (int, int, bool, bool) {
	var count = len(values)
	if count == 0 {
		return 0, 0, false, false
	}
	for index := 0; index < count; index++ {
		if value > values[index] {
			continue
		}
		return values[index], index, value < values[index], false
	}
	return values[0], 0, true, true
}

func isWeekdayMatch(year, month, day int, weekdays map[time.Weekday]bool) bool {
	if len(weekdays) == 0 {
		return true
	}
	var date = time.Date(
		year,
		time.Month(month+1),
		day+1,
		0,
		0,
		0,
		0,
		time.Local,
	)
	var valid, found = weekdays[date.Weekday()]
	return found && valid
}

func determineScheduleIndex(
	start time.Time,
	schedule *schedule,
) (bool, time.Time, error) {
	var increment bool
	var overflow bool
	schedule.year, schedule.yearIndex, increment, overflow = findValueMatch(
		start.Year(),
		schedule.years,
	)
	if overflow {
		return false, start, fmt.Errorf("Invalid schedule configuration: no valid next execution time available")
	} else if increment {
		return false,
			time.Date(
				schedule.year,
				time.January,
				1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	schedule.month, schedule.monthIndex, increment, overflow = findValueMatch(
		int(start.Month())-1,
		schedule.months,
	)
	if overflow {
		return false,
			time.Date(
				start.Year()+1,
				time.January,
				1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	} else if increment {
		return false,
			time.Date(
				schedule.year,
				time.Month(schedule.month+1),
				1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	schedule.day, schedule.dayIndex, increment, overflow = findValueMatch(
		start.Day()-1,
		schedule.days,
	)
	if overflow {
		return false,
			time.Date(
				start.Year(),
				start.Month()+1,
				1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	} else if increment {
		return false,
			time.Date(
				schedule.year,
				time.Month(schedule.month+1),
				schedule.day+1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	if !isWeekdayMatch(
		schedule.year,
		schedule.month,
		schedule.day,
		schedule.weekdays,
	) {
		return false,
			time.Date(
				start.Year(),
				start.Month(),
				start.Day()+1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	schedule.hour, schedule.hourIndex, increment, overflow = findValueMatch(
		start.Hour(),
		schedule.hours,
	)
	if overflow {
		return false,
			time.Date(
				start.Year(),
				start.Month(),
				start.Day()+1,
				0,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	} else if increment {
		return false,
			time.Date(
				schedule.year,
				time.Month(schedule.month+1),
				schedule.day+1,
				schedule.hour,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	schedule.minute, schedule.minuteIndex, increment, overflow = findValueMatch(
		start.Minute(),
		schedule.minutes,
	)
	if overflow {
		return false,
			time.Date(
				start.Year(),
				start.Month(),
				start.Day(),
				start.Hour()+1,
				0,
				0,
				0,
				schedule.timezone,
			),
			nil
	} else if increment {
		return false,
			time.Date(
				schedule.year,
				time.Month(schedule.month+1),
				schedule.day+1,
				schedule.hour,
				schedule.minute,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	schedule.second, schedule.secondIndex, _, overflow = findValueMatch(
		start.Second(),
		schedule.seconds,
	)
	if overflow {
		return false,
			time.Date(
				start.Year(),
				start.Month(),
				start.Day(),
				start.Hour(),
				start.Minute()+1,
				0,
				0,
				schedule.timezone,
			),
			nil
	}
	return true, start, nil
}

func initialiseSchedule(
	start time.Time,
	schedule *schedule,
) error {
	var calcError error
	for completed := false; !completed; {
		completed, start, calcError = determineScheduleIndex(
			start,
			schedule,
		)
		if calcError != nil {
			return calcError
		}
	}
	return nil
}

// Schedule returns a compiled schedule based on all previously configured settings
func (scheduleMaker *scheduleMaker) Schedule() (Schedule, error) {
	var start = time.Now()
	if scheduleMaker.from != nil &&
		scheduleMaker.from.After(start) {
		start = *scheduleMaker.from
	}
	var schedule = constructScheduleTemplate(
		scheduleMaker,
	)
	var calcError = initialiseSchedule(
		start,
		schedule,
	)
	return schedule, calcError
}
