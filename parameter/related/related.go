package related

import "github.com/rickb777/ical2/parameter"

const (
	// The RELATED parameter
	RELATED = "RELATED"
	// The RELTYPE parameter
	RELTYPE = "RELTYPE"
)

// Start specifies the relationship of the alarm trigger with
// respect to the start of the calendar component.
// https://tools.ietf.org/html/rfc5545#section-3.2.14
func Start() parameter.Parameter {
	return parameter.Single(RELATED, "START")
}

// End specifies the relationship of the alarm trigger with
// respect to the end of the calendar component.
// https://tools.ietf.org/html/rfc5545#section-3.2.14
func End() parameter.Parameter {
	return parameter.Single(RELATED, "END")
}

// Parent specifies the type of hierarchical relationship associated
// with the calendar component specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.15
func Parent() parameter.Parameter {
	return RelType("PARENT")
}

// Child specifies the type of hierarchical relationship associated
// with the calendar component specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.15
func Child() parameter.Parameter {
	return RelType("CHILD")
}

// Sibling specifies the type of hierarchical relationship associated
// with the calendar component specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.15
func Sibling() parameter.Parameter {
	return RelType("SIBLING")
}

// RelType specifies the type of hierarchical relationship associated
// with the calendar component specified by the property.
// https://tools.ietf.org/html/rfc5545#section-3.2.15
func RelType(v string) parameter.Parameter {
	return parameter.Single(RELTYPE, v)
}
