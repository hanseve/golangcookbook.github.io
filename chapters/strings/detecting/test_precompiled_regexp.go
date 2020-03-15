package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile("(?i)Wo.*") //case insensitive
	fmt.Println(re.MatchString("Hello world"), err)

	re = regexp.MustCompile("(?i)Wo.*")
	fmt.Println(re.MatchString("Hello world"))
}
