// Package ini implements the ini file format that is used in many of
// the xdg specs.
//
// WARNING: This package is meant for internal use and the API may
// change without warning.
package ini

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// Map is a map of a map of strings. The first string is the header
// section and the second is the key.
type Map struct {
	M map[string]map[string]string
}

// New creates a new Map and returns it.
func New(r io.Reader) (*Map, error) {
	m := new(Map)
	m.M = make(map[string]map[string]string)
	hdr := ""
	m.M[hdr] = make(map[string]string)

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
			m.M[hdr] = make(map[string]string)
		case strings.Contains(line, "="):
			// Key=Value pair.
			p := strings.SplitN(line, "=", 2)
			p[0] = strings.TrimSpace(p[0])
			p[1] = strings.TrimSpace(p[1])
			m.M[hdr][p[0]] = p[1]
		default:
			return nil, errors.New("invalid format")
		}
	}
	return m, nil
}

// String returns the value with the given group and key. This function
// is here because underlying data structure of Map may change.
func (m *Map) String(g, k string) string {
	return m.M[g][k]
}

// Bool returns the value as a bool.
func (m *Map) Bool(g, k string) bool {
	b, _ := strconv.ParseBool(m.String(g, k))
	return b
}

// List returns the value as a slice of strings.
func (m *Map) List(g, k string) []string {
	l := strings.Split(m.String(g, k), ";")
	for i := 0; i < len(l); i++ {
		if l[i] == "" {
			l = append(l[:i], l[i+1:]...)
			i--
		}
	}
	return l
}
