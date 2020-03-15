package main

import (
	"fmt"
	"strings"
)

func main() {
	str3 := []string{"Hello"}
	str3 = append(str3, " world")
	fmt.Println(strings.Join(str3, ""))
}
