package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	chars := []string{"\\]", "\\^", "\\-", "\\\\", "\\[", "\\.", "\\(", "\\)"}
	r := strings.Join(chars, "")
	s := "[Some]-\\(string)^."
	re := regexp.MustCompile("[" + r + "]+")
	s = re.ReplaceAllString(s, "")
	fmt.Println(s)
}
