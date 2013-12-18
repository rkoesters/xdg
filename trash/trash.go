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

func (d *Dir) file2path(s string) string {
	return filepath.Join(d.path, "files", s)
}

func (d *Dir) info2path(s string) string {
	return filepath.Join(d.path, "info", s+".trashinfo")
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

// Stat returns the Info for the given file in the trash.
func (d *Dir) Stat(s string) (*Info, error) {
	f, err := os.Open(d.info2path(s))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewInfo(f)
}

// Erase removes the given file from the trash.
func (d *Dir) Erase(s string) error {
	err := os.Remove(d.file2path(s))
	if err != nil {
		return err
	}
	err = os.Remove(d.info2path(s))
	if err != nil {
		return err
	}
	return nil
}

// EraseAll removes the given file and any children it contains.
func (d *Dir) EraseAll(s string) error {
	err := os.RemoveAll(d.file2path(s))
	if err != nil {
		return err
	}
	err = os.RemoveAll(d.info2path(s))
	if err != nil {
		return nil
	}
	return nil
}

// Empty erases all the files in the trash.
func (d *Dir) Empty() error {
	files, err := d.Files()
	if err != nil {
		return err
	}
	for _, i := range files {
		err = d.EraseAll(i)
		if err != nil {
			return err
		}
	}
	return nil
}
