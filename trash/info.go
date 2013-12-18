package trash

import (
	"github.com/rkoesters/xdg/ini"
	"io"
	"net/url"
)

const tinfo = "Trash Info"

// Info represents a .trashinfo file.
type Info struct {
	Path string
	DeletionDate time.Time
}

// NewInfo creates a new Info using the given io.Reader.
func NewInfo(r io.Reader) (*Info, error) {
	m, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	info := new(Info)
	info.Path = url.QueryUnescape(m.String(tinfo, "Path"))
	info.DeletionDate, err = time.Parse(time.RFC3339, m.String(tinfo, "DeletionDate"))
	if err != nil {
		return nil, err
	}
	return info, nil
}
