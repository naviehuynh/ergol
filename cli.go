package main

import (
	"flag"
	"strings"
)

// StringArgs ...
type StringArgs []string

// String is used by flag parser to pass multiple values arg
func (args *StringArgs) String() string {
	return strings.Join(*args, ", ")
}

// Set is used by flag parser to pass multiple values arg
func (args *StringArgs) Set(commands string) error {
	*args = append(*args, strings.Split(commands, ",")...)
	return nil
}

// ErgolArgs contains cli arguments passed into Ergol
type ErgolArgs struct {
	files                  []string
	highlightPattern       string
	highlightCaseSensitive bool
	lineCount              int
	cmds                   StringArgs
}

// ParseArgs cli argumens into structured format
func ParseArgs() ErgolArgs {
	args := ErgolArgs{}
	flag.StringVar(&args.highlightPattern, "p", "", "highlight text matching this pattern")
	flag.BoolVar(&args.highlightCaseSensitive, "s", false, "highlight string with casing")
	flag.Var(&args.cmds, "e", "Commands to be executed and log STDOUT")
	flag.IntVar(&args.lineCount, "l", -1, "number of tailing lines to read, default to -1 (all)")

	flag.Parse()

	args.files = flag.Args()
	return args
}
