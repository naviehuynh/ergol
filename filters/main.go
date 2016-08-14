package filters

import "github.com/naviehuynh/ergol/types"

// Filter is the interface for all ergol's text filters
type Filter interface {
	Apply(types.Log) types.Log
}

// ApplyLogFilters synchronousely read from src, apply filteres and output to dest
func ApplyLogFilters(allFilters []Filter, src types.Log) types.Log {
	newLog := src
	for _, filter := range allFilters {
		newLog = filter.Apply(newLog)
	}
	return newLog
}

// ApplyCombinedFilters synchronousely read from src, apply filteres and output to dest
func ApplyCombinedFilters(src types.Log) types.Log {
	return src
}
