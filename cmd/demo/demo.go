package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexflint/gallium"
)

func main() {
	f, err := os.OpenFile("/Users/alex/gallium.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	//f, err := os.Create("/Users/alex/gallium.log")
	if err == nil {
		defer f.Close()
		fmt.Fprintf(f, "%v Gallium invoked with args %v\n", time.Now(), strings.Join(os.Args, " "))
	}

	gallium.Run(os.Args)
}
