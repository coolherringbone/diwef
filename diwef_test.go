package diwef

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

// diwef
func defInit() (*Logger, error) {
	w1, err := newDefCliWriter()
	if err != nil {
		return nil, err
	}

	w2, err := newDefFileWriter()
	if err != nil {
		return nil, err
	}

	l, err := Init(w1, w2)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func TestInit(t *testing.T) {
	_, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestErrInit(t *testing.T) {
	_, err := Init()
	if err == nil {
		t.Errorf("incorrect processing error: there must be at least one writer")
	}

	w1, err := newDefCliWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	w2, err := newDefCliWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = Init(w1, w2)
	if err == nil {
		t.Errorf("incorrect processing error: there can be only one cli writer")
	}

	w1, err = newDefFileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	w2, err = newDefFileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = Init(w1, w2)
	if err == nil {
		t.Errorf("incorrect processing error:  there is already such a file writer")
	}
}

func TestDebug(t *testing.T) {
	log, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Debug("debug_test")
}

func TestInfo(t *testing.T) {
	log, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Info("info_test")
}

func TestWarning(t *testing.T) {
	log, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Warning("warning_test")
}

func TestError(t *testing.T) {
	log, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Error("error_test")
}

func TestFatal(t *testing.T) {
	log, err := defInit()
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Fatal("fatal_test")
}

// writer
func newDefFileWriter() (writer, error) {
	w, err := NewFileWriter()
	return w, err
}

func newCastomFileWriter() (writer, error) {
	w, err := NewFileWriter(FileWriter{
		UseLevels: Levels{DebugLevel},
		FileName:  "test",
		Formatter: JSONFormatter,
	})
	return w, err
}

func newCastom2FileWriter() (writer, error) {
	w, err := NewFileWriter(FileWriter{})
	return w, err
}

func newErrFileWriter() (writer, error) {
	w, err := NewFileWriter(FileWriter{
		UseLevels: Levels{DebugLevel},
		FileName:  "test",
		Formatter: JSONFormatter,
	}, FileWriter{
		UseLevels: Levels{ErrorLevel},
		FileName:  "test2",
		Formatter: STRFormatter,
	})
	return w, err
}

func TestNewFileWriter(t *testing.T) {
	_, err := newDefFileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newCastomFileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newCastom2FileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newErrFileWriter()
	if err == nil {
		t.Errorf("incorrect processing error: there can be only one config (or even empty)")
	}
}

func newDefCliWriter() (writer, error) {
	w, err := NewCliWriter()
	return w, err
}

func newCastomCliWriter() (writer, error) {
	w, err := NewCliWriter(CliWriter{
		UseLevels: Levels{DebugLevel},
	})
	return w, err
}

func newCastom2CliWriter() (writer, error) {
	w, err := NewCliWriter(CliWriter{})
	return w, err
}

func newErrCliWriter() (writer, error) {
	w, err := NewCliWriter(CliWriter{
		UseLevels: Levels{DebugLevel},
	}, CliWriter{
		UseLevels: Levels{ErrorLevel},
	})
	return w, err
}

func TestNewCliWriter(t *testing.T) {
	_, err := newDefCliWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newCastomCliWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newCastom2CliWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = newErrCliWriter()
	if err == nil {
		t.Errorf("incorrect processing error: there can be only one config (or even empty)")
	}
}

func TestOpenLogFileError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		path := "log"
		fileName := ":?"

		w, err := NewFileWriter(FileWriter{
			Path:     path,
			FileName: fileName,
		})
		if err != nil {
			t.Errorf(err.Error())
		}

		log, err := Init(w)
		if err != nil {
			t.Errorf(err.Error())
		}

		log.Debug("debug_test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestOpenLogFileError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	e, ok := err.(*exec.ExitError)
	if ok && !e.Success() {
		return
	} else {
		t.Errorf(err.Error())
	}
}

// formatter
func TestStrFormatting(t *testing.T) {
	var level level = DebugLevel
	msg := "test"
	caller := ""
	str := fmt.Sprintf("time=\"%s\"		level=\"%s\"		msg=\"%v\" caller=\"%s\"\n",
		time.Now().Format("02-01-2006 15:04:05"),
		level,
		msg,
		caller)

	if str != strFormatting(level, msg, caller) {
		t.Errorf("incorrect string formatting")
	}
}

func TestJsonsFormatting(t *testing.T) {
	var level level = DebugLevel
	msg := "test"
	caller := ""
	str := fmt.Sprintf(`[
  {
    "time": "%s",
    "level": "%s",
    "msg": "%s",
    "caller": "%s"
  }
]`, time.Now().Format("02-01-2006 15:04:05"),
		level,
		msg,
		caller)

	res, err := jsonsFormatting(level, msg, caller, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	if string(res) != str {
		t.Errorf("incorrect jsons formatting")
	}

	str = fmt.Sprintf(`
		{
		  "date": "%s",
		  "level": "%s",
		  "msg": "%s",
		  "caller": "%s"
		}
	  `, time.Now().Format("02-01-2006 15:04:05"),
		level,
		msg,
		caller)
	_, err = jsonsFormatting(level, msg, caller, []byte(str))
	if err == nil {
		t.Errorf("incorrect processing error")
	}
}

// level
func TestSetLevels(t *testing.T) {
	var levels Levels
	r := setLevels(levels)
	if r[DebugLevel] != false ||
		r[InfoLevel] != false ||
		r[WarningLevel] != false ||
		r[ErrorLevel] != false ||
		r[FatalLevel] != false {
		t.Errorf("error set levels")
	}

	levels = append(levels, InfoLevel, FatalLevel)
	r = setLevels(levels)
	if r[DebugLevel] != false ||
		r[InfoLevel] != true ||
		r[WarningLevel] != false ||
		r[ErrorLevel] != false ||
		r[FatalLevel] != true {
		t.Errorf("error set levels")
	}

	levels = defaultUseLevel
	r = setLevels(levels)
	if r[DebugLevel] != true ||
		r[InfoLevel] != true ||
		r[WarningLevel] != true ||
		r[ErrorLevel] != true ||
		r[FatalLevel] != true {
		t.Errorf("error set levels")
	}
}

// tools
func TestNvl(t *testing.T) {
	as := "test"
	bs := ""

	cs := nvl(as, bs).(string)
	if cs != as {
		t.Errorf("incorrect processing \"\"")
	}

	as = ""
	bs = "test"

	cs = nvl(as, bs).(string)
	if cs != bs {
		t.Errorf("incorrect processing \"\"")
	}

	an := 1
	bn := 0
	cn := nvl(an, bn).(int)
	if cn != an {
		t.Errorf("incorrect processing 0")
	}

	an = 0
	bn = 1
	cn = nvl(an, bn).(int)
	if cn != bn {
		t.Errorf("incorrect processing 0")
	}
}

func TestGetCallerInfo(t *testing.T) {
	caller := getCallerInfo(0)
	if caller == "" {
		t.Errorf("error getting caller info")
	}

	caller = getCallerInfo(3)
	if caller != "" {
		t.Errorf("incorrect processing error when getting caller info")
	}
}
