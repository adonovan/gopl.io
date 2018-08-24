/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package monthdays

import (
	"testing"
	"time"
)

const shortForm = "2006-Jan-02"

var tests = map[string]int{
	"2018-Jan-01": 31,
	"2018-Feb-01": 28,
	"2020-Feb-01": 29, // next leap day
	"2018-Mar-01": 31,
	"2018-Apr-01": 30,
	"2018-May-01": 31,
	"2018-Jun-01": 30,
	"2018-Jul-01": 31,
	"2018-Aug-01": 31,
	"2018-Sep-01": 30,
	"2018-Oct-01": 31,
	"2018-Nov-01": 30,
	"2018-Dec-01": 31,
}

func TestDaysInMonth(t *testing.T) {
	for dateString, want := range tests {
		date, err := time.Parse(shortForm, dateString)
		if err != nil {
			t.Errorf("While parsing %v got %s", dateString, err)
		}
		got := DaysInMonth(date)
		if got != want {
			t.Errorf("DaysInMonth: for %v want %d got %d", date, want, got)
		}
	}
}
