package daycount

import (
	"encoding/json"
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

func Test_Convention_UnmarshalJSON(t *testing.T) {
	type data struct {
		Convention Convention `json:"convention"`
	}
	b := []byte(`{
		"convention": "ActualActual"
	}`)
	var d data
	assert.NoError(t, json.Unmarshal(b, &d))
	assert.Equal(t, ActualActual, d.Convention)
}

func Test_Convention_UnmarshalJSON_InvalidInput(t *testing.T) {
	type data struct {
		Convention Convention `json:"convention"`
	}
	b := []byte(`{
		"convention": 0.01
	}`)
	var d data
	assert.Error(t, json.Unmarshal(b, &d))
}

func Test_Convention_UnmarshalJSON_UnrecognizedConvention(t *testing.T) {
	type data struct {
		Convention Convention `json:"convention"`
	}
	b := []byte(`{
		"convention": "SomeInvalidConvention"
	}`)
	var d data
	assert.Error(t, json.Unmarshal(b, &d))
}
