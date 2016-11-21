package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo CFLAGS: -Idist/include

#include <stdlib.h>
#include "gallium/screen.h"
*/
import "C"
import "fmt"

// A screen represents a rectangular display, normally corresponding to a
// physical display. "Device coordinates" means a position on a screen
// measured from (0, 0) at the bottom left of the device. "Global coordinates"
// means the coordinate system in which each of the screens are positioned
// relative to each other. Global and device coordinates almost always have
// the same scale factor. It is possible for screens to overlap in global
// coordinates (such as when mirroring a display.)
type Screen struct {
	Shape        Rect // the size and position of this screen in global coords
	Usable       Rect // excludes the menubar and dock
	BitsPerPixel int  // color depth of this screen (total of all color components)
	ID           int  // unique identifier for this screen
}

func screenFromC(c *C.gallium_screen_t) Screen {
	return Screen{
		Shape:        rectFromC(C.GalliumScreenShape(c)),
		Usable:       rectFromC(C.GalliumScreenUsable(c)),
		BitsPerPixel: int(C.GalliumScreenBitsPerPixel(c)),
		ID:           int(C.GalliumScreenID(c)),
	}
}

// Screens gets a list of available screens
func Screens() []Screen {
	var screens []Screen
	n := int(C.GalliumScreenCount())
	for i := 0; i < n; i++ {
		c := C.GalliumScreen(C.int(i))
		if c == nil {
			panic(fmt.Sprintf("GalliumScreen returned nil for index %d", i))
		}
		screens = append(screens, screenFromC(c))
	}
	return screens
}

// FocusedScreen gets the screen containing the currently focused window
func FocusedScreen() Screen {
	c := C.GalliumFocusedScreen()
	if c == nil {
		panic("GalliumFocusedScreen returned nil")
	}
	return screenFromC(c)
}
