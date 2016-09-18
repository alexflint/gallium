package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "lib/api/gallium.h"
#include "lib/api/menu.h"

// It does not seem that we can import "_cgo_export.h" from here
extern void cgo_onMenuClicked(void*);

// This is a wrapper around NSMenu_AddMenuItem that adds the function pointer
// argument, since this does not seem to be possible from Go directly.
static inline gallium_nsmenuitem_t* helper_NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* shortcutKey,
	gallium_modifier_t shortcutModifier,
	void *callbackArg) {

	return NSMenu_AddMenuItem(menu, title, shortcutKey, shortcutModifier, &cgo_onMenuClicked, callbackArg);
}

*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type MenuEntry interface {
	menu()
}

type MenuItem struct {
	Title    string
	Shortcut string
	OnClick  func()
}

func (MenuItem) menu() {}

type Menu struct {
	Title   string
	Entries []MenuEntry
}

func (Menu) menu() {}

var menu *menuManager

func SetMenu(menus []Menu) {
	menu = newMenuManager()
	root := C.NSMenu_New(C.CString("<root>"))
	for _, m := range menus {
		menu.add(m, root)
	}
	C.NSApplication_SetMainMenu(C.NSApplication_SharedApplication(), root)
}

//export cgo_onMenuClicked
func cgo_onMenuClicked(data unsafe.Pointer) {
	logger.Println("in cgo_onMenuClicked")

	if menu == nil {
		logger.Println("onMenuClicked called but menu manager was nil")
		return
	}

	if data == nil {
		logger.Println("onMenuClicked called but data parameter was nil")
		return
	}

	id := *(*int)(data)
	logger.Printf("cgo_onMenuClicked: id=%d", id)

	item, found := menu.items[id]
	if !found {
		logger.Printf("onMenuClicked received non-existent ID %d", id)
		return
	}

	if item.OnClick == nil {
		logger.Printf("onMenuClicked found %s but OnClick was nil", item.Title)
		return
	}

	item.OnClick()
}

type menuManager struct {
	items map[int]MenuItem
}

func newMenuManager() *menuManager {
	return &menuManager{make(map[int]MenuItem)}
}

func (m *menuManager) add(menu MenuEntry, parent *C.gallium_nsmenu_t) {
	switch menu := menu.(type) {
	case Menu:
		item := C.NSMenu_AddMenuItem(parent, C.CString(menu.Title), nil, 0, nil, nil)
		submenu := C.NSMenu_New(C.CString(menu.Title))
		C.NSMenuItem_SetSubmenu(item, submenu)
		for _, entry := range menu.Entries {
			m.add(entry, submenu)
		}
	case MenuItem:
		id := len(m.items)
		m.items[id] = menu

		callbackArg := C.malloc(C.sizeof_int)
		*(*C.int)(callbackArg) = C.int(id)

		key, modifiers, _ := parseShortcut(menu.Shortcut)

		C.helper_NSMenu_AddMenuItem(
			parent,
			C.CString(menu.Title),
			C.CString(key),
			C.gallium_modifier_t(modifiers),
			callbackArg)
	default:
		logger.Printf("unexpected menu entry: %T", menu)
	}
}

func parseShortcut(s string) (key string, modifiers int, err error) {
	parts := strings.Split(s, "+")
	if len(parts) == 0 {
		return "", 0, fmt.Errorf("empty shortcut")
	}
	key = parts[len(parts)-1]
	if len(key) == 0 {
		return "", 0, fmt.Errorf("empty key")
	}
	for _, part := range parts[:len(parts)-1] {
		switch strings.ToLower(part) {
		case "cmd":
			modifiers |= int(C.GalliumCmdModifier)
		case "ctrl":
			modifiers |= int(C.GalliumCmdModifier)
		case "cmdctrl":
			modifiers |= int(C.GalliumCmdOrCtrlModifier)
		case "alt":
			modifiers |= int(C.GalliumAltOrOptionModifier)
		case "option":
			modifiers |= int(C.GalliumAltOrOptionModifier)
		case "fn":
			modifiers |= int(C.GalliumFunctionModifier)
		case "shift":
			modifiers |= int(C.GalliumShiftModifier)
		default:
			return "", 0, fmt.Errorf("unknown modifier: %s", part)
		}
	}
	return
}

func RunApplication() {
	logger.Println("in RunApplication")
	C.NSApplication_Run(C.NSApplication_SharedApplication())
}
