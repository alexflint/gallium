#ifndef GALLIUM_EXPORT_H_
#define GALLIUM_EXPORT_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

// GALLIUM_EXPORT is the macro for exporting symbols
#define GALLIUM_EXPORT __attribute__ ((visibility ("default")))

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_error {
	const char* msg;
} gallium_error_t;

// gallium_modifier represents a modifier mask for a keyboard shortcut
typedef enum GALLIUM_EXPORT gallium_modifier {
	GalliumCmdModifier = 1 << 0,
	GalliumCtrlModifier = 1 << 1,
	GalliumCmdOrCtrlModifier = 1 << 2,
	GalliumAltOrOptionModifier = 1 << 3,
	GalliumFunctionModifier = 1 << 4,
	GalliumShiftModifier = 1 << 5,
} gallium_modifier_t;

// gallium_rect represents a rectangle
typedef struct GALLIUM_EXPORT {
	int width;
	int height;
	int left;
	int bottom;
} gallium_rect_t;

#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_EXPORT_H_

