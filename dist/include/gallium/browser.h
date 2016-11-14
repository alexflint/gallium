#ifndef GALLIUM_BROWSER_H_
#define GALLIUM_BROWSER_H_

#include <stdbool.h>

#include "gallium/export.h"
#include "gallium/rect.h"

#ifdef __cplusplus
extern "C" {
#endif

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

  
// GalliumWindowGetWidth gets the width of a window
GALLIUM_EXPORT gallium_rect_t GalliumWindowGetShape(gallium_window_t* window);

// GalliumWindowGetWidth gets the width of a window
GALLIUM_EXPORT void GalliumWindowSetShape(gallium_window_t* window,
                                          int width,
                                          int height,
                                          int left,
                                          int top);

// GalliumWindowGetURL gets the URL that the window is currently at
GALLIUM_EXPORT const char* GalliumWindowGetURL(gallium_window_t* window);

// GalliumWindowLoadURL causes the window to load the given URL
GALLIUM_EXPORT void GalliumWindowLoadURL(gallium_window_t* window, const char* url);

// GalliumWindowReload reloads the current page
GALLIUM_EXPORT void GalliumWindowReload(gallium_window_t* window);

// GalliumWindowReload reloads the current page, ignoring any cached resources
GALLIUM_EXPORT void GalliumWindowReloadNoCache(gallium_window_t* window);

// GalliumWindowOpen opens the window (only for use after GalliumWindowClose)
GALLIUM_EXPORT void GalliumWindowOpen(gallium_window_t* window);

// GalliumWindowClose closes the window
GALLIUM_EXPORT void GalliumWindowClose(gallium_window_t* window);

// GalliumWindowClose miniaturizes the window
GALLIUM_EXPORT void GalliumWindowMiniaturize(gallium_window_t* window);

// GalliumWindowUndo undoes the last text editing action
GALLIUM_EXPORT void GalliumWindowUndo(gallium_window_t* window);

// GalliumWindowRedo redoes the last text editing action
GALLIUM_EXPORT void GalliumWindowRedo(gallium_window_t* window);

// GalliumWindow cuts the current text selection to the pastboard
GALLIUM_EXPORT void GalliumWindowCut(gallium_window_t* window);

// GalliumWindow copies the current text selection to the pasteboard
GALLIUM_EXPORT void GalliumWindowCopy(gallium_window_t* window);

// GalliumWindow pastes from the pasteboard
GALLIUM_EXPORT void GalliumWindowPaste(gallium_window_t* window);

// GalliumWindow pastes from the pasteboard, matching style to the current element
GALLIUM_EXPORT void GalliumWindowPasteAndMatchStyle(gallium_window_t* window);

// GalliumWindow deletes the current text selection
GALLIUM_EXPORT void GalliumWindowDelete(gallium_window_t* window);

// GalliumWindow selects all text in the current element
GALLIUM_EXPORT void GalliumWindowSelectAll(gallium_window_t* window);

// GalliumWindowUnselect unselects any text selection
GALLIUM_EXPORT void GalliumWindowUnselect(gallium_window_t* window);

// GalliumWindowOpenDevTools opens the developer tools for the given window
GALLIUM_EXPORT void GalliumWindowOpenDevTools(gallium_window_t* window);

// GalliumWindowCloseDevTools closes the developer tools for the given window
GALLIUM_EXPORT void GalliumWindowCloseDevTools(gallium_window_t* window);
  
// GalliumWindowDevToolsVisible returns true if the developer tools are currently visible for the given window
GALLIUM_EXPORT bool GalliumWindowDevToolsAreOpen(gallium_window_t* window);
  
// GalliumWindowNativeWindow gets a native handle for this window (NSWindow*).
GALLIUM_EXPORT void* GalliumWindowNativeWindow(gallium_window_t* window);

// GalliumWindowNativeWindow gets a native handle for the window controller (NSWindowController*).
GALLIUM_EXPORT void* GalliumWindowNativeController(gallium_window_t* window);

// GalliumWindowNativeWindow gets a native handle for the window content (content::WebContent*).
GALLIUM_EXPORT void* GalliumWindowNativeContent(gallium_window_t* window);
  
#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_BROWSER_H_
