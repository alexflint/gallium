# basic build & run script

# does not do any embedding yet

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# grab deps
go get -u github.com/gopherjs/gopherjs
go get -u github.com/gopherjs/jquery
go get -u github.com/gopherjs/websocket
go get golang.org/x/net/websocket

cd $DIR/browser
gopherjs build ex.go
cp ex.html $DIR/webserver
cp ex.js $DIR/webserver

cd $DIR/webserver
go generate             #embeds assets (html & js)


# bundle
go build ex.go bindata.go
gallium-bundle -o ex.app ex
open ex.app



