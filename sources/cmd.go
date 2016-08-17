package sources

import (
	"strconv"

	"github.com/naviehuynh/ergol/types"
)

// CmdReader return a Log with content redirected from a process Stdout and Stderr
func CmdReader(command string, orderNo int) types.Log {
	textChan, pid := readCmd(command)
	return types.Log{
		Text:       textChan,
		SourceType: types.Cmd,
		OrderNo:    orderNo,
		SourceID:   strconv.Itoa(pid) + "|" + command,
	}
}

func closeChannel(count *int, channel chan string) {
	*count--
	if *count == 0 {
		close(channel)
	}
}
