# wsrpc


This is just a simple example an application with a Frontend (running in the embedded browser) and a Backend running in the same process serving http.
Features:
- Automatic embedding of resources in backend Server
    - In dev mode allows live reload.
- Gopherjs setup to use if you want.
- Backend setup


## Building

## Get Deps
````
chmod +x *.sh
./t-getdeps.sh
`````

### Development Build

Accesses assets from disk directly:
./t-build-dev.sh

### Production Build

All assets are statically embedded in the binary, so it can run standalone in any folder:
./t-build-prod.sh




