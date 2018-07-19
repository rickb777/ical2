// Package cuvalue enumerates values for calendar user types.
package cuvalue

import "github.com/rickb777/ical2/parameter"

// CUTYPE is the key for a calendar user type parameter.
const CUTYPE = "CUTYPE"

const (
	// An individual
	INDIVIDUAL = "INDIVIDUAL"
	// A group of individuals
	GROUP = "GROUP"
	// A physical resource
	RESOURCE = "RESOURCE"
	// A room resource
	ROOM = "ROOM"
	// Otherwise not known
	UNKNOWN = "UNKNOWN"
)

// CUType identifies the type of calendar user specified by the property.
func CUType(v string) parameter.Parameter {
	return parameter.Single(CUTYPE, v)
}
