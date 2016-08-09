package main

import (
	"github.com/naviehuynh/ergol/displayers"
	"github.com/naviehuynh/ergol/filters"
	"github.com/naviehuynh/ergol/sources"
	"github.com/naviehuynh/ergol/types"
)

func main() {
	paths := Parse().files
	if len(paths) == 0 {
		return
	}

	bufferStream := make(chan types.Log)
	inputStreams := make([]chan types.Log, len(paths))
	for i, path := range paths {
		inputStreams[i] = sources.FileReader(path)
	}

	for _, channel := range inputStreams {
		go filters.ApplyFilters(channel, bufferStream)
	}
	displayers.Print(bufferStream)
}
