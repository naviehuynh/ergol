package filters

import (
	"regexp"
	"strings"

	"github.com/naviehuynh/ergol/types"
	"github.com/naviehuynh/ergol/utils"
)

// Highlighter filter
type Highlighter struct {
	CaseSensitive bool
	KeepUnmatched bool
	Pattern       string

	// TODO: Implement these options
	// IsRegexp      bool
}

//Apply text filtering on a log.Text
func (f Highlighter) Apply(log types.Log) types.Log {
	newText := make(chan string)
	go func() {
		defer close(newText)
		for line := range log.Text {
			highlightedLine, _ := highlightMatches(line, f.Pattern, f.CaseSensitive)
			newText <- highlightedLine
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
	count := 0
	// TODO: find a better way to count
	highlightedText := reg.ReplaceAllStringFunc(text, func(oriText string) string {
		count++
		return utils.Colored(oriText, 0)
	})
	return highlightedText, count
}
