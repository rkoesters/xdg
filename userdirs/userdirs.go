// Package userdirs provides easy access to "well known" user
// directories. For more information, see:
// https://www.freedesktop.org/wiki/Software/xdg-user-dirs/
package userdirs

import (
	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/keyfile"
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

// New creates a new UserDirs struct by reading from the given
// io.Reader.
func New(r io.Reader) (*UserDirs, error) {
	kf, err := keyfile.New(r)
	if err != nil {
		return nil, err
	}

	dirs := new(UserDirs)

	dirs.Desktop, err = parse(kf.String("", "XDG_DESKTOP_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Documents, err = parse(kf.String("", "XDG_DOCUMENTS_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Download, err = parse(kf.String("", "XDG_DOWNLOAD_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Music, err = parse(kf.String("", "XDG_MUSIC_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Pictures, err = parse(kf.String("", "XDG_PICTURES_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.PublicShare, err = parse(kf.String("", "XDG_PUBLICSHARE_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Templates, err = parse(kf.String("", "XDG_TEMPLATES_DIR"))
	if err != nil {
		return nil, err
	}
	dirs.Videos, err = parse(kf.String("", "XDG_VIDEOS_DIR"))
	if err != nil {
		return nil, err
	}

	return dirs, nil
}

// parse takes a given string and returns it as a path.
func parse(s string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	s = strings.Trim(s, "\"")
	if strings.HasPrefix(s, "$HOME") {
		s = filepath.Join(basedir.Home, strings.TrimPrefix(s, "$HOME"))
	}
	if s == "" {
		s = basedir.Home
	}
	return s, nil
}
