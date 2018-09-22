// Package display enumerates values for the display parameter.
package display

import "github.com/rickb777/ical2/parameter"

// DISPLAY is the key for a display parameter.
const DISPLAY = "DISPLAY"

// Badge identifies the property should be displayed as a badge.
func Badge() parameter.Parameter {
	return parameter.Single(DISPLAY, "BADGE")
}

// Graphic identifies the property should be displayed as a graphic.
func Graphic() parameter.Parameter {
	return parameter.Single(DISPLAY, "GRAPHIC")
}

// Fullsize identifies the property should be displayed fullsize.
func Fullsize() parameter.Parameter {
	return parameter.Single(DISPLAY, "FULLSIZE")
}

// Thumbnail identifies the property should be displayed as a smaller variant of fullsize
// when space for the image is constrained.
func Thumbnail() parameter.Parameter {
	return parameter.Single(DISPLAY, "THUMBNAIL")
}

// Display identifies how the property should be displayed, via the DISPLAY parameter.
func Display(v string) parameter.Parameter {
	return parameter.Single(DISPLAY, v)
}
