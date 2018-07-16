package display

import "github.com/rickb777/ical2/parameter"

// DISPLAY is the key for a display parameter.
const DISPLAY = "DISPLAY"

const (
	// BADGE is image inline with the title of the event.
	BADGE = "BADGE"

	// DisplayGraphic is a full image replacement for the event itself.
	GRAPHIC = "GRAPHIC"

	// DisplayFullSize is an image that is used to enhance the event.
	FULLSIZE = "FULLSIZE"

	// DisplayThumbnail is a smaller variant of "FULLSIZE" to be used when
	// space for the image is constrained.
	THUMBNAIL = "THUMBNAIL"
)

// Display identifies how the property should be displayed.
func Display(v string) parameter.Parameter {
	return parameter.Single(DISPLAY, v)
}
