//go:generate go-bindata-assetfs static/...

package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	// gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}

func init() {
	http.Handle("/", http.FileServer(assetFS()))
}

func onReady(app *gallium.App) {
	app.NewWindow("", "Here is a window")
}
