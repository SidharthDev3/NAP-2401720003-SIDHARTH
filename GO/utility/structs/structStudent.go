package structs

import "fmt"

func StructStudent() {
	s1 := Student{
		Name:  "Sidharth",
		Age:   19,
		Marks: 90,
	}
	s2 := Student{
		Name:  "John",
		Age:   19,
		Marks: 89,
	}
	s3 := Student{
		Name:  "David",
		Age:   19,
		Marks: 70,
	}
	s4 := Student{
		Name:  "George",
		Age:   19,
		Marks: 89,
	}
	fmt.Printf("Name: %v \n", s1.Name)
	fmt.Printf("Age: %v \n", s1.Age)
	fmt.Printf("Marks: %v \n", s1.Marks)

	fmt.Printf("Name: %v \n", s2.Name)
	fmt.Printf("Age: %v \n", s2.Age)
	fmt.Printf("Marks: %v \n", s2.Marks)

	fmt.Printf("Name: %v \n", s3.Name)
	fmt.Printf("Age: %v \n", s3.Age)
	fmt.Printf("Marks: %v \n", s3.Marks)

	fmt.Printf("Name: %v \n", s4.Name)
	fmt.Printf("Age: %v \n", s4.Age)
	fmt.Printf("Marks: %v \n", s4.Marks)
}
