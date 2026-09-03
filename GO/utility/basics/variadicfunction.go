package utility

func VariadicFunction(nums ...int) int {
	total := 0
	for _, n := range nums {
		total = total + n
	}
	return total
}
