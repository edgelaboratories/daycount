package daycount

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

const epsilon = 1.0e-6

func TestYearFractionDiff(t *testing.T) {
	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	dayDiff := 211.0
	thirtyThreeSixtyDiff := (threeSixtyDays*0.0 + 30.0*5.0 + 29.0 + 30.0) / threeSixtyDays

	testCases := []struct {
		convention Convention
		expected   float64
	}{
		{
			ActualActual,
			dayDiff / threeSixtyFiveDays,
		},
		{
			ActualActualAFB,
			dayDiff / threeSixtyFiveDays,
		},
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
			thirtyThreeSixtyDiff,
		},
		{
			ThirtyThreeSixtyEuropean,
			thirtyThreeSixtyDiff,
		},
		{
			ThirtyThreeSixtyItalian,
			thirtyThreeSixtyDiff,
		},
		{
			ThirtyThreeSixtyGerman,
			thirtyThreeSixtyDiff,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.convention.String(), func(t *testing.T) {
			assert.InEpsilon(t, tc.expected, YearFractionDiff(from, to, tc.convention), epsilon)
		})
	}
}

func TestYearFractionDiffDefaultConvention(t *testing.T) {
	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	expected := 211.0 / threeSixtyFiveDays
	assert.InEpsilon(t, expected, YearFractionDiff(from, to, Convention(-1)), epsilon)
}

func TestYearFractionActualActual(t *testing.T) {
	testCases := []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2019.01.01 to 2019.01.01",
			date.New(2019, time.January, 1),
			date.New(2019, time.January, 1),
			0.0,
		},
		{
			"from 2019.07.01 to 2019.01.01",
			date.New(2019, time.July, 1),
			date.New(2019, time.January, 1),
			-181.0 / 365.0,
		},
		{
			"from 2019.01.01 to 2019.07.01",
			date.New(2019, time.January, 1),
			date.New(2019, time.July, 1),
			181.0 / 365.0,
		},
		{
			"from 2019.01.01 to 2020.01.01",
			date.New(2019, time.January, 1),
			date.New(2020, time.January, 1),
			1.0,
		},
		{
			"from 2020.01.01 to 2021.01.01",
			date.New(2020, time.January, 1),
			date.New(2021, time.January, 1),
			1.0,
		},
		{
			"from 2019.01.01 to 2021.01.01",
			date.New(2019, time.January, 1),
			date.New(2021, time.January, 1),
			2.0,
		},
		{
			"from 2019.03.04 to 2023.06.01",
			date.New(2019, time.March, 4),
			date.New(2023, time.June, 1),
			303.0/365.0 + 3.0 + 151.0/365.0,
		},
		{
			"from 2020.02.10 to 2021.07.02",
			date.New(2020, time.February, 10),
			date.New(2021, time.July, 2),
			326.0/366.0 + 182.0/365.0,
		},
		{
			"from 2016.03.04 to 2023.06.01",
			date.New(2016, time.March, 4),
			date.New(2023, time.June, 1),
			303.0/366.0 + 6.0 + 151.0/365.0,
		},
		{
			"from 2016.03.04 to 2116.03.04",
			date.New(2016, time.March, 4),
			date.New(2116, time.March, 4),
			100.0,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, tc.expected, yearFractionActualActual(tc.from, tc.to), epsilon)
		})
	}
}

func TestYearFractionActualActualAFB(t *testing.T) {
	testCases := []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2019.01.01 to 2019.01.01",
			date.New(2019, time.January, 1),
			date.New(2019, time.January, 1),
			0.0,
		},
		{
			"from 2008.02.28 to 2007.12.28",
			date.New(2008, time.February, 28),
			date.New(2007, time.December, 28),
			-62.0 / 365.0,
		},
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			62.0 / 365.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			63.0 / 365.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			1.0 + 30.0/365.0,
		},
		{
			"from 2008.02.01 to 2009.05.31",
			date.New(2008, time.February, 1),
			date.New(2009, time.May, 31),
			1.0 + 120.0/366.0,
		},
		{
			"from 2012.12.28 to 2013.02.28",
			date.New(2012, time.December, 28),
			date.New(2013, time.February, 28),
			62.0 / 365.0,
		},
		{
			"from 2012.02.28 to 2015.01.28",
			date.New(2012, time.February, 28),
			date.New(2015, time.January, 28),
			2.0 + 335.0/366.0,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, tc.expected, yearFractionActualActualAFB(tc.from, tc.to), epsilon)
		})
	}
}

func TestYearFractionActualThreeSixty(t *testing.T) {
	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	assert.InEpsilon(t, 211.0/360.0, yearFractionActualThreeSixty(from, to), epsilon)
}

func TestYearFractionActualThreeSixtyFiveFixed(t *testing.T) {
	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
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