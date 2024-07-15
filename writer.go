package diwef

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

type writer interface {
	debug(msg any)
	info(msg any)
	warning(msg any)
	error(msg any)
	fatal(msg any)
}

type fileWriter struct {
	useLevels map[level]bool
	path      string
	fileName  string
	liveTime  int
}

type FileWriter struct {
	UseLevels Levels
	Path      string
	FileName  string
	LiveTime  int
}

type cliWriter struct {
	useLevels map[level]bool
}

type CliWriter struct {
	UseLevels Levels
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
		f.useLevels = setLevels(nvl(config[0].UseLevels, defaultUseLevel).(Levels))
		f.path = nvl(config[0].Path, defaultPath).(string)
		f.fileName = nvl(config[0].FileName, defaultFileName).(string)
		f.liveTime = nvl(config[0].LiveTime, defaultLiveTime).(int)
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		f.useLevels = setLevels(defaultUseLevel)
		f.path = defaultPath
		f.fileName = defaultFileName
		f.liveTime = defaultLiveTime
	}

	if err := os.MkdirAll(f.path, 0744); err != nil {
		return nil, err
	}

	w = f

	return w, nil
}

func NewCliWriter(config ...CliWriter) (writer, error) {
	var w writer
	var cli = &cliWriter{}

	if len(config) == 1 {
		cli.useLevels = setLevels(nvl(config[0].UseLevels, defaultUseLevel).(Levels))
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		cli.useLevels = setLevels(defaultUseLevel)
	}

	w = cli

	return w, nil
}

func (f *fileWriter) debug(msg any) {
	f.writing(DebugLevel, msg)
}

func (cli *cliWriter) debug(msg any) {
	cli.writing(DebugLevel, msg)
}

func (f *fileWriter) info(msg any) {
	f.writing(InfoLevel, msg)
}

func (cli *cliWriter) info(msg any) {
	cli.writing(InfoLevel, msg)
}

func (f *fileWriter) warning(msg any) {
	f.writing(WarningLevel, msg)
}

func (cli *cliWriter) warning(msg any) {
	cli.writing(WarningLevel, msg)
}

func (f *fileWriter) error(msg any) {
	f.writing(ErrorLevel, msg)
}

func (cli *cliWriter) error(msg any) {
	cli.writing(ErrorLevel, msg)
}

func (f *fileWriter) fatal(msg any) {
	f.writing(FatalLevel, msg)
}

func (cli *cliWriter) fatal(msg any) {
	cli.writing(FatalLevel, msg)
}

func (f *fileWriter) writing(level level, msg any) {
	if !f.useLevels[level] {
		return
	}

	file, _ := f.openLogFile()
	defer file.Close()

	logStr := stylingLogStr(level, msg)

	file.WriteString(logStr)

	_ = f.clearingLogs()
}

func (cli *cliWriter) writing(level level, msg any) {
	if !cli.useLevels[level] {
		return
	}

	logStr := stylingLogStr(level, msg)

	fmt.Print(logStr)
}

func (f *fileWriter) openLogFile() (*os.File, error) {
	fullName := fmt.Sprintf("%s/%s-%s.log",
		f.path,
		f.fileName,
		time.Now().Format("02-01-2006"))

	file, err := os.OpenFile(fullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *fileWriter) clearingLogs() error {
	if f.liveTime == 0 {
		return nil
	}

	dateCutoff := time.Now().Add(-24 * time.Duration(f.liveTime) * time.Hour)
	dateMask := `\d{1,2}-\d{1,2}-\d{4}`
	dateMaskRe, _ := regexp.Compile(dateMask)

	path, err := os.Open(f.path)
	if err != nil {
		return err
	}
	defer path.Close()

	pathFileList, err := path.Readdir(-1)
	if err != nil {
		return err
	}

	for _, file := range pathFileList {
		matched, err := regexp.MatchString(f.fileName+`-`+dateMask+`.log`, file.Name())
		if err != nil {
			return err
		}

		if matched {
			dateStr := dateMaskRe.FindAllString(file.Name(), -1)
			date, err := time.Parse("02-01-2006", dateStr[0])
			if err != nil {
				return err
			}

			if dateCutoff.After(date) {
				fullName := fmt.Sprintf("%s/%s", f.path, file.Name())
				if err = os.Remove(fullName); err != nil {
					return err
				}
			}
		}
	}

	return nil
}