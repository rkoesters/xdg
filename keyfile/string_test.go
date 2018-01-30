package keyfile

import (
	"testing"
)

func TestUnescapeString(t *testing.T) {
	goodData := map[string]string{
		"":              "",
		"asdf":          "asdf",
		"asdf\\\\asdf":  "asdf\\asdf",
		"\\sasdf":       " asdf",
		"hello\\nworld": "hello\nworld",
		"asdf\\tasdf":   "asdf\tasdf",
		"asdf\\rasdf":   "asdf\rasdf",
	}
	for input, expected := range goodData {
		str, err := unescapeString(input)
		if err != nil {
			t.Error(err)
		}
		if str != expected {
			t.Errorf("error escapeing '%v'", input)
		}
	}

	badData := map[string]error{
		"\\":  ErrUnexpectedEndOfString,
		"\\p": ErrBadEscapeSequence,
	}
	for input, expected := range badData {
		_, err := unescapeString(input)
		if err != expected {
			t.Error(err)
		}
	}
}
