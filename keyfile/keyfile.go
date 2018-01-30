// Package keyfile implements the ini file format that is used in many
// of the xdg specs.
//
// WARNING: This package is meant for internal use and the API may
// change without warning.
package keyfile

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// KeyFile is a map of a map of strings. The first string is the header
// section and the second is the key.
type KeyFile struct {
	m map[string]map[string]string
}

// New creates a new KeyFile and returns it.
func New(r io.Reader) (*KeyFile, error) {
	kf := new(KeyFile)
	kf.m = make(map[string]map[string]string)
	hdr := ""
	kf.m[hdr] = make(map[string]string)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		switch {
		case len(line) == 0:
			// Empty line.
		case strings.HasPrefix(line, "#"):
			// Comment.
		case line[0:1] == "[" && line[len(line)-1:] == "]":
			// Group header.
			hdr = line[1 : len(line)-1]
			kf.m[hdr] = make(map[string]string)
		case strings.Contains(line, "="):
			// Entry.
			p := strings.SplitN(line, "=", 2)
			p[0] = strings.TrimSpace(p[0])
			p[1] = strings.TrimSpace(p[1])
			kf.m[hdr][p[0]] = p[1]
		default:
			return nil, errors.New("invalid format")
		}
	}
	return kf, nil
}

// Value returns the raw string for group 'g' and key 'k'.
func (kf *KeyFile) Value(g, k string) string {
	return kf.m[g][k]
}

// List returns the value as a slice of strings.
func (kf *KeyFile) List(g, k string) []string {
	l := strings.Split(kf.Value(g, k), ";")
	for i := 0; i < len(l); i++ {
		if l[i] == "" {
			l = append(l[:i], l[i+1:]...)
			i--
		}
	}
	return l
}
