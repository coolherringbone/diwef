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

func stylingLogStr(level level, msg any) string {
	logStr := fmt.Sprintf("time=\"%s\"		level=\"%s\"		msg=\"%v\"\n",
		time.Now().Format("02-01-2006 15:04:05"),
		level,
		msg)

	return logStr
}
