Write desktop UI applications in Go using embedded Chromium.

Gallium lets you create windows, menus, dock icons, menubar icons, etc from Go. The idea is then to build all the UI components inside the window in HTML / javascript.

### Warning

This is an extremely early version of Gallium. Most APIs will probably change
before the 1.0 release, and much of the functionality that is already implemented
remains unstable.

### Installation

First install git large file storage, then install Gallium:
```shell
$ brew install git-lfs
$ git lfs install
$ go get github.com/alexflint/gallium  # will not work without git lfs!
```

This will fetch a 92MB framework containing a binary distribution
of the Chromium content module, so it may take a few moments. This
is also why git large file storage must be installed (github has
a limit on file size.)

### Quickstart

```go
package main

import (
  "os"
  "runtime"

  "github.com/alexflint/gallium"
)

func main() {
  runtime.LockOSThread()         // must be the first statement in main - see below
  gallium.Loop(os.Args, OnReady) // must be called from main function
}

func OnReady(app *gallium.App) {
  app.NewWindow("http://example.com/", "Window title!")
}
```

To run the example as a full-fledged UI applicaiton, you need to build
an app bundle:
```shell
$ go build ./example
$ go install github.com/alexflint/gallium/cmd/gallium-bundle
$ gallium-bundle -o example.app example
$ open example.app
```

![Result of the example](https://cloud.githubusercontent.com/assets/640247/18623245/c71c2d26-7def-11e6-9ad3-1a5541d7fc86.png)

Alternatively, you can run the executable directly, but the window
will initially appear behind all other windows, and it will also not
appear in the dock or the switcher, so you will have to find it manually:
```shell
$ go run example.go
```

### Rationale

The goal of Gallium is to make it possible to write cross-platform
desktop UI applications in Go.

### Common pitfalls

- When you run an app bundle with `open app.bundle`, OSX launch services
  discards standard output and standard error. If you need to see
  this output for debugging purposes, use a redirect:
  ```
  gallium.RedirectStdoutStderr("output.log")
  ```
- When you run an app bundle with `open app.bundle`, OSX launch services
  will only start your app if there is not already another instance
  of the same application running, so if your app refuses to start then
  try checking the activity monitor for an already running instance.

### UI thread issues and runtime.LockOSThread

It is very important that the first statement in your main function
be `runtime.LockOSThread()`. The reason is that gallium calls
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
That framework in turn contains `libchromiumcontent.dylib`, which is a 
shared library containing the chromium content module and is distributed
in binary form by the same folks responsible for the excellent Electron
framework. When you build your Go executable, the directives in
`Gallium.framework` instruct the linker to set up the executable to look for
`Gallium.framework` in two places at runtime:
 1. `<dir containing executable>/../Frameworks/Gallium.framework`: this
     will resolve correctly if you choose to build and run your app as a
     bundle (and also means you can distribute the app bundle as a
     self-contained unit).
 2. `$GOPATH/src/github.com/alexflint/dist/Gallium.framework`: this will
     resolve if you choose to run your executable directly.

