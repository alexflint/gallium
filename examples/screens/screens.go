package main

import (
	"os"
	"runtime"

	"github.com/alexflint/gallium"
	"github.com/kr/pretty"
)

func onReady(app *gallium.App) {
	pretty.Println(gallium.FocusedScreen())
	pretty.Println(gallium.Screens())
	x
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}
