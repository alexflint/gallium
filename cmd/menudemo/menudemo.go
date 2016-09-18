package main

import "github.com/alexflint/gallium"

func main() {
	gallium.SetMenu([]gallium.Menu{
		gallium.Menu{
			Title: "menudemo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{Title: "Quit"},
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

	gallium.RunApplication()
}
