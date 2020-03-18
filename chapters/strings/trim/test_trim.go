package main

import (
	"fmt"
	"strings"
)

func main() {
	s := " Hello world  \n\t\r"
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", strings.TrimRight(s, " \n\t\r"))
	fmt.Printf("%q\n", strings.TrimSpace(s))
}
