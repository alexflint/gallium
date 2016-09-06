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
#include "lib/common/gallium.h"
*/
import "C"
import "os"

func Run(args []string) {
	cerr := (*C.struct_gallium_error)(C.malloc(C.sizeof_struct_gallium_error))
	defer C.free(cerr)
	os.Exit(int(C.GalliumLoop(C.CString(args[0]), &cerr)))
	// for _, arg := range os.Args {
	// 	C.AddArg(C.CString(arg))
	// }
	// os.Exit(int(C.RunGallium()))
}
