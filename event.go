package ical

import (
	"io"
	"time"
	"fmt"
)

const (
	tstampLayout   = "20060102T150405Z"
	dateLayout     = "20060102"
	dateTimeLayout = "20060102T150405"
)

// Party is a person, typically an organiser or attendee.
type Party struct {
	Name  string
	Email string
}

func (p Party) String() string {
	if p.Name == "" {
		return fmt.Sprintf(":MAILTO=%s", p.Email)
	}
	return fmt.Sprintf(";CN=%s:MAILTO=%s", p.Name, p.Email)
}

// VEvent captures a calendar event
type VEvent struct {
	UID         string
	DTSTAMP     time.Time
	DTSTART     time.Time
	DTEND       time.Time
	ORGANIZER   Party
	ATTENDEE    []Party
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

	if len(e.TZID) != 0 && e.TZID != "UTC" {
		if err := b.WriteLine("TZID:", e.TZID); err != nil {
			return err
		}
	}

	if e.ORGANIZER.Email != "" {
		if err := b.WriteLine("ORGANIZER", e.ORGANIZER.String()); err != nil {
			return err
		}
	}

	if len(e.ATTENDEE) > 0 {
		for _, attendee := range e.ATTENDEE {
			if err := b.WriteLine("ATTENDEE", attendee.String()); err != nil {
				return err
			}
		}
	}

	if e.SEQUENCE != "" {
		if err := b.WriteLine("SEQUENCE:", e.SEQUENCE); err != nil {
			return err
		}
	}

	if e.STATUS != "" {
		if err := b.WriteLine("STATUS:", e.STATUS); err != nil {
			return err
		}
	}

	if err := b.WriteLine("SUMMARY:", e.SUMMARY); err != nil {
		return err
	}

	if e.DESCRIPTION != "" {
		if err := b.WriteLine("DESCRIPTION:", e.DESCRIPTION); err != nil {
			return err
		}
	}

	if e.LOCATION != "" {
		if err := b.WriteLine("LOCATION:", e.LOCATION); err != nil {
			return err
		}
	}

	if e.TRANSP != "" {
		if err := b.WriteLine("TRANSP:", e.TRANSP); err != nil {
			return err
		}
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
