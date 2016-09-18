package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
// #cgo LDFLAGS: -Flib/build/Debug
// #cgo LDFLAGS: -framework Gallium
// #cgo LDFLAGS: -Wl,-rpath,@executable_path/../Frameworks
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "lib/api/gallium.h"
#include "lib/api/menu.h"
*/
import "C"
import "log"

type Event struct{}

type MenuEntry interface {
	menu()
}

type MenuItem struct {
	Title    string
	Shortcut string
	OnClick  func(*Event)
}

func (MenuItem) menu() {}

type Menu struct {
	Title   string
	Entries []MenuEntry
}

func (Menu) menu() {}

func SetMenu(menus []Menu) {
	root := C.NSMenu_New(C.CString("<root>"))
	for _, menu := range menus {
		buildMenu(root, menu)
	}
	C.NSApplication_SetMainMenu(C.NSApplication_SharedApplication(), root)
}

func buildMenu(parent *C.gallium_nsmenu_t, menu MenuEntry) {
	switch menu := menu.(type) {
	case Menu:
		item := C.NSMenu_AddMenuItem(parent, C.CString(menu.Title), C.CString(""))
		submenu := C.NSMenu_New(C.CString(menu.Title))
		C.NSMenuItem_SetSubmenu(item, submenu)
		for _, entry := range menu.Entries {
			buildMenu(submenu, entry)
		}
	case MenuItem:
		C.NSMenu_AddMenuItem(parent, C.CString(menu.Title), C.CString(menu.Shortcut))
	default:
		log.Printf("unexpected menu entry: %T", menu)
	}
}

func RunApplication() {
	logger.Println("in RunApplication")
	C.NSApplication_Run(C.NSApplication_SharedApplication())
}
