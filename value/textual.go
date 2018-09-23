package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter"
	"github.com/rickb777/ical2/parameter/value"
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
		Parameters: parameter.Parameters{value.URI()},
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
func (v URIValue) IsAttachable() {}

// URIs constructs a new []URIValue slice. These will be rendered on
// separate lines.
func URIs(vv ...string) []URIValue {
	s := make([]URIValue, len(vv))
	for i, v := range vv {
		s[i] = URI(v)
	}
	return s
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

// Texts constructs a new []TextValue slice. These will be rendered on
// separate lines.
func Texts(vv ...string) []TextValue {
	s := make([]TextValue, len(vv))
	for i, v := range vv {
		s[i] = Text(v)
	}
	return s
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

// Lists constructs one or more list values, grouping the strings provided so that they
// span across multiple lines rather than using the line folding algorithm.
func Lists(v ...string) []ListValue {
	s := make([]ListValue, 0)
	n := 0
	i := 0
	for i = 0; i < len(v); i++ {
		lv := len(v[i])
		n += lv + 1
		if n > 65 {
			s = append(s, List(v[:i]...))
			v = v[i:]
			i = 0
			n = 0
		}
	}
	if n > 0 {
		s = append(s, List(v...))
	}
	return s
}

//-------------------------------------------------------------------------------------------------

// Public is a PUBLIC event. Use this for the Class.
func Public() TextValue {
	return Text("PUBLIC")
}

// Private is a PRIVATE event. Use this for the Class.
func Private() TextValue {
	return Text("PRIVATE")
}

// Confidential is a CONFIDENTIAL event. Use this for the Class.
func Confidential() TextValue {
	return Text("CONFIDENTIAL")
}

//-------------------------------------------------------------------------------------------------

// Publish specifies a PUBLISH event. Use this for the Method.
func Publish() TextValue {
	return Text("PUBLISH")
}

// Request specifies an event that is a REQUEST to attend. Use this for the Method.
func Request() TextValue {
	return Text("REQUEST")
}

//-------------------------------------------------------------------------------------------------

// Tentative specifies an event with TENTATIVE status. Use this for the Status.
func Tentative() TextValue {
	return Text("TENTATIVE")
}

// Confirmed specifies an event with confirmed status. Use this for the Status.
func Confirmed() TextValue {
	return Text("CONFIRMED")
}

// Cancelled specifies an event, a to-do or a journal with cancelled status. Use this for the Status.
func Cancelled() TextValue {
	return Text("CANCELLED")
}

// Completed specifies a to-do with COMPLETED status. Use this for the Status.
func Completed() TextValue {
	return Text("COMPLETED")
}

// NeedsAction specifies a to-do with NEEDS-ACTION status. Use this for the Status.
func NeedsAction() TextValue {
	return Text("NEEDS-ACTION")
}

// InProcess specifies a to-do with IN-PROCESS status. Use this for the Status.
func InProcess() TextValue {
	return Text("IN-PROCESS")
}

// Draft specifies a journal with DRAFT status. Use this for the Status.
func Draft() TextValue {
	return Text("DRAFT")
}

// Final specifies a journal with FINAL status. Use this for the Status.
func Final() TextValue {
	return Text("FINAL")
}

//-------------------------------------------------------------------------------------------------

// Transparent specifies event is TRANSPARENT, i.e. when the event does not block other events.
func Transparent() TextValue {
	return Text("TRANSPARENT")
}

// Opaque specifies event is OPAQUE, i.e. when the event blocks other events.
func Opaque() TextValue {
	return Text("OPAQUE")
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
