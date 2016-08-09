package utils

// Check generic error checking
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
