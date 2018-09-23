// Package value enumerates values of the type options for parameters (the VALUE parameter).
// https://tools.ietf.org/html/rfc5545#section-3.2.20
package value

import "github.com/rickb777/ical2/parameter"

// VALUE is the key for a type value parameter.
const VALUE = "VALUE"

const DATE_TIME = "DATE-TIME"

// Binary specifies the value type format for a property value is BINARY.
func Binary() parameter.Parameter {
	return parameter.Single(VALUE, "BINARY")
}

// Boolean specifies the value type format for a property value is BOOLEAN.
func Boolean() parameter.Parameter {
	return parameter.Single(VALUE, "BOOLEAN")
}

// CalAddress specifies the value type format for a property value is CAL-ADDRESS.
func CalAddress() parameter.Parameter {
	return parameter.Single(VALUE, "CAL-ADDRESS")
}

// Date specifies the value type format for a property value is DATE.
func Date() parameter.Parameter {
	return parameter.Single(VALUE, "DATE")
}

// DateTime specifies the value type format for a property value is DATE-TIME.
func DateTime() parameter.Parameter {
	return parameter.Single(VALUE, "DATE-TIME")
}

// Duration specifies the value type format for a property value is DURATION.
func Duration() parameter.Parameter {
	return parameter.Single(VALUE, "DURATION")
}

// Float specifies the value type format for a property value is FLOAT.
func Float() parameter.Parameter {
	return parameter.Single(VALUE, "FLOAT")
}

// Integer specifies the value type format for a property value is INTEGER.
func Integer() parameter.Parameter {
	return parameter.Single(VALUE, "INTEGER")
}

// Period specifies the value type format for a property value is PERIOD.
func Period() parameter.Parameter {
	return parameter.Single(VALUE, "PERIOD")
}

// Recur specifies the value type format for a property value is RECUR.
func Recur() parameter.Parameter {
	return parameter.Single(VALUE, "RECUR")
}

// Text specifies the value type format for a property value is TEXT.
func Text() parameter.Parameter {
	return parameter.Single(VALUE, "TEXT")
}

// Time specifies the value type format for a property value is TIME.
func Time() parameter.Parameter {
	return parameter.Single(VALUE, "TIME")
}

// URI specifies the value type format for a property value is URI.
func URI() parameter.Parameter {
	return parameter.Single(VALUE, "URI")
}

// UTCOffset specifies the value type format for a property value is UTC-OFFSET.
func UTCOffset() parameter.Parameter {
	return parameter.Single(VALUE, "UTC-OFFSET")
}

// Type explicitly specifies the value type format for a property value.
func Type(v string) parameter.Parameter {
	return parameter.Single(VALUE, v)
}
