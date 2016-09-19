package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, OnReady)
}

func handleMenuQuit() {
	log.Println("quit clicked")
	os.Exit(0)
}

func OnReady(browser *gallium.App) {
	browser.NewWindow("http://example.com/", "Here is a window")
	gallium.SetMenu([]gallium.Menu{
		gallium.Menu{
			Title: "menudemo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "Cmd+q",
					OnClick:  handleMenuQuit,
				},
			},
		},
	})
}
