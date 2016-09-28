# Prod - Build and run 
# DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# build
go generate
go build

# bundle
gallium-bundle -o wsrpc.app wsrpc
open wsrpc.app



