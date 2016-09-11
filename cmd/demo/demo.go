package main

import (
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.Loop(os.Args, OnReady)
}

func OnReady(browser *gallium.Browser) {
	browser.CreateWindow("http://example.com/", "Here is a window")
	browser.CreateWindow("http://httpbin.org/", "Here is another window")
}
