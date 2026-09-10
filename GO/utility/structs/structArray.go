package structs

import "fmt"

func StructArray() {
	class := []Student{
		{"Sidharth", 19, 90},
		{"David", 19, 89},
	}
	for _, s := range class {
		fmt.Printf("Name: %v \n Grade %v \n", s.Name, s.Marks)
	}
}
