package utils

// Check generic error checking
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// MaxInt return the max of Ints
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
