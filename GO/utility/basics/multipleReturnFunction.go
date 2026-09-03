package utility

func MultipleReturnFunction(a, b, c int) (int, string) {
	var avg int = (a + b + c) / 3

	if avg < 40 {
		return avg, "Failed"
	}
	return avg, "Pass!!"
}
