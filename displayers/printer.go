package displayers

import (
	"fmt"

	"github.com/naviehuynh/ergol/types"
)

// Print prints content to STDOUT
func Print(buffer chan types.Log) {
	for log := range buffer {
		fmt.Printf("%s: %s\n", log.Source(), log.Text())
	}
}
