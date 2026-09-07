package arrays

import "fmt"

func DynamicLists() {
	s := []int{80, 87, 92}

	sum := 0

	for _, v := range s {
		sum = sum + v
	}

	avg := sum / 3

	fmt.Println(avg)

}
