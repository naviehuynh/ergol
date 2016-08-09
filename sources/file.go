package sources

import (
	"github.com/hpcloud/tail"
	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// FileReader reads and follow files
func FileReader(path string) (stream chan types.Log) {
	t, err := tail.TailFile(path, tail.Config{Follow: true})
	utils.Check(err)
	channel := make(chan types.Log)
	go func() {
		for line := range t.Lines {
			log := types.Log{}
			log.SetText(line.Text)
			log.SetSource(path)
			channel <- log
		}
	}()

	return channel
}
