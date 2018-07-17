package parameter

import (
	"fmt"
)

// RFC-5545 parameters

// ALTREP is the key for an alternate text representation parameter.
const ALTREP = "ALTREP"

// AltRep specifies an alternate text representation for the property value.
func AltRep(v string) Parameter {
	return Single(ALTREP, v)
}

// CN is the key for a common name parameter.
const CN = "CN"

// CommonName specifies the common name to be associated with the
// calendar user specified by the property.
func CommonName(v string) Parameter {
	return Single(CN, v)
}

// DELEGATED_FROM is the key for a delegator parameter.
const DELEGATED_FROM = "DELEGATED-FROM"

// Delegator names the calendar users that have delegated their
// participation to the calendar user(s) specified by the property.
func Delegator(v ...string) Parameter {
	return Multiple(DELEGATED_FROM, v)
}

// DELEGATED_TO is the key for a delegatee parameter.
const DELEGATED_TO = "DELEGATED-TO"

// Delegatee names the calendar users to whom the calendar user
// specified by the property has delegated participation.
func Delegatee(v ...string) Parameter {
	return Multiple(DELEGATED_TO, v)
}

// DIR is the key for a directory parameter.
const DIR = "DIR"

// Dir specifies reference to a directory entry associated with
// the calendar user specified by the property.
func Dir(v string) Parameter {
	return Single(DIR, v)
}

// ENCODING is the key for an encoding parameter.
const ENCODING = "ENCODING"

// Encoding specifies an alternate inline encoding for the property value.
func Encoding(base64 bool) Parameter {
	return either(ENCODING, base64, "BASE64", "8BIT")
}

// FMTYPE is the key for a media type parameter.
const FMTTYPE = "FMTTYPE"

// FmtType specifies the content type of a referenced object, e.g. "image/png".
func FmtType(typeName, subTypeName string) Parameter {
	return Single(FMTTYPE, fmt.Sprintf("%s/%s", typeName, subTypeName))
}

// LANGUAGE is the key for a language parameter.
const LANGUAGE = "LANGUAGE"

// Language specifies the language for text values in a property or
// property parameter. See https://tools.ietf.org/html/rfc5646
func Language(v string) Parameter {
	return Single(LANGUAGE, v)
}

// MEMBER is the key for a member parameter.
const MEMBER = "MEMBER"

// Member specifies the group or list membership of the calendar
// user specified by the property.
func Member(v ...string) Parameter {
	return Multiple(MEMBER, v)
}

/// TODO RANGE ; Recurrence identifier range
/// TODO RELATED ; Alarm trigger relationship
/// TODO RELTYPE ; Relationship type

// RSVP is the key for a RSVP parameter.
const RSVP = "RSVP"

// Rsvp specifies whether there is an expectation of a reply from
// the calendar user specified by the property value.
func Rsvp(yes bool) Parameter {
	return either(RSVP, yes, "TRUE", "FALSE")
}

// SENT_BY is the key for a sent-by parameter.
const SENT_BY = "SENT-BY"

// SentBy specifies the calendar user that is acting on behalf of
// the calendar user specified by the property.
func SentBy(v string) Parameter {
	return Single(SENT_BY, v)
}

// TZID is the key for a timezone ID parameter.
const TZID = "TZID"

// TZid specifies the identifier for the time zone definition for
// a time component in the property value.
func TZid(v string) Parameter {
	return Single("TZID", v)
}
