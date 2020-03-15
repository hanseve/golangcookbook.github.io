package main

import (
	"fmt"
	"regexp"
)

func main()  {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("paranormal", 1))

	fmt.Println(re.FindAllStringIndex("paranormal", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", 2))
	fmt.Println(re.FindAllStringIndex("paranormal", 1))
}