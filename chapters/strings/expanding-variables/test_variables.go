package main

import (
	"flag"
	"fmt"
)

func main() {
	var num = flag.Int("num", 0, "test flag int")
	var num2 int
	flag.IntVar(&num2, "num2", 0, "test flag int")
	flag.Parse()
	fmt.Println(*num)
	fmt.Println(num2)
}
