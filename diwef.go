package diwef

type writer interface {
	debug(msg string)
	info(msg string)
	warning(msg string)
	error(msg string)
	fatal(msg string)
}

type logger struct {
	writers []writer
}

func Init(writers ...writer) *logger {
	return &logger{
		writers: writers,
	}
}

func (l *logger) Debug(msg string) {
	for _, w := range l.writers {
		w.debug(msg)
	}
}

func (l *logger) Info(msg string) {
	for _, w := range l.writers {
		w.info(msg)
	}
}

func (l *logger) Warning(msg string) {
	for _, w := range l.writers {
		w.warning(msg)
	}
}

func (l *logger) Error(msg string) {
	for _, w := range l.writers {
		w.error(msg)
	}
}

func (l *logger) Fatal(msg string) {
	for _, w := range l.writers {
		w.fatal(msg)
	}
}
