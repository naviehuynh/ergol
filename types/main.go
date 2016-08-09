package types

// Log has a log line with its meta data
type Log struct {
	text   string
	source string
}

//SetText set text of log
func (l *Log) SetText(text string) string {
	l.text = text
	return text
}

//Text return log text
func (l *Log) Text() string {
	return l.text
}

//SetSource  set source of log
func (l *Log) SetSource(source string) string {
	l.source = source
	return source
}

//Source return log source
func (l *Log) Source() string {
	return l.source
}
