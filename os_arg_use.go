package main

import (
	// "flag"
	"fmt"
	"os"
)

var p = fmt.Println

func main() {
	p(os.Args)
	if len(os.Args) > 1 {
		args := os.Args[1]
		switch string(args) {
		case "-v", "-version":
			p("Version is ", 123)
			os.Exit(0)
		case "-h", "-help":
			p("help msg is no 咬我啊")
			os.Exit(0)
		}
	}
}
