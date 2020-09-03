package daycount

import "time"

// BusinessDaysBetween returns the number of business days
// between a starting and an ending time.
func BusinessDaysBetween(start, end time.Time) int {
	return len(BusinessTimes(start, end)) - 1
}
