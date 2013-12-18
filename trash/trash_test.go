package trash

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	files, err := Files()
	if err != nil {
		t.Error(err)
	}
	for _, i := range files {
		t.Log(i)
	}
}

func TestStat(t *testing.T) {
	files, err := Files()
	if err != nil {
		t.Error(err)
	}
	for _, i := range files {
		info, err := Stat(i)
		if err != nil {
			t.Error(err)
		}
		t.Log(info)
	}
}
