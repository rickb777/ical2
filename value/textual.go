package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/parameter/valuetype"
	"strings"
)

type baseValue struct {
	Parameters parameter.Parameters
	Value      string
	Others     []string
	escape     func(string) string
}

// IsDefined tests whether the value has been explicitly defined or is default.
func (v baseValue) IsDefined() bool {
	return v.Value != ""
}

// WriteTo writes the value to the writer.
// This is part of the Valuer interface.
func (v baseValue) WriteTo(w ics.StringWriter) (err error) {
	v.Parameters.WriteTo(w)
	w.WriteString(":")
	_, err = w.WriteString(v.escape(v.Value))
	for _, o := range v.Others {
		w.WriteByte(',')
		_, err = w.WriteString(v.escape(o))
	}
	return err
}

//-------------------------------------------------------------------------------------------------

// URIValue holds a URI.
type URIValue struct {
	baseValue
}

var _ Attachable = URIValue{}

// URI returns a new URIValue.
func URI(uri string) URIValue {
	return URIValue{baseValue{
		Parameters: parameter.Parameters{valuetype.Type(valuetype.URI)},
		Value:      uri,
		escape:     noOp,
	}}
}

// CalAddress returns a new CalAddressValue.
func CalAddress(mailto string) URIValue {
	if !strings.HasPrefix(mailto, "mailto:") {
		mailto = "mailto:" + mailto
	}
	return URIValue{baseValue{Value: mailto, escape: noOp}}
}

// With appends parameters to the value.
func (v URIValue) With(params ...parameter.Parameter) URIValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

// IsAttachable indicates that URI values can be used as images or attachments.
func (v URIValue) IsAttachable() {
}

//-------------------------------------------------------------------------------------------------

// TextValue holds a value that is a string and which will be escaped
// when written out.
type TextValue struct {
	baseValue
}

// Text constructs a new text value.
func Text(v string) TextValue {
	return TextValue{baseValue{Value: v, escape: escapeText}}
}

// With appends parameters to the value.
func (v TextValue) With(params ...parameter.Parameter) TextValue {
	v.Parameters = v.Parameters.Append(params...)
	return v
}

//-------------------------------------------------------------------------------------------------

// ListValue holds a list of one or more text values.
type ListValue struct {
	baseValue
}

// List constructs a new list value.
//
// Recommended categories values include
// "ANNIVERSARY", "APPOINTMENT", "BUSINESS", "EDUCATION", "HOLIDAY",
// "MEETING", "MISCELLANEOUS", "NON-WORKING HOURS", "NOT IN OFFICE",
// "PERSONAL", "PHONE CALL", "SICK DAY", "SPECIAL OCCASION",
// "TRAVEL", "VACATION".
//
// Recommended resources values include
// "CATERING", "CHAIRS", "COMPUTER PROJECTOR", "EASEL",
// "OVERHEAD PROJECTOR", "SPEAKER PHONE", "TABLE", "TV", "VCR",
// "VIDEO PHONE", "VEHICLE".
func List(v ...string) ListValue {
	return ListValue{baseValue{Value: v[0], Others: v[1:], escape: escapeText}}
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

// StatusValue holds a status value.
type StatusValue struct {
	TextValue
}

// Status constructs a new status value.
func Status(v string) StatusValue {
	return StatusValue{Text(v)}
}

// TentativeStatus specifies an event with tentative status.
func TentativeStatus() StatusValue {
	return Status("TENTATIVE")
}

// ConfirmedStatus specifies an event with confirmed status.
func ConfirmedStatus() StatusValue {
	return Status("CONFIRMED")
}

// CancelledStatus specifies an event, a to-do or a journal with cancelled status.
func CancelledStatus() StatusValue {
	return Status("CANCELLED")
}

// CompletedStatus specifies a to-do with completed status.
func CompletedStatus() StatusValue {
	return Status("COMPLETED")
}

// NeedsActionStatus specifies a to-do with needs-action status.
func NeedsActionStatus() StatusValue {
	return Status("NEEDS-ACTION")
}

// InProcessStatus specifies a to-do with in-process status.
func InProcessStatus() StatusValue {
	return Status("IN-PROCESS")
}

// DraftStatus specifies a journal with draft status.
func DraftStatus() StatusValue {
	return Status("DRAFT")
}

// FinalStatus specifies a journal with final status.
func FinalStatus() StatusValue {
	return Status("FINAL")
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

func noOp(s string) string {
	return s
}

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
