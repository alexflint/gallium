package gallium

/*
#cgo CFLAGS: -Flib/build/Debug
#cgo LDFLAGS: -framework Gallium
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo LDFLAGS: -Flib/build/Debug
#cgo LDFLAGS: -framework Gallium
#cgo LDFLAGS: -Wl,-rpath -Wl,@executable_path/../Frameworks
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "lib/api/gallium.h"
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	logging "log"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var log *logging.Logger

func init() {
	f, err := os.OpenFile("/Users/alex/Library/Logs/Gallium.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		ioutil.WriteFile("/Users/alex/gallium.log", []byte(err.Error()), 0777)
		return
	}
	fmt.Fprint(f, "\n=== GALLIUM INITIALIZED ===\n")

	// Create the log and write header
	log = logging.New(f, "", logging.Ldate|logging.Ltime|logging.Lshortfile)

	// redirect stdout to the log file
	err = syscall.Dup2(int(f.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		log.Println(err)
	}

	// redirect stderr to the log file
	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Println(err)
	}
}

// Browser is the handle that allows you to create windows
type Browser struct{}

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
func Loop(args []string, onready func(*Browser)) error {
	log.Println("=== gallium.Run ===")
	cerr := newCerr()
	defer cerr.free()
	go func() {
		time.Sleep(time.Second) // TODO: find out when the browser is actually ready
		onready(&Browser{})
	}()
	C.GalliumLoop(C.CString(args[0]), &cerr.st)
	return cerr.err()
}

func (b *Browser) CreateWindow(url, title string) error {
	log.Println("=== gallium.CreateWindow ===")
	cerr := newCerr()
	defer cerr.free()
	C.GalliumCreateWindow(C.CString(url), C.CString(title), &cerr.st)
	return nil
}
