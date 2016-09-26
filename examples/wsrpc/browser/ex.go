package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gopherjs/jquery"

	"github.com/alexflint/gallium/examples/wsrpc/common/browser"

	"github.com/alexflint/gallium/examples/wsrpc/shared"
)

//convenience:
var jQuery = jquery.NewJQuery

//aa
const (
	INPUT   = "button"
	OUTPUT  = "#output"
	OUTPUT2 = "#output2"
)

//GUI is
type GUI struct{}

//Write is
func (g *GUI) Write(args *shared.Args, reply *int) error {
	//show welcome message:
	jQuery(OUTPUT2).SetText("string from server:" + args.C)
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	
	fmt.Println("browser on http://localhost:7000")

	b, err := browser.New("localhost:7000", new(GUI))
	if err != nil {
		log.Fatal(err)
	}
	//	defer b.Close()
	i := 0
	jQuery(INPUT).On(jquery.CLICK, func(e jquery.Event) {
		go func() {
			i++
			args := shared.Args{A: i, B: i}
			var reply int
			err = b.Call("Arith.Multiply", args, &reply)
			if err != nil {
				log.Fatal("arith error:", err)
			}
			//show welcome message:
			ii := strconv.Itoa(i)
			jQuery(OUTPUT).SetText("result of " + ii + "x" + ii + " from server:" + strconv.Itoa(reply))
		}()
	})

}
