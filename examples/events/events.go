package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv("$HOME/Library/Logs/Gallium.log"))
	e := &events{}
	gallium.Loop(os.Args, e.onReady)
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		<p>Events:</p>
		<ul id="events"></ul>
		<script type="text/javascript">
		function append(txt) {
			var ul = document.getElementById("events");
			var li = document.createElement('li');
			li.innerHTML = txt;
			ul.appendChild(li);
		}

		append("connecting");

		var source = new EventSource("/gallium.events");

		source.onopen = function () { append("connected"); };
		source.onerror = function () { append("connection failed"); };

		source.addEventListener("menu", function (event) {
			append("Clicked: "+event.data);
		});
		</script>
		`)
	})
}

type events struct {
	app *gallium.App
}

func (e *events) aaa() {
	log.Println(e.app.Emit("menu", "aaa"))
}

func (e *events) bbb() {
	log.Println(e.app.Emit("menu", "bbb"))
}

func (e *events) ccc() {
	log.Println(e.app.Emit("menu", "ccc"))
}

func (e *events) onReady(app *gallium.App) {
	e.app = app
	app.SetMenu([]gallium.Menu{
		{
			Title: "Events",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: "cmd+q",
					OnClick:  func() { os.Exit(0) },
				},
				gallium.MenuItem{
					Title:    "New Window",
					Shortcut: "cmd+n",
					OnClick:  func() { app.OpenWindow("", gallium.FramedWindow) },
				},
			},
		},
		{
			Title: "View",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:    "AAA",
					Shortcut: "cmd+shift+a",
					OnClick:  e.aaa,
				},
				gallium.MenuItem{Title: "BBB", OnClick: e.bbb},
				gallium.MenuItem{Title: "CCC", OnClick: e.ccc},
			},
		},
	})
	app.OpenWindow("", gallium.FramedWindow)
}
