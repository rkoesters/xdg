xdg
===

Package xdg provides access to the XDG specs.

[![Build Status](https://travis-ci.org/rkoesters/xdg.svg?branch=master)](https://travis-ci.org/rkoesters/xdg)
[![Go Report Card](https://goreportcard.com/badge/github.com/rkoesters/xdg)](https://goreportcard.com/report/github.com/rkoesters/xdg)

Documentation
-------------

* [xdg](https://godoc.org/github.com/rkoesters/xdg) - Provides xdg.Open
  function to call `xdg-open` command.
* [xdg/basedir](https://godoc.org/github.com/rkoesters/xdg/basedir) -
  Provides access to the xdg basedir spec.
* [xdg/desktop](https://godoc.org/github.com/rkoesters/xdg/desktop) -
  Read desktop files (w/ localization support).
* [xdg/keyfile](https://godoc.org/github.com/rkoesters/xdg/keyfile) -
  Provides access to xdg key file format (w/ localization support).
* [xdg/trash](https://godoc.org/github.com/rkoesters/xdg/trash) -
  Provides access to xdg trash spec.
* [xdg/userdirs](https://godoc.org/github.com/rkoesters/xdg/userdirs) -
  Provides access to common user directories.

Testing
-------

Tests can be run with `go test`.

The tests for the [xdg/trash](trash) package expect the trash to exist
(`$XDG_DATA_HOME/Trash/files` (or `$HOME/.local/share/Trash/files` if
`$XDG_DATA_HOME` is undefined)).

The tests for the [xdg/userdirs](userdirs) package require the
`xdg-user-dir` command.

TODO
----

- autostart
- desktop.Launch (in progress on desktop-launch branch)
- trash.New (requires checking for a "sticky bit" on the filesystem of
  the drive on which the trash exists; in addition to other (easier)
  checks)

License
-------

See [LICENSE](LICENSE).
