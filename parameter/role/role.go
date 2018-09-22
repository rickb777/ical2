// Package role enumerates values for the participation role parameter.
// https://tools.ietf.org/html/rfc5545#section-3.2.16
package role

import "github.com/rickb777/ical2/parameter"

// ROLE is the key for a participation role parameter.
const ROLE = "ROLE"

// Chair specifies the participation role for the calendar user
// specified by the property is CHAIR.
func Chair(v string) parameter.Parameter {
	return Other("CHAIR")
}

// ReqParticipant specifies the participation role for the calendar user
// specified by the property is a required participant, REQ-PARTICIPANT.
func ReqParticipant() parameter.Parameter {
	return Other("REQ-PARTICIPANT")
}

// OptParticipant specifies the participation role for the calendar user
// specified by the property is an optional participant, OPT-PARTICIPANT.
func OptParticipant() parameter.Parameter {
	return Other("OPT-PARTICIPANT")
}

// NonParticipant specifies the participation role for the calendar user
// specified by the property is a non-participant, NON-PARTICIPANT.
func NonParticipant() parameter.Parameter {
	return Other("NON-PARTICIPANT")
}

// Other specifies some other participation role for the calendar user
// specified by the property.
func Other(v string) parameter.Parameter {
	return parameter.Single(ROLE, v)
}
