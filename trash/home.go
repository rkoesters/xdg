package trash

import (
	"github.com/rkoesters/basedir"
	"os"
	"path/filepath"
)

// The main trash directory.
var Home *Dir

func init() {
	Home = &Dir{path: filepath.Join(basedir.DataHome, "Trash")}

	_, err := os.Stat(Home.Path())
	if err != nil {
		err = os.MkdirAll(Home.Path(), 777)
		if err != nil {
			Home = nil
		}
	}
}

// FindTrash returns the trash directory for the given path.
func FindTrash(path string) (Dir, error) {
	panic("TODO")
}
