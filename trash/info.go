package trash

import (
	"fmt"
	"github.com/rkoesters/xdg/ini"
	"io"
	"net/url"
	"strings"
	"time"
)

const tinfo = "Trash Info"

// time.RFC3339 isn't working right, so we use our own.
const rfc3339 = "2006-01-02T15:04:05"

// Info represents a .trashinfo file.
type Info struct {
	Path         string
	DeletionDate time.Time
}

// NewInfo creates a new Info using the given io.Reader.
func NewInfo(r io.Reader) (*Info, error) {
	m, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	info := new(Info)
	info.Path, err = url.QueryUnescape(m.String(tinfo, "Path"))
	if err != nil {
		return nil, err
	}
	info.DeletionDate, err = time.Parse(rfc3339, m.String(tinfo, "DeletionDate"))
	if err != nil {
		return nil, err
	}
	return info, nil
}

// String returns Info as a string in the INI format.
func (i *Info) String() string {
	return fmt.Sprintf(
		"[Trash Info]\nPath=%v\nDeletionDate=%v\n",
		queryEscape(i.Path),
		i.DeletionDate.Format(rfc3339),
	)
}

// queryEscape is a wrapper function around url.QueryEscape that doesn't
// escape '/'.
func queryEscape(s string) string {
	// The first for loop is the workaround for "/".
	a := strings.Split(s, "/")
	for i := 0; i < len(a); i++ {
		// The second for loop is the workaround for " ".
		b := strings.Split(a[i], " ")
		for j := 0; j < len(b); j++ {
			b[j] = url.QueryEscape(b[j])
		}
		a[i] = strings.Join(b, "%20")
	}
	return strings.Join(a, "/")
}
