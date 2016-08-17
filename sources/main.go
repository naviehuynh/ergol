package sources

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/naviehuynh/ergol/utils"
)

// readCmd forks a new process, returning its STDOUT, STDERR channel and pid
func readCmd(command string) (output chan string, pid int) {
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
	count := len(scanners)
	for _, scanner := range scanners {
		go func(scanner *bufio.Scanner) {
			defer closeChannel(&count, textChan)
			for scanner.Scan() {
				textChan <- scanner.Text()
			}
		}(scanner)
	}

	return textChan, cmd.Process.Pid
}
