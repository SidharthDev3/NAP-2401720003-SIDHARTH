package pointers_address

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float32
}

func bonus(s *Student, graceMarks float32) {
	s.Marks = s.Marks + graceMarks
}

func PointersStudent() {
	s1 := Student{
		Name:  "Sidharth",
		Age:   20,
		Marks: 85,
	}
	fmt.Printf("Marks before bonus: %v \n", s1.Marks)
	bonus(&s1, 10)
	fmt.Printf("Marks after bonus: %v \n", s1.Marks)
}
