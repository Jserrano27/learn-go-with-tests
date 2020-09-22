package iteration

// Repeat provided string X times (negative ints are converted to positive)
func Repeat(s string, times int) (r string) {
	if times < 0 {
		times *= -1
	}

	for i := 0; i < times; i++ {
		r += s
	}
	return
}
