/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package monthdays

import "time"

// DaysInMonth counts the days in the month of the given date. It correctly
// handles leap years.
func DaysInMonth(t time.Time) int {
	curMonth := t.Month()
	nextMonth := curMonth + 1
	if nextMonth > 12 {
		nextMonth = time.January
	}
	// note about Date(): by setting the day to zero will get normalized to the
	// day before the 1st of this month.
	testDate := time.Date(t.Year(), nextMonth,
		0, // 0th day of next month -> last day of this month
		0, 0, 0, 0, time.UTC,
	)
	return testDate.Day()
}
