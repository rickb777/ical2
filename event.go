package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/value"
)

// VEvent captures a calendar event
type VEvent struct {
	UID          value.TextValue
	DTStamp      value.DateTimeValue
	DTStart      value.DateTimeValue
	DTEnd        value.DateTimeValue
	LastModified value.DateTimeValue
	Organizer    value.CalAddressValue
	Attendee     []value.CalAddressValue
	Contact      value.TextValue
	Summary      value.TextValue
	Description  value.TextValue
	Class        value.TextValue // PUBLIC, PRIVATE, CONFIDENTIAL
	Comment      value.TextValue
	RelatedTo    value.TextValue
	TZID         value.TextValue
	Sequence     ics.Valuer
	Status       value.TextValue
	ALARM        value.TextValue
	Location     value.TextValue
	Transparency value.TextValue
	Color        value.TextValue // CSS3 color name

	// TODO (RFC5545) CREATED GEO PRIORITY RECURRENCE-ID EXDATE RDATE RRULE
	// TODO (RFC7986) []CONFERENCE
}

func (e *VEvent) AllDay() *VEvent {
	e.DTStart = e.DTStart.AsDate()
	e.DTEnd = e.DTEnd.AsDate()
	return e
}

func (e *VEvent) EncodeIcal(b *ics.Buffer) error {

	if !ics.IsDefined(e.DTStamp) {
		return fmt.Errorf("DTstamp is required")
	}

	if !ics.IsDefined(e.UID) {
		return fmt.Errorf("UID is required")
	}

	tzIsDefined := ics.IsDefined(e.TZID) && e.TZID.Value != "UTC"

	if tzIsDefined {
		e.DTStart.Parameters = e.DTStart.Parameters.Prepend(parameter.TZID(e.TZID.Value))
		e.DTEnd.Parameters = e.DTEnd.Parameters.Prepend(parameter.TZID(e.TZID.Value))
	} else if !ics.IsDefined(e.TZID) || e.TZID.Value == "UTC" {
		e.DTStart = e.DTStart.UTC()
		e.DTEnd = e.DTEnd.UTC()
	}

	b.WriteLine("BEGIN:VEVENT")

	b.IfWriteValuerLine(true, "DTSTAMP", e.DTStamp)
	b.IfWriteValuerLine(ics.IsDefined(e.LastModified), "LAST-MODIFIED", e.LastModified)
	b.IfWriteValuerLine(true, "UID", e.UID)
	b.IfWriteValuerLine(tzIsDefined, "TZID", e.TZID)
	b.IfWriteValuerLine(ics.IsDefined(e.Organizer), "ORGANIZER", e.Organizer)

	for _, attendee := range e.Attendee {
		b.IfWriteValuerLine(true, "ATTENDEE", attendee)
	}

	b.IfWriteValuerLine(ics.IsDefined(e.Contact), "CONTACT", e.Contact)
	b.IfWriteValuerLine(ics.IsDefined(e.Sequence), "SEQUENCE", e.Sequence)
	b.IfWriteValuerLine(ics.IsDefined(e.Status), "STATUS", e.Status)
	b.IfWriteValuerLine(ics.IsDefined(e.Summary), "SUMMARY", e.Summary)
	b.IfWriteValuerLine(ics.IsDefined(e.Description), "DESCRIPTION", e.Description)
	b.IfWriteValuerLine(ics.IsDefined(e.Class), "CLASS", e.Class)
	b.IfWriteValuerLine(ics.IsDefined(e.Comment), "COMMENT", e.Comment)
	b.IfWriteValuerLine(ics.IsDefined(e.Location), "LOCATION", e.Location)
	b.IfWriteValuerLine(ics.IsDefined(e.RelatedTo), "RELATED-TO", e.RelatedTo)
	b.IfWriteValuerLine(ics.IsDefined(e.Transparency), "TRANSP", e.Transparency)
	b.IfWriteValuerLine(ics.IsDefined(e.DTStart), "DTSTART", e.DTStart)
	b.IfWriteValuerLine(ics.IsDefined(e.DTEnd), "DTEND", e.DTEnd)

	if ics.IsDefined(e.ALARM) {
		b.WriteLine("BEGIN:VALARM")
		b.IfWriteValuerLine(true, "TRIGGER", e.ALARM)
		b.WriteLine("ACTION:DISPLAY")
		b.WriteLine("END:VALARM")
	}

	b.WriteLine("END:VEVENT")

	return b.Flush()
}
