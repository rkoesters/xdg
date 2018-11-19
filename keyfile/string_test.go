package keyfile

import (
	"strings"
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

func TestBadStringList(t *testing.T) {
	const badList = `list=asdf;asdf\`

	kf, err := New(strings.NewReader(badList))
	if err != nil {
		t.Fail()
	} else {
		_, err = kf.StringList("", "list")
		if err != ErrUnexpectedEndOfString {
			t.Fail()
		}
	}

	const badStringInList = `list=asdf;as\qasd;`

	kf, err = New(strings.NewReader(badStringInList))
	if err != nil {
		t.Fail()
	} else {
		_, err = kf.StringList("", "list")
		if err != ErrBadEscapeSequence {
			t.Fail()
		}
	}
}
