package partstat

import "github.com/rickb777/ical2/parameter"

// PARTSTAT is the key for a participation status parameter.
const PARTSTAT = "PARTSTAT"

const (
	NEEDS_ACTION = "NEEDS-ACTION"
	ACCEPTED     = "ACCEPTED"
	DECLINED     = "DECLINED"
	TENTATIVE    = "TENTATIVE"
	DELEGATED    = "DELEGATED"
)

// PartStat specifies the participation status for the calendar user
// specified by the property.
func PartStat(v string) parameter.Parameter {
	return parameter.Single(PARTSTAT, v)
}
