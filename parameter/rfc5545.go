package parameter

// RFC-5545 parameters

// ALTREP is the key for an alternate text representation parameter.
const ALTREP = "ALTREP"

// AltRep specifies an alternate text representation for the property value.
// https://tools.ietf.org/html/rfc5545#section-3.2.1
func AltRep(v string) Parameter {
	return Single(ALTREP, v)
}

// CN is the key for a common name parameter.
const CN = "CN"

// CommonName specifies the common name to be associated with the
// calendar user specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.2
func CommonName(v string) Parameter {
	return Single(CN, v)
}

// DELEGATED_FROM is the key for a delegator parameter.
const DELEGATED_FROM = "DELEGATED-FROM"

// Delegator names the calendar users that have delegated their
// participation to the calendar user(s) specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.4
func Delegator(v ...string) Parameter {
	return Multiple(DELEGATED_FROM, v...)
}

// DELEGATED_TO is the key for a delegatee parameter.
const DELEGATED_TO = "DELEGATED-TO"

// Delegatee names the calendar users to whom the calendar user
// specified by the property has delegated participation.
// https://tools.ietf.org/html/rfc5545#section-3.2.5
func Delegatee(v ...string) Parameter {
	return Multiple(DELEGATED_TO, v...)
}

// DIR is the key for a directory parameter.
const DIR = "DIR"

// Dir specifies reference to a directory entry associated with
// the calendar user specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.6
func Dir(v string) Parameter {
	return Single(DIR, v)
}

// ENCODING is the key for an encoding parameter.
const ENCODING = "ENCODING"

// Encoding specifies an alternate inline encoding for the property value.
// https://tools.ietf.org/html/rfc5545#section-3.2.7
func Encoding(base64 bool) Parameter {
	return either(ENCODING, base64, "BASE64", "8BIT")
}

// FMTTYPE is the key for a media type parameter.
const FMTTYPE = "FMTTYPE"

// FmtType specifies the format type (a.k.a content type, media type) of a
// referenced object, e.g. "image/png".
// https://tools.ietf.org/html/rfc5545#section-3.2.8
func FmtType(mediaType string) Parameter {
	return Single(FMTTYPE, mediaType)
}

// LANGUAGE is the key for a language parameter.
const LANGUAGE = "LANGUAGE"

// Language specifies the language for text values in a property or
// property parameter.
// https://tools.ietf.org/html/rfc5545#section-3.2.10
func Language(v string) Parameter {
	return Single(LANGUAGE, v)
}

// MEMBER is the key for a member parameter.
const MEMBER = "MEMBER"

// Member specifies the group or list membership of the calendar
// user specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.11
func Member(v ...string) Parameter {
	return Multiple(MEMBER, v...)
}

const RANGE = "RANGE"

// RangeThisandfuture specifies the effective range of recurrence instances from
// the instance specified by the recurrence identifier specified by
// the property. The only allowed value is "THISANDFUTURE".
// https://tools.ietf.org/html/rfc5545#section-3.2.13
func RangeThisandfuture() Parameter {
	return Single(RANGE, "THISANDFUTURE")
}

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
