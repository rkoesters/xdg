// Package basedir provides access to XDG base directory spec.
package basedir

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	// The home directory.
	Home string

	// The path to the directory where user data files should be
	// written.
	DataHome string

	// The path to the directory where user configuration files
	// should be written.
	ConfigHome string

	// The path to the directory where non-essential (cached) data
	// should be written.
	CacheHome string

	// The path to the directory where runtime files should be
	// placed.
	RuntimeDir string

	// A slice of paths that should be searched for data files.
	DataDirs []string

	// A slice of paths that should be searched for configuration
	// files.
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

	DataHome = getenv("XDG_DATA_HOME", filepath.Join(Home, ".local/share"))
	ConfigHome = getenv("XDG_CONFIG_HOME", filepath.Join(Home, ".config"))
	CacheHome = getenv("XDG_CACHE_HOME", filepath.Join(Home, ".cache"))
	RuntimeDir = getenv("XDG_RUNTIME_DIR", CacheHome)

	DataDirs = filepath.SplitList(os.Getenv("XDG_DATA_DIRS"))
	if len(DataDirs) == 0 {
		DataDirs = []string{"/usr/local/share", "/usr/share"}
	}

	ConfigDirs = filepath.SplitList(os.Getenv("XDG_CONFIG_DIRS"))
	if len(ConfigDirs) == 0 {
		ConfigDirs = []string{"/etc/xdg"}
	}
}

func getenv(env, def string) string {
	x := os.Getenv(env)
	if x == "" {
		return def
	}
	return x
}
