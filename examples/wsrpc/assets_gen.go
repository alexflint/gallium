// +build generate

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(assets, vfsgen.Options{
		BuildTags: "!dev",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
