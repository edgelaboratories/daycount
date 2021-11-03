package daycount

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

const epsilon = 1.0e-6

func Test_YearFraction(t *testing.T) {
	t.Parallel()

	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	dayDiff := 211.0
	thirtyThreeSixtyDiff := (threeSixtyDays*0.0 + 30.0*5.0 + 29.0 + 30.0) / threeSixtyDays

	for _, tc := range []struct {
		name       string
		convention Convention
		expected   float64
	}{
		{
			"actual actual",
			ActualActual,
			dayDiff / threeSixtyFiveDays,
		},
		{
			"actual actual afb",
			ActualActualAFB,
			dayDiff / threeSixtyFiveDays,
		},
		{
			"actual 360",
			ActualThreeSixty,
			dayDiff / threeSixtyDays,
		},
		{
			"actual 365 fixed",
			ActualThreeSixtyFiveFixed,
			dayDiff / threeSixtyFiveDays,
		},
		{
			"30 360 us",
			ThirtyThreeSixtyUS,
			(30.0*5.0 + 29.0 + 30.0 + 1.0) / threeSixtyDays,
		},
		{
			"30 360 european",
			ThirtyThreeSixtyEuropean,
			thirtyThreeSixtyDiff,
		},
		{
			"30 360 italian",
			ThirtyThreeSixtyItalian,
			thirtyThreeSixtyDiff,
		},
		{
			"30 360 german",
			ThirtyThreeSixtyGerman,
			thirtyThreeSixtyDiff,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InEpsilon(t, tc.expected, YearFraction(from, to, tc.convention), epsilon)
		})
	}
}

func Test_YearFraction_ConsistencyWithDayCounter(t *testing.T) {
	t.Parallel()

	// This test verifies that the convenience function YearFraction
	// is consistent with the closure implemented by a DayCounter
	// on the same convention.

	const testDates = 100

	origin := date.New(2021, time.October, 1)

	for _, tc := range []struct {
		name       string
		convention Convention
	}{
		{
			"ActualActual",
			ActualActual,
		},
		{
			"ActualActualAFB",
			ActualActualAFB,
		},
		{
			"ActualThreeSixty",
			ActualThreeSixty,
		},
		{
			"ActualThreeSixtyFiveFixed",
			ActualThreeSixtyFiveFixed,
		},
		{
			"ThirtyThreeSixtyUS",
			ThirtyThreeSixtyUS,
		},
		{
			"ThirtyThreeSixtyEuropean",
			ThirtyThreeSixtyEuropean,
		},
		{
			"ThirtyThreeSixtyItalian",
			ThirtyThreeSixtyItalian,
		},
		{
			"ThirtyThreeSixtyGerman",
			ThirtyThreeSixtyGerman,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Fuzz test dates by shifting the origin by a random duration.
			f := fuzz.New().Funcs(
				func(d *date.Date, c fuzz.Continue) {
					*d = date.NewAt(origin.UTC().Add(time.Duration(c.Int63())))
				},
				func(d *time.Duration, c fuzz.Continue) {
					*d = time.Duration(c.Int63())
				},
			)

			from, lag := date.Date{}, time.Duration(0)
			for i := 0; i < testDates; i++ {
				f.Fuzz(&from)
				f.Fuzz(&lag)

				to := date.NewAt(from.UTC().Add(lag))

				assert.Equal(t,
					NewDayCounter(tc.convention)(from, to),
					YearFraction(from, to, tc.convention),
				)
			}
		})
	}
}

func Test_YearFraction_DefaultConvention(t *testing.T) {
	t.Parallel()

	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	expected := 211.0 / threeSixtyFiveDays
	assert.InEpsilon(t, expected, YearFraction(from, to, Convention(-1)), epsilon)
}

func Test_YearFraction_EqualDates(t *testing.T) {
	t.Parallel()

	from := date.New(2018, time.January, 1)
	to := from
	assert.Equal(t, 0.0, YearFraction(from, to, ActualActual))
}

func Test_YearFraction_InvertedDates(t *testing.T) {
	t.Parallel()

	from := date.New(2018, time.July, 31)
	to := date.New(2018, time.January, 1)
	expected := -211.0 / threeSixtyFiveDays
	assert.InEpsilon(t, expected, YearFraction(from, to, ActualActual), epsilon)
}

func Test_yearFractionActualActual(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
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
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionActualActual(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionActualActualAFB(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
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
		{
			"from 1996.01.01 to 1997.01.01",
			date.New(1996, time.February, 1),
			date.New(1997, time.January, 1),
			(366.0 - 31.0) / 366.0,
		},
		{
			"from 2004.02.28 to 2004.03.02",
			date.New(2004, time.February, 28),
			date.New(2004, time.March, 2),
			3.0 / 366.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionActualActualAFB(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionActualThreeSixty(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			62.0 / 360.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			63.0 / 360.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			396.0 / 360.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			485.0 / 360.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionActualThreeSixty(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionActualThreeSixtyFiveFixed(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
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
			396.0 / 365.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			485.0 / 365.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionActualThreeSixtyFiveFixed(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionThirtyThreeSixtyUS(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			60.0 / 360.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			61.0 / 360.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			390.0 / 360.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			480.0 / 360.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionThirtyThreeSixtyUS(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionThirtyThreeSixtyEuropean(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			60.0 / 360.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			61.0 / 360.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			390.0 / 360.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			479.0 / 360.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionThirtyThreeSixtyEuropean(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionThirtyThreeSixtyItalian(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			62.0 / 360.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			62.0 / 360.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			390.0 / 360.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			479.0 / 360.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionThirtyThreeSixtyItalian(tc.from, tc.to), epsilon)
		})
	}
}

func Test_yearFractionThirtyThreeSixtyGerman(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		from     date.Date
		to       date.Date
		expected float64
	}{
		{
			"from 2007.12.28 to 2008.02.28",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			60.0 / 360.0,
		},
		{
			"from 2007.12.28 to 2008.02.29",
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			62.0 / 360.0,
		},
		{
			"from 2007.10.31 to 2008.11.30",
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			390.0 / 360.0,
		},
		{
			"from 2007.02.01 to 2008.05.31",
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			479.0 / 360.0,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, yearFractionThirtyThreeSixtyGerman(tc.from, tc.to), epsilon)
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
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
	} {
		assert.Equal(t, tc.expected, isLeapYear(tc.year), tc.year)
	}
}

func Test_daysPerYear(t *testing.T) {
	t.Parallel()

	assert.Equal(t, threeSixtyFiveDays, daysPerYear(2015))
	assert.Equal(t, threeSixtySixDays, daysPerYear(2000))
}
