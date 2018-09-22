// Package freebusy enumerates values for the free-busy parameter.
// https://tools.ietf.org/html/rfc5545#section-3.2.9
package freebusy

import "github.com/rickb777/ical2/parameter"

// FBTYPE is the key for a free-busy type parameter.
const FBTYPE = "FBTYPE"

// Free specifies the free/busy time type is FREE.
func Free() parameter.Parameter {
	return parameter.Single(FBTYPE, "FREE")
}

// Busy specifies the free/busy time type is FREE.
func Busy() parameter.Parameter {
	return parameter.Single(FBTYPE, "BUSY")
}

// BusyUnavailable specifies the free/busy time type is BUSY-UNAVAILABLE.
func BusyUnavailable() parameter.Parameter {
	return parameter.Single(FBTYPE, "BUSY-UNAVAILABLE")
}

// BusyTentative specifies the free/busy time type is BUSY-TENTATIVE.
func BusyTentative() parameter.Parameter {
	return parameter.Single(FBTYPE, "BUSY-TENTATIVE")
}

// Other specifies some other free or busy time type.
func Other(v string) parameter.Parameter {
	return parameter.Single(FBTYPE, v)
}
