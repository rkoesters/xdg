package xdg

import (
	"testing"
)

func TestOpen(t *testing.T) {
	err := Open("http://www.golang.org/")
	if err != nil {
		t.Error(err)
	}
}

func TestOpenMissing(t *testing.T) {
	err := Open("non existent file")
	if err == nil {
		t.Fail()
	}
}
