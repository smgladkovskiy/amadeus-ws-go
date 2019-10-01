package nilLogger

type LogWriter struct{}

func Init() *LogWriter {
	return &LogWriter{}
}

func (*LogWriter) WriteLog(direction string, action string, content string) error {
	return nil
}
