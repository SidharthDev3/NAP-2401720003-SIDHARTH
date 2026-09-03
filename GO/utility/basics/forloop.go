package utility

import "fmt"

func ForLoop() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
