package arrayslices

// Sum all integers of a given slice
func Sum(slice []int) (sum int) {
	for _, number := range slice {
		sum += number
	}
	return
}

func SumAll(s ...[]int) (sum []int) {
	for _, numbers := range s {
		sum = append(sum, Sum(numbers))
	}
	return
}
