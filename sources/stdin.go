package sources

import (
	"bufio"
	"os"

	"github.com/naviehuynh/ergol/types"
)

// StdinReader return a Log from STDIN
func StdinReader(orderNo int) types.Log {
	scanner := bufio.NewScanner(os.Stdin)
	textChan := make(chan string)

	go func() {
		for scanner.Scan() {
			textChan <- scanner.Text()
		}
		close(textChan)
	}()

	return types.Log{Text: textChan, SourceType: types.Stdin, OrderNo: orderNo, SourceID: "STDIN"}
}
