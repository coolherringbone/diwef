package diwef

import (
	"fmt"
)

type cliWriter struct {
	debugLevel   level
	infoLevel    level
	warningLevel level
	errorLevel   level
	fatalLevel   level
}

func NewCliWriter() writer {
	var w writer
	var cli = &cliWriter{}

	cli.SetLevel(DebugLevel, InfoLevel, WarningLevel, ErrorLevel, FatalLevel)

	w = cli

	return w
}

func (cli *cliWriter) SetLevel(level ...level) {
	cli.debugLevel.activ = false
	cli.infoLevel.activ = false
	cli.warningLevel.activ = false
	cli.errorLevel.activ = false
	cli.fatalLevel.activ = false

	for _, l := range level {
		switch l {
		case DebugLevel:
			cli.debugLevel = l
		case InfoLevel:
			cli.infoLevel = l
		case WarningLevel:
			cli.warningLevel = l
		case ErrorLevel:
			cli.errorLevel = l
		case FatalLevel:
			cli.fatalLevel = l
		}

	}
}

func (cli *cliWriter) debug(msg string) {
	cli.writing(cli.debugLevel, msg)
}

func (cli *cliWriter) info(msg string) {
	cli.writing(cli.infoLevel, msg)
}

func (cli *cliWriter) warning(msg string) {
	cli.writing(cli.warningLevel, msg)
}

func (cli *cliWriter) error(msg string) {
	cli.writing(cli.errorLevel, msg)
}

func (cli *cliWriter) fatal(msg string) {
	cli.writing(cli.fatalLevel, msg)
}

func (cli *cliWriter) writing(level level, msg string) {
	if !level.activ {
		return
	}

	logStr := stylingLogStr(level.name, msg)
	fmt.Print(logStr)
}
