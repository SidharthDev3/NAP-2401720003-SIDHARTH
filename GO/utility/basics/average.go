package utility

import "fmt"

func Average() {
	var a int = 35
	var b int = 29
	var c int = 28
	var total int = (a + b + c)
	if avg := total / 3; avg >= 40 {
		fmt.Println("Pass ", avg)
	} else {
		fmt.Println("Fail", avg)
	}

}
