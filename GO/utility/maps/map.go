package maps

import "fmt"

func Map() {
	ages := map[string]int{}
	ages["George"] = 20
	ages["Sidharth"] = 19
	fmt.Println(ages["George"])
	fmt.Println(ages["Sidharth"])

	fmt.Println(ages)
	delete(ages, "Sidharth")

	fmt.Println(ages)

}
