package pointers_address

import "fmt"

func PointersBasic() {
	x := 10
	var ptr *int = &x

	fmt.Println("Value of x:       ", x)      //Prints Value
	fmt.Println("Address of x (&x): ", &x)    //Prints Address
	fmt.Println("Value in ptr:      ", ptr)   //Prints Address
	fmt.Println("Dereferenced (*ptr):", *ptr) //Prints Value
	*ptr = 25
	fmt.Println("New value of x:   ", x)
}
