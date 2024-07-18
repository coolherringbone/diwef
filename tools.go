package diwef

import (
	"fmt"
	"runtime"
)

func nvl(a, b any) any {
	if a == "" || a == 0 {
		return b
	} else {
		return a
	}
}

func getCallerInfo(skip int) string {
	_, file, no, ok := runtime.Caller(skip + 1)
	if ok {
		return fmt.Sprintf("file:%s line:%d", file, no)
	} else {
		return ""
	}
}
