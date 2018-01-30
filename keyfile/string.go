package keyfile

import (
	"bytes"
)

// String returns the value with the given group and key. This function
// is here because underlying data structure of KeyFile may change.
func (kf *KeyFile) String(g, k string) (string, error) {
	return unescapeString(kf.Value(g, k))
}

func unescapeString(s string) (string, error) {
	var buf bytes.Buffer
	var isEscaped bool
	var err error

	for _, r := range s {
		if isEscaped {
			switch r {
			case 's':
				_, err = buf.WriteRune(' ')
			case 'n':
				_, err = buf.WriteRune('\n')
			case 't':
				_, err = buf.WriteRune('\t')
			case 'r':
				_, err = buf.WriteRune('\r')
			case '\\':
				_, err = buf.WriteRune('\\')
			default:
				err = ErrBadEscapeSequence
			}

			if err != nil {
				return "", err
			}

			isEscaped = false
		} else {
			if r == '\\' {
				isEscaped = true
			} else {
				buf.WriteRune(r)
			}
		}
	}
	if isEscaped {
		return "", ErrUnexpectedEndOfString
	}
	return buf.String(), nil
}
