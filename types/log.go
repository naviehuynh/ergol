package types

// Log contains log meta data and log text
type Log struct {
	OrderNo    int
	Path       string
	SourceType int
	Text       chan string
}

// SourceType represents different input sources
type SourceType int

// input source types
const (
	File SourceType = iota
	CMD
	STDIN
	SSH
)
