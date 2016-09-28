#ifndef GALLIUM_API_GALLIUM_H_
#define GALLIUM_API_GALLIUM_H_

#ifdef __cplusplus
extern "C" {
#endif

#define GALLIUM_EXPORT __attribute__ ((visibility ("default")))

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_error {
	const char* msg;
} gallium_error_t;

// gallium_view represents the contents of a window. It contains the top-level
// chromium objects corresponding to a window, but not the Cocoa objects for
// that window
typedef struct GALLIUM_EXPORT gallium_view gallium_view_t;

// GalliumLoop runs the chromium browser loop
GALLIUM_EXPORT int GalliumLoop(
	int app_id,
	const char* argv0,
	void(*on_ready)(int),
	struct gallium_error** err);

// GalliumCreateWindow creates a window pointed at the given url
GALLIUM_EXPORT void GalliumCreateWindow(
	const char* url);

// GalliumView_New creates a new chromium view
GALLIUM_EXPORT gallium_view_t* GalliumView_New();

// GalliumView_LoadURL loads a URL in the given view
GALLIUM_EXPORT void GalliumView_LoadURL(
	gallium_view_t* view,
	const char* url);

#ifdef __cplusplus
}
#endif

#endif
