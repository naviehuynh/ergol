package filters

import "github.com/naviehuynh/ergol/types"

// ApplyFilters synchronousely read from src, apply filteres and output to dest
func ApplyFilters(src, dest chan types.Log) {
	for str := range src {
		dest <- str
	}
}
