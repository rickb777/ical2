package ical2_test

import (
	"fmt"
	"github.com/rickb777/ical2"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/value"
	"time"
)

func ExampleVEventWithTimezone() {
	const tz = "Europe/Paris"
	zone, _ := time.LoadLocation(tz)
	dt := time.Date(2014, time.Month(1), 1, 7, 0, 0, 0, zone)
	ds := dt.Add(time.Hour)
	de := ds.Add(5 * time.Hour)

	event := &ical2.VEvent{
		UID:          value.Text("123"),
		DTStamp:      value.TStamp(dt),
		Start:        value.DateTime(ds).With(parameter.TZID(tz)),
		End:          value.DateTime(de).With(parameter.TZID(tz)),
		Organizer:    value.CalAddress("ht@throne.com").With(parameter.CommonName("H.Tudwr")),
		Attendee:     []value.URIValue{value.CalAddress("ann.blin@example.com").With(parameter.Role(parameter.REQ_PARTICIPANT), parameter.CommonName("Ann Blin"))},
		Contact:      value.Text("T.Moore, Esq."),
		Summary:      value.Text("Event summary"),
		Description:  value.Text("This describes the event."),
		RelatedTo:    value.Text("19960401-080045-4000F192713-0052@example.com"),
		Location:     value.Text("South Bank, London SE1 9PX"),
		Transparency: value.Transparent(),
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
	// ORGANIZER;CN=H.Tudwr:mailto:ht@throne.com
	// ATTENDEE;ROLE=REQ-PARTICIPANT;CN=Ann Blin:mailto:ann.blin@example.com
	// CONTACT:T.Moore\, Esq.
	// SUMMARY:Event summary
	// DESCRIPTION:This describes the event.
	// LOCATION:South Bank\, London SE1 9PX
	// RELATED-TO:19960401-080045-4000F192713-0052@example.com
	// TRANSP:TRANSPARENT
	// END:VEVENT
	// END:VCALENDAR
}

func ExampleMeetingVEvent() {
	dt := time.Date(2014, time.Month(1), 1, 8, 0, 0, 0, time.UTC)
	ds := dt.Add(48 * time.Hour)
	de := ds.Add(72 * time.Hour)

	cath1 := value.CalAddress("cath.dragon@example.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Cath Dragon"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	ann1 := value.CalAddress("anne.bollin@example.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Anne Bollin"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	jane := value.CalAddress("jane.seemoor@example.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Jane Seemoor"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	ann2 := value.CalAddress("anne@cleves.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Anne Cleves"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	cath2 := value.CalAddress("cath@thehowards.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Cath Howard"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	cath3 := value.CalAddress("catherine.parr@respectable.com").
		With(
			parameter.CUType(parameter.INDIVIDUAL),
			parameter.Role(parameter.REQ_PARTICIPANT),
			parameter.PartStat(parameter.NEEDS_ACTION),
			parameter.RSVP(true),
			parameter.CommonName("Cath Parr"),
			parameter.Parameter{"X-NUM-GUESTS", "0"},
		)

	event := &ical2.VEvent{
		UID:          value.Text("0ibinszut0oiksq0sa0ac98d46@google.com"),
		DTStamp:      value.TStamp(dt),
		Created:      value.TStamp(dt.Add(-2 * time.Hour)),
		LastModified: value.TStamp(dt.Add(-1 * time.Hour)),
		Sequence:     value.Integer(0),
		Status:       value.Text("CONFIRMED"),
		Start:        value.TStamp(ds),
		End:          value.TStamp(de),
		Organizer:    value.CalAddress("ht@throne.com").With(parameter.CommonName("H.Tudwr")),
		Attendee:     []value.URIValue{cath1, ann1, jane, ann2, cath2, cath3},
		Summary:      value.Text("Meet the family"),
		Description:  value.Text("This is a great chance to meet each other!"),
		Location:     value.Text("South Bank, London SE1 9PX"),
		Transparency: value.Opaque(),
	}

	c := ical2.NewVCalendar("-//My App//Event Calendar//EN").With(event)
	c.Method = value.Request()
	fmt.Printf(c.String())

	// Output:
	// BEGIN:VCALENDAR
	// VERSION:2.0
	// PRODID:-//My App//Event Calendar//EN
	// CALSCALE:GREGORIAN
	// METHOD:REQUEST
	// BEGIN:VEVENT
	// DTSTART:20140103T080000Z
	// DTEND:20140106T080000Z
	// DTSTAMP:20140101T080000Z
	// UID:0ibinszut0oiksq0sa0ac98d46@google.com
	// ORGANIZER;CN=H.Tudwr:mailto:ht@throne.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Cath Dragon;X-NUM-GUESTS=0:mailto:cath.dragon@example.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Anne Bollin;X-NUM-GUESTS=0:mailto:anne.bollin@example.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Jane Seemoor;X-NUM-GUESTS=0:mailto:jane.seemoor@example.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Anne Cleves;X-NUM-GUESTS=0:mailto:anne@cleves.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Cath Howard;X-NUM-GUESTS=0:mailto:cath@thehowards.com
	// ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=
	//  TRUE;CN=Cath Parr;X-NUM-GUESTS=0:mailto:catherine.parr@respectable.com
	// SUMMARY:Meet the family
	// DESCRIPTION:This is a great chance to meet each other!
	// LOCATION:South Bank\, London SE1 9PX
	// CREATED:20140101T060000Z
	// LAST-MODIFIED:20140101T070000Z
	// SEQUENCE:0
	// STATUS:CONFIRMED
	// TRANSP:OPAQUE
	// END:VEVENT
	// END:VCALENDAR
}
