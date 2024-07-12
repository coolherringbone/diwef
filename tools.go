package diwef

import (
	"fmt"
	"time"
)

func nvl(a, b any) any {
	if a == nil || a == "" || a == 0 {
		return b
	} else {
		return a
	}
}

func stylingLogStr(level, msg string) string {
	logStr := fmt.Sprintf("time=\"%s\"		level=\"%s\"		msg=\"%s\"\n",
		time.Now().Format("15:04:05"),
		level,
		msg)

	return logStr
}
