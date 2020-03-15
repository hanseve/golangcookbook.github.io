package main

import (
	"bytes"
	"fmt"
)

func main() {
	var str bytes.Buffer
	str.WriteString("Hello")
	str.WriteString(" world")
	fmt.Println(str.String())
}
