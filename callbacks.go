package gallium

import (
	"log"
	"unsafe"
)

import "C"

// This file contains all Go functions that are exported to cgo because
// the presence of an export means that the C prelude gets copied into
// two locations.

//export cgo_onReady
func cgo_onReady(appId int) {
	log.Printf("cgo_onReady called with %d", appId)
	// do not actually call the user function from here because that would
	// block the UI loop
	apps.get(appId).ready <- struct{}{}
}

//export cgo_onMenuClicked
func cgo_onMenuClicked(data unsafe.Pointer) {
	log.Println("in cgo_onMenuClicked")

	if menuMgr == nil {
		log.Println("onMenuClicked called but menu manager was nil")
		return
	}

	if data == nil {
		log.Println("onMenuClicked called but data parameter was nil")
		return
	}

	id := *(*int)(data)
	log.Printf("cgo_onMenuClicked: id=%d", id)

	item, found := menuMgr.items[id]
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
