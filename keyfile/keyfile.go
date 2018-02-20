// Package keyfile implements the ini file format that is used in many
// of the xdg specs.
//
// WARNING: This package is meant for internal use and the API may
// change without warning.
package keyfile

import (
	"bufio"
	"bytes"
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

// Groups returns a slice of groups that exist for the KeyFile.
func (kf *KeyFile) Groups() []string {
	groups := make([]string, 0)
	for k := range kf.m {
		groups = append(groups, k)
	}
	return groups
}

// Keys returns a slice of keys that exist for the given group 'g'.
func (kf *KeyFile) Keys(g string) []string {
	keys := make([]string, 0)
	for k := range kf.m[g] {
		keys = append(keys, k)
	}
	return keys
}

// ValueExists returns a bool indicating whether the given group 'g' and
// key 'k' have a value.
func (kf *KeyFile) ValueExists(g, k string) bool {
	_, exists := kf.m[g][k]
	return exists
}

// Value returns the raw string for group 'g' and key 'k'.
func (kf *KeyFile) Value(g, k string) string {
	return kf.m[g][k]
}

// ValueList returns a slice of raw strings for group 'g' and key 'k'.
func (kf *KeyFile) ValueList(g, k string) ([]string, error) {
	var buf bytes.Buffer
	var isEscaped bool
	var list []string
	var err error

	for _, r := range kf.Value(g, k) {
		if isEscaped {
			if r == ';' {
				_, err = buf.WriteRune(';')
				if err != nil {
					return nil, err
				}
			} else {
				// The escape sequence isn't '\;', so we want to copy it
				// over as is.
				_, err = buf.WriteRune('\\')
				if err != nil {
					return nil, err
				}
				_, err = buf.WriteRune(r)
				if err != nil {
					return nil, err
				}
			}
			isEscaped = false
		} else {
			switch r {
			case '\\':
				isEscaped = true
			case ';':
				list = append(list, buf.String())
				buf.Reset()
			default:
				buf.WriteRune(r)
			}
		}
	}
	if isEscaped {
		return nil, ErrUnexpectedEndOfString
	}

	last := buf.String()
	if last != "" {
		list = append(list, last)
	}

	return list, nil
}
