package golang

type Logger struct {
	data map[string]int
}

func LoggerConstructor() Logger {
	return Logger{data: make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	allow, ok := l.data[message]

	if !ok || timestamp >= allow {
		l.data[message] = timestamp + 10

		return true
	}

	return false
}
