# Packages

We need use some cross platform Golang Packages for some of the common functionality we need outside of the main Browser window.

## Embedded assets
https://github.com/shurcooL/vfsgen
- efficient one.
- useful for the assets needed to be served inside the browser over http File system.
- Also has links to all the other ones.

https://github.com/hyangah/mgodoc
- Not of interest for desktop, but for mobile using webviews.
- The pont is that it uses the golang.org/x/tools/godoc/vfs/zipfs like others, but does not return a reader, but instead a []byte to it can be used on mobile, whilst allow golang templates etc.


## Systems Paths
"github.com/kardianos/osext

## desktop notifications
https://github.com/getlantern/systray
- cross platform
- well supported as its a big project using it.


## Packager / Installer
Some projects demand a proper installer for users to install the software.

The golang library itself has a very nice installer.
It uses golang with templates, etc
https://github.com/golang/build/tree/master/cmd/release
- windows wix
- osx pck
- linux debian


## Updater
Also you  want to be able to do remote updates of the software too.

https://github.com/inconshreveable/go-update
- simple and works

https://github.com/keybase/go-updater
- complex and handles lots more edge cases.
- should be well supported as Keybase is a heavy golang user.
