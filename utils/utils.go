package utils

import (
	"os"

	"github.com/fatih/color"
)

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

// Colored colors the string, the color is determined by number
func Colored(text string, number int) string {
	return color.New(color.Attribute(number%6+31), color.Bold).SprintFunc()(text)
}

// HasStdin return true if there are some data pumping into STDIN
func HasStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
