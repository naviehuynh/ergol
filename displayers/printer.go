package displayers

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/naviehuynh/ergol/types"
)

// Print prints content to STDOUT
func Print(logs []types.Log) {
	buffers := make(chan string)
	for _, log := range logs {
		go func(log types.Log) {
			for str := range log.Text {
				buffers <- fmt.Sprintf("%s:  %s", colored(log.Path, log.OrderNo), str)
			}
		}(log)
	}
	for txt := range buffers {
		fmt.Printf("%s\n", txt)
	}
}

func colored(text string, number int) string {
	return color.New(color.Attribute(number%6+31), color.Bold).SprintFunc()(text)
}
