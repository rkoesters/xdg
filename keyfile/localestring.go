package keyfile

import (
	"fmt"
)

// LocaleString returns the value for group 'g' and key 'k' using the
// system's locale.
func (kf *KeyFile) LocaleString(g, k string) (string, error) {
	return kf.LocaleStringWithLocale(g, k, DefaultLocale())
}

// LocaleStringWithLocale returns the value for group 'g', key 'k', and
// locale 'l'.
func (kf *KeyFile) LocaleStringWithLocale(g, k string, l *Locale) (string, error) {
	for _, locale := range l.Variants() {
		key := fmt.Sprintf("%v[%v]", k, locale)
		if kf.ValueExists(g, key) {
			return kf.String(g, key)
		}
	}
	return kf.String(g, k)
}

// TODO LocaleStringList
// TODO LocaleStringListWithLocale
