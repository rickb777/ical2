package ical2

import (
	"bytes"
	. "github.com/rickb777/ical2/parameter"
	. "github.com/rickb777/ical2/value"
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

	event := &VEvent{
		UID:          Text("123"),
		DTStamp:      TStamp(dt),
		Start:        DateTime(ds).With(TZID(tz)),
		End:          DateTime(de).With(TZID(tz)),
		Organizer:    CalAddress("ht@throne.com").With(CommonName("H.Tudwr")),
		Attendee:     []URIValue{CalAddress("ann.blin@example.com").With(Role("REQ-PARTICIPANT"), CommonName("Ann Blin"))},
		Contact:      Text("T.Moore, Esq."),
		Summary:      Text("summary, with punctuation"),
		Description:  Text("Lorem ipsum dolor sit amet, consectetµr adipiscing elit, sed do eiusmod tempor incididµnt µt labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		RelatedTo:    Text("19960401-080045-4000F192713-0052@example.com"),
		Location:     Text("South Bank, London SE1 9PX"),
		Transparency: Transparent(),
	}

	b, err := testSetup(tz, event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Europe/Paris
BEGIN:VEVENT
DTSTART;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T080000
DTEND;VALUE=DATE-TIME;TZID=Europe/Paris:20140101T130000
DTSTAMP:20140101T060000Z
UID:123
ORGANIZER;CN=H.Tudwr:mailto:ht@throne.com
ATTENDEE;ROLE=REQ-PARTICIPANT;CN=Ann Blin:mailto:ann.blin@example.com
CONTACT:T.Moore\, Esq.
SUMMARY:summary\, with punctuation
DESCRIPTION:Lorem ipsum dolor sit amet\, consectetµr adipiscing elit\, sed
  do eiusmod tempor incididµnt µt labore et dolore magna aliqua. Ut enim a
 d minim veniam\, quis nostrud exercitation ullamco laboris nisi ut aliquip 
 ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptat
 e velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaeca
 t cupidatat non proident\, sunt in culpa qui officia deserunt mollit anim i
 d est laborum.
LOCATION:South Bank\, London SE1 9PX
RELATED-TO:19960401-080045-4000F192713-0052@example.com
TRANSP:TRANSPARENT
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeAllDayTrue(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	event := (&VEvent{
		UID:          Text("123"),
		DTStamp:      TStamp(d),
		Start:        DateTime(d).With(TZID(tz)),
		End:          DateTime(d).With(TZID(tz)),
		Summary:      Text("summary"),
		Transparency: Opaque(),
	}).AllDay()

	b, err := testSetup(tz, event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
BEGIN:VEVENT
DTSTART;TZID=Asia/Tokyo;VALUE=DATE:20140101
DTEND;TZID=Asia/Tokyo;VALUE=DATE:20140101
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
TRANSP:OPAQUE
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeDraftProperties(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	event := &VEvent{
		UID:     Text("123"),
		DTStamp: TStamp(d),
		Start:   DateTime(d).With(TZID(tz)),
		End:     DateTime(d).With(TZID(tz)),
		Summary: Text("summary"),
	}

	b, err := testSetupWithDraft(tz, event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
METHOD:PUBLISH
NAME:name
DESCRIPTION:desc
URL:http://my.calendar/url
COLOR:#34AF10
REFRESH-INTERVAL;VALUE=DURATION:PT12H
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
X-PUBLISHED-TTL:PT12H
BEGIN:VEVENT
DTSTART;VALUE=DATE-TIME;TZID=Asia/Tokyo:20140101T000000
DTEND;VALUE=DATE-TIME;TZID=Asia/Tokyo:20140101T000000
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeNoTzid(t *testing.T) {
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, time.Local)

	event := &VEvent{
		UID:     Text("123"),
		DTStamp: TStamp(d),
		Start:   DateTime(d),
		End:     DateTime(d),
		Summary: Text("summary"),
	}

	b, err := testSetup("", event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
BEGIN:VEVENT
DTSTART;VALUE=DATE-TIME:20140101T000000
DTEND;VALUE=DATE-TIME:20140101T000000
DTSTAMP:20140101T000000Z
UID:123
SUMMARY:summary
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeUtcTzid(t *testing.T) {
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	event := &VEvent{
		UID:     Text("123"),
		DTStamp: TStamp(d),
		Start:   DateTime(d),
		End:     DateTime(d),
		Summary: Text("summary"),
	}

	b, err := testSetup("UTC", event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:UTC
BEGIN:VEVENT
DTSTART;VALUE=DATE-TIME:20140101T000000Z
DTEND;VALUE=DATE-TIME:20140101T000000Z
DTSTAMP:20140101T000000Z
UID:123
SUMMARY:summary
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeNoTzidAllDay(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	event := (&VEvent{
		UID:     Text("123"),
		DTStamp: TStamp(d),
		Start:   DateTime(d),
		End:     DateTime(d),
		Summary: Text("summary"),
	}).AllDay()

	b, err := testSetup(tz, event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
BEGIN:VEVENT
DTSTART;VALUE=DATE:20140101
DTEND;VALUE=DATE:20140101
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func TestEncodeUtcTzidAllDay(t *testing.T) {
	const tz = "Asia/Tokyo"
	zone := time.FixedZone(tz, 60*60*9)
	d := time.Date(2014, time.Month(1), 1, 0, 0, 0, 0, zone)

	event := (&VEvent{
		UID:     Text("123"),
		DTStamp: TStamp(d),
		Start:   DateTime(d),
		End:     DateTime(d),
		Summary: Text("summary"),
	}).AllDay()

	b, err := testSetup(tz, event)
	if err != nil {
		t.Error("got err:", err)
	}

	expect := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:prodid
CALSCALE:GREGORIAN
X-WR-CALNAME:name
X-WR-CALDESC:desc
X-WR-TIMEZONE:Asia/Tokyo
BEGIN:VEVENT
DTSTART;VALUE=DATE:20140101
DTEND;VALUE=DATE:20140101
DTSTAMP:20131231T150000Z
UID:123
SUMMARY:summary
END:VEVENT
END:VCALENDAR
`
	expect = unixToDOSLineEndings(expect)

	if s := b.String(); s != expect {
		t.Errorf("got %s", s)
	}
}

func unixToDOSLineEndings(input string) string {
	return strings.Replace(input, "\n", "\r\n", -1)
}

func testSetup(tz string, vComponents ...VComponent) (bytes.Buffer, error) {
	c := NewVCalendar("prodid")
	c.Extend("X-WR-CALNAME", Text("name"))
	c.Extend("X-WR-CALDESC", Text("desc"))
	if tz != "" {
		c.Extend("X-WR-TIMEZONE", Text(tz))
	}

	c.VComponent = vComponents

	var b bytes.Buffer
	if err := c.Encode(&b); err != nil {
		return b, err
	}

	return b, nil
}

func testSetupWithDraft(tz string, vComponents ...VComponent) (bytes.Buffer, error) {
	c := NewVCalendar("prodid")
	c.URL = Text("http://my.calendar/url")
	c.Name = Text("name")
	c.Extend("X-WR-CALNAME", Text("name"))
	c.Description = Text("desc")
	c.Extend("X-WR-CALDESC", Text("desc"))
	c.Extend("X-WR-TIMEZONE", Text(tz))
	c.RefreshInterval = Duration("PT12H")
	c.Extend("X-PUBLISHED-TTL", Text("PT12H"))
	c.Color = Text("#34AF10")
	c.Method = Publish()

	c.VComponent = vComponents

	var b bytes.Buffer
	if err := c.Encode(&b); err != nil {
		return b, err
	}

	return b, nil
}
