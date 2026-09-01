package main

import (
	"fmt"
	"task1/mathutil"
)

func main() {
	var a int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	sum := mathutil.Add(a, b)
	fmt.Println(sum)
}
