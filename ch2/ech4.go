package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var p = flag.String("p", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *p))
	if !*n {
		fmt.Println()
	}
}
