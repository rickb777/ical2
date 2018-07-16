package parameter

// RFC7986 additions

// EMAIL is the key for an email parameter.
const EMAIL = "EMAIL"

// Email specifies an email address that is used to identify or
// contact an organizer or attendee.
func Email(v string) Parameter {
	return Single(EMAIL, v)
}

// LABEL is the key for a label parameter.
const LABEL = "LABEL"

// Label provides a human-readable label.
func Label(v string) Parameter {
	return Single(LABEL, v)
}
