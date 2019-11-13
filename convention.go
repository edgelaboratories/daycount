package daycount

// Convention is the daycounting convention
type Convention int

const (
	// ActualActual is an actual/actual convention
	// This convention splits up the actual number of days falling in leap years and in non-leap years.
	// The year fraction is the sum of the actual number of days falling in leap years divided by 366 and the
	// actual number of days falling in non-leap years divided by 365.
	ActualActual Convention = iota
	// ActualActualISDA is the actual/actual ISDA convention
	ActualActualISDA
	// ActualActualAFB is the actual/actual AFB convention
	// This method first calculates the number of full years counting backwards from the second date.
	// For any resulting stub periods, the numerator is the actual number of days in the period, the denominator
	// being 365 or 366 depending on whether February 29th falls in the stub period.
	ActualActualAFB
	// ActualThreeSixty is the actual/360 convention
	// The actual number of days between two dates is used as the numerator.
	// The denominator is always 360 days.
	ActualThreeSixty
	// ActualThreeSixtyFiveFixed is the actual/365 fixed convention
	// The numerator is the actual number of days between the two dates.
	// The denominator is always 365 days.
	ActualThreeSixtyFiveFixed
	// ThirtyThreeSixtyUSNASD is the 30/360 US (NASD) convention
	// If the first date falls on the 31st, it is changed to the 30th.
	// If the second date falls on the 31st and the first date is earlier than the 30th, then the second date is
	// changed to the 1st of the next month, otherwise it is changed to the 30th.
	ThirtyThreeSixtyUSNASD
	// ThirtyThreeSixtyEuropean is the 30/360 European convention
	// If the first date falls on the 31st, it is changed to the 30th.
	// If the second date falls on the 31th, it is changed to the 30th.
	ThirtyThreeSixtyEuropean
	// ThirtyThreeSixtyItalian is the 30/360 Italian convention
	// If the first date falls on the 31st or if it is February 28th or 29th, then it is changed to the 30th.
	// If the second date falls on the 31st or if it is February 28th or 29th, then it is changed to the 30th.
	ThirtyThreeSixtyItalian
	// ThirtyThreeSixtyGerman is the 30/360 German convention
	// If the first date falls on the 31st or if it is the last day of February, then it is changed to the 30th.
	// If the second date falls on the 31st or if it is the last day of February, then it is changed to the 30th.
	ThirtyThreeSixtyGerman
)

// String returns the convention name
func (d Convention) String() string {
	return [...]string{
		"ActualActual",
		"ActualActualISDA",
		"ActualActualAFB",
		"ActualThreeSixty",
		"ActualThreeSixtyFiveFixed",
		"ThirtyThreeSixtyUSNASD",
		"ThirtyThreeSixtyEuropean",
		"ThirtyThreeSixtyItalian",
		"ThirtyThreeSixtyGerman",
	}[d]
}
