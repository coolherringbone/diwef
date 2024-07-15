package diwef

type level string
type Levels []level

const (
	DebugLevel   level = "debug"
	InfoLevel    level = "info"
	WarningLevel level = "warning"
	ErrorLevel   level = "error"
	FatalLevel   level = "fatal"
)

var (
	defaultUseLevel = Levels{DebugLevel, InfoLevel, WarningLevel, ErrorLevel, FatalLevel}
)

func setLevels(levels Levels) map[level]bool {
	var res = map[level]bool{
		DebugLevel:   false,
		InfoLevel:    false,
		WarningLevel: false,
		ErrorLevel:   false,
		FatalLevel:   false,
	}

	for _, l := range levels {
		res[l] = true
	}

	return res
}
