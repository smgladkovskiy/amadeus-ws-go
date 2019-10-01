package logger

type Service interface {
	Push(direction string, action string, content string) error
}

type LogWriter interface {
	WriteLog(direction string, action string, content string) error
}

func NewLogger(ls LogWriter) Service {
	return &logger{writer: ls}
}

// Internal type for storing the hooks on a r instance.
type logger struct {
	writer LogWriter
}

func (l *logger) Push(direction string, action string, content string) error {
	return l.writer.WriteLog(direction, action, content)
}
