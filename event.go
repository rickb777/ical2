package ical2

import (
	"fmt"
)

// VEvent captures a calendar event
type VEvent struct {
	UID          TextValue
	DTStamp      DateTimeValue
	DTStart      DateTimeValue
	DTEnd        DateTimeValue
	LastModified DateTimeValue
	Organizer    CalAddressValue
	Attendee     []CalAddressValue
	Contact      TextValue
	Summary      TextValue
	Description  TextValue
	Class        TextValue // PUBLIC, PRIVATE, CONFIDENTIAL
	Comment      TextValue
	RelatedTo    TextValue
	TZID         TextValue
	Sequence     Valuer
	Status       TextValue
	ALARM        TextValue
	Location     TextValue
	Transparency TextValue
	Color        TextValue // CSS3 color name

	// TODO (RFC5545) CREATED GEO PRIORITY RECURRENCE-ID EXDATE RDATE RRULE
	// TODO (RFC7986) []CONFERENCE
}

func (e *VEvent) AllDay() *VEvent {
	e.DTStart = e.DTStart.AsDate()
	e.DTEnd = e.DTEnd.AsDate()
	return e
}

func (e *VEvent) EncodeIcal(b *Buffer) error {

	if !IsDefined(e.DTStamp) {
		return fmt.Errorf("DTstamp is required")
	}

	if !IsDefined(e.UID) {
		return fmt.Errorf("UID is required")
	}

	tzIsDefined := IsDefined(e.TZID) && e.TZID.Value != "UTC"

	if tzIsDefined {
		e.DTStart.Parameters = e.DTStart.Parameters.Prepend(TZID(e.TZID.Value))
		e.DTEnd.Parameters = e.DTEnd.Parameters.Prepend(TZID(e.TZID.Value))
	} else if !IsDefined(e.TZID) || e.TZID.Value == "UTC" {
		e.DTStart = e.DTStart.UTC()
		e.DTEnd = e.DTEnd.UTC()
	}

	b.WriteLine("BEGIN:VEVENT")

	b.IfWriteValuerLine(true, "DTSTAMP", e.DTStamp)
	b.IfWriteValuerLine(IsDefined(e.LastModified), "LAST-MODIFIED", e.LastModified)
	b.IfWriteValuerLine(true, "UID", e.UID)
	b.IfWriteValuerLine(tzIsDefined, "TZID", e.TZID)
	b.IfWriteValuerLine(IsDefined(e.Organizer), "ORGANIZER", e.Organizer)

	for _, attendee := range e.Attendee {
		b.IfWriteValuerLine(true, "ATTENDEE", attendee)
	}

	b.IfWriteValuerLine(IsDefined(e.Contact), "CONTACT", e.Contact)
	b.IfWriteValuerLine(IsDefined(e.Sequence), "SEQUENCE", e.Sequence)
	b.IfWriteValuerLine(IsDefined(e.Status), "STATUS", e.Status)
	b.IfWriteValuerLine(IsDefined(e.Summary), "SUMMARY", e.Summary)
	b.IfWriteValuerLine(IsDefined(e.Description), "DESCRIPTION", e.Description)
	b.IfWriteValuerLine(IsDefined(e.Class), "CLASS", e.Class)
	b.IfWriteValuerLine(IsDefined(e.Comment), "COMMENT", e.Comment)
	b.IfWriteValuerLine(IsDefined(e.Location), "LOCATION", e.Location)
	b.IfWriteValuerLine(IsDefined(e.RelatedTo), "RELATED-TO", e.RelatedTo)
	b.IfWriteValuerLine(IsDefined(e.Transparency), "TRANSP", e.Transparency)
	b.IfWriteValuerLine(IsDefined(e.DTStart), "DTSTART", e.DTStart)
	b.IfWriteValuerLine(IsDefined(e.DTEnd), "DTEND", e.DTEnd)

	if IsDefined(e.ALARM) {
		b.WriteLine("BEGIN:VALARM")
		b.IfWriteValuerLine(true, "TRIGGER", e.ALARM)
		b.WriteLine("ACTION:DISPLAY")
		b.WriteLine("END:VALARM")
	}

	b.WriteLine("END:VEVENT")

	return b.Flush()
}
