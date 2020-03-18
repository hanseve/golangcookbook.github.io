package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please input your firstname and lastname:")
	var firstname, lastname string
	fmt.Scanln(&firstname, &lastname)
	fmt.Println("Please input your age:")
	var age int
	fmt.Scanf("%3d", &age)
	fmt.Printf("Hi, I'm %s·%s, %d years old\n", firstname, lastname, age)
}
