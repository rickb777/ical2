// package ical provides a data model for the iCal specification.
//
// See
// https://tools.ietf.org/html/rfc5545
// https://tools.ietf.org/html/rfc6868
// https://tools.ietf.org/html/rfc7986.
//
// Availability (https://tools.ietf.org/html/rfc7953) is not supported.
package ical2

import (
	"io"
	"bytes"
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
	Version  TextValue // 2.0
	ProdId   TextValue // -//My Company//NONSGML Event Calendar//EN
	Method   TextValue // PUBLISH
	CalScale TextValue // GREGORIAN

	// RFC-7986 properties
	Name            TextValue // My Calendar Name
	Description     TextValue // A description of my calendar
	UID             TextValue
	URL             TextValue // http://my.calendar/url
	LastModified    DateTimeValue
	RefreshInterval DurationValue // PT12H
	Color           TextValue     // CSS3 color name
	// TODO CATEGORIES, SOURCE, []IMAGE

	TimezoneId string // Europe/London

	//X_WR_CALNAME string // My Calendar Name
	//X_WR_CALDESC string // A description of my calendar
	//X_WR_TIMEZONE string // Europe/London
	//X_PUBLISHED_TTL  string // PT12H
	Extensions []Extension

	VComponent []VComponent
}

func NewBasicVCalendar(prodId string) *VCalendar {
	return &VCalendar{
		Version:  Text("2.0"),
		ProdId:   Text(prodId),
		CalScale: Text("GREGORIAN"),
	}
}

func (c *VCalendar) Extend(key string, value Valuer) *VCalendar {
	c.Extensions = append(c.Extensions, Extension{key, value})
	return c
}

func (c *VCalendar) With(component VComponent) *VCalendar {
	c.VComponent = append(c.VComponent, component)
	return c
}

// DoEncode encodes the calendar in ICS format, writing it to some Writer. The
// lineEnding can be "" or "\r\n" for normal iCal formatting, or "\n" in other cases.
func (c *VCalendar) DoEncode(w io.Writer, lineEnding string) error {
	b := NewBuffer(w, lineEnding)

	b.WriteLine("BEGIN:VCALENDAR")

	b.IfWriteValuerLine(true, "VERSION", c.Version)
	b.IfWriteValuerLine(true, "PRODID", c.ProdId)
	b.IfWriteValuerLine(IsDefined(c.CalScale), "CALSCALE", c.CalScale)
	b.IfWriteValuerLine(IsDefined(c.Method), "METHOD", c.Method)
	b.IfWriteValuerLine(IsDefined(c.Name), "NAME", c.Name)
	b.IfWriteValuerLine(IsDefined(c.Description), "DESCRIPTION", c.Description)
	b.IfWriteValuerLine(IsDefined(c.UID), "UID", c.UID)
	b.IfWriteValuerLine(IsDefined(c.URL), "URL", c.URL)
	b.IfWriteValuerLine(IsDefined(c.LastModified), "LAST-MODIFIED", c.LastModified)
	b.IfWriteValuerLine(IsDefined(c.Color), "COLOR", c.Color)
	b.IfWriteValuerLine(IsDefined(c.RefreshInterval), "REFRESH-INTERVAL", c.RefreshInterval)

	for _, extension := range c.Extensions {
		b.IfWriteValuerLine(true, extension.Key, extension.Value)
	}

	for _, component := range c.VComponent {
		if err := component.EncodeIcal(b); err != nil {
			return err
		}
	}

	b.WriteLine("END:VCALENDAR")

	return b.Flush()
}

// Encode encodes the calendar in ICS format, writing it to some Writer. The
// line endings are "\r\n" for normal iCal transmission purposes.
func (c *VCalendar) Encode(w io.Writer) error {
	return c.DoEncode(w, "\r\n")
}

// EncodePlain encodes the calendar in ICS format, writing it to some Writer. The
// line ending are "\n" for non-transmission purposes, e.g. for viewing.
func (c *VCalendar) EncodePlain(w io.Writer) error {
	return c.DoEncode(w, "\n")
}

// String returns the ICS formatted content, albeit using "\n" line endings.
func (c *VCalendar) String() string {
	buf := &bytes.Buffer{}
	c.EncodePlain(buf)
	return buf.String()
}

// VComponent is an item that belongs to a calendar.
type VComponent interface {
	EncodeIcal(b *Buffer) error
}

type Extension struct {
	Key   string
	Value Valuer
}
