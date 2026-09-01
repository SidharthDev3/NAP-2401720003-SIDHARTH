package maths

func Power(a, b int) int {
	if a == 0 {
		return 1
	}
	if b == 0 {
		return 1
	}
	result := 1

	for i := 1; i <= b; i++ {
		result = result * a
	}

	return result
}
