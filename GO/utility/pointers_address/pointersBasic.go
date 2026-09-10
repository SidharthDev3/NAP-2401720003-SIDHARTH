package pointers_address

import "fmt"

func PointersBasic() {
	x := 10
	var ptr *int = &x

	fmt.Println("Value of x:       ", x)
	fmt.Println("Address of x (&x): ", &x)
	fmt.Println("Value in ptr:      ", ptr)
	fmt.Println("Dereferenced (*ptr):", *ptr)
	*ptr = 25
	fmt.Println("New value of x:   ", x)
}
