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

func Loop(args []string) {
	log.Println("=== gallium.Run ===")
	cerr := (*C.struct_gallium_error)(C.malloc(C.sizeof_struct_gallium_error))
	defer C.free(unsafe.Pointer(cerr))
	C.GalliumLoop(C.CString(args[0]), &cerr)
}

func CreateWindow(url, title string) error {
	log.Println("=== gallium.CreateWindow ===")
	cerr := (*C.struct_gallium_error)(C.malloc(C.sizeof_struct_gallium_error))
	defer C.free(unsafe.Pointer(cerr))
	C.GalliumCreateWindow(C.CString(url), C.CString(title), &cerr)
	return nil
}
