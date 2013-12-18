package trash

import (
	"github.com/rkoesters/xdg/basedir"
	"path/filepath"
)

var hometrash *Dir

func init() {
	var err error
	hometrash, err = New(filepath.Join(basedir.DataHome, "Trash"))
	if err != nil {
		// TODO: remove this
		panic(err)
	}
}

func Files() ([]string, error) { return hometrash.Files() }

func Stat(s string) (*Info, error) { return hometrash.Stat(s) }
