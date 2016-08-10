package types

// Log contains log meta data and log text
type Log struct {
	OrderNo    int
	Path       string
	SourceType int
	Text       chan string
}

// sourceType can only be FILE or SSH
const (
	SourceTypeFile int = iota
	SourceTypeCMD
	SourceTypeSTDIN
	SourceTypeSSH
)
