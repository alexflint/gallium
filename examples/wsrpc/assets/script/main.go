// +build js

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gopherjs/jquery"
)

//convenience:
var jQuery = jquery.NewJQuery

//aa
const (
	INPUT   = "button"
	OUTPUT  = "#output"
	OUTPUT2 = "#output2"
)


func main() {
	
	fmt.Println("Script is working! Try making some changes to it.")
	fmt.Println("Browser on http://localhost:7000")

}
