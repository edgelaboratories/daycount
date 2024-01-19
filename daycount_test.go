package daycount

import (
	"fmt"
	"testing"
	"time"

	"github.com/edgelaboratories/date"
	fuzz "github.com/google/gofuzz"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
)

const epsilon = 1.0e-15

func Test_YearFraction(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		from     date.Date
		to       date.Date
		expected map[Convention]float64
	}{
		{
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 28),
			map[Convention]float64{
				ActualActualAFB:           62.0 / 365.0,
				ActualThreeSixty:          62.0 / 360.0,
				ActualThreeSixtyFiveFixed: 62.0 / 365.0,
				ThirtyThreeSixtyUS:        60.0 / 360.0,
				ThirtyThreeSixtyEuropean:  60.0 / 360.0,
				ThirtyThreeSixtyItalian:   62.0 / 360.0,
				ThirtyThreeSixtyGerman:    60.0 / 360.0,
			},
		},
		{
			date.New(2007, time.December, 28),
			date.New(2008, time.February, 29),
			map[Convention]float64{
				ActualActualAFB:           63.0 / 365.0,
				ActualThreeSixty:          63.0 / 360.0,
				ActualThreeSixtyFiveFixed: 63.0 / 365.0,
				ThirtyThreeSixtyUS:        61.0 / 360.0,
				ThirtyThreeSixtyEuropean:  61.0 / 360.0,
				ThirtyThreeSixtyItalian:   62.0 / 360.0,
				ThirtyThreeSixtyGerman:    62.0 / 360.0,
			},
		},
		{
			date.New(2007, time.October, 31),
			date.New(2008, time.November, 30),
			map[Convention]float64{
				ActualActualAFB:           1.0 + 30.0/365.0,
				ActualThreeSixty:          396.0 / 360.0,
				ActualThreeSixtyFiveFixed: 396.0 / 365.0,
				ThirtyThreeSixtyUS:        390.0 / 360.0,
				ThirtyThreeSixtyEuropean:  390.0 / 360.0,
				ThirtyThreeSixtyItalian:   390.0 / 360.0,
				ThirtyThreeSixtyGerman:    390.0 / 360.0,
			},
		},
		{
			date.New(2007, time.February, 1),
			date.New(2008, time.May, 31),
			map[Convention]float64{
				ActualThreeSixty:          485.0 / 360.0,
				ActualThreeSixtyFiveFixed: 485.0 / 365.0,
				ThirtyThreeSixtyUS:        480.0 / 360.0,
				ThirtyThreeSixtyEuropean:  479.0 / 360.0,
				ThirtyThreeSixtyItalian:   479.0 / 360.0,
				ThirtyThreeSixtyGerman:    479.0 / 360.0,
			},
		},
		{
			date.New(2008, time.February, 1),
			date.New(2009, time.May, 31),
			map[Convention]float64{
				ActualActualAFB: 1.0 + 120.0/366.0,
			},
		},
		{
			date.New(2012, time.December, 28),
			date.New(2013, time.February, 28),
			map[Convention]float64{
				ActualActualAFB: 62.0 / 365.0,
			},
		},
		{
			date.New(2012, time.February, 28),
			date.New(2015, time.January, 28),
			map[Convention]float64{
				ActualActualAFB: 2.0 + 335.0/366.0,
			},
		},
		{
			date.New(1996, time.February, 1),
			date.New(1997, time.January, 1),
			map[Convention]float64{
				ActualActualAFB: (366.0 - 31.0) / 366.0,
			},
		},
		{
			date.New(2004, time.February, 28),
			date.New(2004, time.March, 2),
			map[Convention]float64{
				ActualActualAFB: 3.0 / 366.0,
			},
		},
		{
			date.New(2019, time.July, 1),
			date.New(2019, time.January, 1),
			map[Convention]float64{
				ActualActual: -181.0 / 365.0,
			},
		},
		{
			date.New(2019, time.January, 1),
			date.New(2019, time.July, 1),
			map[Convention]float64{
				ActualActual: 181.0 / 365.0,
			},
		},
		{
			date.New(2019, time.January, 1),
			date.New(2020, time.January, 1),
			map[Convention]float64{
				ActualActual: 1.0,
			},
		},
		{
			date.New(2020, time.January, 1),
			date.New(2021, time.January, 1),
			map[Convention]float64{
				ActualActual: 1.0,
			},
		},
		{
			date.New(2019, time.January, 1),
			date.New(2021, time.January, 1),
			map[Convention]float64{
				ActualActual: 2.0,
			},
		},
		{
			date.New(2019, time.March, 4),
			date.New(2023, time.June, 1),
			map[Convention]float64{
				ActualActual: 303.0/365.0 + 3.0 + 151.0/365.0,
			},
		},
		{
			date.New(2020, time.February, 10),
			date.New(2021, time.July, 2),
			map[Convention]float64{
				ActualActual: 326.0/366.0 + 182.0/365.0,
			},
		},
		{
			date.New(2016, time.March, 4),
			date.New(2023, time.June, 1),
			map[Convention]float64{
				ActualActual: 303.0/366.0 + 6.0 + 151.0/365.0,
			},
		},
		{
			date.New(2016, time.March, 4),
			date.New(2116, time.March, 4),
			map[Convention]float64{
				ActualActual: 100.0,
			},
		},
	} {
		tc := tc

		for convention, expected := range tc.expected {
			convention, expected := convention, expected

			name := fmt.Sprintf("%s to %s/%s", tc.from, tc.to, convention)
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assert.InEpsilon(t, expected, YearFraction(tc.from, tc.to, convention), epsilon)
			})
		}
	}
}

