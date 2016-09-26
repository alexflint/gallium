# basic build sript

# does not do any embedding yet

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $DIR/browser
go get  
gopherjs build ex.go

exit

cp ex.html $DIR/webserver
cp ex.js $DIR/webserver

cd $DIR/webserver

go run ex.go

