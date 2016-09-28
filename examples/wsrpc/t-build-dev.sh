# Dev - Build and run 
#DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# build
go build -tags=dev
./wsrpc

# bundle (we dont want to)
#gallium-bundle -o wsrpc.app wsrpc
#open wsrpc.app






