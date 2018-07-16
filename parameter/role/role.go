package role

import "github.com/rickb777/ical2/parameter"

// ROLE is the key for a participation role parameter.
const ROLE = "ROLE"

const (
	CHAIR            = "CHAIR"
	REQ_PARTICIPANT  = "REQ-PARTICIPANT"
	OPT_PARTICIPANT  = "OPT-PARTICIPANT"
	NON_PARTICIPANT  = "NON-PARTICIPANT"
)

// Role specifies the participation role for the calendar user
// specified by the property.
func Role(v string) parameter.Parameter {
	return parameter.Single(ROLE, v)
}
