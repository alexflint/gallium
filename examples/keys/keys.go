package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func handleMenuFirst() {
	log.Println("menu shortcut: first")
}

func handleMenuSecond() {
	log.Println("menu shortcut: second")
}

func handleMenuQuit() {
	os.Exit(0)
}

func onReady(app *gallium.App) {
	gallium.AddGlobalShortcut(gallium.MustParseKeys("shift cmd u"), func() {
		log.Println("global shortcut: U")
	})
	gallium.AddGlobalShortcut(gallium.MustParseKeys("shift cmd o"), func() {
		log.Println("global shortcut: O")
	})
	app.SetMenu([]gallium.Menu{
		{
			Title: "Shortcut",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "First",
					Shortcut: gallium.MustParseKeys("cmd 1"),
					OnClick:  handleMenuFirst,
				},
				gallium.MenuItem{
					Title:    "Second",
					Shortcut: gallium.MustParseKeys("cmd 2"),
					OnClick:  handleMenuSecond,
				},
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: gallium.MustParseKeys("cmd q"),
					OnClick:  handleMenuQuit,
				},
			},
		},
	})
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}
