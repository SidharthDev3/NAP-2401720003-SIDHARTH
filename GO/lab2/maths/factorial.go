package maths

func Factorial(a int) int {
	if a == 0 {
		return 1
	}
	if a == 1 {
		return 1
	}
	return a * Factorial(a-1)
}
