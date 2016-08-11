package filters

import (
	"regexp"
	"strings"

	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// Grep filter
type Grep struct {
	CaseSensitive bool
	keepUnmatched bool
	Pattern       string
	// TODO: Implement these options
	// IsRegexp      bool
	// KeepBefore int
	// KeepAfter int
	// Inverse int
}

//Apply text filtering on a log.Text
func (f Grep) Apply(log types.Log) types.Log {
	newText := make(chan string)
	go func() {
		for line := range log.Text {
			highlightedLine, count := highlightMatches(line, f.Pattern, f.CaseSensitive)
			if count > 0 || f.keepUnmatched {
				newText <- highlightedLine
			}
		}
	}()
	newLog := log.Clone()
	newLog.Text = newText
	return newLog
}

// highlightMatches return a string with highlighted matches and matches count
func highlightMatches(text, pattern string, CaseSensitive bool) (string, int) {
	regexpString := strings.Replace(pattern, " ", "[ \\._-]", -1)
	if !CaseSensitive {
		regexpString = "(?i)" + regexpString
	}
	reg, err := regexp.Compile(regexpString)
	utils.Check(err)
	coloredPattern := utils.Colored(pattern, 0)
	count := 0
	// TODO: find a better way to count
	highlightedText := reg.ReplaceAllStringFunc(text, func(_ string) string {
		count++
		return coloredPattern
	})
	return highlightedText, count
}
