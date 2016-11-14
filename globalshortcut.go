package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo CFLAGS: -Idist/include

#include <stdlib.h>
#include "gallium/globalshortcut.h"
*/
import "C"

func RegisterGlobalShortcut() {
	C.RegisterGlobalShortcut()
}
