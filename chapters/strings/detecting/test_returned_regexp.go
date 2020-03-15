package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindString("paranormal"))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))

	re = regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
