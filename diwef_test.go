package diwef

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	initEmptyConfig(t)
	initOneConfig(t)
	initManyConfig(t)

	errorCreatPath(t)
}

func initEmptyConfig(t *testing.T) {
	log, err := Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	if log.config.Path != DefaultPath {
		t.Errorf("config.Path did not accept default value")
	}

	if log.config.FileName != DefaultFileName {
		t.Errorf("config.FileName did not accept default value")
	}

	if log.config.LiveTime != DefaultLiveTime {
		t.Errorf("config.LiveTime did not accept default value")
	}

	_, err = os.Stat(log.config.Path)
	if err == nil {
		os.Remove(log.config.Path)
	} else if os.IsNotExist(err) {
		t.Errorf("Path was not created")
	} else {
		t.Errorf(err.Error())
	}
}

func initOneConfig(t *testing.T) {
	path := "test_path"
	fileName := "test_file"
	liveTime := 1

	log, err := Init(Config{
		Path:     path,
		FileName: fileName,
		LiveTime: liveTime,
	})
	if err != nil {
		t.Errorf(err.Error())
	}

	if log.config.Path != path {
		t.Errorf("config.Path did not accept the specified value")
	}

	if log.config.FileName != fileName {
		t.Errorf("config.FileName did not accept the specified value")
	}

	if log.config.LiveTime != liveTime {
		t.Errorf("config.LiveTime did not accept the specified value")
	}

	_, err = os.Stat(path)
	if err == nil {
		os.Remove(path)
	} else if os.IsNotExist(err) {
		t.Errorf("Path was not created")
	} else {
		t.Errorf(err.Error())
	}
}

func initManyConfig(t *testing.T) {
	path := "test_path"
	fileName := "test_file"
	liveTime := 1

	if _, err := Init(Config{
		Path:     path,
		FileName: fileName,
		LiveTime: liveTime,
	}, Config{
		Path:     path,
		FileName: fileName,
		LiveTime: liveTime,
	}); err == nil {
		t.Errorf("no error: there can be only one config (or even empty)")
	}
}

func errorCreatPath(t *testing.T) {
	path := ":?"

	if _, err := Init(Config{
		Path: path,
	}); err == nil {
		t.Errorf("no error: the directory name is invalid.")
	}
}
