package main

import (
	"log"
	"os"

	"github.com/alexflint/gallium"
)

func menuQuit_OnClick() {
	os.Exit(0)
}

func menuAAA_OnClick() {
	log.Println("you clicked AAA")
}

func main() {
	gallium.SetMenu([]gallium.Menu{
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

	gallium.RunApplication()
}
