package diwef

type writer interface {
	debug(msg string)
	info(msg string)
	warning(msg string)
	error(msg string)
	fatal(msg string)
}

type logger struct {
	writer writer
}

func Init(writer writer) *logger {
	return &logger{
		writer: writer,
	}
}

func (l *logger) Debug(msg string) {
	l.writer.debug(msg)
}

func (l *logger) Info(msg string) {
	l.writer.info(msg)
}

func (l *logger) Warning(msg string) {
	l.writer.warning(msg)
}

func (l *logger) Error(msg string) {
	l.writer.error(msg)
}

func (l *logger) Fatal(msg string) {
	l.writer.fatal(msg)
}
