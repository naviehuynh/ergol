package main

import "os"

// ErgolArgs contains cli arguments passed into Ergol
type ErgolArgs struct {
	files []string
}

// Parse cli argumens into structured format
func Parse() ErgolArgs {
	rawArgs := os.Args[1:]
	return ErgolArgs{files: rawArgs}
}
