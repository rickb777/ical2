package freebusy

import "github.com/rickb777/ical2/parameter"

// FBTTYPE is the key for a free-busy type parameter.
const FBTTYPE = "FBTTYPE"

const (
	FREE              = "FREE"
	BUSY              = "BUSY"
	BUSY_UNAVAILABLE  = "BUSY-UNAVAILABLE"
	BUSY_TENTATIVE    = "BUSY-TENTATIVE"
)

// FbtType specifies the free or busy time type.
func FbtType(v string) parameter.Parameter {
	return parameter.Single(FBTTYPE, v)
}
