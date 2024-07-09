package diwef

type App struct {
}

func Init() *App {
	app := &App{}

	return app
}

func (a *App) Debug(msg string) {
}

func (a *App) Info(msg string) {
}

func (a *App) Warning(msg string) {
}

func (a *App) Error(msg string) {
}

func (a *App) Fatal(msg string) {
}
