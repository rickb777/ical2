package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"strconv"
	"time"
)

type simpleValue struct {
	Parameters parameter.Parameters
	Value      string
}

// IsDefined tests whether the value has been explicitly defined or is default.
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

// With appends parameters to the value.
func (v CalAddressValue) With(params ...parameter.Parameter) CalAddressValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v CalAddressValue) WriteTo(w ics.StringWriter) error {
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
	return DurationValue{simpleValue{Value: d}}.With(parameter.Type("DURATION"))
}

// With appends parameters to the value.
func (v DurationValue) With(params ...parameter.Parameter) DurationValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v DurationValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value)
	return e
}

//-------------------------------------------------------------------------------------------------

// IntegerValue holds an integer.
type IntegerValue struct {
	Parameters parameter.Parameters
	Value      int
	defined    bool
}

// Integer returns a new IntegerValue.
func Integer(d int) IntegerValue {
	return IntegerValue{Value: d, defined: true}
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v IntegerValue) IsDefined() bool {
	return v.defined
}

// With appends parameters to the value.
func (v IntegerValue) With(params ...parameter.Parameter) IntegerValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v IntegerValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(strconv.Itoa(v.Value))
	return e
}

//-------------------------------------------------------------------------------------------------

const (
	// Calendar class property
	PUBLIC       = "PUBLIC"
	PRIVATE      = "PRIVATE"
	CONFIDENTIAL = "CONFIDENTIAL"

	// Event Transparency - does not block other events
	TRANSPARENT = "TRANSPARENT"
	// Event Transparency - blocks other events
	OPAQUE = "OPAQUE"

	// Event status
	TENTATIVE = "TENTATIVE"
	// Event status
	CONFIRMED = "CONFIRMED"
	// Event and To-do status
	CANCELLED = "CANCELLED"

	// To-do status
	NEEDS_ACTION = "NEEDS-ACTION"
	// To-do status
	COMPLETED = "COMPLETED"
	// To-do status
	IN_PROCESS = "IN-PROCESS"
)

// TextValue holds a value that is a string and which will be escaped
// when written out.
type TextValue struct {
	simpleValue
}

// Text constructs a new text value.
func Text(v string) TextValue {
	return TextValue{simpleValue{Value: v}}
}

// With appends parameters to the value.
func (v TextValue) With(params ...parameter.Parameter) TextValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v TextValue) WriteTo(w ics.StringWriter) error {
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
	Parameters parameter.Parameters
	Value      time.Time
	format     string
}

// Date constructs a new date value, i.e. without time.
func Date(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateLayout}.With(parameter.Type(parameter.DATE))
}

// DateTime constructs a new date+time value.
func DateTime(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateTimeLayout}.With(parameter.Type(parameter.DATE_TIME))
}

// TStamp constructs a date-time value using UTC.
func TStamp(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t.UTC(), format: dateTimeLayout + "Z"}
}

// AsDate converts a date+time value to a date-only value.
func (v DateTimeValue) AsDate() DateTimeValue {
	v.format = dateLayout
	return v.With(parameter.Type(parameter.DATE))
}

// UTC converts a date+time value to be shown as UTC ('zulu'). It does not alter the
// time value, which should already be UTC.
func (v DateTimeValue) UTC() DateTimeValue {
	if v.format == dateTimeLayout {
		v.format = dateTimeLayout + "Z"
	}
	return v
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v DateTimeValue) IsDefined() bool {
	return !v.Value.IsZero()
}

// With appends parameters to the value.
func (v DateTimeValue) With(params ...parameter.Parameter) DateTimeValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v DateTimeValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value.Format(v.format))
	return e
}
