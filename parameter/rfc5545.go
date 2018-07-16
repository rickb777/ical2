package parameter

import (
	"fmt"
)

// RFC-5545 parameters

const ALTREP = "ALTREP"

// AltRep specifies an alternate text representation for the property value.
func AltRep(v string) Parameter {
	return Single(ALTREP, v)
}

const CN = "CN"

// CommonName specifies the common name to be associated with the
// calendar user specified by the property.
func CommonName(v string) Parameter {
	return Single(CN, v)
}

const DELEGATED_FROM = "DELEGATED-FROM"

// Delegator names the calendar users that have delegated their
// participation to the calendar user(s) specified by the property.
func Delegator(v ...string) Parameter {
	return Parameter{DELEGATED_FROM, v}
}

const DELEGATED_TO = "DELEGATED-TO"

// Delegatee names the calendar users to whom the calendar user
// specified by the property has delegated participation.
func Delegatee(v ...string) Parameter {
	return Parameter{DELEGATED_TO, v}
}

const DIR = "DIR"

// Dir specifies reference to a directory entry associated with
// the calendar user specified by the property.
func Dir(v string) Parameter {
	return Single(DIR, v)
}

const ENCODING = "ENCODING"

// Encoding specifies an alternate inline encoding for the property value.
func Encoding(base64 bool) Parameter {
	return either(ENCODING, base64, "BASE64", "8BIT")
}

const FMTTYPE = "FMTTYPE"

// FmtType specifies the content type of a referenced object, e.g. "image/png".
func FmtType(typeName, subTypeName string) Parameter {
	return Single(FMTTYPE, fmt.Sprintf("%s/%s", typeName, subTypeName))
}

const LANGUAGE = "LANGUAGE"

// Language specifies the language for text values in a property or
// property parameter. See https://tools.ietf.org/html/rfc5646
func Language(v string) Parameter {
	return Single(LANGUAGE, v)
}

const MEMBER = "MEMBER"

// Member specifies the group or list membership of the calendar
// user specified by the property.
func Member(v ...string) Parameter {
	return Parameter{MEMBER, v}
}

/// TODO RANGE ; Recurrence identifier range
/// TODO RELATED ; Alarm trigger relationship
/// TODO RELTYPE ; Relationship type

const RSVP = "RSVP"

// Rsvp specifies whether there is an expectation of a reply from
// the calendar user specified by the property value.
func Rsvp(yes bool) Parameter {
	return either(RSVP, yes, "TRUE", "FALSE")
}

const SENT_BY = "SENT-BY"

// SentBy specifies the calendar user that is acting on behalf of
// the calendar user specified by the property.
func SentBy(v string) Parameter {
	return Single(SENT_BY, v)
}

const TZID = "TZID"

// TZid specifies the identifier for the time zone definition for
// a time component in the property value.
func TZid(v string) Parameter {
	return Single("TZID", v)
}
