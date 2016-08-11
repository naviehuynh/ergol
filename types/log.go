package types

import "strings"

// Log contains log meta data and log text
type Log struct {
	OrderNo    int
	SourceID   string
	SourceType SourceType
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

// ShortName returns a readble name that often identifies a Log object
func (l *Log) ShortName() string {
	switch l.SourceType {
	case File:
		chunks := strings.Split(l.SourceID, "/")
		return chunks[len(chunks)-1]
	default:
		return l.SourceID
	}
}

// Clone return a shallow copy of log
func (l *Log) Clone() Log {
	return Log{
		OrderNo:    l.OrderNo,
		SourceID:   l.SourceID,
		SourceType: l.SourceType,
		Text:       l.Text,
	}
}

// Close gracefully closes a Log
func (l *Log) Close() bool {
	switch l.SourceType {
	case File:
		close(l.Text)
	default:
	}
	return true
}
