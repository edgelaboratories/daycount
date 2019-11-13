package daycount

// Convention is the daycounting convention
type Convention int

const (
	// ActualActual is commonly used for all sterling bonds, Euro denominated bonds, US Treasury bonds and
	// for some USD interest rate swaps.
	// In this case, the day-count fraction is the number of days in the period in a normal year over 365
	// or the number of days in the period in a leap year over 366.
	ActualActual Convention = iota
	// ActualActualISDA accounts for days in the period based on the portion in a leap year and the portion
	// in a non-leap year.
	// The year fraction is the sum of the actual number of days falling in leap years divided by 366 and the
	// actual number of days falling in non-leap years divided by 365.
	ActualActualISDA
	// ActualActualAFB is the actual/actual AFB convention
	// This method first calculates the number of full years counting backwards from the second date.
	// For any resulting stub periods, the numerator is the actual number of days in the period, the denominator
	// being 365 or 366 depending on whether February 29th falls in the stub period.
	ActualActualAFB
	// ActualThreeSixty is commonly used for all Eurocurrency LIBOR rates, except sterling.
	// The day count fraction is defined as the actual number of days in the period over 360.
	ActualThreeSixty
	// ActualThreeSixtyFiveFixed is commonly used for all sterling interest rates, including LIBOR.
	// The day count fraction is defined as the actual number of days in the period over 365.
	// It is also used for money markets in Australia, Canada and New Zealand.
	ActualThreeSixtyFiveFixed
	// ThirtyThreeSixtyUS is commonly used for corporate bonds, municipal bonds, and agency bonds in the U.S.
	// If the first date falls on the 31st, it is changed to the 30th.
	// If the second date falls on the 31st and the first date is earlier than the 30th, then the second date is
	// changed to the 1st of the next month, otherwise it is changed to the 30th.
	ThirtyThreeSixtyUS
	// ThirtyThreeSixtyEuropean is used for calculating accrued interest on some legacy currency pre Euro
	// Eurobonds and on bonds in Sweden and Switzerland.
	// This method assumes that all months have 30 days, even February, and that a year is 360 days.
	// Effectively if the start date d1 is 31 then it changes to 30, and if the second date d2 is 31 it too changes to 30.
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
		"ThirtyThreeSixtyUS",
		"ThirtyThreeSixtyEuropean",
		"ThirtyThreeSixtyItalian",
		"ThirtyThreeSixtyGerman",
	}[d]
}
