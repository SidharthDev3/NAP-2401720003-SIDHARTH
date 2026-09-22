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

func QuestionPointers() {
	x := 10
	var ptr *int = &x

	fmt.Printf("Value of x: %v \n", x)             //Prints Value
	fmt.Printf("Address of x (&x): %v \n", &x)     //Prints Address
	fmt.Printf("Value in ptr: %v \n", ptr)         //Prints Address
	fmt.Printf("Dereferenced (*ptr): %v \n", *ptr) //Prints Value

	num := 10
	fmt.Printf("Before Modifying: %v \n", num)
	PointersModify(&num)
	fmt.Printf("After Modifying: %v \n", num)

	s1 := new(Student)
	fmt.Printf("Struct before Modifying: %v \n", *s1)
	s1.Name = "Sidharth"
	s1.Age = 20
	s1.Marks = 90
	fmt.Printf("Struct after Modifying: %v \n", *s1)

}
