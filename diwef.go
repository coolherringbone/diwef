package diwef

import (
	"os"
)

type logger struct {
	config Config
}

func Init(config ...Config) (*logger, error) {
	logger := &logger{}

	if err := setConfig(logger, config); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(logger.config.Path, 0744); err != nil {
		return nil, err
	}

	return logger, nil
}

func (l *logger) Debug(msg string) {
	l.writer("debug", msg)
}

func (l *logger) Info(msg string) {
	l.writer("info", msg)
}

func (l *logger) Warning(msg string) {
	l.writer("warning", msg)
}

func (l *logger) Error(msg string) {
	l.writer("error", msg)
}

func (l *logger) Fatal(msg string) {
	l.writer("fatal", msg)
}
