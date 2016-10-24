//go:generate go-bindata -o bindata.go icon.png

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
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

func onReady(app *gallium.App) {
	img, err := gallium.ImageFromPNG(MustAsset("icon.png"))
	if err != nil {
		fmt.Println("unable to decode icon-gray.png:", err)
		os.Exit(1)
	}

	app.AddStatusItem(gallium.StatusItemOptions{
		Image:     img,
		Width:     27.5,
		Highlight: true,
		Menu: []gallium.MenuEntry{
			gallium.MenuItem{Title: "Do something", OnClick: handleDoSomething},
			gallium.MenuItem{Title: "Do something else", OnClick: handleDoSomethingElse},
		},
	})
}
