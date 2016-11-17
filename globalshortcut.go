package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo CFLAGS: -DGALLIUM_DIR=${SRCDIR}
#cgo CFLAGS: -Idist/include

#include <stdlib.h>
#include "gallium/globalshortcut.h"

extern void cgo_onGlobalShortcut(int64_t);

// This is a wrapper around NSMenu_AddMenuItem that adds the function pointer
// argument, since this does not seem to be possible from Go directly.
static inline void helper_AddGlobalShortcut(
	int ID,
	const char* shortcutKey,
	gallium_modifier_t shortcutModifier) {

	GalliumAddGlobalShortcut(
		ID,
		shortcutKey,
		shortcutModifier,
		&cgo_onGlobalShortcut);
}
*/
import "C"
import (
	"fmt"
	"log"
)

var (
	nextID   int
	handlers = make(map[int]func())
)

//export cgo_onGlobalShortcut
func cgo_onGlobalShortcut(id int) {
	fmt.Println("cgo_onGlobalShortcut received ID", id)
	if handler, found := handlers[id]; found {
		handler()
	} else {
		log.Println("no handler for global shortcut ID", id)
	}
}

func AddGlobalShortcut(keys string, handler func()) {
	id := nextID
	nextID++
	handlers[id] = handler

	key, modifier, _ := parseShortcut(keys)
	C.helper_AddGlobalShortcut(
		C.int(id),
		C.CString(key),
		C.gallium_modifier_t(modifier))
}
