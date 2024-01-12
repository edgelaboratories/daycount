package daycount

import (
	"math"
	"time"

	"github.com/fxtlabs/date"
)

// DayCounter computes the year fraction between a from and a to date
// according to a predefined day-count convention.
// All DayCounter functions assume that from is never later than to.
type DayCounter func(from, to date.Date) float64

// NewDayCounter returns a DayCounter based on the input convention.
func NewDayCounter(convention Convention) DayCounter {
	switch convention {
	case ActualActualAFB:
		return yearFractionActualActualAFB

	case ActualThreeSixty:
		return yearFractionActualThreeSixty

	case ActualThreeSixtyFiveFixed:
		return yearFractionActualThreeSixtyFiveFixed

	case ThirtyThreeSixtyUS:
		return yearFractionThirtyThreeSixtyUS

	case ThirtyThreeSixtyEuropean:
		return yearFractionThirtyThreeSixtyEuropean

	case ThirtyThreeSixtyItalian:
		return yearFractionThirtyThreeSixtyItalian

	case ThirtyThreeSixtyGerman:
		return yearFractionThirtyThreeSixtyGerman

	case ActualActual, outOfRangeConvention:
		fallthrough

	default:
		return yearFractionActualActual
	}
}

// YearFraction returns the year fraction difference between two dates
// according to the input convention.
// If the convention is not recognized, it defaults to ActualActual.
func YearFraction(from, to date.Date, convention Convention) float64 {
	return NewDayCounter(convention)(from, to)
}

const (
	threeSixtyDays     = 360.0
	threeSixtyFiveDays = 365.0
	threeSixtySixDays  = 366.0
)

func yearFractionActualActual(from, to date.Date) float64 {
	fromYear, toYear := from.Year(), to.Year()
	if fromYear == toYear {
		return float64(to.Sub(from)) / daysPerYear(fromYear)
	}

	firstFraction := float64(date.New(fromYear+1, time.January, 1).Sub(from)) / daysPerYear(fromYear)
	lastFraction := float64(to.Sub(date.New(toYear, time.January, 1))) / daysPerYear(toYear)

	return firstFraction + lastFraction + float64(toYear-fromYear-1)
}

func yearFractionActualActualAFB(from, to date.Date) float64 {
	nbFullYears := 0

	remaining := to

	// This function uses backward induction from the later date,
	// 'from' and 'to' are swapped if 'to' falls before 'from'.
	sign := 1.0
	if to.Before(from) {
		sign = -1.0

		remaining = from
		from, to = to, from
	}

	for tmp := to; tmp.After(from); {
		tmp = tmp.AddDate(-1, 0, 0)
		year, month, day := tmp.Date()

		if day == 28 && month == time.February && isLeapYear(year) {
			tmp = tmp.Add(1)
		}

		if !tmp.Before(from) {
			nbFullYears++
			remaining = tmp
		}
	}

	return sign * (float64(nbFullYears) + float64(remaining.Sub(from))/computeYearDurationAFB(from, remaining))
}

func computeYearDurationAFB(from, remaining date.Date) float64 {
	remainingYear := remaining.Year()
	if isLeapYear(remainingYear) {
		date := date.New(remainingYear, time.February, 29)
		if remaining.After(date) && !from.After(date) {
			return threeSixtySixDays
		}
	}

	fromYear := from.Year()
	if isLeapYear(fromYear) {
		date := date.New(fromYear, time.February, 29)
		if remaining.After(date) && !from.After(date) {
			return threeSixtySixDays
		}
	}

	return threeSixtyFiveDays
}

func yearFractionActualThreeSixty(from, to date.Date) float64 {
	return float64(to.Sub(from)) / threeSixtyDays
}

func yearFractionActualThreeSixtyFiveFixed(from, to date.Date) float64 {
	return float64(to.Sub(from)) / threeSixtyFiveDays
}

func yearFractionThirtyThreeSixtyUS(from, to date.Date) float64 {
	if to.Day() == 31 && from.Day() < 30 {
		to = to.Add(1)
	}

	return yearFractionThirtyThreeSixty(from, to, 0.0)
}

func yearFractionThirtyThreeSixtyEuropean(from, to date.Date) float64 {
	return yearFractionThirtyThreeSixty(from, to, 0.0)
}

func yearFractionThirtyThreeSixtyItalian(from, to date.Date) float64 {
	shift := func(d date.Date) int {
		_, month, day := d.Date()
		if month == time.February && day > 27 {
			return 30 - day
		}

		return 0
	}

	dayShift := shift(from) + shift(to)

	return yearFractionThirtyThreeSixty(from, to, dayShift)
}

func yearFractionThirtyThreeSixtyGerman(from, to date.Date) float64 {
	shift := func(d date.Date) int {
		tmp := d.Add(1)
		_, month, day := tmp.Date()
		if month == time.March && day == 1 {
			return 1
		}

		return 0
	}

	dayShift := shift(from) + shift(to)

	return yearFractionThirtyThreeSixty(from, to, dayShift)
}

func yearFractionThirtyThreeSixty(from, to date.Date, dayShift int) float64 {
	fromYear, fromMonth, fromDay := from.Date()
	toYear, toMonth, toDay := to.Date()

	yearDiff := float64(360 * (toYear - fromYear))
	monthDiff := float64(30 * (toMonth - fromMonth - 1))
	dayDiff := math.Max(0, float64(30-fromDay)) + math.Min(30, float64(toDay))

	return (yearDiff + monthDiff + dayDiff + float64(dayShift)) / threeSixtyDays
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func daysPerYear(year int) float64 {
	if isLeapYear(year) {
		return threeSixtySixDays
	}

	return threeSixtyFiveDays
}
