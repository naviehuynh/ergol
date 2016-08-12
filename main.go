package main

import (
	"github.com/naviehuynh/ergol/filters"
	"github.com/naviehuynh/ergol/printers"
	"github.com/naviehuynh/ergol/sources"
	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

func main() {
	paths := Parse().files
	logsCount := len(paths)
	hasStdin := utils.HasStdin()
	if hasStdin {
		logsCount++
	}
	if logsCount == 0 {
		return
	}

	// start to read logs
	filteredLogs := make([]types.Log, 0, logsCount)
	if hasStdin {
		filteredLogs = append(filteredLogs, filters.ApplyLogFilters(sources.StdinReader(len(filteredLogs))))
	}
	for _, path := range paths {
		filteredLogs = append(filteredLogs, filters.ApplyLogFilters(sources.FileReader(path, len(filteredLogs))))
	}

	printers.Print(filteredLogs)
}
