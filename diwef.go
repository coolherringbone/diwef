package diwef

import (
	"errors"
	"os"
)

type App struct {
	config Config
}

func Init(config ...Config) (*App, error) {
	app := &App{}

	if len(config) == 1 {
		app.config = config[0]
	} else if len(config) > 1 {
		return nil, errors.New("there can be only one config (or even empty)")
	} else {
		app.config.Path = DefaultPath
		app.config.FileName = DefaultFileName
		app.config.LiveTime = DefaultLiveTime
	}

	if err := os.MkdirAll(app.config.Path, 0744); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Debug(msg string) {
	a.writer("debug", msg)
}

func (a *App) Info(msg string) {
	a.writer("info", msg)
}

func (a *App) Warning(msg string) {
	a.writer("warning", msg)
}

func (a *App) Error(msg string) {
	a.writer("error", msg)
}

func (a *App) Fatal(msg string) {
	a.writer("fatal", msg)
}
