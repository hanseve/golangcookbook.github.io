package main

import (
	"strings"
	"fmt"
)

func main() {
	var str2 strings.Builder
	str2.WriteString("Hello ")
	str2.WriteString("world")
	fmt.Println(str2.String())
}
