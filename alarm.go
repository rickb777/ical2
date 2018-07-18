package ical2

import (
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
)

// VAlarm is an alarm component.
type VAlarm interface {
	VComponent
	IsAlarm()
}

// VAudioAlarm captures a calendar event
type VAudioAlarm struct {
	Trigger  value.Trigger // required
	Duration value.DurationValue
	Repeat   value.IntegerValue
	Attach   value.Attachable // optional
}

// VDisplayAlarm captures a calendar event
type VDisplayAlarm struct {
	Description value.TextValue // required
	Trigger     value.Trigger   // required
	Duration    value.DurationValue
	Repeat      value.IntegerValue
}

// VEmailAlarm captures a calendar event
type VEmailAlarm struct {
	Description value.TextValue  // required
	Trigger     value.Trigger    // required
	Summary     value.TextValue  // required
	Attendee    []value.URIValue // required one or more
	Duration    value.DurationValue
	Repeat      value.IntegerValue
	Attach      []value.Attachable // optional
}

// IsAlarm marks this type.
func (e *VAudioAlarm) IsAlarm() {}

// EncodeIcal serialises the event to the buffer in iCal ics format
func (e *VAudioAlarm) EncodeIcal(b *ics.Buffer, method value.MethodValue) error {

	if !ics.IsDefined(e.Trigger) {
		return fmt.Errorf("Trigger is required")
	}

	if ics.IsDefined(e.Duration) != ics.IsDefined(e.Repeat) {
		return fmt.Errorf("Duration and Repeat must both be present or absent")
	}

	b.WriteLine("BEGIN:VALARM")
	b.WriteLine("ACTION:AUDIO")

	b.WriteValuerLine(true, "TRIGGER", e.Trigger)
	b.WriteValuerLine(ics.IsDefined(e.Duration), "DURATION", e.Duration)
	b.WriteValuerLine(ics.IsDefined(e.Repeat), "REPEAT", e.Repeat)
	b.WriteValuerLine(ics.IsDefined(e.Attach), "ATTACH", e.Attach)

	b.WriteLine("END:VALARM")

	return b.Flush()
}

// IsAlarm marks this type.
func (e *VDisplayAlarm) IsAlarm() {}

// EncodeIcal serialises the event to the buffer in iCal ics format
func (e *VDisplayAlarm) EncodeIcal(b *ics.Buffer, method value.MethodValue) error {

	if !ics.IsDefined(e.Description) {
		return fmt.Errorf("Description is required")
	}

	if !ics.IsDefined(e.Trigger) {
		return fmt.Errorf("Trigger is required")
	}

	if ics.IsDefined(e.Duration) != ics.IsDefined(e.Repeat) {
		return fmt.Errorf("Duration and Repeat must both be present or absent")
	}

	b.WriteLine("BEGIN:VALARM")
	b.WriteLine("ACTION:DISPLAY")

	b.WriteValuerLine(true, "DESCRIPTION", e.Description)
	b.WriteValuerLine(true, "TRIGGER", e.Trigger)
	b.WriteValuerLine(ics.IsDefined(e.Duration), "DURATION", e.Duration)
	b.WriteValuerLine(ics.IsDefined(e.Repeat), "REPEAT", e.Repeat)

	b.WriteLine("END:VALARM")

	return b.Flush()
}

// IsAlarm marks this type.
func (e *VEmailAlarm) IsAlarm() {}

// EncodeIcal serialises the event to the buffer in iCal ics format
func (e *VEmailAlarm) EncodeIcal(b *ics.Buffer, method value.MethodValue) error {

	if !ics.IsDefined(e.Description) {
		return fmt.Errorf("Description is required")
	}

	if !ics.IsDefined(e.Trigger) {
		return fmt.Errorf("Trigger is required")
	}

	if !ics.IsDefined(e.Summary) {
		return fmt.Errorf("Summary is required")
	}

	if len(e.Attendee) < 1 {
		return fmt.Errorf("At least one attendee is required")
	}

	if ics.IsDefined(e.Duration) != ics.IsDefined(e.Repeat) {
		return fmt.Errorf("Duration and Repeat must both be present or absent")
	}

	b.WriteLine("BEGIN:VALARM")
	b.WriteLine("ACTION:EMAIL")

	b.WriteValuerLine(true, "DESCRIPTION", e.Description)
	b.WriteValuerLine(true, "TRIGGER", e.Trigger)
	b.WriteValuerLine(true, "SUMMARY", e.Summary)
	for _, attendee := range e.Attendee {
		b.WriteValuerLine(true, "ATTENDEE", attendee)
	}
	b.WriteValuerLine(ics.IsDefined(e.Duration), "DURATION", e.Duration)
	b.WriteValuerLine(ics.IsDefined(e.Repeat), "REPEAT", e.Repeat)
	for _, attach := range e.Attach {
		b.WriteValuerLine(ics.IsDefined(attach), "ATTACH", attach)
	}

	b.WriteLine("END:VALARM")

	return b.Flush()
}
