package ical2

import (
	"time"
	"fmt"
)

func ExampleVEvent() {
	const tz = "Europe/Paris"
	zone := time.FixedZone(tz, 60*60*1)
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, zone)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	event := &VEvent{
		UID:          Text("123"),
		DTStamp:      TStamp(dt),
		DTStart:      DateTime(ds),
		DTEnd:        DateTime(de),
		Organizer:    CalAddress("ht@throne.com").With(CommonName("H.Tudwr")),
		Attendee:     []CalAddressValue{CalAddress("ann.blin@example.com").With(Role("REQ-PARTICIPANT"), CommonName("Ann Blin"))},
		Contact:      Text("T.Moore, Esq."),
		Summary:      Text("Event summary"),
		Description:  Text("This describes the event."),
		RelatedTo:    Text("19960401-080045-4000F192713-0052@example.com"),
		TZID:         Text(tz),
		Location:     Text("South Bank, London SE1 9PX"),
		Transparency: Text(TRANSPARENT),
	}

	c := NewBasicVCalendar("-//My App//Event Calendar//EN").With(event)
	// usually you'd Encode to some io.Writer
	//c.Encode(w, "")
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// VERSION:2.0
	// PRODID:-//My App//Event Calendar//EN
	// CALSCALE:GREGORIAN
	// BEGIN:VEVENT
	// DTSTAMP:20140101T060000Z
	// UID:123
	// TZID:Europe/Paris
	// ORGANIZER;CN=H.Tudwr:mailto:ht@throne.com
	// ATTENDEE;ROLE=REQ-PARTICIPANT;CN=Ann Blin:mailto:ann.blin@example.com
	// CONTACT:T.Moore\, Esq.
	// SUMMARY:Event summary
	// DESCRIPTION:This describes the event.
	// LOCATION:South Bank\, London SE1 9PX
	// RELATED-TO:19960401-080045-4000F192713-0052@example.com
	// TRANSP:TRANSPARENT
	// DTSTART;TZID=Europe/Paris;VALUE=DATE-TIME:20140101T080000
	// DTEND;TZID=Europe/Paris;VALUE=DATE-TIME:20140101T130000
	// END:VEVENT
	// END:VCALENDAR
}
