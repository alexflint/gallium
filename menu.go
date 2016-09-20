package gallium

/*
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
	"log"
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
	if menu == nil {
		menu = newMenuManager()
	}
	root := C.NSMenu_New(C.CString("<root>"))
	for _, m := range menus {
		menu.add(m, root)
	}
	C.NSApplication_SetMainMenu(root)
}

func AddStatusItem(width int, title string, highlight bool, entries ...MenuEntry) {
	if menu == nil {
		menu = newMenuManager()
	}

	root := C.NSMenu_New(C.CString("<statusbar>"))
	for _, m := range entries {
		menu.add(m, root)
	}
	C.NSStatusBar_AddItem(C.int(width), C.CString(title), C.bool(highlight), root)
}

//export cgo_onMenuClicked
func cgo_onMenuClicked(data unsafe.Pointer) {
	log.Println("in cgo_onMenuClicked")

	if menu == nil {
		log.Println("onMenuClicked called but menu manager was nil")
		return
	}

	if data == nil {
		log.Println("onMenuClicked called but data parameter was nil")
		return
	}

	id := *(*int)(data)
	log.Printf("cgo_onMenuClicked: id=%d", id)

	item, found := menu.items[id]
	if !found {
		log.Printf("onMenuClicked received non-existent ID %d", id)
		return
	}

	if item.OnClick == nil {
		log.Printf("onMenuClicked found %s but OnClick was nil", item.Title)
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
		log.Printf("unexpected menu entry: %T", menu)
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
	log.Println("in RunApplication")
	C.NSApplication_Run()
}
