// +build dev

package main

import (
	"net/http"

	"github.com/shurcooL/go/gopherjs_http"
	"github.com/shurcooL/httpfs/union"
)

var assets = union.New(map[string]http.FileSystem{
	"/assets": gopherjs_http.NewFS(http.Dir("assets")),
})
