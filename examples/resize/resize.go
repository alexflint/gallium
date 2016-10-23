package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

var html = `
<!doctype html>
<html><head></head><body>
<a href="/resize">Resize</a>
</body></html>
`

type app struct {
	window *gallium.Window
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, html)
}

func (a *app) handleResize(w http.ResponseWriter, r *http.Request) {
	shape := a.window.Shape()
	log.Println("got shape:", shape)
	shape.Width += 100
	shape.Height += 100
	log.Println("setting shape:", shape)
	a.window.SetShape(shape)
	log.Println("done setting shape")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *app) onReady(app *gallium.App) {
	var err error
	opts := gallium.FramedWindow
	opts.Shape.Width = 300
	opts.Shape.Height = 200
	a.window, err = app.OpenWindow("http://127.0.0.1:9478/", opts)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))

	var app app

	http.HandleFunc("/", app.handleIndex)
	http.HandleFunc("/resize", app.handleResize)
	go http.ListenAndServe(":9478", nil)

	gallium.Loop(os.Args, app.onReady)
}
