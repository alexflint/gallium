package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func menuQuit_OnClick() {
	os.Exit(0)
}

func menuAAA_OnClick() {
	log.Println("you clicked AAA")
}

func main() {
	runtime.LockOSThread()
	gallium.Loop(os.Args, onReady)

}

func onReady(app *gallium.App) {
	app.NewWindow("http://example.com/", "Here is a window")

	app.SetMenu([]gallium.Menu{
		{
			Title: "menudemo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "cmd+q",
					OnClick:  menuQuit_OnClick,
				},
			},
		},
		{
			Title: "View",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "AAA",
					Shortcut: "cmd+shift+a",
					OnClick:  menuAAA_OnClick,
				},
				gallium.MenuItem{Title: "BBB"},
				gallium.MenuItem{Title: "CCC"},
			},
		},
		{
			Title: "Help",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{Title: "What"},
				gallium.MenuItem{Title: "Is"},
				gallium.MenuItem{Title: "This?"},
			},
		},
	})

	log.Println("menu loaded")
}
