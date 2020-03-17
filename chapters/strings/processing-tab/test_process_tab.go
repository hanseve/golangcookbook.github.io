package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("%q\n",strings.ReplaceAll("oink\toink\toink", "\t", ""))
	fmt.Printf("%q\n",strings.ReplaceAll("foo foo foo", " ", "\t"))

	re := regexp.MustCompile("\t")
	fmt.Printf("%q\n",re.ReplaceAllString("oink\toink\toink", ""))
	re = regexp.MustCompile(" ")
	fmt.Printf("%q\n",re.ReplaceAllString("foo foo foo", "\t"))
}
