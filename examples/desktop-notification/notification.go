//go:generate go-bindata -o bindata.go gopher.png

package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()

	// TODO: add to gallium core a cross platform way to find correct folders for logs, settings / configs & data. These are files that persist between installs & upgrades.
	// "$HOME/Library/Logs/Gallium.log"

	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {

	// consider more modern bindata implemenations that work on desktop, browser and mobile.
	img, err := gallium.ImageFromPNG(MustAsset("gopher.png"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app.Post(gallium.Notification{
		Title:    "Wow this is a notification",
		Subtitle: "The subtitle",
		Image:    img,
	})
}
