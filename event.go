package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
)

// VEvent captures a calendar event
type VEvent struct {
	UID            value.TextValue
	DTStamp        value.DateTimeValue
	Start          value.DateTimeValue
	End            value.DateTimeValue
	Created        value.DateTimeValue
	LastModified   value.DateTimeValue
	ExceptionDate  value.DateTimeValue
	RecurrenceDate value.Temporal // DateTime or Period
	Organizer      value.URIValue
	Attendee       []value.URIValue
	Conference     []value.URIValue
	Contact        value.TextValue
	Summary        value.TextValue
	Description    value.TextValue
	Class          value.ClassValue // PUBLIC, PRIVATE, CONFIDENTIAL
	Comment        []value.TextValue
	RelatedTo      value.TextValue
	Categories     value.ListValue
	Resources      value.ListValue
	Sequence       value.IntegerValue
	Priority       value.IntegerValue // in the range 0 to 9; 0 is undefined; 1 is highest; 9 is lowest
	Status         value.StatusValue
	Location       value.TextValue
	Geo            value.GeoValue
	Transparency   value.TransparencyValue
	Color          value.TextValue // CSS3 color name
	Attach         []value.Attachable
	Image          []value.Attachable
	Alarm          []VAlarm

	// TODO (RFC5545) RECURRENCE-ID EXDATE RDATE RRULE
}

// AllDay changes the start and end to represent dates without time.
// If they are already configured as dates only, this has no effect.
func (e *VEvent) AllDay() *VEvent {
	e.Start = e.Start.AsDate()
	e.End = e.End.AsDate()
	return e
}

// EncodeIcal serialises the event to the buffer in iCalendar ics format
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
	for _, conference := range e.Conference {
		b.WriteValuerLine(true, "CONFERENCE", conference)
	}
	b.WriteValuerLine(ics.IsDefined(e.Contact), "CONTACT", e.Contact)
	b.WriteValuerLine(ics.IsDefined(e.Summary), "SUMMARY", e.Summary)
	b.WriteValuerLine(ics.IsDefined(e.Description), "DESCRIPTION", e.Description)
	b.WriteValuerLine(ics.IsDefined(e.Location), "LOCATION", e.Location)
	b.WriteValuerLine(ics.IsDefined(e.Geo), "GEO", e.Geo)
	b.WriteValuerLine(ics.IsDefined(e.Class), "CLASS", e.Class)
	for _, comment := range e.Comment {
		b.WriteValuerLine(ics.IsDefined(comment), "COMMENT", comment)
	}
	b.WriteValuerLine(ics.IsDefined(e.Created), "CREATED", e.Created)
	b.WriteValuerLine(ics.IsDefined(e.LastModified), "LAST-MODIFIED", e.LastModified)
	b.WriteValuerLine(ics.IsDefined(e.ExceptionDate), "EXDATE", e.ExceptionDate)
	b.WriteValuerLine(ics.IsDefined(e.RelatedTo), "RELATED-TO", e.RelatedTo)
	b.WriteValuerLine(ics.IsDefined(e.Categories), "CATEGORIES", e.Categories)
	b.WriteValuerLine(ics.IsDefined(e.Resources), "RESOURCES", e.Resources)
	b.WriteValuerLine(ics.IsDefined(e.Sequence), "SEQUENCE", e.Sequence)
	b.WriteValuerLine(ics.IsDefined(e.Priority), "PRIORITY", e.Priority)
	b.WriteValuerLine(ics.IsDefined(e.Status), "STATUS", e.Status)
	b.WriteValuerLine(ics.IsDefined(e.Transparency), "TRANSP", e.Transparency)
	b.WriteValuerLine(ics.IsDefined(e.Color), "COLOR", e.Color)
	for _, attachment := range e.Attach {
		b.WriteValuerLine(true, "ATTACH", attachment)
	}
	for _, image := range e.Image {
		b.WriteValuerLine(true, "IMAGE", image)
	}
	for _, alarm := range e.Alarm {
		alarm.EncodeIcal(b, method)
	}

	b.WriteLine("END:VEVENT")

	return b.Flush()
}
