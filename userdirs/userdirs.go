// Package userdirs provides easy access to "well known" user
// directories.
package userdirs

import (
	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/ini"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// The default set of userdirs. Most people will only need to use this.
var (
	Desktop     string
	Documents   string
	Download    string
	Music       string
	Pictures    string
	PublicShare string
	Templates   string
	Videos      string
)

func init() {
	f, err := os.Open(filepath.Join(basedir.ConfigHome, "user-dirs.dirs"))
	if err != nil {
		return
	}
	defer f.Close()

	dirs, err := New(f)
	if err != nil {
		return
	}

	Desktop = dirs.Desktop
	Documents = dirs.Documents
	Download = dirs.Download
	Music = dirs.Music
	Pictures = dirs.Pictures
	PublicShare = dirs.PublicShare
	Templates = dirs.Templates
	Videos = dirs.Videos
}

// UserDirs is a set of user directories that are common in graphical
// environments.
type UserDirs struct {
	Desktop     string
	Documents   string
	Download    string
	Music       string
	Pictures    string
	PublicShare string
	Templates   string
	Videos      string
}

// New creates a new UserDirs struct buy reading from the given
// io.Reader.
func New(r io.Reader) (*UserDirs, error) {
	m, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	return &UserDirs{
		Desktop:     parse(m.String("", "XDG_DESKTOP_DIR")),
		Documents:   parse(m.String("", "XDG_DOCUMENTS_DIR")),
		Download:    parse(m.String("", "XDG_DOWNLOAD_DIR")),
		Music:       parse(m.String("", "XDG_MUSIC_DIR")),
		Pictures:    parse(m.String("", "XDG_PICTURES_DIR")),
		PublicShare: parse(m.String("", "XDG_PUBLICSHARE_DIR")),
		Templates:   parse(m.String("", "XDG_TEMPLATES_DIR")),
		Videos:      parse(m.String("", "XDG_VIDEOS_DIR")),
	}, nil
}

// parse takes a given string and returns it as a path.
func parse(s string) string {
	s = strings.Trim(s, "\"")
	if strings.HasPrefix(s, "$HOME") {
		return filepath.Join(basedir.Home, strings.TrimPrefix(s, "$HOME"))
	}
	return s
}
