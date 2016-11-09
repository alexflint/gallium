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
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, OnReady)
}

func OnReady(app *gallium.App) {
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
				gallium.Separator,
				gallium.MenuItem{Title: "DDD"},
				gallium.MenuItem{Title: "EEE"},
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
}
