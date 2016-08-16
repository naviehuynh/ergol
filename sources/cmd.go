package sources

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"

	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// CmdReader return a Log with content redirected from a process Stdout and Stderr
func CmdReader(command string, orderNo int) types.Log {
	// TODO: better shell command parsing here
	cmdChunks := strings.Split(command, " ")
	cmd := exec.Command(cmdChunks[0], cmdChunks[1:]...)
	stderr, err := cmd.StdoutPipe()
	utils.Check(err)
	stdout, err := cmd.StderrPipe()
	utils.Check(err)
	err = cmd.Start()
	utils.Check(err)

	textChan := make(chan string)
	scanners := []*bufio.Scanner{bufio.NewScanner(stdout), bufio.NewScanner(stderr)}
	count := 2
	for _, scanner := range scanners {
		go func(scanner *bufio.Scanner) {
			defer closeChannel(&count, textChan)
			for scanner.Scan() {
				textChan <- scanner.Text()
			}
		}(scanner)
	}

	return types.Log{Text: textChan, SourceType: types.Cmd, OrderNo: orderNo, SourceID: strconv.Itoa(cmd.Process.Pid) + "|" + command}
}

func closeChannel(count *int, channel chan string) {
	*count--
	if *count == 0 {
		close(channel)
	}
}
