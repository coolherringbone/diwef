package diwef

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func setConfig(logger *logger, config []Config) error {
	if len(config) == 1 {
		logger.config.Path = nvl(config[0].Path, DefaultPath).(string)
		logger.config.FileName = nvl(config[0].FileName, DefaultFileName).(string)
		logger.config.LiveTime = nvl(config[0].LiveTime, DefaultLiveTime).(int)
	} else if len(config) > 1 {
		return errors.New("there can be only one config (or even empty)")
	} else {
		logger.config.Path = DefaultPath
		logger.config.FileName = DefaultFileName
		logger.config.LiveTime = DefaultLiveTime
	}

	return nil
}

func (l *logger) writer(level, msg string) {
	logStr := l.stylingLogStr(level, msg)
	file := l.openLogFile()
	defer file.Close()

	file.WriteString(logStr)
}

func (l *logger) openLogFile() *os.File {
	fullName := fmt.Sprintf("%s/%s-%s.log",
		l.config.Path,
		l.config.FileName,
		time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}

	return file
}

func (l *logger) stylingLogStr(level, msg string) string {
	logStr := fmt.Sprintf("time=\"%s\"		level=\"%s\"		msg=\"%s\"\n",
		time.Now().Format("15:04:05"),
		level,
		msg)

	return logStr
}

func nvl(a, b any) any {
	if a == nil || a == "" || a == 0 {
		return b
	} else {
		return a
	}
}
