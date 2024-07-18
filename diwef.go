package diwef

type Logger struct {
	writers []writer
}

func Init(writers ...writer) *Logger {
	return &Logger{
		writers: writers,
	}
}

func (l *Logger) Debug(msg any) {
	for _, w := range l.writers {
		w.debug(msg)
	}
}

func (l *Logger) Info(msg any) {
	for _, w := range l.writers {
		w.info(msg)
	}
}

func (l *Logger) Warning(msg any) {
	for _, w := range l.writers {
		w.warning(msg)
	}
}

func (l *Logger) Error(msg any) {
	for _, w := range l.writers {
		w.error(msg)
	}
}

func (l *Logger) Fatal(msg any) {
	for _, w := range l.writers {
		w.fatal(msg)
	}
}
