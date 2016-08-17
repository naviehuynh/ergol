package main

import (
	"github.com/naviehuynh/ergol/filters"
	"github.com/naviehuynh/ergol/printers"
	"github.com/naviehuynh/ergol/sources"
	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

func main() {
	// reads args
	args := ParseArgs()
	paths := args.files
	cmds := args.cmds
	logsCount := len(paths) + len(cmds)
	hasStdin := utils.HasStdin()

	if hasStdin {
		logsCount++
	}
	if logsCount == 0 {
		return
	}

	// make filter instances
	filterInstances := []filters.Filter{}
	if len(args.grepPattern) > 0 {
		grepFilter := filters.Grep{
			Pattern:       args.grepPattern,
			KeepUnmatched: args.grepKeepUnmatched,
			CaseSensitive: args.grepCaseSensitive,
		}
		filterInstances = append(filterInstances, grepFilter)
	}

	// start to read logs
	logs := make([]types.Log, 0, logsCount)
	if hasStdin {
		stdinLog := sources.StdinReader(len(logs))
		logs = append(logs, stdinLog)
	}
	for _, path := range paths {
		fileLog := sources.FileReader(path, len(logs), args.grepLineCount)
		logs = append(logs, fileLog)
	}
	for _, cmd := range cmds {
		logs = append(logs, sources.CmdReader(cmd, len(logs)))
	}

	// apply filters
	filteredLogs := make([]types.Log, 0, logsCount)
	for _, log := range logs {
		filteredLogs = append(filteredLogs, filters.ApplyLogFilters(filterInstances, log))
	}

	// finally print filtered logs to screen
	printers.Print(filteredLogs)
}
