package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
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
