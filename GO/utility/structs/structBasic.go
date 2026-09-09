package structs

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func brithday(s Student) {
	s.Age++
	fmt.Printf("Inside function: %d \n", s.Age)
}

func StructBasic() {
	sidharth := Student{
		Name:  "Sidharth",
		Age:   19,
		Marks: 90.7,
	}

	brithday(sidharth)
	fmt.Printf("Outside function: %v", sidharth.Age)
}
