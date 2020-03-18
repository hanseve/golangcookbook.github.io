package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please input a string:")
	var input, err = reader.ReadString('\n');
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Your input: %s\n", input)

	fmt.Println("Please input another string:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan()  {
		fmt.Printf("Your input: %s\n", scanner.Text())
	}
}
