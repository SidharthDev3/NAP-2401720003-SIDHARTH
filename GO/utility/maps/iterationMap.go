package maps

import "fmt"

func IterationMap() {
	ages := map[string]int{}
	ages["George"] = 20
	ages["Sidharth"] = 19
	ages["Hemraj"] = 23
	for k, v := range ages {
		fmt.Println(k, v)
	}
}
