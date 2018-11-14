package keyfile

import (
	"strings"
	"testing"
)

const localeStringExample = `
[Header 1]
Key=Value 0
Key[en_US]=Value 1
Key[en_UK]=Value 2
`

func TestLocaleString(t *testing.T) {
	kf, err := New(strings.NewReader(localeStringExample))
	if err != nil {
		t.Error(err)
	}
	t.Log(kf)

	s, err := kf.LocaleString("Header 1", "Key")
	if err != nil {
		t.Error(err)
	}
	if s != "Value 0" {
		t.Errorf("expected=Value 0 real=%v", s)
	}


	s, err = kf.LocaleStringWithLocale("Header 1", "Key", &Locale{})
	if err != nil {
		t.Error(err)
	}
	if s != "Value 0" {
		t.Errorf("expected=Value 0 real=%v", s)
	}

	locale, err := ParseLocale("en_US")
	if err != nil {
		t.Error(err)
	}
	s, err = kf.LocaleStringWithLocale("Header 1", "Key", locale)
	if err != nil {
		t.Error(err)
	}
	if s != "Value 1" {
		t.Errorf("expected=Value 1 real=%v", s)
	}

	locale, err = ParseLocale("en_UK")
	if err != nil {
		t.Error(err)
	}
	t.Log(locale)
	s, err = kf.LocaleStringWithLocale("Header 1", "Key", locale)
	if err != nil {
		t.Error(err)
	}
	if s != "Value 2" {
		t.Errorf("expected=Value 2 real=%v", s)
	}
}
