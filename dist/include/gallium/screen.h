#ifndef GALLIUM_SCREEN_H_
#define GALLIUM_SCREEN_H_

#include <stdbool.h>
#include <stdint.h>

#include "gallium/export.h"
#include "gallium/rect.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct GALLIUM_EXPORT gallium_screen gallium_screen_t;

// GalliumScreenCount gets the number of screens
GALLIUM_EXPORT int GalliumScreenCount();

// GalliumScreen gets the i-th screen
GALLIUM_EXPORT gallium_screen_t* GalliumScreen(int index);

// GalliumFocusedScreen gets the screen containing the currently focussed window.
GALLIUM_EXPORT gallium_screen_t* GalliumFocusedScreen();

// GalliumScreenShape gets the shape of a screen
GALLIUM_EXPORT gallium_rect_t GalliumScreenShape(gallium_screen_t*);

// GalliumScreenShape gets the usable area of a screen (excludes dock and menubar)
GALLIUM_EXPORT gallium_rect_t GalliumScreenUsable(gallium_screen_t*);

// GalliumScreenBitsPerPixel gets the bits per pixels for a screen
GALLIUM_EXPORT int GalliumScreenBitsPerPixel(gallium_screen_t*);

// GalliumScreenID gets the unique ID for a screen
GALLIUM_EXPORT int GalliumScreenID(gallium_screen_t*);

#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_SCREEN_H_
