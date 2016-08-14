package main

import "flag"

// ErgolArgs contains cli arguments passed into Ergol
type ErgolArgs struct {
	files             []string
	grepPattern       string
	grepAfter         int
	grepBefore        int
	grepKeepUnmatched bool
}

// ParseArgs cli argumens into structured format
func ParseArgs() ErgolArgs {
	args := ErgolArgs{}
	flag.StringVar(&args.grepPattern, "g", "", "grep pattern")
	flag.IntVar(&args.grepAfter, "A", 0, "keep X lines before a match")
	flag.IntVar(&args.grepBefore, "B", 0, "keep X lines after a match")
	flag.BoolVar(&args.grepKeepUnmatched, "K", false, "keep lines that doesn't match pattern")
	grepC := flag.Int("C", 0, "keep X lines before and after a match")
	flag.Parse()
	if *grepC != 0 {
		args.grepAfter = *grepC
		args.grepBefore = *grepC
	}
	args.files = flag.Args()
	return args
}
