package diwef

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type fileWriter struct {
	config FileWriter
}

type FileWriter struct {
	Path      string
	FileName  string
	UseLevels []Level
	LiveTime  int
}

const (
	DefaultPath     string = "log"
	DefaultFileName string = "app"
	DefaultLiveTime int    = 0
)

func NewFileWriter(config ...FileWriter) (writer, error) {
	var w writer
	var f = &fileWriter{}

	if len(config) == 1 {
		f.config.Path = nvl(config[0].Path, DefaultPath).(string)
		f.config.FileName = nvl(config[0].FileName, DefaultFileName).(string)
		f.config.UseLevels = nvl(config[0].UseLevels, DefaultUseLevels).([]Level)
		f.config.LiveTime = nvl(config[0].LiveTime, DefaultLiveTime).(int)
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		f.config.Path = DefaultPath
		f.config.FileName = DefaultFileName
		f.config.UseLevels = DefaultUseLevels
		f.config.LiveTime = DefaultLiveTime
	}

	if err := os.MkdirAll(f.config.Path, 0744); err != nil {
		return nil, err
	}

	w = f

	return w, nil
}

func (f *fileWriter) debug(msg string) {
	logStr := stylingLogStr("debug", msg)
	file := f.openLogFile()
	defer file.Close()

	file.WriteString(logStr)

	f.clearingLogs()
}

func (f *fileWriter) info(msg string) {
	logStr := stylingLogStr("info", msg)
	file := f.openLogFile()
	defer file.Close()

	file.WriteString(logStr)
	f.clearingLogs()
}

func (f *fileWriter) warning(msg string) {
	logStr := stylingLogStr("warning", msg)
	file := f.openLogFile()
	defer file.Close()

	file.WriteString(logStr)
	f.clearingLogs()
}

func (f *fileWriter) error(msg string) {
	logStr := stylingLogStr("error", msg)
	file := f.openLogFile()
	defer file.Close()

	file.WriteString(logStr)
	f.clearingLogs()
}

func (f *fileWriter) fatal(msg string) {
	logStr := stylingLogStr("fatal", msg)
	file := f.openLogFile()
	defer file.Close()

	file.WriteString(logStr)
	f.clearingLogs()
}

func (f *fileWriter) openLogFile() *os.File {
	fullName := fmt.Sprintf("%s/%s-%s.log",
		f.config.Path,
		f.config.FileName,
		time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}

	return file
}

func (f *fileWriter) clearingLogs() {
	if f.config.LiveTime == 0 {
		return
	}

	timeCutoff := time.Now().Add(-24 * time.Duration(f.config.LiveTime) * time.Hour)

	path, err := os.Open(f.config.Path)
	if err != nil {
		panic(err)
	}
	defer path.Close()

	pathFileList, err := path.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, file := range pathFileList {
		if timeCutoff.After(file.ModTime()) {
			fullName := fmt.Sprintf("%s/%s", f.config.Path, file.Name())
			os.Remove(fullName)
		}
	}
}
