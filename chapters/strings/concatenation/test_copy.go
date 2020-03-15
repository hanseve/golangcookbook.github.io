package main

import "fmt"

func main() {
	bs := make([]byte, 20)
	n := copy(bs, "Hello")
	copy(bs[n:], " world")
	fmt.Println(string(bs))
}
