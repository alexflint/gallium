package main

import (
	"os"
	"runtime"
	"time"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()
	go Main()
	gallium.Loop(os.Args)
}

func Main() {
	time.Sleep(time.Second)
	gallium.CreateWindow("Here is a window")
	time.Sleep(time.Second)
	gallium.CreateWindow("Here is another window")
}
