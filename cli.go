package diwef

import (
	"errors"
	"fmt"
)

type cliWriter struct {
	config CliWriter
}

type CliWriter struct {
	UseLevels []Level
}

func NewCliWriter(config ...CliWriter) (writer, error) {
	var w writer
	var cli = &cliWriter{}

	if len(config) == 1 {
		cli.config.UseLevels = nvl(config[0].UseLevels, DefaultUseLevels).([]Level)
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		cli.config.UseLevels = DefaultUseLevels
	}

	w = cli

	return w, nil
}

func (cli *cliWriter) debug(msg string) {
	cli.writing(DebugLevel, msg)
}

func (cli *cliWriter) info(msg string) {
	cli.writing(InfoLevel, msg)
}

func (cli *cliWriter) warning(msg string) {
	cli.writing(WarningLevel, msg)
}

func (cli *cliWriter) error(msg string) {
	cli.writing(ErrorLevel, msg)
}

func (cli *cliWriter) fatal(msg string) {
	cli.writing(FatalLevel, msg)
}

func (cli *cliWriter) writing(level Level, msg string) {
	logStr := stylingLogStr(level.name, msg)
	fmt.Print(logStr)
}
