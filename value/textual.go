package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"strings"
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

// URIValue holds a URI.
type URIValue struct {
	simpleValue
}

// URI returns a new URIValue.
func URI(uri string) URIValue {
	return URIValue{simpleValue{Value: uri}}
}

// CalAddress returns a new CalAddressValue.
func CalAddress(mailto string) URIValue {
	if !strings.HasPrefix(mailto, "mailto:") {
		mailto = "mailto:" + mailto
	}
	return URIValue{simpleValue{Value: mailto}}
}

// With appends parameters to the value.
func (v URIValue) With(params ...parameter.Parameter) URIValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v URIValue) WriteTo(w ics.StringWriter) error {
	v.Parameters.WriteTo(w)
	w.WriteString(":")
	_, e := w.WriteString(v.Value)
	return e
}

//-------------------------------------------------------------------------------------------------

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

// ClassValue holds a classification value.
type ClassValue struct {
	TextValue
}

// Class constructs a new classification value.
func Class(v string) ClassValue {
	return ClassValue{Text(v)}
}

// Public is an event visible publicly.
func Public() ClassValue {
	return Class("PUBLIC")
}

// Private is a private event.
func Private() ClassValue {
	return Class("PRIVATE")
}

// Confidential is a confidential event.
func Confidential() ClassValue {
	return Class("CONFIDENTIAL")
}

//-------------------------------------------------------------------------------------------------

// MethodValue holds a transparency value.
type MethodValue struct {
	TextValue
}

// Method constructs a new transparency value.
func Method(v string) MethodValue {
	return MethodValue{Text(v)}
}

// Publish specifies a publish event.
func Publish() MethodValue {
	return Method("PUBLISH")
}

// Request specifies an event that is a request to attend.
func Request() MethodValue {
	return Method("REQUEST")
}

//-------------------------------------------------------------------------------------------------

// TransparencyValue holds a transparency value.
type TransparencyValue struct {
	TextValue
}

// Transparency constructs a new transparency value.
func Transparency(v string) TransparencyValue {
	return TransparencyValue{Text(v)}
}

// Transparent specifies event transparency when the event does not block other events.
func Transparent() TransparencyValue {
	return Transparency("TRANSPARENT")
}

// Opaque specifies event transparency when the event blocks other events.
func Opaque() TransparencyValue {
	return Transparency("OPAQUE")
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
