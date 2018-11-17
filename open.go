// Package xdg provides access to the XDG specs. Most of the
// functionality can be found in the subpackages.
package xdg

import (
	"errors"
	"os/exec"
)

// These errors can be returned by Open.
var (
	ErrSyntax       = errors.New("error in command line syntax")
	ErrFileNotFound = errors.New("one of the files passed on the command line did not exist")
	ErrToolNotFound = errors.New("a required tool could not be found")
	ErrFailed       = errors.New("the action failed")
)

// Open runs the command xdg-open with the given uri as an argument.
func Open(uri string) error {
	c := exec.Command("xdg-open", uri)
	err := c.Run()
	if err != nil {
		switch err.Error() {
		case "exit status 1":
			return ErrSyntax
		case "exit status 2":
			return ErrFileNotFound
		case "exit status 3":
			return ErrToolNotFound
		case "exit status 4":
			return ErrFailed
		default:
			return err
		}
	}
	return nil
}
