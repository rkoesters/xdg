package basedir

import (
	"os"
	"testing"
)

func TestBaseDir(t *testing.T) {
	if Home == "" {
		t.Error("Home not set")
	}
	if DataHome == "" {
		t.Error("DataHome not set")
	}
	if ConfigHome == "" {
		t.Error("ConfigHome not set")
	}
	if CacheHome == "" {
		t.Error("CacheHome not set")
	}
	if RuntimeDir == "" {
		t.Error("RuntimeDir not set")
	}
	if len(DataDirs) == 0 {
		t.Error("DataDirs not set")
	}
	if len(ConfigDirs) == 0 {
		t.Error("ConfigDirs not set")
	}
}

func TestGetpath(t *testing.T) {
	const notSet = "not set"
	if getpath("HOME", notSet) == notSet {
		t.Error("Couldn't get HOME")
	}
	if getpath("does_not_exist", notSet) != notSet {
		t.Error("does_not_exist exists")
	}
	if getpath("USER", notSet) != notSet {
		t.Error("USER appears to be an absolute path")
	}
}

func TestGetpathlist(t *testing.T) {
	if getpathlist("PATH", nil) == nil {
		t.Error("Couldn't get PATH")
	}
	if getpathlist("does_not_exist", nil) != nil {
		t.Error("does_not_exist exists")
	}
	err := os.Setenv("xdg_test_var", "/a:c:/a/b:d/f")
	if err != nil {
		t.Error(err)
	}
	testVar := getpathlist("xdg_test_var", nil)
	if testVar[0] != "/a" || testVar[1] != "/a/b" {
		t.Error("getpathlist returned relative paths")
	}
}
