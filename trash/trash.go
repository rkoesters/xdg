package trash

import (
	"os"
	"path/filepath"
)

// Dir represents a trash directory.
type Dir struct {
	path string
}

// New creates a new Trash directory.
func New(root string) (*Dir, error) {
	dir := &Dir{root}
	// TODO: We need to check permissions and stuff on creation of a new
	// trash. This means we need to check for a sticky bit and a bunch of
	// other stuff.
	return dir, nil
}

// Files returns a slice of the files in the trash.
func (d *Dir) Files() ([]string, error) {
	dir, err := os.Open(filepath.Join(d.path, "files"))
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	return dir.Readdirnames(0)
}

// Stat returns the trash Info for the given file in the trash.
func (d *Dir) Stat(s string) (*Info, error) {
	f, err := os.Open(filepath.Join(d.path, "info", s+".trashinfo"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewInfo(f)
}
