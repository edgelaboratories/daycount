package daycount

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_sub(t *testing.T) {
	testCases := []struct {
		startTime time.Time
		endTime   time.Time
		expected  int
	}{
		{
			time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			365,
		},
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			366,
		},
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
			31,
		},
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.February, 1, 12, 0, 0, 0, time.UTC),
			31,
		},
		{
			time.Date(2020, time.January, 1, 6, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC),
			0,
		},
		{
			time.Date(2020, time.January, 1, 6, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 2, 18, 0, 0, 0, time.UTC),
			1,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, sub(tc.startTime, tc.endTime))
	}
}

func Test_BusinessTimes(t *testing.T) {
	testCases := []struct {
		startTime time.Time
		endTime   time.Time
		expected  []time.Time
	}{
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			[]time.Time{
				time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 6, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 9, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			time.Date(2020, time.April, 9, 12, 0, 0, 0, time.UTC),
			time.Date(2020, time.April, 17, 18, 0, 0, 0, time.UTC),
			[]time.Time{
				time.Date(2020, time.April, 9, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 10, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 13, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 14, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 16, 12, 0, 0, 0, time.UTC),
				time.Date(2020, time.April, 17, 12, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, BusinessTimes(tc.startTime, tc.endTime))
	}
}
