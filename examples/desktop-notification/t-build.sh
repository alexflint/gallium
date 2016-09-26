# build script

go build notification.go bindata.go
gallium-bundle -o notification.app notification
open notification.app