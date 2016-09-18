package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.Loop(os.Args, OnReady)
}

func handleMenuQuit() {
	log.Println("quit clicked")
	os.Exit(0)
}

func OnReady(browser *gallium.Browser) {
	browser.CreateWindow("http://example.com/", "Here is a window")
	gallium.SetMenu([]gallium.Menu{
		gallium.Menu{
			Title: "menudemo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "cmd+q",
					OnClick:  handleMenuQuit,
				},
			},
		},
		gallium.Menu{
			Title: "View",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{Title: "AAA"},
				gallium.MenuItem{Title: "BBB"},
				gallium.MenuItem{Title: "CCC"},
			},
		},
		gallium.Menu{
			Title: "Help",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{Title: "What"},
				gallium.MenuItem{Title: "Is"},
				gallium.MenuItem{Title: "This?"},
			},
		},
	})
}
