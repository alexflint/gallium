Write desktop UI applications in Go using embedded Chromium.

### Installation

This is an extremely early version of Gallium. All APIs will likely change
significantly before verison 1.0, and much of the functionality that has'
been implemented is unstable.

### Installation

First install git large file storage (you will be downloading a 90MB C library):
```shell
$ brew install git-lfs
$ git lfs install
```

Install gallium:
```shell
$ go get github.com/alexflint/gallium
```

This will fetch a 92MB framework containing a binary distribution
of the Chromium content module, so it may take a few moments.

### Quick start

```go
package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()  // must be the first statement in main - see below
	gallium.Loop(os.Args, OnReady)  // must be called from main function
}

func OnReady(browser *gallium.Browser) {
	browser.CreateWindow("http://example.com/", "Example Title")
}
```

To run the example as a full-fledged UI applicaiton, you need to build
an app bundle:
```shell
$ go install github.com/alexflint/gallium/cmd/gallium-bundle
$ go build ./example
$ gallium-bundle -o example.app example
$ open example.app
```

You can also run the example directly, but the window will initially appear
behind all other windows, and will not appear in the dock or the switcher, so you
will have to find it manually:
```shell
$ go run example.go
```

### Rationale

The goal of Gallium is to make it possible to write cross-platform
desktop UI applications in Go.

### runtime.LockOSThread and the "main" thread

It is very important that the first statement in your main function
be `runtime.LockOSThread()`. The reason for this is that gallium calls
out to various C functions in order to create and manage OSX UI elements,
and many of these are required to be called from the first thread
created by the process. But the Go runtime creates many threads and any
one piece of Go code could end up running on any thread. The solution
is `runtime.LockOSThread`, which tells the Go scheduler to lock the
current goroutine so that it will only ever run on the current thread.
Since the main function always starts off on the main thread, this wil
guarantee that the later call to `gallium.Loop` will also be on the main
thread. At this point gallium takes ownership of this thread for its main
event loop and calls the `OnReady` callback in a separate goroutine.
From this point forward it is safe to call gallium functions from any
goroutine.

### Shared libraries and linking issues

Gallium is based on Chromium, which it accesses via `Gallium.framework`.
That frame contains `libchromiumcontent.dylib`, which is a shared library
containing the chromium content module and is distributed in binary form
by the same folks responsible for the excellent Electron framework. When
you build your Go executable, the directives in `Gallium.framework` 
instruct the linker to set up the binary to look for `Gallium.framework`
in two places at runtime:
 1. `<dir containing executable>/../Frameworks/Gallium.framework`: this
     will resolve correctly if you choose to build and run your app as a
     bundle (and also means you can distribute the app bundle as a
     self-contained unit).
 2. `$GOPATH/src/github.com/alexflint/dist/Gallium.framework`: this will
     resolve if you choose to run your executable directly.

