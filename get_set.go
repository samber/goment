package goment

import (
	"math"
	"time"
)

// Get is a string getter using the units. Returns 0 if unsupported property.
func (g *Goment) Get(units string) int {
	switch units {
	case "y", "year", "years":
		return g.Year()
	case "M", "month", "months":
		return g.Month()
	case "D", "date", "dates":
		return g.Date()
	case "h", "hour", "hours":
		return g.Hour()
	case "m", "minute", "minutes":
		return g.Minute()
	case "s", "second", "seconds":
		return g.Second()
	case "ms", "millisecond", "milliseconds":
		return g.Millisecond()
	case "ns", "nanosecond", "nanoseconds":
		return g.Nanosecond()
	}
	return 0
}

// Nanosecond gets the nanoseconds.
func (g *Goment) Nanosecond() int {
	return g.ToTime().Nanosecond()
}

// Millisecond gets the milliseconds.
func (g *Goment) Millisecond() int {
	return g.Second() * 1000
}

// Second gets the seconds.
func (g *Goment) Second() int {
	return g.ToTime().Second()
}

// Minute gets the minutes.
func (g *Goment) Minute() int {
	return g.ToTime().Minute()
}

// Hour gets the hour.
func (g *Goment) Hour() int {
	return g.ToTime().Hour()
}

// Date gets the day of the month.
func (g *Goment) Date() int {
	return g.ToTime().Day()
}

// Day gets the day of the week (Sunday = 0...).
func (g *Goment) Day() int {
	return int(g.ToTime().Weekday())
}

// Weekday gets the day of the week according to the locale.
func (g *Goment) Weekday() int {
	return 0
}

// ISOWeekday gets the ISO day of the week with 1 being Monday and 7 being Sunday.
func (g *Goment) ISOWeekday() int {
	wd := g.Day()
	if wd == 0 {
		wd = 7
	}
	return wd
}

// DayOfYear gets the day of the year.
func (g *Goment) DayOfYear() int {
	return g.ToTime().YearDay()
}

// Week gets the week of the year according to the locale.
func (g *Goment) Week() int {
	return 0
}

// ISOWeek gets the ISO week of the year.
func (g *Goment) ISOWeek() int {
	_, week := g.ToTime().ISOWeek()
	return week
}

// Month gets the month (January = 1...).
func (g *Goment) Month() int {
	return int(g.ToTime().Month())
}

// Quarter gets the quarter (1 to 4).
func (g *Goment) Quarter() int {
	return int(math.Ceil(float64(g.Month()) / 3))
}

// Year gets the year.
func (g *Goment) Year() int {
	return g.ToTime().Year()
}

// WeekYear gets the week-year according to the locale.
func (g *Goment) WeekYear() int {
	return 0
}

// ISOWeekYear gets the ISO week-year.
func (g *Goment) ISOWeekYear() int {
	year, _ := g.ToTime().ISOWeek()
	return year
}

// WeeksInYear gets the number of weeks according to locale in the current Goment's year.
func (g *Goment) WeeksInYear() int {
	return 0
}

// ISOWeeksInYear gets the number of weeks in the current Goment's year, according to ISO weeks.
func (g *Goment) ISOWeeksInYear() int {
	return 0
}

// Set is a generic setter, accepting units as the first argument, and value as the second.
func (g *Goment) Set(units string, value int) *Goment {
	switch units {
	case "y", "year", "years":
		return g.SetYear(value)
	case "M", "month", "months":
		return g.SetMonth(value)
	case "D", "date", "dates":
		return g.SetDate(value)
	case "h", "hour", "hours":
		return g.SetHour(value)
	case "m", "minute", "minutes":
		return g.SetMinute(value)
	case "s", "second", "seconds":
		return g.SetSecond(value)
	case "ms", "millisecond", "milliseconds":
		return g.SetMillisecond(value)
	case "ns", "nanosecond", "nanoseconds":
		return g.SetNanosecond(value)
	}
	return g
}

