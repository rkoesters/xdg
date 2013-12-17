// Package ini implements the ini file format that is used in many of
// the xdg specs.
package ini

import (
	"io"
	"bufio"
	"strings"
	"errors"
)

// File is a map of a map of strings. The first string is the header
// section and the second is the key.
type Ini map[string]map[string]string

// New creates a new File and returns it.
func New(r io.Reader) (Ini, error) {
	m := make(Ini)
	hdr := "default"
	m[hdr] = make(map[string]string)

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
			hdr = line[1:len(line)-1]
			m[hdr] = make(map[string]string)
		case strings.Contains(line, "="):
			// Key=Value pair.
			p := strings.SplitN(line, "=", 2)
			p[0] = strings.TrimSpace(p[0])
			p[1] = strings.TrimSpace(p[1])
			m[hdr][p[0]] = p[1]
		default:
			return nil, errors.New("invalid format")
		}
	}
	return m, nil
}
