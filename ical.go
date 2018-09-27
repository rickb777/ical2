// Package ical2 provides a data model for the iCalendar specification. Marshalling
// to the textual iCalendar ics format is implemented. Unmarshalling is not currently
// supported.
//
// See
// https://tools.ietf.org/html/rfc5545
// https://tools.ietf.org/html/rfc6868
// https://tools.ietf.org/html/rfc7986.
//
// Availability (https://tools.ietf.org/html/rfc7953) is not supported.
package ical2

import (
	"bytes"
	"fmt"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/value"
	"io"
	"os"
)

/*
  http://tools.ietf.org/html/draft-daboo-icalendar-extensions-09
  http://stackoverflow.com/a/17187346/195141

  BEGIN:VCALENDAR
  VERSION:2.0
  PRODID:-//My Company//NONSGML Event Calendar//EN
  URL:http://my.calendar/url
  NAME:My Calendar Name
  X-WR-CALNAME:My Calendar Name
  DESCRIPTION:A description of my calendar
  X-WR-CALDESC:A description of my calendar
  TIMEZONE-ID:Europe/London
  X-WR-TIMEZONE:Europe/London
  REFRESH-INTERVAL;VALUE=DURATION:PT12H
  X-PUBLISHED-TTL:PT12H
  COLOR:34:50:105
  CALSCALE:GREGORIAN
  METHOD:PUBLISH
*/

// VCalendar is a calendar as per RFC-5545 https://tools.ietf.org/html/rfc5545.
type VCalendar struct {
	// RFC-5545 properties
	Version  value.TextValue // 2.0
	ProdId   value.TextValue // -//My Company//NONSGML Event Calendar//EN
	Method   value.TextValue // PUBLISH, REQUEST,
	CalScale value.TextValue // GREGORIAN

	// RFC-7986 properties
	Name            value.TextValue     // My Calendar Name
	Description     value.TextValue     // A description of my calendar
	URL             value.TextValue     // http://my.calendar/url
	LastModified    value.DateTimeValue // can also be specified per VComponent
	RecurrenceId    value.DateTimeValue
	RefreshInterval value.DurationValue // PT12H
	Color           value.TextValue     // CSS3 color name
	// TODO CATEGORIES, SOURCE, []IMAGE

	//X_WR_CALNAME string // My Calendar Name
	//X_WR_CALDESC string // A description of my calendar
	//X_WR_TIMEZONE string // Europe/London
	//X_PUBLISHED_TTL  string // PT12H
	Extensions []Extension

	VComponent []VComponent
}

// NewVCalendar constructs a new VCalendar with the required properties set.
// The version is set to 2.0 and the calendar scale is Gregorian.
func NewVCalendar(prodId string) *VCalendar {
	return &VCalendar{
		Version:  value.Text("2.0"),
		ProdId:   value.Text(prodId),
		CalScale: value.Text("GREGORIAN"),
	}
}

// Extend adds an extension property to the calendar.
// The VCalendar modified and is returned.
func (c *VCalendar) Extend(key string, value ics.Valuer) *VCalendar {
	c.Extensions = append(c.Extensions, Extension{key, value})
	return c
}

// With associates a component with the calendar.
// The VCalendar modified and is returned.
func (c *VCalendar) With(component VComponent) *VCalendar {
	c.VComponent = append(c.VComponent, component)
	return c
}

// doEncode encodes the calendar in ICS format, writing it to some Writer. The
// lineEnding can be "" or "\r\n" for normal iCalendar formatting, or "\n" in other cases.
func (c *VCalendar) doEncode(w io.Writer, lineEnding string) error {
	b := ics.NewBuffer(w, lineEnding)

	b.WriteLine("BEGIN:VCALENDAR")

	b.WriteValuerLine(true, "PRODID", c.ProdId)
	b.WriteValuerLine(true, "VERSION", c.Version)
	b.WriteValuerLine(ics.IsDefined(c.CalScale), "CALSCALE", c.CalScale)
	b.WriteValuerLine(ics.IsDefined(c.Method), "METHOD", c.Method)
	b.WriteValuerLine(ics.IsDefined(c.Name), "NAME", c.Name)
	b.WriteValuerLine(ics.IsDefined(c.Description), "DESCRIPTION", c.Description)
	b.WriteValuerLine(ics.IsDefined(c.URL), "URL", c.URL)
	b.WriteValuerLine(ics.IsDefined(c.LastModified), "LAST-MODIFIED", c.LastModified)
	b.WriteValuerLine(ics.IsDefined(c.RecurrenceId), "RECURRENCE-ID", c.RecurrenceId)
	b.WriteValuerLine(ics.IsDefined(c.Color), "COLOR", c.Color)
	b.WriteValuerLine(ics.IsDefined(c.RefreshInterval), "REFRESH-INTERVAL", c.RefreshInterval)

	for _, extension := range c.Extensions {
		b.WriteValuerLine(true, extension.Key, extension.Value)
	}

	for _, component := range c.VComponent {
		if err := component.EncodeIcal(b, c.Method); err != nil {
			return err
		}
	}

	b.WriteLine("END:VCALENDAR")

	return b.Flush()
}

// Encode encodes the calendar in ICS format, writing it to some Writer. The
// line endings are "\r\n" for normal iCalendar transmission purposes.
func (c *VCalendar) Encode(w io.Writer) error {
	return c.doEncode(w, "\r\n")
}

// EncodePlain encodes the calendar in ICS format, writing it to some Writer. The
// line ending are "\n" for non-transmission purposes, e.g. for viewing.
func (c *VCalendar) EncodePlain(w io.Writer) error {
	return c.doEncode(w, "\n")
}

// String returns the ICS formatted content, albeit using "\n" line endings.
func (c *VCalendar) String() string {
	buf := &bytes.Buffer{}
	err := c.EncodePlain(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return buf.String()
}

// VComponent is an item that belongs to a calendar.
type VComponent interface {
	EncodeIcal(b *ics.Buffer, method value.TextValue) error
}

// Extension is a key/value struct for any additional non-standard or unsupported
// calendar properties.
type Extension struct {
	Key   string
	Value ics.Valuer
}
