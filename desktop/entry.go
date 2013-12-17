package desktop

import (
	"errors"
	"fmt"
	"github.com/rkoesters/xdg/ini"
	"io"
)

var (
	ErrMissingType = errors.New("missing entry type")
	ErrMissingName = errors.New("missing entry name")
	ErrMissingURL  = errors.New("missing entry url")
)

// Entry represents a desktop entry file.
//
// TODO: extensively comment this struct.
type Entry struct {
	Type    Type
	Version string

	Name        string
	GenericName string
	Comment     string
	Icon        string
	URL         string

	NoDisplay  bool
	Hidden     bool
	OnlyShowIn []string
	NotShowIn  []string

	DBusActivatable bool
	TryExec         string
	Exec            string
	Path            string
	Terminal        bool

	Actions    []*Action
	MimeType   []string
	Categories []string
	Keywords   []string

	StartupNotify  bool
	StartupWMClass string

	// Extended pairs (X-PRODUCT-Key).
	X map[string]map[string]string
}

const dent = "Desktop Entry"

func New(r io.Reader) (*Entry, error) {
	m, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	// Create the entry.
	e := &Entry{
		Type:            ParseType(m.Get(dent, "Type")),
		Version:         m.Get(dent, "Version"),
		Name:            m.Get(dent, "Name"),
		GenericName:     m.Get(dent, "GenericName"),
		Comment:         m.Get(dent, "Comment"),
		Icon:            m.Get(dent, "Icon"),
		URL:             m.Get(dent, "URL"),
		NoDisplay:       m.Bool(dent, "NoDisplay"),
		Hidden:          m.Bool(dent, "Hidden"),
		OnlyShowIn:      m.List(dent, "OnlyShowIn"),
		NotShowIn:       m.List(dent, "NotShowIn"),
		DBusActivatable: m.Bool(dent, "DBusActivatable"),
		TryExec:         m.Get(dent, "TryExec"),
		Exec:            m.Get(dent, "Exec"),
		Path:            m.Get(dent, "Path"),
		Terminal:        m.Bool(dent, "Terminal"),
		Actions:         getActions(m),
		MimeType:        m.List(dent, "MimeType"),
		Categories:      m.List(dent, "Categories"),
		Keywords:        m.List(dent, "Keywords"),
		StartupNotify:   m.Bool(dent, "StartupNotify"),
		StartupWMClass:  m.Get(dent, "StartupWMClass"),
		X:               make(map[string]map[string]string),
	}

	// Validate the entry.
	if e.Type == None {
		return nil, ErrMissingType
	}
	if e.Type > None && e.Type < Unknown && e.Name == "" {
		return nil, ErrMissingName
	}
	if e.Type == Link && e.URL == "" {
		return nil, ErrMissingURL
	}

	// Search for extended keys.
	var prod string
	var key string
	// TODO: add support for extend groups.
	for _, gv := range m.M {
		for k, v := range gv {
			n, _ := fmt.Sscanf(k, "X-%v-%v", &prod, &key)
			if n != 2 {
				continue
			}
			e.X[prod][key] = v
		}
	}

	return e, nil
}

// Action is an Action group.
type Action struct {
	Name string
	Icon string
	Exec string
}

func getActions(m *ini.Map) []*Action {
	var acts []*Action

	for _, a := range m.List(dent, "Actions") {
		g := "Desktop Action " + a
		acts = append(acts, &Action{
			Name: m.Get(g, "Name"),
			Icon: m.Get(g, "Icon"),
			Exec: m.Get(g, "Exec"),
		})
	}
	return acts
}
