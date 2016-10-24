package main

import (
	"fmt"
	"net/http"
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
	opt := gallium.FramedWindow
	opt.Title = "Framed Window"
	_, err := app.OpenWindow("", opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opt = gallium.FramelessWindow
	opt.Title = "Frameless Window"
	_, err = app.OpenWindow("http://example.com/", opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<p>loaded local gallium server</p>")
	})
}
