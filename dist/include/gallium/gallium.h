#ifndef GALLIUM_API_GALLIUM_H_
#define GALLIUM_API_GALLIUM_H_

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

#define GALLIUM_EXPORT __attribute__ ((visibility ("default")))

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_error {
	const char* msg;
} gallium_error_t;

// gallium_window represents a window
typedef struct GALLIUM_EXPORT gallium_window gallium_window_t;

// GalliumLoop runs the chromium browser loop
GALLIUM_EXPORT int GalliumLoop(
	int app_id,
	const char* argv0,
	void(*on_ready)(int),
	struct gallium_error** err);

// GalliumCreateWindow creates a window pointed at the given url
GALLIUM_EXPORT gallium_window_t* GalliumOpenWindow(const char* url,
                                      const char* title,
                                      int width,
                                      int height,
                                      int x,
                                      int y,
                                      bool titleBar,
                                      bool frame,
                                      bool resizable,
                                      bool closeButton,
                                      bool minButton,
                                      bool fullScreenButton);

#ifdef __cplusplus
}
#endif

#endif
