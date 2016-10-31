//go:generate gallium-main

package main

import (
	"log"
	"os"

	"github.com/alexflint/gallium"
)

type simple struct{}

func New() *simple { return &simple{} }

func handleMenuQuit() {
	log.Println("quit clicked")
	os.Exit(0)
}

func handleDoSomething() {
	log.Println("do something")
}

func handleDoSomethingElse() {
	log.Println("do something else")
}

func (*simple) Start(app *gallium.App) {
	app.OpenWindow("http://example.com/", gallium.FramedWindow)
	app.SetMenu([]gallium.Menu{
		{
			Title: "demo",
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
