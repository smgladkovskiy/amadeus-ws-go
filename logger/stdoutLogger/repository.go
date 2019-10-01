package stdoutLogger

import (
	"fmt"
	"time"
)

type LogWriter struct{}

func Init() *LogWriter {
	return &LogWriter{}
}

func (*LogWriter) WriteLog(direction string, action string, content string) error {
	_, err := fmt.Printf("%s_%s_%s\n%s\n", time.Now().Format(time.RFC3339), action, direction, content)
	return err
}
