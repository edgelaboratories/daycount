package daycount

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Convention_String(t *testing.T) {
	t.Parallel()

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
		{
			"Unsupported",
			Convention(666),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.name, tc.convention.String())
		})
	}
}

func Test_Parse(t *testing.T) {
	t.Parallel()

	for _, convention := range []Convention{
		ActualActual,
		ActualActualAFB,
		ActualThreeSixty,
		ActualThreeSixtyFiveFixed,
		ThirtyThreeSixtyUS,
		ThirtyThreeSixtyEuropean,
		ThirtyThreeSixtyItalian,
		ThirtyThreeSixtyGerman,
	} {
		output, err := Parse(convention.String())
		require.NoError(t, err)
		assert.Equal(t, convention, output)
	}
}

func Test_Parse_UnknownConvention(t *testing.T) {
	t.Parallel()

	_, err := Parse("UnknownConvention")
	require.Error(t, err)
}

func Test_Convention_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	var d struct {
		Convention Convention `json:"convention"`
	}
	require.NoError(t, json.Unmarshal([]byte(`{"convention": "ActualActual"}`), &d))
	assert.Equal(t, ActualActual, d.Convention)
}

func Test_Convention_UnmarshalJSON_Invalid(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		input string
	}{
		{
			"invalid input",
			`{"convention": 0.01}`,
		},
		{
			"unrecognized convention",
			`{"convention": "SomeInvalidConvention"}`,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var d struct {
				Convention Convention `json:"convention"`
			}
			require.Error(t, json.Unmarshal([]byte(tc.input), &d))
		})
	}
}

func Test_Convention_MarshalJSON(t *testing.T) {
	t.Parallel()

	output, err := json.Marshal(struct {
		Convention Convention `json:"convention"`
	}{
		Convention: ThirtyThreeSixtyEuropean,
	})
	require.NoError(t, err)
	assert.Equal(t, []byte(`{"convention":"ThirtyThreeSixtyEuropean"}`), output)
}
