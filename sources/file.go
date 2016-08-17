package sources

import (
	"github.com/hpcloud/tail"
	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// FileReader reads and follow files
func FileReader(path string, orderNo int) (Log types.Log) {
	t, err := tail.TailFile(path, tail.Config{
		Follow: true,
		ReOpen: true,
	})
	utils.Check(err)
	channel := make(chan string)
	go func() {
		for line := range t.Lines {
			channel <- line.Text
		}
	}()

	return types.Log{Text: channel, SourceType: types.File, OrderNo: orderNo, SourceID: path}
}
