package value

import (
	"encoding/base64"
	"github.com/rickb777/date/timespan"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/parameter/valuetype"
	"strconv"
	"time"
)

// Attachable marks values that are attachable.
type Attachable interface {
	ics.Valuer
	IsAttachable()
}

//-------------------------------------------------------------------------------------------------

const (
	dateLayout     = "20060102"
	dateTimeLayout = "20060102T150405"
	//timeLayout     = "150405"
)

// DateTimeValue holds a date/time and its formatting decision.
// See https://tools.ietf.org/html/rfc5545#section-3.3.5
type DateTimeValue struct {
	Parameters  parameter.Parameters
	Value       time.Time
	includeTime bool
	zulu        bool
}

// DateTime constructs a new date-time value. This is represented as a "floating"
// local time. It has VALUE=DATE-TIME.
func DateTime(t time.Time) DateTimeValue {
	return DateTimeValue{
		Parameters:  parameter.Parameters{valuetype.Type(valuetype.DATE_TIME)},
		Value:       t,
		includeTime: true,
	}
}

// Date constructs a new date value, i.e. without time. It has VALUE=DATE.
func Date(t time.Time) DateTimeValue {
	return DateTimeValue{
		Parameters:  parameter.Parameters{valuetype.Type(valuetype.DATE)},
		Value:       t,
		includeTime: false,
	}
}

// TStamp constructs a date-time value using UTC. It has no VALUE parameter; the type the default
// and is obvious from the rendered value.
func TStamp(t time.Time) DateTimeValue {
	return DateTimeValue{
		Value:       t.UTC(),
		includeTime: true,
		zulu:        true,
	}
}

// AsDate converts a date-time value to a date-only value.
func (v DateTimeValue) AsDate() DateTimeValue {
	v.includeTime = false
	v.Parameters = v.Parameters.RemoveByKey(valuetype.DATE_TIME).Append(valuetype.Type(valuetype.DATE))
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
	format := dateLayout
	if v.includeTime {
		format = dateTimeLayout
		if v.zulu {
			format = dateTimeLayout + "Z"
		}
	}

	// when the date-time is UTC, remove the TZID parameter and add Zulu "Z" instead
	if zone, _ := v.Value.Zone(); zone == "UTC" && v.includeTime {
		v.Parameters = v.Parameters.RemoveByKey(parameter.TZID, valuetype.DATE_TIME)
		format = dateTimeLayout + "Z"
	}

	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value.Format(format))
	return e
}

//-------------------------------------------------------------------------------------------------

// FreeBusyValue holds a date/time and its formatting decision.
// See https://tools.ietf.org/html/rfc5545#section-3.3.5
type FreeBusyValue struct {
	Parameters parameter.Parameters
	Value      timespan.TimeSpan
}

// FreeBusy constructs a new timespan value. The time should be UTC.
// It has VALUE=PERIOD.
func FreeBusy(ts timespan.TimeSpan) FreeBusyValue {
	return FreeBusyValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.PERIOD)},
		Value:      ts,
	}
}

// FreeBusyOf constructs a new timespan value. The time should be UTC.
// It has VALUE=PERIOD.
func FreeBusyOf(t time.Time, d time.Duration) FreeBusyValue {
	return FreeBusy(timespan.TimeSpanOf(t, d))
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v FreeBusyValue) IsDefined() bool {
	return !v.Value.Start().IsZero()
}

// With appends parameters to the value.
func (v FreeBusyValue) With(params ...parameter.Parameter) FreeBusyValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v FreeBusyValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	_, e := w.WriteString(v.Value.FormatRFC5545(true))
	return e
}

//-------------------------------------------------------------------------------------------------

// DurationValue holds a time duration. This should be in ISO-8601 form
// (https://en.wikipedia.org/wiki/ISO_8601#Durations);
// see github.com/rickb777/date/period for a compatible duration API.
type DurationValue struct {
	baseValue
}

// Duration returns a new DurationValue. It has VALUE=DURATION.
func Duration(d string) DurationValue {
	return DurationValue{baseValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.DURATION)},
		Value:      d,
	}}
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

// Integer returns a new IntegerValue. It has VALUE=INTEGER.
func Integer(number int) IntegerValue {
	return IntegerValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.INTEGER)},
		Value:      number,
		defined:    true,
	}
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

// GeoValue holds an integer.
type GeoValue struct {
	Parameters parameter.Parameters
	Lat, Lon   float64
	defined    bool
}

// Geo returns a new GeoValue.
// Values for latitude and longitude are expressed as decimal
// fractions of degrees.  Whole degrees of latitude are
// represented by a decimal number ranging from 0 through
// 90.  Whole degrees of longitude are represented by a decimal
// number ranging from 0 through 180. Each can be positive or negative.
//
// It has VALUE=FLOAT.
//
// See https://tools.ietf.org/html/rfc5545#section-3.8.1.6
func Geo(lat, lon float64) GeoValue {
	return GeoValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.FLOAT)},
		Lat:        lat,
		Lon:        lon,
		defined:    true,
	}
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v GeoValue) IsDefined() bool {
	return v.defined
}

// With appends parameters to the value.
func (v GeoValue) With(params ...parameter.Parameter) GeoValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v GeoValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	w.WriteString(strconv.FormatFloat(v.Lat, 'G', -1, 64))
	w.WriteByte(';')
	_, e := w.WriteString(strconv.FormatFloat(v.Lon, 'G', -1, 64))
	return e
}

//-------------------------------------------------------------------------------------------------

// BinaryValue holds binary data, such as an inline image.
type BinaryValue struct {
	Parameters parameter.Parameters
	Value      []byte
}

// Binary returns a new BinaryValue.
func Binary(data []byte) BinaryValue {
	return BinaryValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.BINARY), parameter.Encoding(true)},
		Value:      data,
	}
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v BinaryValue) IsDefined() bool {
	return len(v.Value) > 0
}

// IsAttachable indicates that binary values can be used as images or attachments.
func (v BinaryValue) IsAttachable() {
}

// With appends parameters to the value.
func (v BinaryValue) With(params ...parameter.Parameter) BinaryValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Writable interface.
func (v BinaryValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteByte(':')
	// RFC5545 requires 'standard' encoding (using alphanum, +, /) with padding.
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	encoder.Write(v.Value)
	return encoder.Close()
}
