package daycount

import (
	"math"
	"time"
)

// YearFractionDiff returns the year fraction difference between two dates
// according to the input convention.
// If the convention is not recognized, it defaults to ActualActual.
func YearFractionDiff(from, to time.Time, convention Convention) float64 {
	switch convention {
	case ActualActual:
		return 0.0 // yearFractionActualActual(from, to)
	case ActualActualISDA:
		return 0.0
	case ActualActualAFB:
		return 0.0
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
		return 0.0 // yearFractionActualActual(from, to)
	}
}

const (
	threeSixtyDays     = 360.0
	threeSixtyFiveDays = 365.0
	threeSixtySixDays  = 366.0
	hoursPerDay        = 24.0
)

// func yearFractionActualActual(from, to time.Time) float64 {
// 	fromYear, toYear := from.Year(), to.Year()
// 	if fromYear == toYear {
// 		return float64(to.Sub(from)) / daysPerYear(fromYear)
// 	}
// 	firstFraction := float64(date.New(fromYear+1, time.January, 1).Sub(from)) / daysPerYear(fromYear)
// 	lastFraction := float64(to.Sub(date.New(toYear, time.January, 1))) / daysPerYear(toYear)
// 	return firstFraction + lastFraction + float64(toYear-fromYear-1)
// }

func yearFractionActualThreeSixty(from, to time.Time) float64 {
	return to.Sub(from).Hours() / (threeSixtyDays * hoursPerDay)
}

func yearFractionThirtyThreeSixtyUS(from, to time.Time) float64 {
	yearDiff := float64(360 * (to.Year() - from.Year()))
	monthDiff := float64(30 * (to.Month() - from.Month() - 1))
	dayDiff := math.Max(0, float64(30-from.Day())) + math.Min(30, float64(to.Day()))
	// TODO: neglecting time granularity here!
	return (yearDiff + monthDiff + dayDiff) / threeSixtyDays
}

func yearFractionActualThreeSixtyFiveFixed(from, to time.Time) float64 {
	return to.Sub(from).Hours() / (threeSixtyFiveDays * hoursPerDay)
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
