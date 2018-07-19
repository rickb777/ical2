package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"strings"
	"testing"
	"time"
)

func TestRecur(t *testing.T) {
	ics.MaxLineLength = 120
	dec24 := time.Date(1997, time.Month(12), 24, 0, 0, 0, 0, time.UTC)
	jan31 := time.Date(2000, time.Month(1), 31, 14, 0, 0, 0, time.UTC)

	// These test cases are all from RFC5545.

	// Daily for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.Count = 10
		return rv
	}, "FREQ=DAILY;COUNT=10")

	// Daily until December 24, 1997
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.Until = dec24
		return rv
	}, "FREQ=DAILY;UNTIL=19971224T000000Z")

	// Every other day - forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.Interval = 2
		return rv
	}, "FREQ=DAILY;INTERVAL=2")

	// Every 10 days, 5 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.Interval = 10
		rv.Count = 5
		return rv
	}, "FREQ=DAILY;INTERVAL=10;COUNT=5")

	// Every day in January, for 3 years
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.Until = jan31
		rv.ByMonth = []uint{1}
		rv.ByDay = []WeekDayNum{SU, MO, TU, WE, TH, FR, SA}
		return rv
	}, "FREQ=YEARLY;UNTIL=20000131T140000Z;BYMONTH=1;BYDAY=SU,MO,TU,WE,TH,FR,SA")

	// Every day in January, for 3 years - alternative
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.Until = jan31
		rv.ByMonth = []uint{1}
		return rv
	}, "FREQ=DAILY;UNTIL=20000131T140000Z;BYMONTH=1")

	// Weekly for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Count = 10
		return rv
	}, "FREQ=WEEKLY;COUNT=10")

	// Weekly until December 24, 1997
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Until = dec24
		return rv
	}, "FREQ=WEEKLY;UNTIL=19971224T000000Z")

	// Every other week - forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Interval = 2
		rv.WeekStart = Sunday
		return rv
	}, "FREQ=WEEKLY;INTERVAL=2;WKST=SU")

	// Weekly on Tuesday and Thursday for five weeks
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Until = dec24
		rv.ByDay = []WeekDayNum{TU, TH}
		rv.WeekStart = Sunday
		return rv
	}, "FREQ=WEEKLY;UNTIL=19971224T000000Z;BYDAY=TU,TH;WKST=SU")

	// Weekly on Tuesday and Thursday for five weeks - alternative
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Count = 10
		rv.ByDay = []WeekDayNum{TU, TH}
		rv.WeekStart = Sunday
		return rv
	}, "FREQ=WEEKLY;COUNT=10;BYDAY=TU,TH;WKST=SU")

	// Every other week on Monday, Wednesday, and Friday until December
	// 24, 1997, starting on Monday, September 1, 1997
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Interval = 2
		rv.Until = dec24
		rv.ByDay = []WeekDayNum{MO, WE, FR}
		rv.WeekStart = Sunday
		return rv
	}, "FREQ=WEEKLY;INTERVAL=2;UNTIL=19971224T000000Z;BYDAY=MO,WE,FR;WKST=SU")

	// Every other week on Tuesday and Thursday, for 8 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Interval = 2
		rv.Count = 8
		rv.ByDay = []WeekDayNum{TU, TH}
		rv.WeekStart = Sunday
		return rv
	}, "FREQ=WEEKLY;INTERVAL=2;COUNT=8;BYDAY=TU,TH;WKST=SU")

	// Monthly on the first Friday for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Count = 10
		rv.ByDay = []WeekDayNum{{1, Friday}}
		return rv
	}, "FREQ=MONTHLY;COUNT=10;BYDAY=1FR")

	// Monthly on the first Friday until December 24, 1997
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Until = dec24
		rv.ByDay = []WeekDayNum{{1, Friday}}
		return rv
	}, "FREQ=MONTHLY;UNTIL=19971224T000000Z;BYDAY=1FR")

	// Every other month on the first and last Sunday of the month for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Interval = 2
		rv.Count = 10
		rv.ByDay = []WeekDayNum{{1, Sunday}, {-1, Sunday}}
		return rv
	}, "FREQ=MONTHLY;INTERVAL=2;COUNT=10;BYDAY=1SU,-1SU")

	// Monthly on the second-to-last Monday of the month for 6 months
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Count = 6
		rv.ByDay = []WeekDayNum{{-2, Monday}}
		return rv
	}, "FREQ=MONTHLY;COUNT=6;BYDAY=-2MO")

	// Monthly on the third-to-the-last day of the month, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.ByMonthDay = []int{-3}
		return rv
	}, "FREQ=MONTHLY;BYMONTHDAY=-3")

	// Monthly on the 2nd and 15th of the month for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Count = 10
		rv.ByMonthDay = []int{2, 15}
		return rv
	}, "FREQ=MONTHLY;COUNT=10;BYMONTHDAY=2,15")

	// Monthly on the first and last day of the month for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Count = 10
		rv.ByMonthDay = []int{1, -1}
		return rv
	}, "FREQ=MONTHLY;COUNT=10;BYMONTHDAY=1,-1")

	// Every 18 months on the 10th thru 15th of the month for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Interval = 18
		rv.Count = 10
		rv.ByMonthDay = []int{10, 11, 12, 13, 14, 15}
		return rv
	}, "FREQ=MONTHLY;INTERVAL=18;COUNT=10;BYMONTHDAY=10,11,12,13,14,15")

	// Every Tuesday, every other month
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Interval = 2
		rv.ByDay = []WeekDayNum{TU}
		return rv
	}, "FREQ=MONTHLY;INTERVAL=2;BYDAY=TU")

	// Yearly in June and July for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.Count = 10
		rv.ByMonth = []uint{6, 7}
		return rv
	}, "FREQ=YEARLY;COUNT=10;BYMONTH=6,7")

	// Every other year on January, February, and March for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.Interval = 2
		rv.Count = 10
		rv.ByMonth = []uint{1, 2, 3}
		return rv
	}, "FREQ=YEARLY;INTERVAL=2;COUNT=10;BYMONTH=1,2,3")

	// Every third year on the 1st, 100th, and 200th day for 10 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.Interval = 3
		rv.Count = 10
		rv.ByYearDay = []int{1, 100, 200}
		return rv
	}, "FREQ=YEARLY;INTERVAL=3;COUNT=10;BYYEARDAY=1,100,200")

	// Every 20th Monday of the year, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.ByDay = []WeekDayNum{{20, Monday}}
		return rv
	}, "FREQ=YEARLY;BYDAY=20MO")

	// Monday of week number 20 (where the default start of the week is Monday), forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.ByDay = []WeekDayNum{MO}
		rv.ByWeekNo = []int{20}
		return rv
	}, "FREQ=YEARLY;BYWEEKNO=20;BYDAY=MO")

	// Every Thursday in March, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.ByMonth = []uint{3}
		rv.ByDay = []WeekDayNum{TH}
		return rv
	}, "FREQ=YEARLY;BYMONTH=3;BYDAY=TH")

	// Every Thursday, but only during June, July, and August, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.ByMonth = []uint{6, 7, 8}
		rv.ByDay = []WeekDayNum{TH}
		return rv
	}, "FREQ=YEARLY;BYMONTH=6,7,8;BYDAY=TH")

	// Every Friday the 13th, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.ByDay = []WeekDayNum{FR}
		rv.ByMonthDay = []int{13}
		return rv
	}, "FREQ=MONTHLY;BYDAY=FR;BYMONTHDAY=13")

	// The first Saturday that follows the first Sunday of the month, forever
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.ByDay = []WeekDayNum{SA}
		rv.ByMonthDay = []int{7, 8, 9, 10, 11, 12, 13}
		return rv
	}, "FREQ=MONTHLY;BYDAY=SA;BYMONTHDAY=7,8,9,10,11,12,13")

	// Every 4 years, the first Tuesday after a Monday in November,
	// forever (U.S. Presidential Election day):
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(YEARLY)
		rv.Interval = 4
		rv.ByMonth = []uint{11}
		rv.ByDay = []WeekDayNum{TU}
		rv.ByMonthDay = []int{2, 3, 4, 5, 6, 7, 8}
		return rv
	}, "FREQ=YEARLY;INTERVAL=4;BYMONTH=11;BYDAY=TU;BYMONTHDAY=2,3,4,5,6,7,8")

	// The third instance into the month of one of Tuesday, Wednesday, or
	// Thursday, for the next 3 months
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.Count = 3
		rv.ByDay = []WeekDayNum{TU, WE, TH}
		rv.BySetPos = []int{3}
		return rv
	}, "FREQ=MONTHLY;COUNT=3;BYDAY=TU,WE,TH;BYSETPOS=3")

	// The second-to-last weekday of the month
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MONTHLY)
		rv.ByDay = []WeekDayNum{MO, TU, WE, TH, FR}
		rv.BySetPos = []int{-2}
		return rv
	}, "FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-2")

	// Every 3 hours from 9:00 AM to 5:00 PM on a specific day
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(HOURLY)
		rv.Interval = 3
		rv.Until = dec24
		return rv
	}, "FREQ=HOURLY;INTERVAL=3;UNTIL=19971224T000000Z")

	// Every 15 minutes for 6 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MINUTELY)
		rv.Interval = 15
		rv.Count = 6
		return rv
	}, "FREQ=MINUTELY;INTERVAL=15;COUNT=6")

	// Every hour and a half for 4 occurrences
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MINUTELY)
		rv.Interval = 90
		rv.Count = 4
		return rv
	}, "FREQ=MINUTELY;INTERVAL=90;COUNT=4")

	// Every 20 minutes from 9:00 AM to 4:40 PM every day
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.ByHour = []uint{9, 10, 11, 12, 13, 14, 15, 16}
		rv.ByMinute = []uint{0, 20, 40}
		return rv
	}, "FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYMINUTE=0,20,40")

	// Every 20 seconds from 9:00 AM to 4:40 PM every day
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(DAILY)
		rv.ByHour = []uint{9, 10, 11, 12, 13, 14, 15, 16}
		rv.BySecond = []uint{0, 20, 40}
		return rv
	}, "FREQ=DAILY;BYHOUR=9,10,11,12,13,14,15,16;BYSECOND=0,20,40")

	// Every 20 minutes from 9:00 AM to 4:40 PM every day (alternative)
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(MINUTELY)
		rv.Interval = 20
		rv.ByHour = []uint{9, 10, 11, 12, 13, 14, 15, 16}
		return rv
	}, "FREQ=MINUTELY;INTERVAL=20;BYHOUR=9,10,11,12,13,14,15,16")

	// An example where the days generated makes a difference because of WKST
	doTestRecur(t, func() RecurrenceValue {
		rv := Recurrence(WEEKLY)
		rv.Interval = 2
		rv.Count = 4
		rv.ByDay = []WeekDayNum{TU, SU}
		rv.WeekStart = Monday
		return rv
	}, "FREQ=WEEKLY;INTERVAL=2;COUNT=4;BYDAY=TU,SU;WKST=MO")
}

