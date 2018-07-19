// Package role enumerates values for the participation role parameter.
package role

import "github.com/rickb777/ical2/parameter"

// ROLE is the key for a participation role parameter.
const ROLE = "ROLE"

const (
	// The chair
	CHAIR           = "CHAIR"
	// A required participant
	REQ_PARTICIPANT = "REQ-PARTICIPANT"
	// An optional participant
	OPT_PARTICIPANT = "OPT-PARTICIPANT"
	// A non-participant who is copied-in for information only.
	NON_PARTICIPANT = "NON-PARTICIPANT"
)

// Role specifies the participation role for the calendar user
// specified by the property.
func Role(v string) parameter.Parameter {
	return parameter.Single(ROLE, v)
}
