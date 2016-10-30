package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

/*
#cgo CFLAGS: -x objective-c
#cgo CFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework Cocoa

#include <Cocoa/Cocoa.h>
#include <dispatch/dispatch.h>

void SetAlpha(void* window, float alpha) {
  // Cocoa requires that all UI operations happen on the main thread. Since
  // gallium.Loop will have initiated the Cocoa event loop, we can can use
  // dispatch_async to run code on the main thread.
  dispatch_async(dispatch_get_main_queue(), ^{
	NSWindow* w = (NSWindow*)window;
	[w setAlphaValue:alpha];
  });
}
*/
import "C"

func onReady(ui *gallium.App) {
	window, err := ui.OpenWindow("http://example.com/", gallium.FramedWindow)
	if err != nil {
		log.Fatal(err)
	}
	C.SetAlpha(window.NativeWindow(), 0.5)
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}
