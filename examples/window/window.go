package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
	_, err := app.OpenWindow(func(win *gallium.WindowOptions) {
		win.X = 0
		win.Y = 0
		win.Width = 800
		win.Height = 600
		win.Title = "Regular"
		win.URL = "http://example.com/"
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
