package utility

import "fmt"

func IfElseCondition() {
	var score int
	fmt.Print("Enter your score: ")
	fmt.Scan(&score)

	if score >= 60 {
		fmt.Println("PASS")
	} else if score == 0 {
		fmt.Println("GET OUT")
	} else {
		fmt.Println("FAIL")
	}
}
