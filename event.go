package ical

import (
	"io"
	"time"
)

const (
	tstampLayout   = "20060102T150405Z"
	dateLayout     = "20060102"
	dateTimeLayout = "20060102T150405"
)

type Individual struct {
	Name  string
	Dir   string
	Email string
}

func (p Individual) WriteParams(b *buffer) (err error) {
	if _, err = b.IfWriteString(p.Name != "", ";CN=", p.Name); err != nil {
		return err
	}
	if _, err = b.IfWriteString(p.Dir != "", ";DIR=", p.Dir); err != nil {
		return err
	}
	return nil
}

func (p Individual) WriteLine(b *buffer, property string) (err error) {
	if p.Email == "" {
		return nil
	}
	if _, err = b.WriteString(property); err != nil {
		return err
	}
	if err = p.WriteParams(b); err != nil {
		return err
	}
	return b.WriteLine(":mailto=", p.Email)
}

//-------------------------------------------------------------------------------------------------

type Attendee struct {
	Individual
	Role string // CHAIR, REQ-PARTICIPANT, OPT-PARTICIPANT, NON-PARTICIPANT etc
}

func (p Attendee) WriteParams(b *buffer) (err error) {
	if _, err = b.IfWriteString(p.Role != "", ";ROLE=", p.Role); err != nil {
		return err
	}
	if err = p.Individual.WriteParams(b); err != nil {
		return err
	}
	return nil
}

func (p Attendee) WriteLine(b *buffer, property string) (err error) {
	if p.Email == "" {
		return nil
	}
	if _, err = b.WriteString(property); err != nil {
		return err
	}
	if err = p.WriteParams(b); err != nil {
		return err
	}
	return b.WriteLine(":mailto=", p.Email)
}

//-------------------------------------------------------------------------------------------------

// VEvent captures a calendar event
type VEvent struct {
	UID         string
	DTSTAMP     time.Time
	DTSTART     time.Time
	DTEND       time.Time
	ORGANIZER   Individual
	ATTENDEE    []Attendee
	CONTACT     string
	SUMMARY     string
	DESCRIPTION string
	TZID        string
	SEQUENCE    string
	STATUS      string
	ALARM       string
	LOCATION    string
	TRANSP      string

	AllDay bool
}

func (e *VEvent) EncodeIcal(w io.Writer) error {

	var timeStampLayout, timeStampType, tzidTxt string

	if e.AllDay {
		timeStampLayout = dateLayout
		timeStampType = "DATE"
	} else {
		timeStampLayout = dateTimeLayout
		timeStampType = "DATE-TIME"
		if len(e.TZID) == 0 || e.TZID == "UTC" {
			timeStampLayout = timeStampLayout + "Z"
		}
	}

	if len(e.TZID) != 0 && e.TZID != "UTC" {
		tzidTxt = "TZID=" + e.TZID + ";"
	}

	b := newBuffer(w)
	if err := b.WriteLine("BEGIN:VEVENT"); err != nil {
		return err
	}

	if err := b.WriteLine("DTSTAMP:", e.DTSTAMP.UTC().Format(tstampLayout)); err != nil {
		return err
	}

	if err := b.WriteLine("UID:", e.UID); err != nil {
		return err
	}

	if err := b.IfWriteLine(len(e.TZID) != 0 && e.TZID != "UTC", "TZID:", e.TZID); err != nil {
		return err
	}

	if err := e.ORGANIZER.WriteLine(b, "ORGANIZER"); err != nil {
		return err
	}

	for _, attendee := range e.ATTENDEE {
		if err := attendee.WriteLine(b, "ATTENDEE"); err != nil {
			return err
		}
	}

	if err := b.IfWriteLine(e.CONTACT != "", "CONTACT:", e.CONTACT); err != nil {
		return err
	}

	if err := b.IfWriteLine(e.SEQUENCE != "", "SEQUENCE:", e.SEQUENCE); err != nil {
		return err
	}

	if err := b.IfWriteLine(e.STATUS != "", "STATUS:", e.STATUS); err != nil {
		return err
	}

	if err := b.WriteLine("SUMMARY:", e.SUMMARY); err != nil {
		return err
	}

	if err := b.IfWriteLine(e.DESCRIPTION != "", "DESCRIPTION:", e.DESCRIPTION); err != nil {
		return err
	}

	if err := b.IfWriteLine(e.LOCATION != "", "LOCATION:", e.LOCATION); err != nil {
		return err
	}

	if err := b.IfWriteLine(e.TRANSP != "", "TRANSP:", e.TRANSP); err != nil {
		return err
	}

	if err := b.WriteLine("DTSTART;", tzidTxt, "VALUE=", timeStampType, ":", e.DTSTART.Format(timeStampLayout)); err != nil {
		return err
	}

	if err := b.WriteLine("DTEND;", tzidTxt, "VALUE=", timeStampType, ":", e.DTEND.Format(timeStampLayout)); err != nil {
		return err
	}

	if e.ALARM != "" {
		if err := b.WriteLine("BEGIN:VALARM"); err != nil {
			return err
		}
		if err := b.WriteLine("TRIGGER:", e.ALARM); err != nil {
			return err
		}
		if err := b.WriteLine("ACTION:DISPLAY"); err != nil {
			return err
		}
		if err := b.WriteLine("END:VALARM"); err != nil {
			return err
		}
	}

	if err := b.WriteLine("END:VEVENT"); err != nil {
		return err
	}

	return b.Flush()
}
