package trash

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Dir represents a trash directory.
type Dir struct {
	path string
}

// New creates a new Trash directory.
func New(root string) (*Dir, error) {
	// TODO: We need to check permissions and stuff on creation of a
	// new trash. This means we need to check for a sticky bit and a
	// bunch of other stuff.
	panic("TODO")
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

// Trash moves the file at the given path to the trash.
func (d *Dir) Trash(p string) error {
	// Find an open file name.
	tname := filepath.Base(p)
	for i := 2; d.exists(tname); i++ {
		tname = filepath.Base(p) + "." + strconv.Itoa(i)
	}

	// First, write the trashinfo file.
	abs, err := filepath.Abs(p)
	if err != nil {
		return err
	}
	info := &Info{abs, time.Now()}
	err = ioutil.WriteFile(d.info2path(tname), []byte(info.String()), os.ModePerm)
	if err != nil {
		return err
	}

	// Next, move the file to the trash.
	return os.Rename(p, d.file2path(tname))
}

// Restore moves the file from the trash to its original location.
func (d *Dir) Restore(s string) error {
	info, err := d.Stat(s)
	if err != nil {
		return err
	}

	return d.RestoreTo(s, info.Path)
}

// RestoreTo moves the file from the trash to the specified location.
func (d *Dir) RestoreTo(s, p string) error {
	_, err := os.Stat(p)
	if err == nil {
		return os.ErrExist
	}

	return os.Rename(d.file2path(s), p)
}

func (d *Dir) exists(s string) bool {
	_, err := os.Stat(d.file2path(s))
	if os.IsNotExist(err) {
		return false
	}
	return true
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

// IsEmpty returns whether or not the trash is empty.
func (d *Dir) IsEmpty() bool {
	files, _ := d.Files()
	return len(files) == 0
}
