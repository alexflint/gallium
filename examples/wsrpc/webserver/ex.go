//go:generate go-bindata -o bindata.go ex.html ex.js

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/alexflint/gallium"

	"github.com/alexflint/gallium/examples/wsrpc/common/webserver"

	"github.com/alexflint/gallium/examples/wsrpc/shared"
)

func main() {
	runtime.LockOSThread()

	// start backend on goroutine.
	// TODO can be done better to ahndle log, errors, signalling...
	go runWebServer()

	// start frontend
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

//Arith is
type Arith struct{}

//Multiply is
func (t *Arith) Multiply(args *shared.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func runWebServer() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	fmt.Println("webserver running on http://localhost:7000/")

	ws, err := webserver.New("localhost:7000", new(Arith))

	if err != nil {
		log.Fatal(err)
	}
	var reply int
	for i := 0; i < 10; i++ {
		log.Println("writing", i, "to browser")
		if err := ws.Call("GUI.Write", &shared.Args{C: strconv.Itoa(i)}, &reply); err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}
}
