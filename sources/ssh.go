package sources

import (
	"fmt"
	"strings"

	"github.com/naviehuynh/ergol/types"
)

// SSHReader reads files over ssh
func SSHReader(path string, orderNo, numLines int) (Log types.Log) {
	args := []string{"-F"}
	if numLines == -1 {
		// read whole file
		args = append(args, "-n +1")
	} else {
		args = append(args, fmt.Sprintf("-n %d", numLines))
	}

	tailCommand := fmt.Sprintf("tail %s %s", path, strings.Join(args, " "))
	textChan, _ := readCmd(tailCommand)

	return types.Log{
		Text:       textChan,
		SourceType: types.File,
		OrderNo:    orderNo,
		SourceID:   path,
	}
}
