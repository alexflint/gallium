package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo LDFLAGS: -F${SRCDIR}/dist
#cgo LDFLAGS: -framework Gallium
#cgo LDFLAGS: -Wl,-rpath,@executable_path/../Frameworks
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/dist
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "lib/api/gallium.h"
#include "lib/api/menu.h"
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
	st *C.struct_gallium_error
}

// newCerr allocates a new error struct. It must be freed explicitly.
func newCerr() cerr {
	return cerr{
		st: (*C.struct_gallium_error)(C.malloc(C.sizeof_struct_gallium_error)),
	}
}

func (e cerr) free() {
	C.free(unsafe.Pointer(e.st))
}

func (e *cerr) err() error {
	// TODO
	return fmt.Errorf("C error")
}

// Loop starts the browser loop and does not return unless there is an initialization error
func Loop(args []string, onready func(*App)) error {
	log.Println("=== gallium.Loop ===")
	cerr := newCerr()
	defer cerr.free()
	go func() {
		time.Sleep(time.Second) // TODO: find out when the browser is actually ready
		onready(&App{})
	}()
	//C.SetUIApplication()
	C.GalliumLoop(C.CString(args[0]), &cerr.st)
	return cerr.err()
}

// App is the handle that allows you to create windows and menus
type App struct{}

// NewWindow creates a window that will oad the given URL and will display
// the given title
func (b *App) NewWindow(url, title string) error {
	log.Println("=== gallium.NewWindow ===")
	cerr := newCerr()
	defer cerr.free()
	C.GalliumCreateWindow(C.CString(url), C.CString(title), &cerr.st)
	return nil
}
