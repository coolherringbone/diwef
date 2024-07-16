package diwef

import (
	"encoding/json"
	"fmt"
	"time"
)

type formatter string

type jsonLog struct {
	Time  string `json:"time"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

const (
	STRFormatter  formatter = "str"
	JSONFormatter formatter = "json"
)

func strFormatting(level level, msg any) string {
	res := fmt.Sprintf("time=\"%s\"		level=\"%s\"		msg=\"%v\"\n",
		time.Now().Format("02-01-2006 15:04:05"),
		level,
		msg)

	return res
}

func jsonsFormatting(level level, msg any, data []byte) ([]byte, error) {
	var logs []jsonLog

	if err := json.Unmarshal(data, &logs); err != nil && err.Error() != "unexpected end of JSON input" {
		return nil, err
	}

	log := jsonLog{
		Time:  time.Now().Format("02-01-2006 15:04:05"),
		Level: string(level),
		Msg:   fmt.Sprintf("%v", msg),
	}

	logs = append(logs, log)

	res, err := json.MarshalIndent(&logs, "", "  ")
	if err != nil {
		return nil, err
	}

	return res, nil
}