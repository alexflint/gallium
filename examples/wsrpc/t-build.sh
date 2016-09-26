# basic build sript

# does not do any embedding yet

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $DIR/browser
go get  
gopherjs build ex.go
cp ex.html $DIR/webserver
cp ex.js $DIR/webserver

cd $DIR/webserver
#go run ex.go

#TODO: embedd assets

go build ex.go
gallium-bundle -o ex.app ex
open ex.app



