package diwef

import "testing"

func TestInit(t *testing.T) {
	writer, err := NewFileWriter()
	if err != nil {
		t.Errorf(err.Error())
	}

	Init(writer)
}
