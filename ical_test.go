package ical

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	const tz = "Europe/Paris"
	zone := time.FixedZone(tz, 60*60*1)
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, zone)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	vComponents := []VComponent{
		&VEvent{
			UID:         "123",
			DTSTAMP:     dt,
			DTSTART:     ds,
			DTEND:       de,
			ORGANIZER:   Individual{"H.Tudwr", "", "ht@throne.com"},
			ATTENDEE:    []Attendee{{Individual{"Ann Blin", "", "ann.blin@example.com"}, "REQ-PARTICIPANT"}},
			CONTACT:     "T.Moore",
			SUMMARY:     "summary",
			DESCRIPTION: "description",
			TZID:        tz,
			LOCATION:    "Paris",
			TRANSP:      "TRANSPARENT",
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Europe/Paris
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20140101T060000Z
UID:123
TZID:Europe/Paris
ORGANIZER;CN=H.Tudwr:mailto=ht@throne.com
ATTENDEE;ROLE=REQ-PARTICIPANT;CN=Ann Blin:mailto=ann.blin@example.com
CONTACT:T.Moore
SUMMARY:summary
DESCRIPTION:description
LOCATION:Paris
TRANSP:TRANSPARENT
DTSTART;TZID=Europe/Paris;VALUE=DATE-TIME:20140101T080000
DTEND;TZID=Europe/Paris;VALUE=DATE-TIME:20140101T130000
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %v", s)
	}
}

func TestEncodeAllDayTrue(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
			TZID:    "Asia/Tokyo",

			AllDay: true,
			TRANSP: "OPAQUE",
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
TZID:Asia/Tokyo
SUMMARY:summary
TRANSP:OPAQUE
DTSTART;TZID=Asia/Tokyo;VALUE=DATE:20140101
DTEND;TZID=Asia/Tokyo;VALUE=DATE:20140101
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func TestEncodeDraftProperties(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
			TZID:    "Asia/Tokyo",
		},
	}

	b, err := testSetupWithDraft(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
URL:http://my.calendar/url
NAME:name
X-WR-CALNAME:name
DESCRIPTION:desc
X-WR-CALDESC:desc
TIMEZONE-ID:Asia/Tokyo
X-WR-TIMEZONE:Asia/Tokyo
REFRESH-INTERVAL;VALUE=DURATION:PT12H
X-PUBLISHED-TTL:PT12H
COLOR:34:50:105
CALSCALE:GREGORIAN
METHOD:PUBLISH
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
TZID:Asia/Tokyo
SUMMARY:summary
DTSTART;TZID=Asia/Tokyo;VALUE=DATE-TIME:20140101T000000
DTEND;TZID=Asia/Tokyo;VALUE=DATE-TIME:20140101T000000
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func TestEncodeNoTzid(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
DTSTART;VALUE=DATE-TIME:20140101T000000Z
DTEND;VALUE=DATE-TIME:20140101T000000Z
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func TestEncodeUtcTzid(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
			TZID:    "UTC",
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
DTSTART;VALUE=DATE-TIME:20140101T000000Z
DTEND;VALUE=DATE-TIME:20140101T000000Z
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func TestEncodeNoTzidAllDay(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
			AllDay:  true,
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
DTSTART;VALUE=DATE:20140101
DTEND;VALUE=DATE:20140101
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func TestEncodeUtcTzidAllDay(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	vComponents := []VComponent{
		&VEvent{
			UID:     "123",
			DTSTAMP: d,
			DTSTART: d,
			DTEND:   d,
			SUMMARY: "summary",
			TZID:    "UTC",
			AllDay:  true,
		},
	}

	b, err := testSetup(tz, vComponents)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
DTSTART;VALUE=DATE:20140101
DTEND;VALUE=DATE:20140101
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("should %v. but got %v", expect, s)
	}
}

func unixToDOSLineEndings(input string) string {
	return strings.Replace(input, "\n", "\r\n", -1)
}

func testSetup(tz string, vComponents []VComponent) (bytes.Buffer, error) {
	c := NewBasicVCalendar()
	c.PRODID = "prodid"
	c.X_WR_TIMEZONE = tz
	c.X_WR_CALNAME = "name"
	c.X_WR_CALDESC = "desc"

	c.VComponent = vComponents

	var b bytes.Buffer
	if err := c.Encode(&b); err != nil {
		return b, err
	}

	return b, nil
}

func testSetupWithDraft(tz string, vComponents []VComponent) (bytes.Buffer, error) {
	c := NewBasicVCalendar()
	c.PRODID = "prodid"
	c.URL = "http://my.calendar/url"
	c.NAME = "name"
	c.X_WR_CALNAME = "name"
	c.DESCRIPTION = "desc"
	c.X_WR_CALDESC = "desc"
	c.TIMEZONE_ID = tz
	c.X_WR_TIMEZONE = tz
	c.REFRESH_INTERVAL = "PT12H"
	c.X_PUBLISHED_TTL = "PT12H"
	c.COLOR = "34:50:105"
	c.CALSCALE = "GREGORIAN"
	c.METHOD = "PUBLISH"

	c.VComponent = vComponents

	var b bytes.Buffer
	if err := c.Encode(&b); err != nil {
		return b, err
	}

	return b, nil
}
