package basedir

import (
	"testing"
)

func TestGetenv(t *testing.T) {
	if getenv("PATH", "not set") == "not set" {
		t.Error("Couldn't get PATH")
	}
	if getenv("does_not_exist", "not set") != "not set" {
		t.Error("does_not_exist exists")
	}
}
