// Package partstat enumerates values for the participation status parameter.
// https://tools.ietf.org/html/rfc5545#section-3.2.12
package partstat

import "github.com/rickb777/ical2/parameter"

// PARTSTAT is the key for a participation status parameter.
const PARTSTAT = "PARTSTAT"

// NeedsAction specifies the participation status for the calendar user
// specified by the property is NEEDS-ACTION.
func NeedsAction() parameter.Parameter {
	return parameter.Single(PARTSTAT, "NEEDS-ACTION")
}

// Accepted specifies the participation status for the calendar user
// specified by the property is ACCEPTED.
func Accepted() parameter.Parameter {
	return parameter.Single(PARTSTAT, "ACCEPTED")
}

// Declined specifies the participation status for the calendar user
// specified by the property is DECLINED.
func Declined() parameter.Parameter {
	return parameter.Single(PARTSTAT, "DECLINED")
}

// Tentative specifies the participation status for the calendar user
// specified by the property is TENTATIVE.
func Tentative() parameter.Parameter {
	return parameter.Single(PARTSTAT, "TENTATIVE")
}

// Delegated specifies the participation status for the calendar user
// specified by the property is DELEGATED.
func Delegated() parameter.Parameter {
	return parameter.Single(PARTSTAT, "DELEGATED")
}

// Other specifies some other participation status for the calendar user
// specified by the property.
func Other(v string) parameter.Parameter {
	return parameter.Single(PARTSTAT, v)
}
