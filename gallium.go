package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo LDFLAGS: -Fdist
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
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var logger *log.Logger

func init() {
	f, err := os.OpenFile("/Users/alex/Library/Logs/Gallium.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		ioutil.WriteFile("/Users/alex/gallium.log", []byte(err.Error()), 0777)
		return
	}
	fmt.Fprint(f, "\n=== GALLIUM INITIALIZED ===\n")

	// Create the log and write header
	logger = log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)

	// redirect stdout to the log file
	err = syscall.Dup2(int(f.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		logger.Println(err)
	}

	// redirect stderr to the log file
	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		logger.Println(err)
	}
}

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
	logger.Println("=== gallium.Run ===")
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
	logger.Println("=== gallium:CreateWindow ===")
	cerr := newCerr()
	defer cerr.free()
	C.GalliumCreateWindow(C.CString(url), C.CString(title), &cerr.st)
	return nil
}
