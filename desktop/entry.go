package desktop

import (
	"errors"
	"github.com/rkoesters/xdg/ini"
	"io"
)

var (
	ErrMissingType = errors.New("missing entry type")
	ErrMissingName = errors.New("missing entry name")
	ErrMissingURL  = errors.New("missing entry url")
)

const dent = "Desktop Entry"

// Entry represents a desktop entry file.
type Entry struct {
	m ini.Map
}

func New(r io.Reader) (*Entry, error) {
	dfile, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	e := &Entry{dfile}

	// Check that the desktop file is valid.
	_, ok := e.m[dent]["Type"]
	if !ok {
		return nil, ErrMissingType
	}
	t := e.Type()
	switch e.Type() {
	case Link:
		_, ok = e.m[dent]["URL"]
		if !ok {
			return nil, ErrMissingURL
		}
		fallthrough
	case Application, Directory:
		_, ok = e.m[dent]["Name"]
		if !ok {
			return nil, ErrMissingName
		}
	}
	return e, nil
}

func (e *Entry) Type() Type { return ParseType(e.m.Get(dent, "Type")) }

func (e *Entry) Version() string { return e.m.Get(dent, "Version") }

func (e *Entry) Name() string { return e.m.Get(dent, "Name") }

func (e *Entry) GenericName() string { return e.m.Get(dent, "GenericName") }

func (e *Entry) NoDisplay() bool { return e.m.Bool(dent, "NoDisplay") }

func (e *Entry) Comment() string { return e.m.Get(dent, "Comment") }

func (e *Entry) Icon() string { return e.m.Get(dent, "Icon") }

func (e *Entry) Hidden() bool { return e.m.Bool(dent, "Hidden") }

func (e *Entry) OnlyShowIn() []string { return e.m.List(dent, "OnlyShowIn") }

func (e *Entry) NotShowIn() []string { return e.m.List(dent, "NotShowIn") }

func (e *Entry) DBusActivatable() bool { return e.m.Bool(dent, "DBusActivatable") }

func (e *Entry) TryExec() string { return e.m.Get(dent, "TryExec") }

func (e *Entry) Exec() string { return e.m.Get(dent, "Exec") }

func (e *Entry) Path() string { return e.m.Get(dent, "Path") }

func (e *Entry) Terminal() bool { return e.m.Bool(dent, "Terminal") }

func (e *Entry) MimeType() []string { return e.m.List(dent, "MimeType") }

func (e *Entry) Categories() []string { return e.m.List(dent, "Categories") }

func (e *Entry) Keywords() []string { return e.m.List(dent, "Keywords") }

func (e *Entry) StartupNotify() bool { return e.m.Bool(dent, "StartupNotify") }

func (e *Entry) StartupWMClass() string { return e.m.Get(dent, "StartupWMClass") }

func (e *Entry) URL() string { return e.m.Get(dent, "URL") }
