package daycount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConventionString(t *testing.T) {
	testCases := []struct {
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
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.name, tc.convention.String())
		})
	}
}

func Test_Parse(t *testing.T) {
	testConventions := []Convention{
		ActualActual,
		ActualActualAFB,
		ActualThreeSixty,
		ActualThreeSixtyFiveFixed,
		ThirtyThreeSixtyUS,
		ThirtyThreeSixtyEuropean,
		ThirtyThreeSixtyItalian,
		ThirtyThreeSixtyGerman,
	}
	for _, convention := range testConventions {
		output, err := Parse(convention.String())
		assert.NoError(t, err)
		assert.Equal(t, convention, output)
	}
}

func Test_Parse_UnknownConvention(t *testing.T) {
	_, err := Parse("UnknownConvention")
	assert.Error(t, err)
}
