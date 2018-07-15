package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
)

// VEvent captures a calendar event
type VEvent struct {
	UID          value.TextValue
	DTStamp      value.DateTimeValue
	Start        value.DateTimeValue
	End          value.DateTimeValue
	Created      value.DateTimeValue
	LastModified value.DateTimeValue
	Organizer    value.CalAddressValue
	Attendee     []value.CalAddressValue
	Contact      value.TextValue
	Summary      value.TextValue
	Description  value.TextValue
	Class        value.TextValue // PUBLIC, PRIVATE, CONFIDENTIAL
	Comment      value.TextValue
	RelatedTo    value.TextValue
	Sequence     value.IntegerValue
	Status       value.TextValue
	Location     value.TextValue
	Transparency value.TransparencyValue
	Color        value.TextValue // CSS3 color name

	// TODO (RFC5545) CREATED GEO PRIORITY RECURRENCE-ID EXDATE RDATE RRULE
	// TODO (RFC7986) []CONFERENCE
}

// AllDay changes the start and end to represent dates without time.
// If they are already configured as dates only, this has no effect.
func (e *VEvent) AllDay() *VEvent {
	e.Start = e.Start.AsDate()
	e.End = e.End.AsDate()
	return e
}

// EncodeIcal serialises the event to the buffer in iCal ics format
// (a VComponent method).
func (e *VEvent) EncodeIcal(b *ics.Buffer, method value.MethodValue) error {

	if !ics.IsDefined(e.DTStamp) {
		return fmt.Errorf("DTstamp is required")
	}

	if !ics.IsDefined(e.UID) {
		return fmt.Errorf("UID is required")
	}

	if !ics.IsDefined(method) && !ics.IsDefined(e.Start) {
		return fmt.Errorf("When Method is undefined, Start is required")
	}

	b.WriteLine("BEGIN:VEVENT")

	b.WriteValuerLine(ics.IsDefined(e.Start), "DTSTART", e.Start)
	b.WriteValuerLine(ics.IsDefined(e.End), "DTEND", e.End)
	b.WriteValuerLine(true, "DTSTAMP", e.DTStamp)
	b.WriteValuerLine(true, "UID", e.UID)
	b.WriteValuerLine(ics.IsDefined(e.Organizer), "ORGANIZER", e.Organizer)

	for _, attendee := range e.Attendee {
		b.WriteValuerLine(true, "ATTENDEE", attendee)
	}

	b.WriteValuerLine(ics.IsDefined(e.Contact), "CONTACT", e.Contact)
	b.WriteValuerLine(ics.IsDefined(e.Summary), "SUMMARY", e.Summary)
	b.WriteValuerLine(ics.IsDefined(e.Description), "DESCRIPTION", e.Description)
	b.WriteValuerLine(ics.IsDefined(e.Location), "LOCATION", e.Location)
	b.WriteValuerLine(ics.IsDefined(e.Class), "CLASS", e.Class)
	b.WriteValuerLine(ics.IsDefined(e.Comment), "COMMENT", e.Comment)
	b.WriteValuerLine(ics.IsDefined(e.Created), "CREATED", e.Created)
	b.WriteValuerLine(ics.IsDefined(e.LastModified), "LAST-MODIFIED", e.LastModified)
	b.WriteValuerLine(ics.IsDefined(e.RelatedTo), "RELATED-TO", e.RelatedTo)
	b.WriteValuerLine(ics.IsDefined(e.Sequence), "SEQUENCE", e.Sequence)
	b.WriteValuerLine(ics.IsDefined(e.Status), "STATUS", e.Status)
	b.WriteValuerLine(ics.IsDefined(e.Transparency), "TRANSP", e.Transparency)

	//if ics.IsDefined(e.ALARM) {
	//	b.WriteLine("BEGIN:VALARM")
	//	b.WriteValuerLine(true, "TRIGGER", e.ALARM)
	//	b.WriteLine("ACTION:DISPLAY")
	//	b.WriteLine("END:VALARM")
	//}

	b.WriteLine("END:VEVENT")

	return b.Flush()
}
