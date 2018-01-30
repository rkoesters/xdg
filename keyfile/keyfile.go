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
	"strconv"
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
			// New section.
			hdr = line[1 : len(line)-1]
			kf.m[hdr] = make(map[string]string)
		case strings.Contains(line, "="):
			// Key=Value pair.
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

// String returns the value with the given group and key. This function
// is here because underlying data structure of KeyFile may change.
func (kf *KeyFile) String(g, k string) string {
	return kf.m[g][k]
}

// Bool returns the value as a bool.
func (kf *KeyFile) Bool(g, k string) bool {
	b, _ := strconv.ParseBool(kf.String(g, k))
	return b
}

// List returns the value as a slice of strings.
func (kf *KeyFile) List(g, k string) []string {
	l := strings.Split(kf.String(g, k), ";")
	for i := 0; i < len(l); i++ {
		if l[i] == "" {
			l = append(l[:i], l[i+1:]...)
			i--
		}
	}
	return l
}
