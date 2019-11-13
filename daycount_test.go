package daycount

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const epsilon = 1.0e-6

func TestYearFractionDiff(t *testing.T) {
	from := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2018, time.July, 31, 0, 0, 0, 0, time.UTC)
	dayDiff := 211.0

	testCases := []struct {
		convention Convention
		expected   float64
	}{
		// {
		// 	ActualActual,
		// 	0.0,
		// },
		// {
		// 	ActualActualISDA,
		// 	0.0,
		// },
		// {
		// 	ActualActualAFB,
		// 	0.0,
		// },
		{
			ActualThreeSixty,
			dayDiff / threeSixtyDays,
		},
		{
			ActualThreeSixtyFiveFixed,
			dayDiff / threeSixtyFiveDays,
		},
		{
			ThirtyThreeSixtyUS,
			(threeSixtyDays*0.0 + 30.0*5.0 + 29.0 + 30.0) / threeSixtyDays,
		},
		// {
		// 	ThirtyThreeSixtyEuropean,
		// 	0.0,
		// },
		// {
		// 	ThirtyThreeSixtyItalian,
		// 	0.0,
		// },
		// {
		// 	ThirtyThreeSixtyGerman,
		// 	0.0,
		// },
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.convention.String(), func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, YearFractionDiff(from, to, tc.convention), epsilon)
		})
	}
}

func TestYearFractionActualThreeSixty(t *testing.T) {
	from := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2018, time.July, 31, 0, 0, 0, 0, time.UTC)
	assert.InEpsilon(t, 211.0/360.0, yearFractionActualThreeSixty(from, to), epsilon)
}

func TestYearFractionActualThreeSixtyFiveFixed(t *testing.T) {
	from := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2018, time.July, 31, 0, 0, 0, 0, time.UTC)
	assert.InEpsilon(t, 211.0/365.0, yearFractionActualThreeSixtyFiveFixed(from, to), epsilon)
}

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		year     int
		expected bool
	}{
		{
			2012,
			true,
		},
		{
			2015,
			false,
		},
		{
			2016,
			true,
		},
		{
			2021,
			false,
		},
		{
			2100,
			false,
		},
		{
			2000,
			true,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, isLeapYear(tc.year), tc.year)
	}
}

func TestDaysPerYear(t *testing.T) {
	assert.Equal(t, threeSixtyFiveDays, daysPerYear(2015))
	assert.Equal(t, threeSixtySixDays, daysPerYear(2000))
}
