package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo CFLAGS: -Idist/include
#cgo LDFLAGS: -F${SRCDIR}/dist
#cgo LDFLAGS: -framework Gallium
#cgo LDFLAGS: -Wl,-rpath,@executable_path/../Frameworks
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/dist
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "gallium/browser.h"
#include "gallium/cocoa.h"

// It does not seem that we can import "_cgo_export.h" from here
extern void cgo_onReady(int);

// This is a wrapper around GalliumLoop that adds the function pointer
// argument, since this does not seem to be possible from Go directly.
static inline void helper_GalliumLoop(int app_id, const char* arg0, struct gallium_error** err) {
	GalliumLoop(app_id, arg0, &cgo_onReady, err);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"log"
	"time"
	"unsafe"
)

// cerr holds a C-allocated error, which must be freed explicitly.
type cerr struct {
	c *C.struct_gallium_error
}

// newCerr allocates a new error struct. It must be freed explicitly.
func newCerr() cerr {
	return cerr{
		c: (*C.struct_gallium_error)(C.malloc(C.sizeof_struct_gallium_error)),
	}
}

func (e cerr) free() {
	C.free(unsafe.Pointer(e.c))
}

func (e *cerr) err() error {
	// TODO
	return fmt.Errorf("C error")
}

// Loop starts the browser loop and does not return unless there is an initialization error
func Loop(args []string, onReady func(*App)) error {
	log.Println("\n\n=== gallium.Loop ===")
	cerr := newCerr()
	defer cerr.free()

	app := App{
		ready: make(chan struct{}),
	}

	go func() {
		select {
		case <-app.ready:
			onReady(&app)
		case <-time.After(3 * time.Second):
			log.Fatal("Waited for 3 seconds without ready signal")
		}
	}()

	appId := apps.add(&app)
	C.helper_GalliumLoop(C.int(appId), C.CString(args[0]), &cerr.c)
	return cerr.err()
}

// appManager is the singleton for managing app instances
type appManager []*App

func (m *appManager) add(app *App) int {
	id := len(*m)
	*m = append(*m, app)
	return id
}

func (m *appManager) get(id int) *App {
	return (*m)[id]
}

var apps appManager

// App is the handle that allows you to create windows and menus
type App struct {
	// ready is how the cgo onready callback indicates to the Loop goroutine that
	// chromium is initialized
	ready chan struct{}
}

// Rect represents a rectangular region on the screen
type Rect struct {
	Width  int // Width in pixels
	Height int // Height in pixels
	Left   int // Left is offset from left in pixel
	Top    int // Left is offset from top in pixels
}

// WindowOptions contains options for creating windows
type WindowOptions struct {
	Title            string // String to display in title bar
	Shape            Rect   // Initial size and position of window
	TitleBar         bool   // Whether the window title bar
	Frame            bool   // Whether the window has a frame
	Resizable        bool   // Whether the window border can be dragged to change its shape
	CloseButton      bool   // Whether the window has a close button
	MinButton        bool   // Whether the window has a miniaturize button
	FullScreenButton bool   // Whether the window has a full screen button
	Menu             []MenuEntry
}

// FramedWindow contains options for an "ordinary" window with title bar,
// frame, and min/max/close buttons.
var FramedWindow = WindowOptions{
	Shape: Rect{
		Width:  800,
		Height: 600,
		Left:   100,
		Top:    100,
	},
	TitleBar:         true,
	Frame:            true,
	Resizable:        true,
	CloseButton:      true,
	MinButton:        true,
	FullScreenButton: true,
	Title:            "Gallium",
}

// FramelessWindow contains options for a window with no frame or border, but that
// is still resizable.
var FramelessWindow = WindowOptions{
	Shape: Rect{
		Width:  800,
		Height: 600,
		Left:   100,
		Top:    100,
	},
	Resizable: true,
}

// Window represents a window registered with the native UI toolkit (e.g. NSWindow on macOS)
type Window struct {
	c *C.gallium_window_t
}

var (
	errZeroWidth  = errors.New("window width was zero")
	errZeroHeight = errors.New("window height was zero")
)

// OpenWindow creates a window that will load the given URL.
func (app *App) OpenWindow(url string, opt WindowOptions) (*Window, error) {
	if opt.Shape.Width == 0 {
		return nil, errZeroWidth
	}
	if opt.Shape.Height == 0 {
		return nil, errZeroHeight
	}
	// Create the Cocoa window
	cwin := C.GalliumOpenWindow(
		C.CString(url),
		C.CString(opt.Title),
		C.int(opt.Shape.Width),
		C.int(opt.Shape.Height),
		C.int(opt.Shape.Left),
		C.int(opt.Shape.Top),
		C.bool(opt.TitleBar),
		C.bool(opt.Frame),
		C.bool(opt.Resizable),
		C.bool(opt.CloseButton),
		C.bool(opt.MinButton),
		C.bool(opt.FullScreenButton))

	// TODO: associate menu
	return &Window{
		c: cwin,
	}, nil
}

// Shape gets the current shape of the window.
func (w *Window) Shape() Rect {
	return Rect{
		Width:  int(C.GalliumWindowGetWidth(w.c)),
		Height: int(C.GalliumWindowGetHeight(w.c)),
		Left:   int(C.GalliumWindowGetLeft(w.c)),
		Top:    int(C.GalliumWindowGetTop(w.c)),
	}
}

// Shape gets the current shape of the window.
func (w *Window) SetShape(r Rect) {
	C.GalliumWindowSetShape(w.c, C.int(r.Width), C.int(r.Height), C.int(r.Left), C.int(r.Top))
}

// URL gets the URL that the window is currently at.
func (w *Window) URL() string {
	return C.GoString(C.GalliumWindowGetURL(w.c))
}

// LoadURL causes the window to load the given URL
func (w *Window) LoadURL(url string) {
	C.GalliumWindowLoadURL(w.c, C.CString(url))
}

// Reload reloads the current URL
func (w *Window) Reload() {
	C.GalliumWindowReload(w.c)
}

// Reload reloads the current URL, ignoring cached versions of resources.
func (w *Window) ReloadNoCache() {
	C.GalliumWindowReloadNoCache(w.c)
}

// Open opens the window. This is the default state for a window created
// via OpenWindow, so you only need to call this if you manually closed
// the window.
func (w *Window) Open() {
	C.GalliumWindowOpen(w.c)
}

// Close closes the window, as if the close button had been clicked.
func (w *Window) Close() {
	C.GalliumWindowClose(w.c)
}

// Miniaturize miniaturizes the window, as if the min button had been clicked.
func (w *Window) Miniaturize() {
	C.GalliumWindowMiniaturize(w.c)
}

// OpenDevTools opens the developer tools for this window.
func (w *Window) OpenDevTools() {
	C.GalliumWindowOpenDevTools(w.c)
}

// CloseDevTools closes the developer tools.
func (w *Window) CloseDevTools() {
	C.GalliumWindowCloseDevTools(w.c)
}

// DevToolsVisible returns whether the developer tools are showing
func (w *Window) DevToolsAreOpen() bool {
	return bool(C.GalliumWindowDevToolsAreOpen(w.c))
}
