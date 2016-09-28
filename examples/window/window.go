package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
	w, err := app.OpenWindow(func(win *gallium.WindowOptions) {
		win.X = 100
		win.Y = 100
		win.Width = 800
		win.Height = 800
		win.Title = "Demo Window"
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = w
}
