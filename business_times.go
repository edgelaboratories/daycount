package daycount

import "time"

func sub(startTime, endTime time.Time) int {
	return int(endTime.Sub(startTime).Hours() / 24.0)
}

// BusinessTimes returns the slice of business days between start and end input dates.
// Output times are sorted in chronological order.
func BusinessTimes(startTime, endTime time.Time) []time.Time {
	businessTimes := make([]time.Time, 0, sub(startTime, endTime))
	for currentTime := startTime; !currentTime.After(endTime); currentTime = currentTime.Add(24 * time.Hour) {
		if noHolidayCalendar.IsWorkday(currentTime) {
			businessTimes = append(businessTimes, currentTime)
		}
	}
	return businessTimes
}
