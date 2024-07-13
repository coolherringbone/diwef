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
	logStr := stylingLogStr("debug", msg)
	fmt.Print(logStr)
}

func (cli *cliWriter) info(msg string) {
	logStr := stylingLogStr("info", msg)
	fmt.Print(logStr)
}

func (cli *cliWriter) warning(msg string) {
	logStr := stylingLogStr("warning", msg)
	fmt.Print(logStr)
}

func (cli *cliWriter) error(msg string) {
	logStr := stylingLogStr("error", msg)
	fmt.Print(logStr)
}

func (cli *cliWriter) fatal(msg string) {
	logStr := stylingLogStr("fatal", msg)
	fmt.Print(logStr)
}
