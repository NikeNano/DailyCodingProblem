package ProductAllExceptCurrentIndex

func productExcept(input []int) []int {
	var output []int
	var product int = 1
	for _, value := range input {
		product *= value
	}
	for _, value := range input {
		output = append(output, product/value)
	}
	return output
}
