package diwef

type Config struct {
	Path     string
	FileName string
	LiveTime int
}

const (
	DefaultPath     = "log"
	DefaultFileName = "app"
	DefaultLiveTime = 0
)
