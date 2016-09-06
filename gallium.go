package gallium

/*
#cgo CFLAGS: -Flib/build/Debug
#cgo LDFLAGS: -framework Gallium
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo LDFLAGS: -Flib/build/Debug
#cgo LDFLAGS: -framework Gallium
#cgo LDFLAGS: -Wl,-rpath -Wl,@executable_path/../Frameworks
#cgo LDFLAGS: -mmacosx-version-min=10.8
#include "lib/common/gallium.h"
*/
import "C"
import "os"

func Run(args []string) {
	for _, arg := range os.Args {
		C.AddArg(C.CString(arg))
	}
	os.Exit(int(C.RunGallium()))
}
