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
#include "gallium/gallium.h"
#include "gallium/menu.h"

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
	log.Println("=== gallium.Loop ===")
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

// NewWindow creates a window that will load the given URL and will display
// the given title
func (b *App) NewWindow(url, title string) error {
	log.Println("=== gallium.NewWindow ===")
	cerr := newCerr()
	defer cerr.free()
	C.GalliumCreateWindow(C.CString(url), C.CString(title), &cerr.c)
	return nil
}
