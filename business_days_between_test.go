package daycount

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_BusinessDaysBetween(t *testing.T) {
	for _, tc := range []struct {
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			7,
		},
		{
			time.Date(2019, time.February, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2019, time.March, 2, 0, 0, 0, 0, time.UTC),
			3,
		},
		{
			time.Date(2020, time.February, 26, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.March, 2, 0, 0, 0, 0, time.UTC),
			3,
		},
		{
			time.Date(2020, time.September, 3, 18, 0, 0, 0, time.UTC),
			time.Date(2020, time.September, 4, 0, 0, 0, 0, time.UTC),
			0,
		},
	} {
		assert.Equal(t, tc.expected, BusinessDaysBetween(tc.start, tc.end))
	}
}
