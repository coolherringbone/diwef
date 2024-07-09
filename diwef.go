package diwef

import (
	"errors"
	"os"
)

type logger struct {
	config Config
}

func Init(config ...Config) (*logger, error) {
	logger := &logger{}

	if len(config) == 1 {
		logger.config = config[0]
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		logger.config.Path = DefaultPath
		logger.config.FileName = DefaultFileName
		logger.config.LiveTime = DefaultLiveTime
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
