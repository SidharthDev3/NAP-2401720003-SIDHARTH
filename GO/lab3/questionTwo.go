package lab3

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float32
}

func PointersModify(ptr *int) {
	*ptr = 100
}

func bonus(s *Student, graceMarks float32) {
	s.Marks = s.Marks + graceMarks
}

func QuestionPointers() {
	x := 10
	var ptr *int = &x

	fmt.Println("Value of x:       ", x)      //Prints Value
	fmt.Println("Address of x (&x): ", &x)    //Prints Address
	fmt.Println("Value in ptr:      ", ptr)   //Prints Address
	fmt.Println("Dereferenced (*ptr):", *ptr) //Prints Value

	num := 10
	fmt.Println("Before Modifying: ", num)
	PointersModify(&num)
	fmt.Println("After Modifying: ", num)

	s1 := new(Student)
	fmt.Println("Struct before Modifying: ", *s1)
	s1.Name = "Sidharth"
	fmt.Println("Struct after Modifying: ", *s1)

}
