package keyfile

import (
	"reflect"
	"testing"
)

func TestParseLocale(t *testing.T) {
	var localeStr string
	var locale *Locale
	var err error

	localeStr = ""
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}

	localeStr = "en_US.UTF-8@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}

	localeStr = "en_US@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}

	localeStr = "en_US"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}

	localeStr = "en@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}

	localeStr = "en"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale)
	}
}

func TestDefaultLocale(t *testing.T) {
	locale := DefaultLocale()
	t.Log(locale)
	if locale == nil {
		t.Fail()
	}
}

func TestBadLocale(t *testing.T) {
	_, err := ParseLocale("_US")
	if err != ErrBadLocaleFormat {
		t.Fail()
	}

	_, err = ParseLocale(".UTF-8")
	if err != ErrBadLocaleFormat {
		t.Fail()
	}

	_, err = ParseLocale("@Latn")
	if err != ErrBadLocaleFormat {
		t.Fail()
	}
}

func TestSpecialLocale(t *testing.T) {
	blank, err := ParseLocale("")
	if err != nil {
		t.Fail()
	}

	c, err := ParseLocale("C")
	if err != nil || !reflect.DeepEqual(c, blank) {
		t.Fail()
	}

	posix, err := ParseLocale("POSIX")
	if err != nil || !reflect.DeepEqual(posix, blank) {
		t.Fail()
	}
}
