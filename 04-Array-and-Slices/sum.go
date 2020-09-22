package arrayslices

// Sum all integers of a given slice
func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum += number
	}
	return
}
