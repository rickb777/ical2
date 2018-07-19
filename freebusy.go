package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
)

// VFreeBusy captures a calendar event
type VFreeBusy struct {
	UID       value.TextValue
	DTStamp   value.DateTimeValue
	Start     value.DateTimeValue
	End       value.DateTimeValue
	Organizer value.URIValue
	URL       value.URIValue
	Contact   value.TextValue
	Attendee  []value.URIValue
	Comment   []value.TextValue
	FreeBusy  []value.PeriodValue
	//TODO []rstatus
}

// EncodeIcal serialises the event to the buffer in iCalendar ics format
// (a VComponent method).
func (e *VFreeBusy) EncodeIcal(b *ics.Buffer, method value.MethodValue) error {

	if !ics.IsDefined(e.DTStamp) {
		return fmt.Errorf("DTstamp is required")
	}

	if !ics.IsDefined(e.UID) {
		return fmt.Errorf("UID is required")
	}

	if !ics.IsDefined(method) && !ics.IsDefined(e.Start) {
		return fmt.Errorf("When Method is undefined, Start is required")
	}

	b.WriteLine("BEGIN:VFREEBUSY")

	b.WriteValuerLine(ics.IsDefined(e.Start), "DTSTART", e.Start)
	b.WriteValuerLine(ics.IsDefined(e.End), "DTEND", e.End)
	b.WriteValuerLine(true, "DTSTAMP", e.DTStamp)
	b.WriteValuerLine(true, "UID", e.UID)
	b.WriteValuerLine(ics.IsDefined(e.Organizer), "ORGANIZER", e.Organizer)
	for _, attendee := range e.Attendee {
		b.WriteValuerLine(true, "ATTENDEE", attendee)
	}
	b.WriteValuerLine(ics.IsDefined(e.Contact), "CONTACT", e.Contact)
	b.WriteValuerLine(ics.IsDefined(e.URL), "URL", e.URL)
	for _, comment := range e.Comment {
		b.WriteValuerLine(ics.IsDefined(comment), "COMMENT", comment)
	}
	for _, fb := range e.FreeBusy {
		b.WriteValuerLine(ics.IsDefined(fb), "FREEBUSY", fb)
	}

	b.WriteLine("END:VFREEBUSY")

	return b.Flush()
}
