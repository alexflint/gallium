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

// gallium_window represents a window
struct GALLIUM_EXPORT gallium_window {
	int index;
};

// GalliumLoop runs the chromium browser loop
GALLIUM_EXPORT int GalliumLoop(
	int app_id,
	const char* argv0,
	void(*on_ready)(int),
	struct gallium_error** err);

// GalliumCreateWindow creates a window pointed at the given url
GALLIUM_EXPORT struct gallium_window* GalliumCreateWindow(
	const char* url,
	const char* title,
	struct gallium_error** err);

#ifdef __cplusplus
}
#endif

#endif
