package jobrunner

import (
	"time"
)

type Schedule interface {
	// NextSchedule returns the time instance pointing to when a next schedule should happen;
	//   if no schedule can be made, return nil to indicate the termination of application
	NextSchedule() *time.Time
}

type schedule struct {
	second      int
	secondIndex int
	seconds     []int
	minute      int
	minuteIndex int
	minutes     []int
	hour        int
	hourIndex   int
	hours       []int
	day         int
	dayIndex    int
	days        []int
	month       int
	monthIndex  int
	months      []int
	year        int
	yearIndex   int
	years       []int
	weekdays    map[time.Weekday]bool
	till        *time.Time
}

func moveValueIndex(
	oldIndex int,
	values []int,
	maxValue int,
) (int, int, bool) {
	var newIndex = oldIndex + 1
	if newIndex >= len(values) ||
		values[newIndex] >= maxValue {
		return values[0], 0, true
	}
	return values[newIndex], newIndex, false
}

func getDaysOfMonth(year int, month int) int {
	// plays a trick that time.Date accepts values outside of normal ranges
	//   so setting day to 0 would actually yield the last day of a previous month
	//   that helps calculate the number of days within a given month
	var lastDay = timeDate(
		year,                   // current year so no change
		time.Month(month+2),    // slide to next month, yields back as current month by the trick
		0,                      // the essential part of this trick - 0-day!
		0, 0, 0, 0, time.Local, // other values don't matter, so use defaults
	)
	return lastDay.Day()
}

func constructTimeBySchedule(schedule *schedule) time.Time {
	return timeDate(
		schedule.year,
		time.Month(schedule.month+1),
		schedule.day+1,
		schedule.hour,
		schedule.minute,
		schedule.second,
		0,
		time.Local,
	)
}

func updateScheduleIndex(
	schedule *schedule,
) {
	var reset bool
	// get next second from schedule data
	schedule.second, schedule.secondIndex, reset = moveValueIndexFunc(
		schedule.secondIndex,
		schedule.seconds,
		60,
	)
	if reset {
		// second reset, thus get next minute from schedule data
		schedule.minute, schedule.minuteIndex, reset = moveValueIndexFunc(
			schedule.minuteIndex,
			schedule.minutes,
			60,
		)
		if reset {
			// minute reset, thus get next hour from schedule data
			schedule.hour, schedule.hourIndex, reset = moveValueIndexFunc(
				schedule.hourIndex,
				schedule.hours,
				24,
			)
			if reset {
				// hour reset, thus get next day from schedule data
				schedule.day, schedule.dayIndex, reset = moveValueIndexFunc(
					schedule.dayIndex,
					schedule.days,
					getDaysOfMonthFunc(
						schedule.year,
						schedule.month,
					),
				)
				if reset {
					// day reset, thus get next month from schedule data
					schedule.month, schedule.monthIndex, reset = moveValueIndexFunc(
						schedule.monthIndex,
						schedule.months,
						12,
					)
					if reset {
						// month reset, thus get next year from schedule data
						schedule.year, schedule.yearIndex, _ = moveValueIndexFunc(
							schedule.yearIndex,
							schedule.years,
							9999, // hopefully nobody is still using this library by year 9999?
						)
					}
				}
			}
		}
	}
}

func (schedule *schedule) NextSchedule() *time.Time {
	var currentLocalTime = timeNow()
	if schedule.till != nil &&
		schedule.till.Before(currentLocalTime) {
		// causes the schedule to terminate
		return nil
	}
	// load next schedule time
	var timeNext = constructTimeByScheduleFunc(
		schedule,
	)
	if timeNext.Before(currentLocalTime) {
		// in past means schedule has been reset, so no longer runnable
		return nil
	}
	updateScheduleIndexFunc(
		schedule,
	)
	return &timeNext
}
