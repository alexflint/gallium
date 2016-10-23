package gallium

/*
#include <stdlib.h>
#include "gallium/gallium.h"
#include "gallium/cocoa.h"

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

	return NSMenu_AddMenuItem(
		menu,
		title,
		shortcutKey,
		shortcutModifier,
		&cgo_onMenuClicked,
		callbackArg);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// MenuEntry is the interface for menus and menu items.
type MenuEntry interface {
	menu()
}

// A MenuItem has a title and can be clicked on. It is a leaf node in the menu tree.
type MenuItem struct {
	Title    string
	Shortcut string
	OnClick  func()
}

func (MenuItem) menu() {}

// A Menu has a title and a list of entries. It is a non-leaf node in the menu tree.
type Menu struct {
	Title   string
	Entries []MenuEntry
}

func (Menu) menu() {}

// menuMgr is the singleton that owns the menu
var menuMgr *menuManager

// menuManager translates Menus and MenuItems to their native equivalent (e.g. NSMenuItem on macOS)
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

func (app *App) SetMenu(menus []Menu) {
	if menuMgr == nil {
		menuMgr = newMenuManager()
	}
	root := C.NSMenu_New(C.CString("<root>"))
	for _, m := range menus {
		menuMgr.add(m, root)
	}
	C.NSApplication_SetMainMenu(root)
}

type StatusItemOptions struct {
	Image     *Image      // image to show in the status bar, must be non-nil
	Width     float64     // width of item in pixels (zero means automatic size)
	Highlight bool        // whether to highlight the item when clicked
	Menu      []MenuEntry // the menu to display when the item is clicked
}

func (app *App) AddStatusItem(opts StatusItemOptions) {
	if menuMgr == nil {
		menuMgr = newMenuManager()
	}
	if opts.Image == nil {
		panic("status item image must not be nil")
	}

	menu := C.NSMenu_New(C.CString("<statusbar>"))
	for _, m := range opts.Menu {
		menuMgr.add(m, menu)
	}
	C.NSStatusBar_AddItem(
		opts.Image.c,
		C.float(opts.Width),
		C.bool(opts.Highlight),
		menu)
}

// Image holds a handle to a platform-specific image structure (e.g. NSImage on macOS).
type Image struct {
	c *C.gallium_nsimage_t
}

var (
	ErrImageDecodeFailed = errors.New("image could not be decoded")
)

// ImageFromPNG creates an image from a buffer containing a PNG-encoded image.
func ImageFromPNG(buf []byte) (*Image, error) {
	cbuf := C.CBytes(buf)
	defer C.free(cbuf)
	cimg := C.NSImage_NewFromPNG(cbuf, C.int(len(buf)))
	if cimg == nil {
		return nil, ErrImageDecodeFailed
	}
	return &Image{cimg}, nil
}

// Notification represents a desktop notification
type Notification struct {
	Title             string
	Subtitle          string
	InformativeText   string
	Image             *Image
	Identifier        string
	ActionButtonTitle string
	OtherButtonTitle  string
}

// Post shows the given desktop notification
func (app *App) Post(n Notification) {
	var cimg *C.gallium_nsimage_t
	if n.Image != nil {
		cimg = n.Image.c
	}
	cn := C.NSUserNotification_New(
		C.CString(n.Title),
		C.CString(n.Subtitle),
		C.CString(n.InformativeText),
		cimg,
		C.CString(n.Identifier),
		len(n.ActionButtonTitle) > 0,
		len(n.OtherButtonTitle) > 0,
		C.CString(n.ActionButtonTitle),
		C.CString(n.OtherButtonTitle))

	C.NSUserNotificationCenter_DeliverNotification(cn)
}

// RunApplication is for debugging only. It allows creation of menus and
// desktop notifications without firing up any parts of chromium. It will
// be removed before the 1.0 release.
func RunApplication() {
	C.NSApplication_Run()
}
