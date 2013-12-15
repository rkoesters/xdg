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
func New(root string) (Dir, error) {
	dir := &Dir{Path: root}
	// TODO: We need to check permissions and stuff on creation of a new
	// trash. This means we need to check for a sticky bit and a bunch of
	// other stuff.
	return dir
}

// Path returns the path to the root of the trash directory.
func (d *Dir) Path() string {
	return d.path
}
