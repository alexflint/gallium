//go:generate gallium-main

package main

import (
	"log"
	"os"

	"github.com/alexflint/gallium"
)

type menu struct{}

func New() *menu { return &menu{} }

func (*menu) Quit() {
	os.Exit(0)
}

func (*menu) menuAAA() {
	log.Println("you clicked AAA")
}

func (ex *menu) Start(app *gallium.App) {
	app.SetMenu([]gallium.Menu{
		{
			Title: "menudemo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "cmd+q",
					OnClick:  ex.Quit,
				},
			},
		},
		{
			Title: "View",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "AAA",
					Shortcut: "cmd+shift+a",
					OnClick:  ex.menuAAA,
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
}
