// Package value enumerates values of the type options for parameters (the VALUE parameter).
package value

import "github.com/rickb777/ical2/parameter"

// VALUE is the key for a type value parameter.
const VALUE = "VALUE"

const (
	BINARY      = "BINARY"
	BOOLEAN     = "BOOLEAN"
	CAL_ADDRESS = "CAL-ADDRESS"
	DATE        = "DATE"
	DATE_TIME   = "DATE-TIME"
	DURATION    = "DURATION"
	FLOAT       = "FLOAT"
	INTEGER     = "INTEGER"
	PERIOD      = "PERIOD"
	RECUR       = "RECUR"
	TEXT        = "TEXT"
	TIME        = "TIME"
	URI         = "URI"
	UTC_OFFSET  = "UTC-OFFSET"
)

// Type explicitly specifies the value type format for a property value.
func Type(v string) parameter.Parameter {
	return parameter.Single(VALUE, v)
}
