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

func handleDoSomething() {
	log.Println("do something")
}

func handleDoSomethingElse() {
	log.Println("do something else")
}

func OnReady(app *gallium.App) {
	app.NewWindow("http://example.com/", "Here is a window")
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
	app.AddStatusItem(
		20,
		"demo",
		true,
		gallium.MenuItem{
			Title:   "Do something",
			OnClick: handleDoSomething,
		},
		gallium.MenuItem{
			Title:   "Do something else",
			OnClick: handleDoSomethingElse,
		},
	)
}
