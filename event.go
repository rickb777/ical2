package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
)

// VEvent captures a calendar event.
// https://tools.ietf.org/html/rfc5545#section-3.6.1
type VEvent struct {
	// Start specifies when the calendar component begins.
	// https://tools.ietf.org/html/rfc5545#section-3.8.2.4
	Start value.DateTimeValue

	// End specifies when the calendar component ends, which must be after the start.
	// Use either End or Duration but not both.
	// https://tools.ietf.org/html/rfc5545#section-3.8.2.2
	End value.DateTimeValue

	// Duration specifies a positive duration of time.
	// Use either End or Duration but not both.
	// https://tools.ietf.org/html/rfc5545#section-3.8.2.5
	Duration value.DurationValue

	// https://tools.ietf.org/html/rfc5545#section-3.8.7.1
	Created value.DateTimeValue
	// https://tools.ietf.org/html/rfc5545#section-3.8.7.2
	DTStamp value.DateTimeValue
	// https://tools.ietf.org/html/rfc5545#section-3.8.7.3
	LastModified value.DateTimeValue

	ExceptionDate  []value.DateTimeValue
	RecurrenceDate []value.Temporal // DateTime or Period
	RecurrenceRule value.RecurrenceValue

	// https://tools.ietf.org/html/rfc7986#section-5.11
	Conference []value.URIValue

	// Attendee defines an "Attendee" within a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.1
	Attendee []value.URIValue

	// Organizer defines the organizer for a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.3
	Organizer value.URIValue

	// Contact is used to represent contact information or alternately a reference to contact information
	// associated with the calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.2
	Contact []value.TextValue

	// https://tools.ietf.org/html/rfc5545#section-3.8.1.12
	Summary value.TextValue

	// Description provides a more complete description of the calendar component than
	// that provided by the "SUMMARY" property.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.5
	Description value.TextValue

	// Class defines the access classification for a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.3
	Class value.TextValue // PUBLIC, PRIVATE, CONFIDENTIAL, etc

	// Comment provides non-processing information intended as a comment to the calendar user.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.4
	Comment []value.TextValue

	// RelatedTo is used to represent a relationship or reference between one calendar
	// component and another. It consists of the persistent, globally unique identifier
	// of another calendar component.  This value would be represented in a calendar
	// component by the "UID" property.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.5
	RelatedTo value.TextValue

	// URL defines a Uniform Resource Locator (URL) associated with the iCalendar object.
	// This implementation always includes the "VALUE=URI" parameter, although some others
	// do not.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.6
	URL value.URIValue

	// UID defines the persistent, globally unique identifier for the calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.4.7
	UID value.TextValue

	// Categories specify categories or subtypes of the calendar component.  The categories are useful
	// in searching for a calendar component of a particular type and category.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.2
	// https://tools.ietf.org/html/rfc7986#section-5.6
	Categories []value.ListValue

	// Resources lists the equipment or resources anticipated for an activity specified
	// by a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.10
	Resources []value.ListValue

	// Sequence defines the revision sequence number of the calendar component within a
	// sequence of revisions.
	// https://tools.ietf.org/html/rfc5545#section-3.8.7.4
	Sequence value.IntegerValue

	// In the range 0 to 9; 0 is undefined; 1 is highest; 9 is lowest.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.9
	Priority value.IntegerValue

	// Status defines the overall status or confirmation for the calendar component.
	// Examples: "TENTATIVE", "CONFIRMED".
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.11
	Status value.TextValue

	// Location defines the intended venue for the activity defined by a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.7
	Location value.TextValue

	// Geo specifies information related to the global position for the activity specified
	// by a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.6
	Geo value.GeoValue

	// Transparency defines whether or not an event is transparent to busy time searches.
	// https://tools.ietf.org/html/rfc5545#section-3.8.2.7
	Transparency value.TextValue

	// Color specifies a color used for displaying the event data. The value is CSS3 color name.
	// Also allowed in the enclosing calendar data.
	// https://tools.ietf.org/html/rfc7986#section-5.9
	Color value.TextValue // CSS3 color name

	// Attach provides the capability to associate a document object with a calendar component.
	// https://tools.ietf.org/html/rfc5545#section-3.8.1.1
	Attach []value.Attachable

	// Image specifies an image or images associated with the calendar or a calendar component.
	// https://tools.ietf.org/html/rfc7986#section-5.10
	Image []value.Attachable

	Alarm []VAlarm

	// TODO (RFC5545) RECURRENCE-ID
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
func (e *VEvent) EncodeIcal(b *ics.Buffer, method value.TextValue) error {

	if !ics.IsDefined(e.DTStamp) {
		return fmt.Errorf("DTstamp is required")
	}

	if !ics.IsDefined(e.UID) {
		return fmt.Errorf("UID is required")
	}

	if !ics.IsDefined(method) && !ics.IsDefined(e.Start) {
		return fmt.Errorf("when Method is undefined, Start is required")
	}

	if ics.IsDefined(e.End) && ics.IsDefined(e.Duration) {
		return fmt.Errorf("End and Duration are exclusive; only one can be set")
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
	for _, contact := range e.Contact {
		b.WriteValuerLine(true, "CONTACT", contact)
	}
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
	for _, date := range e.ExceptionDate {
		b.WriteValuerLine(true, "EXDATE", date)
	}
	for _, date := range e.RecurrenceDate {
		b.WriteValuerLine(true, "RDATE", date)
	}
	b.WriteValuerLine(ics.IsDefined(e.RecurrenceRule), "RRULE", e.RecurrenceRule)
	b.WriteValuerLine(ics.IsDefined(e.RelatedTo), "RELATED-TO", e.RelatedTo)
	for _, cat := range e.Categories {
		b.WriteValuerLine(true, "CATEGORIES", cat)
	}
	for _, res := range e.Resources {
		b.WriteValuerLine(true, "RESOURCES", res)
	}
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
