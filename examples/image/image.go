package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alexflint/gallium"
)

func main() {
	buf, err := ioutil.ReadFile("gopher.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img, err := gallium.ImageFromPNG(buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gallium.ImageToPNG(img, "/tmp/gopher.png")
}
