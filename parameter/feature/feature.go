// Package feature enumerates values for the feature parameter.
package feature

import "github.com/rickb777/ical2/parameter"

// FEATURE is the key for a feature parameter.
// https://tools.ietf.org/html/rfc7986#section-6.3
const FEATURE = "FEATURE"

const (
	// AUDIO - audio capability.
	AUDIO = "AUDIO"
	// CHAT - chat or instant messaging.
	CHAT = "CHAT"
	// FEED - blog or Atom feed.
	FEED = "FEED"
	// MODERATOR - moderator dial-in code.
	MODERATOR = "MODERATOR"
	// PHONE - phone conference.
	PHONE = "PHONE"
	// SCREEN - screen sharing.
	SCREEN = "SCREEN"
	// VIDEO - video capability.
	VIDEO = "VIDEO"
)

// Feature specifies one or more features available on a conference.
func Feature(vv ...string) parameter.Parameter {
	return parameter.Multiple(FEATURE, vv)
}
