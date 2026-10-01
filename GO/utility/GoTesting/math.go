package gotesting

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Math() {
	fmt.Println(Add(2, 3))
}
