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

	filteredLogs := make([]types.Log, len(paths))
	for i, path := range paths {
		filteredLogs[i] = filters.ApplyLogFilters(sources.FileReader(path, i))
	}

	displayers.Print(filteredLogs)
}
