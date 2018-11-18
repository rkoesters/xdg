package keyfile

import (
	"bytes"
	"os"
)

// Locale represents a locale for use in parsing translatable strings.
type Locale struct {
	lang     string
	country  string
	encoding string
	modifier string
}

var defaultLocale *Locale

// DefaultLocale returns the locale specified by the environment.
func DefaultLocale() *Locale {
	if defaultLocale == nil {
		var val string
		var err error

		if val = os.Getenv("LANGUAGE"); val != "" {
			defaultLocale, err = ParseLocale(val)
		} else if val = os.Getenv("LC_ALL"); val != "" {
			defaultLocale, err = ParseLocale(val)
		} else if val = os.Getenv("LC_MESSAGES"); val != "" {
			defaultLocale, err = ParseLocale(val)
		} else if val = os.Getenv("LANG"); val != "" {
			defaultLocale, err = ParseLocale(val)
		}
		if err != nil || defaultLocale == nil {
			defaultLocale = &Locale{}
		}
	}
	return defaultLocale
}

// ParseLocale parses a locale in the format:
//
// 	lang_COUNTRY.ENCODING@MODIFIER
//
// and returns a Locale struct.
func ParseLocale(s string) (*Locale, error) {
	var buf bytes.Buffer

	i := 0
	l := new(Locale)

	// lang
	for i < len(s) && s[i] != '_' && s[i] != '.' && s[i] != '@' {
		buf.WriteByte(s[i])
		i++
	}
	l.lang = buf.String()
	buf.Reset()

	// COUNTRY
	if i < len(s) && s[i] == '_' {
		i++
		for i < len(s) && s[i] != '.' && s[i] != '@' {
			buf.WriteByte(s[i])
			i++
		}
		l.country = buf.String()
		buf.Reset()
	}

	// ENCODING
	if i < len(s) && s[i] == '.' {
		i++
		for i < len(s) && s[i] != '@' {
			buf.WriteByte(s[i])
			i++
		}
		l.encoding = buf.String()
		buf.Reset()
	}

	// MODIFIER
	if i < len(s) && s[i] == '@' {
		i++
		for i < len(s) {
			buf.WriteByte(s[i])
			i++
		}
		l.modifier = buf.String()
	}

	return l, nil
}

// String returns the given locale as a formatted string. The returned
// string is in the same format expected by ParseLocale.
func (l *Locale) String() string {
	var buf bytes.Buffer

	buf.WriteString(l.lang)

	if l.country != "" {
		buf.WriteRune('_')
		buf.WriteString(l.country)
	}

	if l.encoding != "" {
		buf.WriteRune('.')
		buf.WriteString(l.encoding)
	}

	if l.modifier != "" {
		buf.WriteRune('@')
		buf.WriteString(l.modifier)
	}

	return buf.String()
}

// Variants returns a sorted slice of *Locales that should be checked
// for when resolving a localestring.
func (l *Locale) Variants() []*Locale {
	variants := make([]*Locale, 0)

	if l.lang != "" && l.country != "" && l.modifier != "" {
		variants = append(variants, &Locale{
			lang:     l.lang,
			country:  l.country,
			modifier: l.modifier,
		})
	}

	if l.lang != "" && l.country != "" {
		variants = append(variants, &Locale{
			lang:    l.lang,
			country: l.country,
		})
	}

	if l.lang != "" && l.modifier != "" {
		variants = append(variants, &Locale{
			lang:     l.lang,
			modifier: l.modifier,
		})
	}

	if l.lang != "" {
		variants = append(variants, &Locale{
			lang: l.lang,
		})
	}

	return variants
}
