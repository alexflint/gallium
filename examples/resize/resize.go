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
<div><a href="/resize">Resize</a></div>
<div><a href="/window">New window</a></div>
<div><a href="/load">Load example.com</a></div>
<div><a href="/opendev">Open Dev Tools</a></div>
<div><a href="/closedev">Close Dev Tools</a></div>
</body></html>
`

type app struct {
	ui     *gallium.App
	window *gallium.Window
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, html)
}

func (a *app) handleResize(w http.ResponseWriter, r *http.Request) {
	shape := a.window.Shape()
	shape.Width += 100
	shape.Height += 100
	a.window.SetShape(shape)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *app) handleWindow(w http.ResponseWriter, r *http.Request) {
	a.ui.OpenWindow("http://www.example.com/", gallium.FramedWindow)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *app) handleLoad(w http.ResponseWriter, r *http.Request) {
	log.Println("at handleLoad")
	a.window.LoadURL("http://www.example.com/")
}

func (a *app) handleOpenDev(w http.ResponseWriter, r *http.Request) {
	a.window.OpenDevTools()
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (a *app) handleCloseDev(w http.ResponseWriter, r *http.Request) {
	a.window.CloseDevTools()
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func onReady(ui *gallium.App) {
	opts := gallium.FramedWindow
	opts.Shape.Width = 300
	opts.Shape.Height = 200
	window, err := ui.OpenWindow("http://127.0.0.1:9478/", opts)
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		ui:     ui,
		window: window,
	}
	http.HandleFunc("/", app.handleIndex)
	http.HandleFunc("/resize", app.handleResize)
	http.HandleFunc("/window", app.handleWindow)
	http.HandleFunc("/load", app.handleLoad)
	http.HandleFunc("/opendev", app.handleOpenDev)
	http.HandleFunc("/closedev", app.handleCloseDev)
	go http.ListenAndServe(":9478", nil)
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, onReady)
}
