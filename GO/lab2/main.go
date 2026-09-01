package main

import (
	"fmt"
	"task2/hello"
)

func main() {
	var input string
	fmt.Print("Enter a word: ")
	fmt.Scan(&input)

	ans := hello.CountVowels(input)
	fmt.Println(ans)
}
