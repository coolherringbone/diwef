package diwef

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

type fileWriter struct {
	config       FileWriter
	debugLevel   level
	infoLevel    level
	warningLevel level
	errorLevel   level
	fatalLevel   level
}

type FileWriter struct {
	Path     string
	FileName string
	LiveTime int
}

const (
	defaultPath     string = "log"
	defaultFileName string = "app"
	defaultLiveTime int    = 0
)

func NewFileWriter(config ...FileWriter) (writer, error) {
	var w writer
	var f = &fileWriter{}

	if len(config) == 1 {
		f.config.Path = nvl(config[0].Path, defaultPath).(string)
		f.config.FileName = nvl(config[0].FileName, defaultFileName).(string)
		f.config.LiveTime = nvl(config[0].LiveTime, defaultLiveTime).(int)
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		f.config.Path = defaultPath
		f.config.FileName = defaultFileName
		f.config.LiveTime = defaultLiveTime
	}

	f.SetLevel(DebugLevel, InfoLevel, WarningLevel, ErrorLevel, FatalLevel)

	if err := os.MkdirAll(f.config.Path, 0744); err != nil {
		return nil, err
	}

	w = f

	return w, nil
}

func (f *fileWriter) SetLevel(level ...level) {
	f.debugLevel.activ = false
	f.infoLevel.activ = false
	f.warningLevel.activ = false
	f.errorLevel.activ = false
	f.fatalLevel.activ = false

	for _, l := range level {
		switch l {
		case DebugLevel:
			f.debugLevel = l
		case InfoLevel:
			f.infoLevel = l
		case WarningLevel:
			f.warningLevel = l
		case ErrorLevel:
			f.errorLevel = l
		case FatalLevel:
			f.fatalLevel = l
		}

	}
}

func (f *fileWriter) debug(msg any) {
	f.writing(f.debugLevel, msg)
}

func (f *fileWriter) info(msg any) {
	f.writing(f.infoLevel, msg)
}

func (f *fileWriter) warning(msg any) {
	f.writing(f.warningLevel, msg)
}

func (f *fileWriter) error(msg any) {
	f.writing(f.errorLevel, msg)
}

func (f *fileWriter) fatal(msg any) {
	f.writing(f.fatalLevel, msg)
}

func (f *fileWriter) writing(level level, msg any) {
	if !level.activ {
		return
	}

	logStr := stylingLogStr(level.name, msg)
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

	dateCutoff := time.Now().Add(-24 * time.Duration(f.config.LiveTime) * time.Hour)
	dateMask := `\d{1,2}-\d{1,2}-\d{4}`
	dateMaskRe, _ := regexp.Compile(dateMask)

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
		matched, err := regexp.MatchString(f.config.FileName+`-`+dateMask+`.log`, file.Name())
		if err != nil {
			panic(err)
		}

		if matched {
			dateStr := dateMaskRe.FindAllString(file.Name(), -1)
			date, err := time.Parse("02-01-2006", dateStr[0])
			if err != nil {
				panic(err)
			}

			if dateCutoff.After(date) {
				fullName := fmt.Sprintf("%s/%s", f.config.Path, file.Name())
				os.Remove(fullName)
			}
		}
	}
}
