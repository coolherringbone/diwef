package diwef

type logger struct {
	writers []writer
}

func Init(writers ...writer) *logger {
	return &logger{
		writers: writers,
	}
}

func (l *logger) Debug(msg any) {
	for _, w := range l.writers {
		w.debug(msg)
	}
}

func (l *logger) Info(msg any) {
	for _, w := range l.writers {
		w.info(msg)
	}
}

func (l *logger) Warning(msg any) {
	for _, w := range l.writers {
		w.warning(msg)
	}
}

func (l *logger) Error(msg any) {
	for _, w := range l.writers {
		w.error(msg)
	}
}

func (l *logger) Fatal(msg any) {
	for _, w := range l.writers {
		w.fatal(msg)
	}
}
