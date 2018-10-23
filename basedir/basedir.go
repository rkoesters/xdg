// Package basedir provides access to XDG base directory spec. For more
// information, please see the spec:
// https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
package basedir

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	// Home is the user's home directory.
	Home string

	// DataHome is the path to the directory where user data files
	// should be written.
	DataHome string

	// ConfigHome is the path to the directory where user
	// configuration files should be written.
	ConfigHome string

	// CacheHome is the path to the directory where non-essential
	// (cached) data should be written.
	CacheHome string

	// RuntimeDir is the path to the directory where runtime files
	// should be placed.
	RuntimeDir string

	// DataDirs is a slice of paths that should be searched for data
	// files.
	DataDirs []string

	// ConfigDirs is a slice of paths that should be searched for
	// configuration files.
	ConfigDirs []string
)

func init() {
	Home = os.Getenv("HOME")
	if Home == "" {
		u, err := user.Current()
		if err == nil {
			Home = u.HomeDir
		} else {
			Home = filepath.Join(os.TempDir(), os.Args[0])
		}
	}

	DataHome = getpath("XDG_DATA_HOME", filepath.Join(Home, ".local/share"))
	ConfigHome = getpath("XDG_CONFIG_HOME", filepath.Join(Home, ".config"))
	CacheHome = getpath("XDG_CACHE_HOME", filepath.Join(Home, ".cache"))
	RuntimeDir = getpath("XDG_RUNTIME_DIR", CacheHome)
	DataDirs = getpathlist("XDG_DATA_DIRS", []string{"/usr/local/share", "/usr/share"})
	ConfigDirs = getpathlist("XDG_CONFIG_DIRS", []string{"/etc/xdg"})
}

func getpath(env, def string) string {
	path := os.Getenv(env)
	if path == "" || !filepath.IsAbs(path) {
		return def
	}
	return path
}

func getpathlist(env string, def []string) []string {
	paths := filepath.SplitList(os.Getenv(env))
	for i := 0; i < len(paths); i++ {
		// If the path isn't absolute, we need to ignore it.
		if !filepath.IsAbs(paths[i]) {
			paths = append(paths[:i], paths[i+1:]...)
		}
	}
	if len(paths) == 0 {
		return def
	}
	return paths
}