// SetNanosecond sets the nanoseconds.
func (g *Goment) SetNanosecond(nanoseconds int) *Goment {
	if nanoseconds >= 0 && nanoseconds <= 999999999 {
		return g.addNanoseconds(nanoseconds - g.Nanosecond())
	}
	return g
}

// SetMillisecond sets the milliseconds.
func (g *Goment) SetMillisecond(milliseconds int) *Goment {
	if milliseconds >= 0 && milliseconds <= 59000 {
		return g.addMilliseconds(milliseconds - g.Millisecond())
	}
	return g
}

// SetSecond sets the seconds.
func (g *Goment) SetSecond(seconds int) *Goment {
	if seconds >= 0 && seconds <= 59 {
		return g.addSeconds(seconds - g.Second())
	}
	return g
}

// SetMinute sets the minutes.
func (g *Goment) SetMinute(minutes int) *Goment {
	if minutes >= 0 && minutes <= 59 {
		return g.addMinutes(minutes - g.Minute())
	}
	return g
}

// SetHour sets the hour.
func (g *Goment) SetHour(hours int) *Goment {
	if hours >= 0 && hours <= 23 {
		return g.addHours(hours - g.Hour())
	}
	return g
}

// SetDate sets the day of the month. If the date passed in is greater than the number of days in the month,
// then the day is set to the last day of the month.
func (g *Goment) SetDate(date int) *Goment {
	if date >= 1 && date <= 31 {
		daysInMonth := g.DaysInMonth()
		if date >= daysInMonth {
			date = daysInMonth
		}
		return g.addDays(date - g.Date())
	}
	return g
}

// SetDay sets the day of the week (Sunday = 0...).
func (g *Goment) SetDay(day int) *Goment {
	if day >= 0 && day <= 6 {
		return g.addDays(day - g.Day())
	}
	return g
}

// SetWeekday sets the day of the week according to the locale.
func (g *Goment) SetWeekday(weekday int) *Goment {
	return g
}

// SetISOWeekday sets the ISO day of the week with 1 being Monday and 7 being Sunday.
func (g *Goment) SetISOWeekday(weekday int) *Goment {
	if weekday >= 1 && weekday <= 7 {
		if weekday == 7 {
			weekday = 0
		}
		return g.SetDay(weekday)
	}
	return g
}

// SetDayOfYear sets the day of the year. For non-leap years, 366 is treated as 365.
func (g *Goment) SetDayOfYear(doy int) *Goment {
	if doy >= 1 && doy <= 366 {
		if !g.IsLeapYear() && doy == 366 {
			doy = 365
		}
		return g.addDays(doy - g.DayOfYear())
	}
	return g
}

// SetWeek sets the week of the year according to the locale.
func (g *Goment) SetWeek(week int) *Goment {
	return g
}

// SetISOWeek sets the ISO week of the year.
func (g *Goment) SetISOWeek(week int) *Goment {
	return g
}

// SetMonth sets the month (January = 1...). If new month has less days than current month,
// the date is pinned to the end of the target month.
func (g *Goment) SetMonth(month int) *Goment {
	if month >= 1 && month <= 12 {
		currentDate := g.Date()
		newDaysInMonth := daysInMonth(month, g.Year())
		if currentDate > newDaysInMonth {
			g.SetDate(newDaysInMonth)
		}
		return g.addMonths(month - g.Month())
	}
	return g
}

// SetQuarter sets the quarter (1 to 4).
func (g *Goment) SetQuarter(quarter int) *Goment {
	if quarter >= 1 && quarter <= 4 {
		return g.addQuarters(quarter - g.Quarter())
	}
	return g
}

// SetYear sets the year.
func (g *Goment) SetYear(year int) *Goment {
	return g.addYears(year - g.Year())
}

// SetWeekYear sets the week-year according to the locale.
func (g *Goment) SetWeekYear(weekYear int) *Goment {
	return g
}

// SetISOWeekYear sets the ISO week-year.
func (g *Goment) SetISOWeekYear(weekYear int) *Goment {
	return g
}

// Format functions.

// DaysInMonth returns the number of days in the set month.
func (g *Goment) DaysInMonth() int {
	return daysInMonth(g.Month(), g.Year())
}

func daysInMonth(month, year int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
