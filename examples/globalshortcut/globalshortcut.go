package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func onReady(app *gallium.App) {
	gallium.AddGlobalShortcut("shift+cmd+o", func() {
		log.Println("got global shortcut: O")
	})
	gallium.AddGlobalShortcut("shift+cmd+u", func() {
		log.Println("got global shortcut: U")
	})
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}
