package desktop

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"unicode"
)

// TODO de.DBusActivatable?
// TODO de.Action[].Launch?

// Launch TODO
func (de *Entry) Launch(uris ...string) error {
	argvs, err := de.expandExec(uris...)
	if err != nil {
		return err
	}

	for _, argv := range argvs {
		cmd := exec.Command(argv[0], argv[1:]...)

		cmd.Dir = de.Path

		// TODO should these be left as is, removed, or made
		// TODO configurable?
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Alternative: Use Start and Wait with a callback
		// (maybe have separate LaunchAsync method?)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// expandExec returns a slice of argvs. The slice will have 1 element
// unless Exec contains %f or %u in which case it'll return a copy of
// argv for each uri.
func (de *Entry) expandExec(args ...string) ([][]string, error) {
	ret := make([][]string, 0)

	// singleIteration defaults to true, but will be set to false if
	// we find '%f' or '%u'
	singleIteration := true
	for i := 0; (i == 0 || !singleIteration) && (i < len(args) || len(args) == 0); i++ {
		ret = append(ret, make([]string, 0))

		if de.Terminal {
			term, err := getTerminalPrefix()
			if err != nil {
				return nil, err
			}
			ret[i] = append(ret[i], term...)
		}

		var buf bytes.Buffer
		var isEscaped, isQuoted, isSpecialField bool
		var quoteRune rune
		var err error

		// Expand the Exec key.
		for _, r := range de.Exec {
			if isEscaped {
				switch r {
				// TODO what escape sequences should we
				// accept? Do we even process escape
				// sequences here?
				case '`':
					fallthrough
				case '$':
					fallthrough
				case '\\':
					fallthrough
				case '"':
					buf.WriteRune(r)
				default:
					err = errors.New("bad escape sequence")
				}
				if err != nil {
					return nil, err
				}
				isEscaped = false
			} else if isQuoted {
				// TODO what are all the quoting rules?
				switch r {
				case quoteRune:
					isQuoted = false
				default:
					buf.WriteRune(r)
				}
			} else if isSpecialField {
				switch r {
				case '%':
					buf.WriteRune('%')
				case 'f':
					// TODO copy remote uris to
					// local machine.
					singleIteration = false
					if len(args) > i {
						ret[i] = append(ret[i], args[i])
					}
				case 'F':
					// TODO copy remote uris to
					// local machine.
					ret[i] = append(ret[i], args...)
				case 'u':
					singleIteration = false
					if len(args) > i {
						ret[i] = append(ret[i], args[i])
					}
				case 'U':
					ret[i] = append(ret[i], args...)
				case 'd':
					// deprecated
				case 'D':
					// deprecated
				case 'n':
					// deprecated
				case 'N':
					// deprecated
				case 'i':
					if buf.Len() != 0 {
						// TODO how should we
						// handle this?
						return nil, errors.New("%i must be on its own")
					}
					if de.Icon != "" {
						ret[i] = append(ret[i], "--icon")
						ret[i] = append(ret[i], de.Icon)
					}
				case 'c':
					buf.WriteString(de.Name)
				case 'k':
					buf.WriteString(de.filename)
				case 'v':
					// deprecated
				case 'm':
					// deprecated
				default:
					err = errors.New("bad special field code")
				}
				if err != nil {
					return nil, err
				}
				isSpecialField = false
			} else {
				switch {
				case r == '"':
					isQuoted = true
					quoteRune = r
				case r == '\\':
					isEscaped = true
				case r == '%':
					isSpecialField = true
				case unicode.IsSpace(r):
					arg := buf.String()
					buf.Reset()

					if arg != "" {
						ret[i] = append(ret[i], arg)
					}
				default:
					buf.WriteRune(r)
				}
			}
		}
		if isEscaped || isQuoted || isSpecialField {
			return nil, errors.New("unexpected end of string")
		}
		arg := buf.String()
		buf.Reset()

		if arg != "" {
			ret[i] = append(ret[i], arg)
		}
	}

	return ret, nil
}

func getTerminalPrefix() ([]string, error) {
	terminals := [][]string{
		{"gnome-terminal", "-x"},
		{"konsole", "-e"},
		{"io.elementary.terminal", "-e"},
		{"pantheon-terminal", "-e"},
		{"nxterm", "-e"},
		{"color-xterm", "-e"},
		{"rxvt", "-e"},
		{"dtterm", "-e"},
		{"xterm", "-e"},
	}

	for _, term := range terminals {
		_, err := exec.LookPath(term[0])
		if err == nil {
			return term, nil
		}
	}
	return nil, errors.New("no terminal emulator found")
}
