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

func TestIsEmpty(t *testing.T) {
	// We can't expect a value for IsEmpty because it depends on the
	// contents of the user's trash which could be empty or not. So
	// we just call the function to make sure it doesn't panic or
	// anything extreme.
	t.Logf("IsEmpty()=%v", IsEmpty())
}
