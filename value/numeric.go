package value

import (
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"strconv"
	"time"
)

const (
	dateLayout     = "20060102"
	dateTimeLayout = "20060102T150405"
	//timeLayout     = "150405"
)

// DateTimeValue holds a date/time and its formatting decision.
// See https://tools.ietf.org/html/rfc5545#section-3.3.5
type DateTimeValue struct {
	Parameters parameter.Parameters
	Value      time.Time
	format     string
}

// DateTime constructs a new date-time value. This is represented as a "floating"
// local time. It has VALUE=DATE-TIME.
func DateTime(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateTimeLayout}.With(parameter.Type(parameter.DATE_TIME_TYPE))
}

// Date constructs a new date value, i.e. without time. It has VALUE=DATE.
func Date(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t, format: dateLayout}.With(parameter.Type(parameter.DATE_TYPE))
}

// TStamp constructs a date-time value using UTC. It has no VALUE parameter; the type the default
// and is obvious from the rendered value.
func TStamp(t time.Time) DateTimeValue {
	return DateTimeValue{Value: t.UTC(), format: dateTimeLayout + "Z"}
}

// AsDate converts a date-time value to a date-only value.
func (v DateTimeValue) AsDate() DateTimeValue {
	v.format = dateLayout
	v.Parameters.Remove(string("DATE-TIME"))
	return v.With(parameter.Type(parameter.DATE_TYPE))
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
	format := v.format

	// when the date-time is UTC, remove the TZID parameter and add Zulu "Z" instead
	if zone, _ := v.Value.Zone(); zone == "UTC" && format == dateTimeLayout {
		v.Parameters = v.Parameters.Remove("TZID", "DATE-TIME")
		format = dateTimeLayout + "Z"
	}

	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value.Format(format))
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
	return DurationValue{simpleValue{Value: d}}.With(parameter.Type(parameter.DURATION_TYPE))
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
