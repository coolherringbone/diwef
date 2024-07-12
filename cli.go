package diwef

import "fmt"

type cliWriter struct {
}

func NewCliWriter() writer {
	var cli = &cliWriter{}

	var w writer = cli

	return w
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
