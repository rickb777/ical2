package ical2_test

import (
	"fmt"
	"github.com/rickb777/date/v2/timespan"
	"github.com/rickb777/ical2"
	"github.com/rickb777/ical2/parameter/freebusy"
	"github.com/rickb777/ical2/value"
	"time"
)

func ExampleVFreeBusy_request() {
	// This is an example of a "VFREEBUSY" calendar component used to request free or busy time information

	dt := time.Date(1997, time.Month(9), 1, 8, 30, 0, 0, time.UTC)
	ds := time.Date(1997, time.Month(10), 15, 5, 0, 0, 0, time.UTC)
	de := time.Date(1997, time.Month(10), 16, 5, 0, 0, 0, time.UTC)

	event := &ical2.VFreeBusy{
		UID:       value.Text("19970901T082949Z-FA43EF@example.com"),
		DTStamp:   value.TStamp(dt),
		Start:     value.DateTime(ds),
		End:       value.DateTime(de),
		Organizer: value.CalAddress("jane_doe@example.com"),
		Attendee:  []value.URIValue{value.CalAddress("john_public@example.com")},
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	c.Method = value.Request()

	// usually you'd Encode to some io.Writer
	//c.Encode(w)
	// but for this example, we'll just stringify
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// PRODID:-//My App//Event Calendar//EN
	// VERSION:2.0
	// CALSCALE:GREGORIAN
	// METHOD:REQUEST
	// BEGIN:VFREEBUSY
	// DTSTART;VALUE=DATE-TIME:19971015T050000Z
	// DTEND;VALUE=DATE-TIME:19971016T050000Z
	// DTSTAMP:19970901T083000Z
	// UID:19970901T082949Z-FA43EF@example.com
	// ORGANIZER:mailto:jane_doe@example.com
	// ATTENDEE:mailto:john_public@example.com
	// END:VFREEBUSY
	// END:VCALENDAR
}

func ExampleVFreeBusy_publish() {
	// This is an example of a "VFREEBUSY" calendar component used to publish busy time information.

	dt := time.Date(1997, time.Month(9), 1, 12, 0, 0, 0, time.UTC)
	ds := time.Date(1998, time.Month(3), 13, 14, 17, 11, 0, time.UTC)
	de := time.Date(1998, time.Month(4), 10, 14, 17, 11, 0, time.UTC)

	t1s := time.Date(1998, time.Month(3), 14, 23, 30, 0, 0, time.UTC)
	t2s := time.Date(1998, time.Month(3), 16, 15, 30, 0, 0, time.UTC)
	t3s := time.Date(1998, time.Month(3), 18, 3, 0, 0, 0, time.UTC)

	event := &ical2.VFreeBusy{
		UID:       value.Text("19970901T115957Z-76A912@example.com"),
		DTStamp:   value.TStamp(dt),
		Start:     value.DateTime(ds),
		End:       value.DateTime(de),
		Organizer: value.CalAddress("jsmith@example.com"),
		URL:       value.URI("http://www.example.com/calendar/busytime/jsmith.ifb"),
		FreeBusy: []value.PeriodValue{
			value.Period(timespan.TimeSpanOf(t1s, time.Hour)).With(freebusy.Busy()),
			value.Period(timespan.TimeSpanOf(t2s, time.Hour)).With(freebusy.BusyTentative()),
			value.Period(timespan.TimeSpanOf(t3s, time.Hour)).With(freebusy.BusyUnavailable()),
		},
		Comment: []value.TextValue{value.Text("Busy time")},
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	c.Method = value.Publish()

	// usually you'd Encode to some io.Writer
	//c.Encode(w)
	// but for this example, we'll just stringify
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// PRODID:-//My App//Event Calendar//EN
	// VERSION:2.0
	// CALSCALE:GREGORIAN
	// METHOD:PUBLISH
	// BEGIN:VFREEBUSY
	// DTSTART;VALUE=DATE-TIME:19980313T141711Z
	// DTEND;VALUE=DATE-TIME:19980410T141711Z
	// DTSTAMP:19970901T120000Z
	// UID:19970901T115957Z-76A912@example.com
	// ORGANIZER:mailto:jsmith@example.com
	// URL;VALUE=URI:http://www.example.com/calendar/busytime/jsmith.ifb
	// COMMENT:Busy time
	// FREEBUSY;VALUE=PERIOD;FBTYPE=BUSY:19980314T233000Z/PT1H
	// FREEBUSY;VALUE=PERIOD;FBTYPE=BUSY-TENTATIVE:19980316T153000Z/PT1H
	// FREEBUSY;VALUE=PERIOD;FBTYPE=BUSY-UNAVAILABLE:19980318T030000Z/PT1H
	// END:VFREEBUSY
	// END:VCALENDAR
}
