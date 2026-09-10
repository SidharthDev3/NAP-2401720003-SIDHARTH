package pointersaddress

import "fmt"

func PointersBasic() {
	var a int = 10
	p := &a
	*p = 20

	fmt.Println(*p)
	fmt.Println(p)
}
