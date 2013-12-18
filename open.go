// Package xdg provides access to the XDG specs. Most of the
// functionality can be found in the subpackages.
package xdg

import (
	"os/exec"
	"errors"
)

// These errors can be returned by Open.
var (
	ErrSyntax = errors.New("Error in command line syntax")
	ErrFileNotFound = errors.New("One of the files passed on the command line did not exist")
	ErrToolNotFound = errors.New("A required tool could not be found")
	ErrFailed = errors.New("The action failed")
)

// Open runs the command xdg-open with the given uri as an argument.
func Open(uri string) error {
	c := exec.Command("xdg-open", uri)
	err := c.Run()
	if err == nil {
		return nil
	}
	switch err.Error() {
	case "exit status 1":
		return ErrSyntax
	case "exit status 2":
		return ErrFileNotFound
	case "exit status 3":
		return ErrToolNotFound
	case "exit status 4":
		return ErrFailed
	}
	return err
}
