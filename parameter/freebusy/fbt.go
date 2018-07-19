// Package freebusy enumerates values for the free-busy parameter.
package freebusy

import "github.com/rickb777/ical2/parameter"

// FBTYPE is the key for a free-busy type parameter.
const FBTYPE = "FBTYPE"

const (
	FREE             = "FREE"
	BUSY             = "BUSY"
	BUSY_UNAVAILABLE = "BUSY-UNAVAILABLE"
	BUSY_TENTATIVE   = "BUSY-TENTATIVE"
)

// FbType specifies the free or busy time type.
func FbType(v string) parameter.Parameter {
	return parameter.Single(FBTYPE, v)
}
