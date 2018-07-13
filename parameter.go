package ical2

import (
	"fmt"
	"strings"
)

// Parameter holds an iCal parameter.
// See https://tools.ietf.org/html/rfc5545#section-3.2
type Parameter struct {
	Key, Value string
}

func quote(v string) string {
	return fmt.Sprintf("%q", v)
}

func quoted(k, v string) Parameter {
	return Parameter{k, quote(v)}
}

func quotedList(k string, vv []string) Parameter {
	qq := make([]string, 0, len(vv))
	for _, s := range vv {
		qq = append(qq, quote(s))
	}
	return Parameter{k, strings.Join(qq, ",")}
}

func plain(k, v string) Parameter {
	return Parameter{k, v}
}

func either(key string, predicate bool, yes, no string) Parameter {
	v := no
	if predicate {
		v = yes
	}
	return plain(key, v)
}

//-------------------------------------------------------------------------------------------------
// RFC-5545 parameters

// AltRep specifies an alternate text representation for the property value.
func AltRep(v string) Parameter {
	return quoted("ALTREP", v)
}

// CommonName specifies the common name to be associated with the
// calendar user specified by the property.
func CommonName(v string) Parameter {
	return plain("CN", v)
}

// CUValue provides values for calendar user types.
type CUValue string

const (
	// An individual
	INDIVIDUAL CUValue = "INDIVIDUAL"
	// A group of individuals
	GROUP CUValue = "GROUP"
	// A physical resource
	RESOURCE CUValue = "RESOURCE"
	// A room resource
	ROOM CUValue = "ROOM"
	// Otherwise not known
	UNKNOWN CUValue = "UNKNOWN"
)

// CUType identifies the type of calendar user specified by the property.
func CUType(v CUValue) Parameter {
	return plain("CUTYPE", string(v))
}

// Delegator names the calendar users that have delegated their
// participation to the calendar user(s) specified by the property.
func Delegator(v ...string) Parameter {
	return quotedList("DELEGATED-FROM", v)
}

// Delegatee names the calendar users to whom the calendar user
// specified by the property has delegated participation.
func Delegatee(v ...string) Parameter {
	return quotedList("DELEGATED-TO", v)
}

// Dir specifies reference to a directory entry associated with
// the calendar user specified by the property.
func Dir(v string) Parameter {
	return quoted("DIR", v)
}

// Encoding specifies an alternate inline encoding for the property value.
func Encoding(base64 bool) Parameter {
	return either("ENCODING", base64, "BASE64", "8BIT")
}

// FmtType specifies the content type of a referenced object.
func FmtType(typeName, subTypeName string) Parameter {
	return plain("FMTTYPE", fmt.Sprintf("%s/%s", typeName, subTypeName))
}

// FbtValue provides values for free or busy time.
type FbtValue string

const (
	FREE             FbtValue = "FREE"
	BUSY             FbtValue = "BUSY"
	BUSY_UNAVAILABLE FbtValue = "BUSY-UNAVAILABLE"
	BUSY_TENTATIVE   FbtValue = "BUSY-TENTATIVE"
)

// FbtType specifies the free or busy time type.
func FbtType(v FbtValue) Parameter {
	return plain("FBTTYPE", string(v))
}

// Language specifies the language for text values in a property or
// property parameter. See https://tools.ietf.org/html/rfc5646
func Language(v string) Parameter {
	return plain("LANGUAGE", v)
}

// Member specifies the group or list membership of the calendar
// user specified by the property.
func Member(v ...string) Parameter {
	return quotedList("Member", v)
}

/// TODO PARTSTAT ; Participation status
/// TODO RANGE ; Recurrence identifier range
/// TODO RELATED ; Alarm trigger relationship
/// TODO RELTYPE ; Relationship type

// RoleValue provides values for participation role.
type RoleValue string

const (
	CHAIR           RoleValue = "CHAIR"
	REQ_PARTICIPANT RoleValue = "REQ-PARTICIPANT"
	OPT_PARTICIPANT RoleValue = "OPT-PARTICIPANT"
	NON_PARTICIPANT RoleValue = "NON-PARTICIPANT"
)

// Role specifies the participation role for the calendar user
// specified by the property.
func Role(v RoleValue) Parameter {
	return plain("ROLE", string(v))
}

// RSVP specifies whether there is an expectation of a reply from
// the calendar user specified by the property value.
func RSVP(yes bool) Parameter {
	return either("RSVP", yes, "TRUE", "FALSE")
}

// SentBy specifies the calendar user that is acting on behalf of
// the calendar user specified by the property.
func SentBy(v string) Parameter {
	return quoted("SENT-BY", v)
}

// TZID specifies the identifier for the time zone definition for
// a time component in the property value.
func TZID(v string) Parameter {
	return plain("TZID", v)
}

// ValueType provides type information for iCal property values.
type ValueType string

const (
	BINARY      ValueType = "BINARY"
	BOOLEAN     ValueType = "BOOLEAN"
	CAL_ADDRESS ValueType = "CAL-ADDRESS"
	DATE        ValueType = "DATE"
	DATE_TIME   ValueType = "DATE-TIME"
	DURATION    ValueType = "DURATION"
	FLOAT       ValueType = "FLOAT"
	INTEGER     ValueType = "INTEGER"
	PERIOD      ValueType = "PERIOD"
	RECUR       ValueType = "RECUR"
	TEXT        ValueType = "TEXT"
	TIME        ValueType = "TIME"
	URI         ValueType = "URI"
	UTC_OFFSET  ValueType = "UTC-OFFSET"
)

// Type explicitly specifies the value type format for a property value.
func Type(v ValueType) Parameter {
	return plain("VALUE", string(v))
}

//-------------------------------------------------------------------------------------------------
// RFC7986 addtitions

// DisplayValue provides values for display.
// https://tools.ietf.org/html/rfc7986#section-6.1
type DisplayValue string

const (
	BADGE     DisplayValue = "BADGE" // the default
	GRAPHIC   DisplayValue = "GRAPHIC"
	FULLSIZE  DisplayValue = "FULLSIZE"
	THUMBNAIL DisplayValue = "THUMBNAIL"
)

// Display specifies different ways in which an image for a calendar
// or component can be displayed.
func Display(v DisplayValue) Parameter {
	return plain("DISPLAY", string(v))
}

// Email specifies an email address that is used to identify or
// contact an organizer or attendee.
func Email(v string) Parameter {
	return plain("EMAIL", v)
}

// FeatureValue provides values for display.
// https://tools.ietf.org/html/rfc7986#section-6.3
type FeatureValue string

const (
	AUDIO     FeatureValue = "AUDIO"
	CHAT      FeatureValue = "CHAT"
	FEED      FeatureValue = "FEED"
	MODERATOR FeatureValue = "MODERATOR"
	PHONE     FeatureValue = "PHONE"
	SCREEN    FeatureValue = "SCREEN"
	VIDEO     FeatureValue = "VIDEO"
)

// Feature specifies a feature or features of a conference or
// broadcast system.
func Feature(v FeatureValue) Parameter {
	return plain("FEATURE", string(v))
}

// Label provides a human-readable label.
func Label(v string) Parameter {
	return plain("LABEL", v)
}
