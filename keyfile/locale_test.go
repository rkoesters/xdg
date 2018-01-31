package keyfile

import (
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
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}

	localeStr = "en_US.UTF-8@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}

	localeStr = "en_US@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}

	localeStr = "en_US"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}

	localeStr = "en@mod"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}

	localeStr = "en"
	locale, err = ParseLocale(localeStr)
	if err != nil {
		t.Error(err)
	}
	if localeStr != locale.String() {
		t.Errorf("'%v' != '%v'", localeStr, locale.String())
	}
}

func TestDefaultLocale(t *testing.T) {
	locale := DefaultLocale()
	if locale == nil {
		t.FailNow()
	}
	t.Log(locale.String())
}
