// Package cutype enumerates calendar user types.
// https://tools.ietf.org/html/rfc5545#section-3.2.3
package cutype

import "github.com/rickb777/ical2/parameter"

// CUTYPE is the key for a calendar user type parameter.
const CUTYPE = "CUTYPE"

// Individual identifies the calendar user specified by the property is an INDIVIDUAL.
func Individual() parameter.Parameter {
	return parameter.Single(CUTYPE, "INDIVIDUAL")
}

// Group identifies the calendar user specified by the property is a GROUP.
func Group() parameter.Parameter {
	return parameter.Single(CUTYPE, "GROUP")
}

// Resource identifies the calendar user specified by the property is a RESOURCE.
func Resource() parameter.Parameter {
	return parameter.Single(CUTYPE, "RESOURCE")
}

// Room identifies the calendar user specified by the property is a ROOM.
func Room() parameter.Parameter {
	return parameter.Single(CUTYPE, "ROOM")
}

// Unknown identifies the calendar user specified by the property is an UNKNOWN.
func Unknown() parameter.Parameter {
	return parameter.Single(CUTYPE, "UNKNOWN")
}

// CUType identifies the calendar user specified by the CUTYPE property with some value.
func CUType(v string) parameter.Parameter {
	return parameter.Single(CUTYPE, v)
}
