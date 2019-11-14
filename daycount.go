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
		return 0.0
	case ThirtyThreeSixtyItalian:
		return 0.0
	case ThirtyThreeSixtyGerman:
		return 0.0
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
	if from == to {
		return 0.0
	}
	if from.After(to) {
		return -yearFractionActualActual(to, from)
	}
	fromYear, toYear := from.Year(), to.Year()
	if fromYear == toYear {
		return float64(to.Sub(from)) / daysPerYear(fromYear)
	}
	firstFraction := float64(date.New(fromYear+1, time.January, 1).Sub(from)) / daysPerYear(fromYear)
	lastFraction := float64(to.Sub(date.New(toYear, time.January, 1))) / daysPerYear(toYear)
	return firstFraction + lastFraction + float64(toYear-fromYear-1)
}

func yearFractionActualActualAFB(from, to date.Date) float64 {
	if from == to {
		return 0.0
	}
	if from.After(to) {
		return -yearFractionActualActualAFB(to, from)
	}
	nbFullYears := 0
	remaining, tmp := to, to
	for tmp.After(from) {
		tmp = tmp.AddDate(-1, 0, 0)
		if tmp.Day() == 28 && tmp.Month() == time.February && isLeapYear(tmp.Year()) {
			tmp = tmp.AddDate(0, 0, 1)
		}
		if !tmp.Before(from) {
			nbFullYears++
			remaining = tmp
		}
	}

	den := threeSixtyFiveDays
	if isLeapYear(remaining.Year()) {
		date := date.New(remaining.Year(), time.February, 29)
		if remaining.After(date) && !from.After(date) {
			den += 1.0
		}
	} else if isLeapYear(from.Year()) {
		date := date.New(from.Year(), time.February, 29)
		if remaining.After(date) && !from.After(date) {
			den += 1.0
		}
	}
	return float64(nbFullYears) + float64(remaining.Sub(from))/den
}

func yearFractionActualThreeSixty(from, to date.Date) float64 {
	if from == to {
		return 0.0
	}
	if from.After(to) {
		return -yearFractionActualThreeSixty(to, from)
	}
	return float64(to.Sub(from)) / threeSixtyDays
}

func yearFractionThirtyThreeSixtyUS(from, to date.Date) float64 {
	if from == to {
		return 0.0
	}
	if from.After(to) {
		return -yearFractionThirtyThreeSixtyUS(to, from)
	}
	yearDiff := float64(360 * (to.Year() - from.Year()))
	monthDiff := float64(30 * (to.Month() - from.Month() - 1))
	dayDiff := math.Max(0, float64(30-from.Day())) + math.Min(30, float64(to.Day()))
	return (yearDiff + monthDiff + dayDiff) / threeSixtyDays
}

func yearFractionActualThreeSixtyFiveFixed(from, to date.Date) float64 {
	if from == to {
		return 0.0
	}
	if from.After(to) {
		return -yearFractionActualThreeSixtyFiveFixed(to, from)
	}
	return float64(to.Sub(from)) / threeSixtyFiveDays
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
