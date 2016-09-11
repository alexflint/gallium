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
	gallium.CreateWindow("http://example.com/", "Here is a window")
	time.Sleep(time.Second)
	gallium.CreateWindow("http://httpbin.org/", "Here is another window")
}
