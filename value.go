package ical2

import (
	"bytes"
	"time"
)

// Valuer holds an iCal value.
type Valuer interface {
	IsDefined() bool
	WriteTo(w StringWriter) error
}

// IsDefined tests whether a valuer is defined.
func IsDefined(v Valuer) bool {
	return v != nil && v.IsDefined()
}

//-------------------------------------------------------------------------------------------------

type simpleValue struct {
	Parameters Parameters
	Value      string
}

func (v simpleValue) IsDefined() bool {
	return v.Value != ""
}

//-------------------------------------------------------------------------------------------------

// CalAddressValue holds calendar address, which is typically an email address.
type CalAddressValue struct {
	simpleValue
}

// CalAddress returns a new CalAddressValue.
func CalAddress(mailto string) CalAddressValue {
	return CalAddressValue{simpleValue{Value: mailto}}
}

func (v CalAddressValue) With(params ...Parameter) CalAddressValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

func (v CalAddressValue) WriteTo(w StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteString(":mailto:")
	_, e := w.WriteString(v.Value)
	return e
}

//-------------------------------------------------------------------------------------------------

// DurationValue holds a time duration. This should be in ISO-8601 form
// (https://en.wikipedia.org/wiki/ISO_8601#Durations);
// see github.com/rickb777/date/period for a compatible duration API.
type DurationValue struct {
	simpleValue
}

// Duration returns a new DurationValue.
func Duration(d string) DurationValue {
	return DurationValue{simpleValue{Value: d}}.With(Type("DURATION"))
}

func (v DurationValue) With(params ...Parameter) DurationValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

func (v DurationValue) WriteTo(w StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value)
	return e
}

//-------------------------------------------------------------------------------------------------

const (
	// Calendar class property
	PUBLIC = "PUBLIC"
	PRIVATE = "PRIVATE"
	CONFIDENTIAL = "CONFIDENTIAL"

	// Event Transparency - does not block other events
	TRANSPARENT  = "TRANSPARENT"
	// Event Transparency - blocks other events
	OPAQUE       = "OPAQUE"

	// Event status
	TENTATIVE    = "TENTATIVE"
	// Event status
	CONFIRMED    = "CONFIRMED"
	// Event and To-do status
	CANCELLED    = "CANCELLED"

	// To-do status
	NEEDS_ACTION = "NEEDS-ACTION"
	// To-do status
	COMPLETED    = "COMPLETED"
	// To-do status
	IN_PROCESS   = "IN-PROCESS"
)

type TextValue struct {
	simpleValue
}

func Text(v string) TextValue {
	return TextValue{simpleValue{Value: v}}
}

func (v TextValue) With(params ...Parameter) TextValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

func (v TextValue) WriteTo(w StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(escapeText(v.Value))
	return e
}

//-------------------------------------------------------------------------------------------------

// escapeText implements the escaping of semicolon, comma, backslash and
// newline. See https://tools.ietf.org/html/rfc5545#section-3.3.11
func escapeText(s string) string {
	if len(s) == 0 {
		return ""
	}

	w := &bytes.Buffer{}

	// treat s as a sequence of bytes, not runes
	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case '\\', ';', ',':
			w.WriteByte('\\')

		case '\n':
			w.WriteByte('\\')
			c = 'n'
		}

		w.WriteByte(c)
	}

	return w.String()
}

//-------------------------------------------------------------------------------------------------

const (
	dateLayout     = "20060102"
	dateTimeLayout = "20060102T150405"
	//timeLayout     = "150405"
)

// DateTimeValue holds a date/time and its formatting decision.
type DateTimeValue struct {
	Parameters Parameters
	Value      time.Time
	format     string
}

func Date(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateLayout}.With(Type(DATE))
}

func DateTime(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateTimeLayout}.With(Type(DATE_TIME))
}

func TStamp(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t.UTC(), format: dateTimeLayout + "Z"}
}

func (v DateTimeValue) AsDate() DateTimeValue {
	v.format = dateLayout
	return v.With(Type(DATE))
}

func (v DateTimeValue) UTC() DateTimeValue {
	if v.format == dateTimeLayout {
		v.format = dateTimeLayout + "Z"
	}
	return v
}

func (v DateTimeValue) IsDefined() bool {
	return !v.Value.IsZero()
}

func (v DateTimeValue) With(params ...Parameter) DateTimeValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

func (v DateTimeValue) WriteTo(w StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value.Format(v.format))
	return e
}

