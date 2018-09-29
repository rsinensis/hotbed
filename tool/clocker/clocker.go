package clocker

import (
	"time"
)

func BeginningOfDay(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfDay(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfYesterday(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Add(-24 * time.Hour)

	return this
}

func EndOfYesterday(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local).Add(-24 * time.Hour).Add(-time.Second)

	return this
}

func BeginningOfTomorrow(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Add(+24 * time.Hour)

	return this
}

func EndOfTomorrow(t time.Time) time.Time {

	year, month, day := t.Date()
	this := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local).Add(+24 * time.Hour).Add(-time.Second)

	return this
}

func BeginningOfWeek(t time.Time, firstDayMonday bool) time.Time {

	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if firstDayMonday {
		weekday = weekday - 1
	}
	this := time.Date(year, month, day-weekday, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfWeek(t time.Time, firstDayMonday bool) time.Time {

	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if firstDayMonday {
		weekday = weekday - 1
	}
	this := time.Date(year, month, day+7-weekday, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfLastWeek(t time.Time, firstDayMonday bool) time.Time {

	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if firstDayMonday {
		weekday = weekday - 1
	}
	this := time.Date(year, month, day-weekday-7, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfLastWeek(t time.Time, firstDayMonday bool) time.Time {

	year, month, day := t.Date()
	weekday := int(t.Weekday())
	if firstDayMonday {
		weekday = weekday - 1
	}
	this := time.Date(year, month, day+7-weekday-7, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	this := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	this := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfLastMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	this := time.Date(year, month-1, 1, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfLastMonth(t time.Time) time.Time {

	year, month, _ := t.Date()
	this := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfYear(t time.Time) time.Time {

	year, _, _ := t.Date()
	this := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfYear(t time.Time) time.Time {

	year, _, _ := t.Date()
	this := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}

func BeginningOfLastYear(t time.Time) time.Time {

	year, _, _ := t.Date()
	this := time.Date(year-1, 1, 1, 0, 0, 0, 0, time.Local)

	return this
}

func EndOfLastYear(t time.Time) time.Time {

	year, _, _ := t.Date()
	this := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	return this
}
