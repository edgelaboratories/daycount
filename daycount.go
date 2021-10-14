package daycount

import (
	"math"
	"time"

	"github.com/fxtlabs/date"
)

// YearFractionDiff returns the year fraction difference between two dates
// according to the input convention.
// If the convention is not recognized, it defaults to ActualActual.
func YearFractionDiff(from, to date.Date, convention Convention) float64 {
	if to.Equal(from) {
		return 0.0
	}

	if from.After(to) {
		return -YearFractionDiff(to, from, convention)
	}

	return yearFractionDiff(from, to, convention)
}

func yearFractionDiff(from, to date.Date, convention Convention) float64 {
	switch convention {
	case ActualActual:
		return yearFractionActualActual(from, to)
	case ActualActualAFB:
		return yearFractionActualActualAFB(from, to)
	case ActualThreeSixty:
		return yearFractionActualThreeSixty(from, to)
	case ActualThreeSixtyFiveFixed:
		return yearFractionActualThreeSixtyFiveFixed(from, to)
	case ThirtyThreeSixtyUS:
		return yearFractionThirtyThreeSixtyUS(from, to)
	case ThirtyThreeSixtyEuropean:
		return yearFractionThirtyThreeSixtyEuropean(from, to)
	case ThirtyThreeSixtyItalian:
		return yearFractionThirtyThreeSixtyItalian(from, to)
	case ThirtyThreeSixtyGerman:
		return yearFractionThirtyThreeSixtyGerman(from, to)
	default:
		return yearFractionActualActual(from, to)
	}
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
	for tmp := to; tmp.After(from); {
		tmp = tmp.AddDate(-1, 0, 0)

		if tmp.Day() == 28 && tmp.Month() == time.February && isLeapYear(tmp.Year()) {
			tmp = tmp.Add(1)
		}

		if !tmp.Before(from) {
			nbFullYears++
			remaining = tmp
		}
	}

	return float64(nbFullYears) + float64(remaining.Sub(from))/computeYearDurationAFB(from, remaining)
}

func computeYearDurationAFB(from, remaining date.Date) float64 {
	if isLeapYear(remaining.Year()) {
		date := date.New(remaining.Year(), time.February, 29)
		if remaining.After(date) && !from.After(date) {
			return threeSixtySixDays
		}
	}

	if isLeapYear(from.Year()) {
		date := date.New(from.Year(), time.February, 29)
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
		if d.Month() == time.February && d.Day() > 27 {
			return 30 - d.Day()
		}

		return 0
	}

	dayShift := shift(from) + shift(to)

	return yearFractionThirtyThreeSixty(from, to, dayShift)
}

func yearFractionThirtyThreeSixtyGerman(from, to date.Date) float64 {
	shift := func(d date.Date) int {
		if tmp := d.Add(1); tmp.Month() == time.March && tmp.Day() == 1 {
			return 1
		}

		return 0
	}

	dayShift := shift(from) + shift(to)

	return yearFractionThirtyThreeSixty(from, to, dayShift)
}

func yearFractionThirtyThreeSixty(from, to date.Date, dayShift int) float64 {
	yearDiff := float64(360 * (to.Year() - from.Year()))
	monthDiff := float64(30 * (to.Month() - from.Month() - 1))
	dayDiff := math.Max(0, float64(30-from.Day())) + math.Min(30, float64(to.Day()))

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