func Test_YearFraction_outOfRangeConvention(t *testing.T) {
	t.Parallel()

	var (
		from = date.Today()
		to   = from.AddDate(1, 0, 0)
	)

	assert.InEpsilon(t,
		YearFraction(from, to, ActualActual),
		YearFraction(from, to, outOfRangeConvention),
		epsilon,
	)
}

func Test_YearFraction_EqualDates(t *testing.T) {
	t.Parallel()

	const tol = 1e-15

	var daysFromOrigin int
	fuzz.NewWithSeed(123).Funcs(
		func(days *int, c fuzz.Continue) {
			*days = -1800 + c.Intn(3600)
		},
	).Fuzz(&daysFromOrigin)

	from := date.New(2018, time.January, 1).Add(daysFromOrigin)

	for _, convention := range []Convention{
		ActualActual,
		ActualActualAFB,
		ActualThreeSixty,
		ActualThreeSixtyFiveFixed,
		ThirtyThreeSixtyEuropean,
		ThirtyThreeSixtyGerman,
		ThirtyThreeSixtyItalian,
		ThirtyThreeSixtyUS,
	} {
		convention := convention

		t.Run(convention.String(), func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, 0.0, YearFraction(from, from, ActualActual), tol)
		})
	}
}

func Test_YearFraction_InvertedDates(t *testing.T) {
	t.Parallel()

	from := date.New(2018, time.July, 31)
	to := date.New(2018, time.January, 1)
	expected := -211.0 / threeSixtyFiveDays
	assert.InEpsilon(t, expected, YearFraction(from, to, ActualActual), epsilon)
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

	const tol = 1e-15

	assert.InDelta(t, threeSixtyFiveDays, daysPerYear(2015), tol)
	assert.InDelta(t, threeSixtySixDays, daysPerYear(2000), tol)
}

func Test_YearFraction_inverted(t *testing.T) {
	t.Parallel()

	const tol = 1e-10

	props := gopter.NewProperties(gopter.DefaultTestParametersWithSeed(1234))

	initialTime := time.Date(1970, 1, 1, 12, 0, 0, 0, time.UTC)
	duration := time.Hour * 24 * 365 * 50

	props.Property("year fraction inversion gives opposite sign", prop.ForAll(
		func(t1, t2 time.Time, conventionInteger int) bool {
			date1 := date.NewAt(t1)
			date2 := date.NewAt(t2)

			yf := YearFraction(date1, date2, Convention(conventionInteger))
			invertedYf := YearFraction(date2, date1, Convention(conventionInteger))

			return assert.InDelta(t, yf, -1.0*invertedYf, tol)
		},
		gen.TimeRange(initialTime, duration),
		gen.TimeRange(initialTime, duration),
		gen.IntRange(0, 7),
	))

	props.TestingRun(t)
}
