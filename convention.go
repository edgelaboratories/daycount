package daycount

import (
	"encoding/json"
	"fmt"
)

// Convention is the daycounting convention.
type Convention int

const (
	// ActualActual is commonly used for all sterling bonds, Euro denominated bonds, US Treasury bonds and
	// for some USD interest rate swaps.
	// In this case, the day-count fraction is the number of days in the period in a normal year over 365
	// or the number of days in the period in a leap year over 366.
	ActualActual Convention = iota
	// ActualActualAFB is the actual/actual AFB convention.
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
	// If the second date falls on the 31st and the first date is earlier than the 30th, then the second date is
	// changed to the 1st of the next month.
	ThirtyThreeSixtyUS
	// ThirtyThreeSixtyEuropean is used for calculating accrued interest on some legacy currency pre Euro
	// Eurobonds and on bonds in Sweden and Switzerland.
	// This method assumes that all months have 30 days, even February, and that a year is 360 days.
	// Effectively if the start date d1 is 31 then it changes to 30, and if the second date d2 is 31 it too changes to 30.
	ThirtyThreeSixtyEuropean
	// ThirtyThreeSixtyItalian is the 30/360 Italian convention.
	// If the first date is February 28th or 29th, then it is changed to the 30th.
	// If the second date is February 28th or 29th, then it is changed to the 30th.
	ThirtyThreeSixtyItalian
	// ThirtyThreeSixtyGerman is the 30/360 German convention.
	// If the first date is the last day of February, then it is changed to the 1st of March.
	// If the second date is the last day of February, then it is changed to the 1st of March.
	ThirtyThreeSixtyGerman

	// outOfRangeConvention is a sentinel value that allows to bound
	// the range of allowed conventions. Add new conventions before it.
	outOfRangeConvention
)

// String returns the convention name.
func (d Convention) String() string {
	if d < ActualActual || d >= outOfRangeConvention {
		return "Unsupported"
	}

	return [...]string{
		"ActualActual",
		"ActualActualAFB",
		"ActualThreeSixty",
		"ActualThreeSixtyFiveFixed",
		"ThirtyThreeSixtyUS",
		"ThirtyThreeSixtyEuropean",
		"ThirtyThreeSixtyItalian",
		"ThirtyThreeSixtyGerman",
	}[d]
}

// Parse maps an input string to a daycount convention.
func Parse(convention string) (Convention, error) {
	switch convention {
	case "ActualActual":
		return ActualActual, nil

	case "ActualActualAFB":
		return ActualActualAFB, nil

	case "ActualThreeSixty":
		return ActualThreeSixty, nil

	case "ActualThreeSixtyFiveFixed":
		return ActualThreeSixtyFiveFixed, nil

	case "ThirtyThreeSixtyUS":
		return ThirtyThreeSixtyUS, nil

	case "ThirtyThreeSixtyEuropean":
		return ThirtyThreeSixtyEuropean, nil

	case "ThirtyThreeSixtyItalian":
		return ThirtyThreeSixtyItalian, nil

	case "ThirtyThreeSixtyGerman":
		return ThirtyThreeSixtyGerman, nil

	default:
		return -1, fmt.Errorf("unrecognized daycount convention %s", convention)
	}
}

// UnmarshalJSON implements the JSON unmarshaler.
func (d *Convention) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("could not unmarshal convention: %w", err)
	}

	res, err := Parse(s)
	if err != nil {
		return err
	}
	*d = res

	return nil
}

// MarshalJSON implements the JSON marshaler.
func (d Convention) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}
