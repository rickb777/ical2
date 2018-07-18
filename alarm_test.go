package ical2_test

import (
	"fmt"
	"github.com/rickb777/ical2"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/value"
	"time"
)

func ExampleVAlarm_audio() {
	const tz = "Europe/Paris"
	zone, _ := time.LoadLocation(tz)
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, zone)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	alarm := &ical2.VAudioAlarm{
		Trigger:  value.DateTime(ds.Add(-time.Hour).In(time.UTC)),
		Duration: value.Duration("PT10M"),
		Repeat:   value.Integer(3),
		Attach: value.URI("http://example.com/clips/poke.aud").
			With(parameter.FmtTypeOf("audio", "basic")),
	}

	event := &ical2.VEvent{
		UID:     value.Text("123"),
		DTStamp: value.TStamp(dt),
		Start:   value.DateTime(ds).With(parameter.TZid(tz)),
		End:     value.DateTime(de).With(parameter.TZid(tz)),
		Status:  value.NeedsActionStatus(),
		Alarm:   []ical2.VAlarm{alarm},
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	// usually you'd Encode to some io.Writer
	//c.Encode(w)
	// but for this example, we'll just stringify
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// VERSION:2.0
	// PRODID:-//My App//Event Calendar//EN
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// DTSTART;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T080000
	// DTEND;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T130000
	// DTSTAMP:20140101T060000Z
	// UID:123
	// STATUS:NEEDS-ACTION
	// BEGIN:VALARM
	// ACTION:AUDIO
	// TRIGGER;VALUE=DATE-TIME:20140101T060000Z
	// DURATION;VALUE=DURATION:PT10M
	// REPEAT;VALUE=INTEGER:3
	// ATTACH;VALUE=URI;FMTTYPE=audio/basic:http://example.com/clips/poke.aud
	// END:VALARM
	// END:VEVENT
	// END:VCALENDAR
}

func ExampleVAlarm_display() {
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, time.UTC)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	alarm := &ical2.VDisplayAlarm{
		Description: value.Text("Wakey wakey"),
		Trigger:     value.Duration("-PT10M"),
	}

	event := &ical2.VEvent{
		UID:     value.Text("123"),
		DTStamp: value.TStamp(dt),
		Start:   value.DateTime(ds),
		End:     value.DateTime(de),
		Alarm:   []ical2.VAlarm{alarm},
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	// usually you'd Encode to some io.Writer
	//c.Encode(w)
	// but for this example, we'll just stringify
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// VERSION:2.0
	// PRODID:-//My App//Event Calendar//EN
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// DTSTART;VALUE=DATE-TIME:20140101T080000Z
	// DTEND;VALUE=DATE-TIME:20140101T130000Z
	// DTSTAMP:20140101T070000Z
	// UID:123
	// BEGIN:VALARM
	// ACTION:DISPLAY
	// DESCRIPTION:Wakey wakey
	// TRIGGER;VALUE=DURATION:-PT10M
	// END:VALARM
	// END:VEVENT
	// END:VCALENDAR
}

func ExampleVAlarm_email() {
	const tz = "Europe/Paris"
	zone, _ := time.LoadLocation(tz)
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, zone)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	alarm := &ical2.VEmailAlarm{
		Description: value.Text("Wakey wakey"),
		Trigger:     value.Duration("-PT10M"),
		Summary:     value.Text("There are things to be done."),
		Attendee:    []value.URIValue{value.CalAddress("john_public@example.com")},
	}

	event := &ical2.VEvent{
		UID:     value.Text("123"),
		DTStamp: value.TStamp(dt),
		Start:   value.DateTime(ds).With(parameter.TZid(tz)),
		End:     value.DateTime(de).With(parameter.TZid(tz)),
		Alarm:   []ical2.VAlarm{alarm},
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	// usually you'd Encode to some io.Writer
	//c.Encode(w)
	// but for this example, we'll just stringify
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// VERSION:2.0
	// PRODID:-//My App//Event Calendar//EN
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// DTSTART;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T080000
	// DTEND;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T130000
	// DTSTAMP:20140101T060000Z
	// UID:123
	// BEGIN:VALARM
	// ACTION:EMAIL
	// DESCRIPTION:Wakey wakey
	// TRIGGER;VALUE=DURATION:-PT10M
	// SUMMARY:There are things to be done.
	// ATTENDEE:mailto:john_public@example.com
	// END:VALARM
	// END:VEVENT
	// END:VCALENDAR
}
