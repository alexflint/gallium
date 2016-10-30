package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

var index = `<!doctype html>
<html><head></head><body>
<textarea rows="10" cols="80"></textarea>
</body></html>`

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, index)
}

func onReady(app *gallium.App) {
	http.HandleFunc("/", handleIndex)
	go http.ListenAndServe(":8967", nil)

	opt := gallium.FramelessWindow
	opt.Title = "Frameless Window"
	_, err := app.OpenWindow("http://127.0.0.1:8967/", opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
