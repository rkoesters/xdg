package desktop

import (
	"errors"
	"github.com/rkoesters/xdg/ini"
	"io"
	"strings"
)

// TODO: Add methods to Entry: Launch(uri string...).

var (
	ErrMissingType = errors.New("missing entry type")
	ErrMissingName = errors.New("missing entry name")
	ErrMissingURL  = errors.New("missing entry url")
)

// Entry represents a desktop entry file.
type Entry struct {
	// The type of desktop entry. It can be: Application, Link, or
	// Directory.
	Type Type
	// The version of spec that the file conforms to.
	Version string

	// The real name of the desktop entry.
	Name string
	// A generic name, for example: Text Editor or Web Browser.
	GenericName string
	// A short comment that describes the desktop entry.
	Comment string
	// The name of an icon that should be used for this desktop
	// entry.  If it is not an absolute path, it should be searched
	// for using the Icon Theme Specification.
	Icon string
	// The URL for a Link type entry.
	URL string

	// Whether or not to display the file in menus.
	NoDisplay bool
	// Whether the use has deleted the desktop entry.
	Hidden bool
	// A list of desktop environments that the desktop entry should
	// only be shown in.
	OnlyShowIn []string
	// A list of desktop environments that the desktop entry should
	// not be shown in.
	NotShowIn []string

	// Whether DBus Activation is supported by this application.
	DBusActivatable bool
	// The path to an executable to test if the program is
	// installed.
	TryExec string
	// Program to execute. TODO: talk about arguments.
	Exec string
	// The path that should be the programs working directory.
	Path string
	// Whether the program should be run in a terminal window.
	Terminal bool

	// A slice of actions.
	Actions []*Action
	// A slice of mimetypes supported by this program.
	MimeType []string
	// A slice of categories that the desktop entry should be shown
	// in in a menu.
	Categories []string
	// A slice of keywords.
	Keywords []string

	// Whether the program will send a "remove" message when started
	// with the DESKTOP_STARTUP_ID env variable is set.
	// TODO: needs be explaination, I don't really know what it
	// means.
	StartupNotify bool
	// The string that the program will set as WM Class or WM name
	// hint.
	StartupWMClass string

	// Extended pairs. These are all of the key=value pairs in which
	// the key follows the format X-PRODUCT-KEY. For example,
	// accessing X-Unity-IconBackgroundColor can be done with:
	//
	//	entry.X["Unity"]["IconBackgroundColor"]
	//
	X map[string]map[string]string
}

const dent = "Desktop Entry"

// New reads an ini.Map formated file from r and returns an Entry that
// represents the Desktop file that was read.
func New(r io.Reader) (*Entry, error) {
	m, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	// Create the entry.
	e := &Entry{
		Type:            ParseType(m.String(dent, "Type")),
		Version:         m.String(dent, "Version"),
		Name:            m.String(dent, "Name"),
		GenericName:     m.String(dent, "GenericName"),
		Comment:         m.String(dent, "Comment"),
		Icon:            m.String(dent, "Icon"),
		URL:             m.String(dent, "URL"),
		NoDisplay:       m.Bool(dent, "NoDisplay"),
		Hidden:          m.Bool(dent, "Hidden"),
		OnlyShowIn:      m.List(dent, "OnlyShowIn"),
		NotShowIn:       m.List(dent, "NotShowIn"),
		DBusActivatable: m.Bool(dent, "DBusActivatable"),
		TryExec:         m.String(dent, "TryExec"),
		Exec:            m.String(dent, "Exec"),
		Path:            m.String(dent, "Path"),
		Terminal:        m.Bool(dent, "Terminal"),
		Actions:         getActions(m),
		MimeType:        m.List(dent, "MimeType"),
		Categories:      m.List(dent, "Categories"),
		Keywords:        m.List(dent, "Keywords"),
		StartupNotify:   m.Bool(dent, "StartupNotify"),
		StartupWMClass:  m.String(dent, "StartupWMClass"),
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
	for k, v := range m.M[dent] {
		a := strings.SplitN(k, "-", 3)
		if a[0] != "X" {
			continue
		}
		if e.X[a[1]] == nil {
			e.X[a[1]] = make(map[string]string)
		}
		e.X[a[1]][a[2]] = v
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
			Name: m.String(g, "Name"),
			Icon: m.String(g, "Icon"),
			Exec: m.String(g, "Exec"),
		})
	}
	return acts
}
