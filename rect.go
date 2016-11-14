package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo CFLAGS: -Idist/include

#include <stdlib.h>
#include "gallium/core.h"
*/
import "C"

// Rect represents a rectangular region on the screen
type Rect struct {
	Width  int // Width in pixels
	Height int // Height in pixels
	Left   int // Left is offset from left in pixel
	Bottom int // Left is offset from top in pixels
}

func rectFromC(c C.gallium_rect_t) Rect {
	return Rect{
		Width:  int(c.width),
		Height: int(c.height),
		Left:   int(c.left),
		Bottom: int(c.bottom),
	}
}
