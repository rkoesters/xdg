package desktop

import (
	"errors"
	"github.com/rkoesters/xdg/keyfile"
	"io"
	"strings"
)

// TODO: Add methods to Entry: Launch(uri string...).

var (
	// ErrMissingType means that the desktop entry is missing the
	// Type key, which is always required.
	ErrMissingType = errors.New("missing entry type")

	// ErrMissingName means that the desktop entry is missing the
	// Name key, which is required by the types Application, Link,
	// and Directory.
	ErrMissingName = errors.New("missing entry name")

	// ErrMissingURL means that the desktop entry is missing the URL
	// key, which is required by the type Link.
	ErrMissingURL = errors.New("missing entry url")
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

// New reads an keyfile.Map formated file from r and returns an Entry that
// represents the Desktop file that was read.
func New(r io.Reader) (*Entry, error) {
	kf, err := keyfile.New(r)
	if err != nil {
		return nil, err
	}

	// Create the entry.
	e := new(Entry)

	e.Type = ParseType(kf.Value(dent, "Type"))
	e.Version, err = kf.String(dent, "Version")
	if err != nil {
		return nil, err
	}
	e.Name, err = kf.String(dent, "Name")
	if err != nil {
		return nil, err
	}
	e.GenericName, err = kf.String(dent, "GenericName")
	if err != nil {
		return nil, err
	}
	e.Comment, err = kf.String(dent, "Comment")
	if err != nil {
		return nil, err
	}
	e.Icon, err = kf.String(dent, "Icon")
	if err != nil {
		return nil, err
	}
	e.URL, err = kf.String(dent, "URL")
	if err != nil {
		return nil, err
	}
	e.NoDisplay, err = kf.Bool(dent, "NoDisplay")
	if err != nil {
		return nil, err
	}
	e.Hidden, err = kf.Bool(dent, "Hidden")
	if err != nil {
		return nil, err
	}
	e.OnlyShowIn, err = kf.StringList(dent, "OnlyShowIn")
	if err != nil {
		return nil, err
	}
	e.NotShowIn, err = kf.StringList(dent, "NotShowIn")
	if err != nil {
		return nil, err
	}
	e.DBusActivatable, err = kf.Bool(dent, "DBusActivatable")
	if err != nil {
		return nil, err
	}
	e.TryExec, err = kf.String(dent, "TryExec")
	if err != nil {
		return nil, err
	}
	e.Exec, err = kf.String(dent, "Exec")
	if err != nil {
		return nil, err
	}
	e.Path, err = kf.String(dent, "Path")
	if err != nil {
		return nil, err
	}
	e.Terminal, err = kf.Bool(dent, "Terminal")
	if err != nil {
		return nil, err
	}
	e.Actions, err = getActions(kf)
	if err != nil {
		return nil, err
	}
	e.MimeType, err = kf.StringList(dent, "MimeType")
	if err != nil {
		return nil, err
	}
	e.Categories, err = kf.StringList(dent, "Categories")
	if err != nil {
		return nil, err
	}
	e.Keywords, err = kf.StringList(dent, "Keywords")
	if err != nil {
		return nil, err
	}
	e.StartupNotify, err = kf.Bool(dent, "StartupNotify")
	if err != nil {
		return nil, err
	}
	e.StartupWMClass, err = kf.String(dent, "StartupWMClass")
	if err != nil {
		return nil, err
	}

	e.X = make(map[string]map[string]string)

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
	for _, k := range kf.Keys(dent) {
		a := strings.SplitN(k, "-", 3)
		if a[0] != "X" || len(a) < 3 {
			continue
		}
		if e.X[a[1]] == nil {
			e.X[a[1]] = make(map[string]string)
		}
		e.X[a[1]][a[2]] = kf.Value(dent, k)
	}

	return e, nil
}

// Action is an Action group.
type Action struct {
	Name string
	Icon string
	Exec string
}

func getActions(kf *keyfile.KeyFile) ([]*Action, error) {
	var acts []*Action
	var act *Action
	var err error
	var list []string

	list, err = kf.StringList(dent, "Actions")
	if err != nil {
		return nil, err
	}
	for _, a := range list {
		g := "Desktop Action " + a

		act = new(Action)

		act.Name, err = kf.String(g, "Name")
		if err != nil {
			return nil, err
		}
		act.Icon, err = kf.String(g, "Icon")
		if err != nil {
			return nil, err
		}
		act.Exec, err = kf.String(g, "Exec")
		if err != nil {
			return nil, err
		}

		acts = append(acts, act)
	}
	return acts, nil
}
