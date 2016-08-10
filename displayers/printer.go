package displayers

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// Print prints content to STDOUT
func Print(logs []types.Log) {
	buffers := make(chan string)

	maxPathLength := -1
	for _, log := range logs {
		maxPathLength = utils.MaxInt(maxPathLength, len(log.SourceID))
	}

	for _, log := range logs {
		go func(log types.Log) {
			prefix := colored(padded(log.SourceID, maxPathLength)+" |   ", log.OrderNo)
			for str := range log.Text {
				buffers <- fmt.Sprintf("%s%s", prefix, str)
			}
		}(log)
	}
	for txt := range buffers {
		fmt.Printf("%s\n", txt)
	}
}

func padded(text string, length int) string {
	buf := new(bytes.Buffer)
	buf.WriteString(text)

	for buf.Len() < length {
		buf.WriteString(" ")
	}
	return buf.String()
}
func colored(text string, number int) string {
	return color.New(color.Attribute(number%6+31), color.Bold).SprintFunc()(text)
}
