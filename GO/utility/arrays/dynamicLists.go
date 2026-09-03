package arrays

import "fmt"

func DynamicLists() {
	s := []int{10, 20}
	s = append(s, 30)

	fmt.Println(s)

	fmt.Println(len(s))

	Nums2 := make([]int, 5, 10)

	fmt.Println(len(Nums2))
}
