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
	"errors"
	"fmt"
	"log"
	"strings"
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

// Modifier represents zero or more modifier keys (control, shift, option, etc)
type Modifier int

// these must be kept in sync with the gallium_modifier_t enum in cocoa.h
const (
	ModifierCmd Modifier = 1 << iota
	ModifierCtrl
	ModifierCmdOrCtrl
	ModifierAltOrOption
	ModifierFn
	ModifierShift
)

// KeyCombination represents a key together with zero or more modifiers. It
// is used to set up keyboard shortcuts.
type KeyCombination struct {
	Key       string
	Modifiers Modifier
}

var (
	errEmptyKey      = errors.New("empty key")
	errEmptyShortcut = errors.New("empty shortcut")
)

// ParseKeys parses a key combination specified as a string like "cmd shift a".
func ParseKeys(s string) (KeyCombination, error) {
	// for backwards compatibility split on both "+" and " "
	parts := strings.Split(s, " ")
	if strings.Contains(s, "+") {
		parts = strings.Split(s, "+")
	}
	if len(parts) == 0 {
		return KeyCombination{}, errEmptyShortcut
	}
	var keys KeyCombination
	keys.Key = parts[len(parts)-1]
	if len(keys.Key) == 0 {
		return KeyCombination{}, errEmptyKey
	}
	for _, part := range parts[:len(parts)-1] {
		switch strings.ToLower(part) {
		case "cmd":
			keys.Modifiers |= ModifierCmd
		case "ctrl":
			keys.Modifiers |= ModifierCtrl
		case "cmdctrl":
			keys.Modifiers |= ModifierCmdOrCtrl
		case "alt":
			keys.Modifiers |= ModifierAltOrOption
		case "option":
			keys.Modifiers |= ModifierAltOrOption
		case "fn":
			keys.Modifiers |= ModifierFn
		case "shift":
			keys.Modifiers |= ModifierShift
		default:
			return KeyCombination{}, fmt.Errorf("unknown modifier: %s", part)
		}
	}
	return keys, nil
}

// MustParseKeys is like ParseKeys but panics on error
func MustParseKeys(s string) KeyCombination {
	keys, err := ParseKeys(s)
	if err != nil {
		panic(err)
	}
	return keys
}

// AddGlobalShortcut calls the handler whenever the key combination is pressed
// in any application.
func AddGlobalShortcut(keys KeyCombination, handler func()) {
	id := nextID
	nextID++
	handlers[id] = handler

	C.helper_AddGlobalShortcut(
		C.int(id),
		C.CString(keys.Key),
		C.gallium_modifier_t(keys.Modifiers))
}
