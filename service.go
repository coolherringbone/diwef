package diwef

import (
	"fmt"
	"os"
	"time"
)

func (a *App) writer(level, msg string) {
	file := a.openLogFile()
	defer file.Close()
}

func (a *App) openLogFile() *os.File {
	fullName := fmt.Sprintf("%s/%s-%s.log",
		a.config.Path,
		a.config.FileName,
		time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}

	return file
}