func doTestRecur(t *testing.T, v func() RecurrenceValue, exp string) {
	t.Helper()

	ics.MaxLineLength = 1000 // disable line folding

	b := &bytes.Buffer{}
	x := ics.NewFoldWriter(b, "\n")

	v().WriteTo(x)

	err := x.(ics.Flusher).Flush()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	s := b.String()
	if !strings.HasPrefix(s, ";VALUE=RECUR:") {
		t.Errorf("expected prefix got %q", s)
	} else {
		// ignore prefix
		s = s[13:]

		if s != exp {
			t.Errorf("expected %q but got %q", exp, s)
		}
	}
}

func TestRecurValidation(t *testing.T) {
	doTestRecurAOK(t, RecurrenceValue{BySecond: []uint{0}})
	doTestRecurAOK(t, RecurrenceValue{BySecond: []uint{60}})
	doTestRecurNAK(t, RecurrenceValue{BySecond: []uint{61}})

	doTestRecurAOK(t, RecurrenceValue{ByMinute: []uint{0}})
	doTestRecurAOK(t, RecurrenceValue{ByMinute: []uint{59}})
	doTestRecurNAK(t, RecurrenceValue{ByMinute: []uint{60}})

	doTestRecurAOK(t, RecurrenceValue{ByHour: []uint{0}})
	doTestRecurAOK(t, RecurrenceValue{ByHour: []uint{23}})
	doTestRecurNAK(t, RecurrenceValue{ByHour: []uint{24}})

	doTestRecurNAK(t, RecurrenceValue{ByMonth: []uint{0}})
	doTestRecurAOK(t, RecurrenceValue{ByMonth: []uint{1}})
	doTestRecurAOK(t, RecurrenceValue{ByMonth: []uint{12}})
	doTestRecurNAK(t, RecurrenceValue{ByMonth: []uint{13}})

	doTestRecurNAK(t, RecurrenceValue{ByMonthDay: []int{-32}})
	doTestRecurAOK(t, RecurrenceValue{ByMonthDay: []int{-31}})
	doTestRecurAOK(t, RecurrenceValue{ByMonthDay: []int{-1}})
	doTestRecurNAK(t, RecurrenceValue{ByMonthDay: []int{0}})
	doTestRecurAOK(t, RecurrenceValue{ByMonthDay: []int{1}})
	doTestRecurAOK(t, RecurrenceValue{ByMonthDay: []int{31}})
	doTestRecurNAK(t, RecurrenceValue{ByMonthDay: []int{32}})

	doTestRecurNAK(t, RecurrenceValue{ByWeekNo: []int{-54}})
	doTestRecurAOK(t, RecurrenceValue{ByWeekNo: []int{-53}})
	doTestRecurAOK(t, RecurrenceValue{ByWeekNo: []int{-1}})
	doTestRecurAOK(t, RecurrenceValue{ByWeekNo: []int{0}})
	doTestRecurAOK(t, RecurrenceValue{ByWeekNo: []int{1}})
	doTestRecurAOK(t, RecurrenceValue{ByWeekNo: []int{53}})
	doTestRecurNAK(t, RecurrenceValue{ByWeekNo: []int{54}})

	doTestRecurNAK(t, RecurrenceValue{ByYearDay: []int{-367}})
	doTestRecurAOK(t, RecurrenceValue{ByYearDay: []int{-366}})
	doTestRecurAOK(t, RecurrenceValue{ByYearDay: []int{-1}})
	doTestRecurNAK(t, RecurrenceValue{ByYearDay: []int{0}})
	doTestRecurAOK(t, RecurrenceValue{ByYearDay: []int{1}})
	doTestRecurAOK(t, RecurrenceValue{ByYearDay: []int{366}})
	doTestRecurNAK(t, RecurrenceValue{ByYearDay: []int{367}})

	doTestRecurNAK(t, RecurrenceValue{BySetPos: []int{-367}})
	doTestRecurAOK(t, RecurrenceValue{BySetPos: []int{-366}})
	doTestRecurAOK(t, RecurrenceValue{BySetPos: []int{-1}})
	doTestRecurNAK(t, RecurrenceValue{BySetPos: []int{0}})
	doTestRecurAOK(t, RecurrenceValue{BySetPos: []int{1}})
	doTestRecurAOK(t, RecurrenceValue{BySetPos: []int{366}})
	doTestRecurNAK(t, RecurrenceValue{BySetPos: []int{367}})

	doTestRecurNAK(t, RecurrenceValue{ByDay: []WeekDayNum{{54, Sunday}}})
	doTestRecurAOK(t, RecurrenceValue{ByDay: []WeekDayNum{{53, Sunday}}})
	doTestRecurAOK(t, RecurrenceValue{ByDay: []WeekDayNum{{0, Sunday}}})
	doTestRecurAOK(t, RecurrenceValue{ByDay: []WeekDayNum{{-53, Sunday}}})
	doTestRecurNAK(t, RecurrenceValue{ByDay: []WeekDayNum{{-54, Sunday}}})
}

func doTestRecurNAK(t *testing.T, v RecurrenceValue) {
	t.Helper()

	b := &bytes.Buffer{}
	x := ics.NewFoldWriter(b, "\n")

	err := v.WriteTo(x)
	x.(ics.Flusher).Flush()

	if err == nil {
		t.Errorf("expected error but got %s", b.String())
	}
}

func doTestRecurAOK(t *testing.T, v RecurrenceValue) {
	t.Helper()

	b := &bytes.Buffer{}
	x := ics.NewFoldWriter(b, "\n")

	err := v.WriteTo(x)

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
}
