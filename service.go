package diwef

import (
	"fmt"
	"os"
	"time"
)

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
