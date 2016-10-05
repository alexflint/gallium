//go:generate gallium-main

package main

import (
	"fmt"
	"os"

	"github.com/alexflint/gallium"
)

type window struct{}

func New() *window { return &window{} }

func (*window) Start(app *gallium.App) {
	opt := gallium.FramedWindow
	opt.Title = "Framed Window"
	_, err := app.OpenWindow("http://example.com/", opt)
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
