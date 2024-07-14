package diwef

type writer interface {
	SetLevel(level ...level)
	debug(msg any)
	info(msg any)
	warning(msg any)
	error(msg any)
	fatal(msg any)
}

type level struct {
	activ bool
	name  string
}

type logger struct {
	writers []writer
}

var (
	DebugLevel = level{
		activ: true,
		name:  "debug",
	}
	InfoLevel = level{
		activ: true,
		name:  "info",
	}
	WarningLevel = level{
		activ: true,
		name:  "warning",
	}
	ErrorLevel = level{
		activ: true,
		name:  "error",
	}
	FatalLevel = level{
		activ: true,
		name:  "fatal",
	}
)

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
