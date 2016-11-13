[![GoDoc](https://godoc.org/github.com/alexflint/gallium?status.svg)](https://godoc.org/github.com/alexflint/gallium)
[![Build Status](https://travis-ci.org/alexflint/gallium.svg?branch=master)](https://travis-ci.org/alexflint/gallium)

Write desktop applications in Go, HTML, Javascript, and CSS.

Gallium is a Go library for managing windows, menus, dock icons, and desktop notifications. Each window contains a webview component, in which you code your UI in HTML. Under the hood, the webview is running Chromium.

### Warning

This is an extremely early version of Gallium. Most APIs will probably change
before the 1.0 release, and much of the functionality that is already implemented
remains unstable.

### Platforms

Only OSX is supported right now. I intend to add support for Windows and Linux
soon.

### Discussion

Join the `#gallium` channel over at the Gophers slack. (You can request an invite to
the Gophers slack team [here](https://gophersinvite.herokuapp.com/).)

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
  gallium.Loop(os.Args, onReady) // must be called from main function
}

func onReady(app *gallium.App) {
  app.OpenWindow("http://example.com/", gallium.FramedWindow)
}
```

To run the example as a full-fledged UI application, you need to build
an app bundle:
```shell
$ go build ./example
$ go install github.com/alexflint/gallium/cmd/gallium-bundle
$ gallium-bundle example
$ open example.app
```

![Result of the example](https://cloud.githubusercontent.com/assets/640247/18623245/c71c2d26-7def-11e6-9ad3-1a5541d7fc86.png)

If you run the executable directly without building an app bundle then
many UI elements, such as menus, will not work correctly.

```shell
$ go run example.go
```

### Menus

```go
func main() {
  runtime.LockOSThread()
  gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
  app.OpenWindow("http://example.com/", gallium.FramedWindow)
  app.SetMenu([]gallium.Menu{
    gallium.Menu{
      Title: "demo",
      Entries: []gallium.MenuEntry{
        gallium.MenuItem{
          Title:    "About",
          OnClick:  handleMenuAbout,
        },
        gallium.Separator,
        gallium.MenuItem{
          Title:    "Quit",
          Shortcut: "Cmd+q",
          OnClick:  handleMenuQuit,
        },
      },
    },
  })
}

func handleMenuAbout() {
  log.Println("about clicked")
  os.Exit(0)
}

func handleMenuQuit() {
  log.Println("quit clicked")
  os.Exit(0)
}
```

![Menu demo](https://cloud.githubusercontent.com/assets/640247/20243830/17fbaa8e-a91d-11e6-8eca-7ae7c1418a7e.png)

### Status Bar

```go
func main() {
  runtime.LockOSThread()
  gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
  app.OpenWindow("http://example.com/", gallium.FramedWindow)
  app.AddStatusItem(
    20,
    "statusbar",
    true,
    gallium.MenuItem{
      Title:   "Do something",
      OnClick: handleDoSomething,
    },
    gallium.MenuItem{
      Title:   "Do something else",
      OnClick: handleDoSomethingElse,
    },
  )
}

func handleDoSomething() {
  log.Println("do something")
}

func handleDoSomethingElse() {
  log.Println("do something else")
}
```

![Statusbar demo](https://cloud.githubusercontent.com/assets/640247/18698431/06e9d88c-7f7f-11e6-9fa5-d6be40a07840.png)

### Desktop Notifications

Note that the OSX Notification Center determines whether or not to show any
given desktop notification, so you may need to open the notification center
and scroll to the bottom in order to see notifications during development.

```go
func main() {
  runtime.LockOSThread()
  gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
  img, err := gallium.ImageFromPNG(pngBuffer)
  if err != nil {
    ...
  }

  app.Post(gallium.Notification{
    Title:    "Wow this is a notification",
    Subtitle: "The subtitle",
    Image:    img,
  })
}
```

### Dock icons

To add a dock icon, create a directory named `myapp.iconset` containing the following files:
```
icon_16x16.png          # 16 x 16
icon_16x16@2x.png       # 32 x 32
icon_32x32.png          # 32 x 32
icon_32x32@2x.png       # 64 x 64
icon_128x128.png        # 128 x 128
icon_128x128@2x.png     # 256 x 256
icon_256x256.png        # 256 x 256
icon_256x256@2x.png     # 512 x 512
icon_512x512.png        # 512 x 512
icon_512x512@2x.png     # 1024 x 1024
```

Then build you app with
```shell
gallium-bundle myapp --icon myapp.iconset
```

Alternatively, if you have a `.icns` file:
```shell
gallium-bundle myapp --icon myapp.icns
```

### Writing native code

You can write C or Objective-C code that interfaces directly with native
windowing APIs using golang's excellent C bridging technology, cgo. The
following example uses the macOS native API `[NSWindow setAlphaValue]` to
create a semi-transparent window.

```go
package main

import (
  "log"
  "os"
  "runtime"

  "github.com/alexflint/gallium"
)

/*
#cgo CFLAGS: -x objective-c
#cgo CFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework Cocoa

#include <Cocoa/Cocoa.h>
#include <dispatch/dispatch.h>

void SetAlpha(void* window, float alpha) {
  // Cocoa requires that all UI operations happen on the main thread. Since
  // gallium.Loop will have initiated the Cocoa event loop, we can can use
  // dispatch_async to run code on the main thread.
  dispatch_async(dispatch_get_main_queue(), ^{
    NSWindow* w = (NSWindow*)window;
    [w setAlphaValue:alpha];
  });
}
*/
import "C"

func onReady(ui *gallium.App) {
  window, err := ui.OpenWindow("http://example.com/", gallium.FramedWindow)
  if err != nil {
    log.Fatal(err)
  }
  C.SetAlpha(window.NativeWindow(), 0.5)
}

func main() {
  runtime.LockOSThread()
  gallium.Loop(os.Args, onReady)
}
```

### Relationship to other projects

[Electron](http://electron.atom.io/) is a well-known framework for writing desktop applications in node.js. Electron and Gallium are similar in that the core UI is developed in HTML and javascript, but with Gallium the "outer layer" of logic is written in Go. Both Electron and Gallium use Chromium under the hood, and some of the C components for Gallium were ported from Electron.

The [Chromium Embedded Framework](https://bitbucket.org/chromiumembedded/cef) is a C framework for embedding Chromium into other applications. I investigated CEF as a basis for Gallium but decided to use [libchromiumcontent](https://github.com/electron/libchromiumcontent) instead.

[cef2go](https://github.com/cztomczak/cef2go) is a Go wrapper for Chromium based on CEF, but so far it still requires some manual steps to use as a library.

### Rationale

The goal of Gallium is to make it possible to write cross-platform
desktop UI applications in Go.

### Common pitfalls

- When you run an app bundle with `open Foo.app`, OSX launch services
  discards standard output and standard error. If you need to see
  this output for debugging purposes, use a redirect:
  ```
  gallium.RedirectStdoutStderr("output.log")
  ```
- When you run an app bundle with `open Foo.app`, OSX launch services
  will only start your app if there is not already another instance
  of the same application running, so if your app refuses to start then
  try checking the activity monitor for an already running instance.
- If you run the binary directly without building an app bundle then
  your menus will not show up, and the window will initially appear
  behind other applications.

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

