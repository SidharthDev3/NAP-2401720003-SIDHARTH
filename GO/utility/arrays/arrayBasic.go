package arrays

import "fmt"

func ArrayBasic() {
	var marks [5]int

	for i := 0; i < len(marks); i++ {
		fmt.Printf("Enter marks of subject %d :", i+1)
		fmt.Scan(&marks[i])

	}
	for i := 0; i < len(marks); i++ {
		fmt.Printf("Mark of subject %d %d \n", i+1, marks[i])
	}
}
