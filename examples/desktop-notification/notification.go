//go:generate gallium-main
//go:generate go-bindata -o bindata.go gopher.png

package main

import (
	"log"
	"os"

	"github.com/alexflint/gallium"
)

type notifications struct{}

func New() *notifications { return &notifications{} }

func (*notifications) Start(app *gallium.App) {
	img, err := gallium.ImageFromPNG(MustAsset("gopher.png"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app.Post(gallium.Notification{
		Title:    "Wow this is a notification",
		Subtitle: "The subtitle",
		Image:    img,
	})
}
