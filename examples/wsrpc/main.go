//go:generate go run assets_gen.go assets.go

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/shurcooL/httpgzip"

	"github.com/alexflint/gallium"
)

var httpFlag = flag.String("http", ":7000", "Listen for HTTP connections on this address.")

func main() {
	runtime.LockOSThread()

	flag.Parse()

	// TODO: Detect if the build tag is "dev", and load only into browser then.

	// First, start backend on goroutine.
	// TODO can be done better to handle log, errors, signalling...
	go mainServer()

	// Second, start frontend
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	gallium.Loop(os.Args, OnReady)

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

// OnReady ...
func OnReady(app *gallium.App) {

	// TODO. use flag instead of hardcoded port.
	app.NewWindow("http://localhost:7000/ex.html", "Here is a window")

	app.SetMenu([]gallium.Menu{
		{
			Title: "demo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "Cmd+q",
					OnClick:  handleMenuQuit,
				},
			},
		},
	})
	app.AddStatusItem(
		20,
		"demo",
		true,
		gallium.MenuItem{
			Title:   "Do something",
			OnClick: handleDoSomething,
		},
		gallium.MenuItem{
			Title:   "Do something else",
			OnClick: handleDoSomethingElse,
		},
	)
}

// All Backend below

func mainServer() {

	printServingAt(*httpFlag)

	// setup main serving
	http.HandleFunc("/", mainHandler)
	http.Handle("/assets/", httpgzip.FileServer(assets, httpgzip.FileServerOptions{ServeError: httpgzip.Detailed}))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(*httpFlag, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}

}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	t, err := loadTemplates()
	if err != nil {
		log.Println("loadTemplates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = struct {
		Animals string
	}{
		Animals: "gophers",
	}

	err = t.ExecuteTemplate(w, "index.html.tmpl", data)
	if err != nil {
		log.Println("t.Execute:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loadTemplates() (*template.Template, error) {
	t := template.New("").Funcs(template.FuncMap{})
	t, err := vfstemplate.ParseGlob(assets, t, "/assets/*.tmpl")
	return t, err
}

func printServingAt(addr string) {
	hostPort := addr
	if strings.HasPrefix(hostPort, ":") {
		hostPort = "localhost" + hostPort
	}
	fmt.Printf("serving at http://%s/\n", hostPort)
}
