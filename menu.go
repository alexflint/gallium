package gallium

/*
#cgo CFLAGS: -mmacosx-version-min=10.8
#cgo LDFLAGS: -mmacosx-version-min=10.8

#include <stdlib.h>
#include "lib/api/gallium.h"
#include "lib/api/menu.h"

//#include "_cgo_export.h"

extern void cgo_onMenuClicked(void*);

static inline gallium_nsmenuitem_t* helper_NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* keyEquivalent,
	int id) {

	int *idholder = (int*)malloc(sizeof(int));
	*idholder = id;
	return NSMenu_AddMenuItem(menu, title, keyEquivalent, &cgo_onMenuClicked, idholder);
}

*/
import "C"
import "unsafe"

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
		logger.Printf("onMenuClicked found item with nil OnClick: %s", item.Title)
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
		item := C.NSMenu_AddMenuItem(parent, C.CString(menu.Title), C.CString(""), nil, nil)
		submenu := C.NSMenu_New(C.CString(menu.Title))
		C.NSMenuItem_SetSubmenu(item, submenu)
		for _, entry := range menu.Entries {
			m.add(entry, submenu)
		}
	case MenuItem:
		id := len(m.items)
		m.items[id] = menu
		C.helper_NSMenu_AddMenuItem(parent, C.CString(menu.Title), C.CString(menu.Shortcut), C.int(id))
	default:
		logger.Printf("unexpected menu entry: %T", menu)
	}
}

func RunApplication() {
	logger.Println("in RunApplication")
	C.NSApplication_Run(C.NSApplication_SharedApplication())
}
