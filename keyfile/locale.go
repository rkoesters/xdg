package keyfile

import (
	"bytes"
	"fmt"
)

// Locale represents a locale for use in parsing localestrings.
type Locale struct {
	lang     string
	country  string
	encoding string
	modifier string
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

	for i < len(s) && s[i] != '_' && s[i] != '.' && s[i] != '@' {
		buf.WriteByte(s[i])
		i++
	}
	l.lang = buf.String()
	buf.Reset()
	if i < len(s) && s[i] == '_' {
		i++
		for i < len(s) && s[i] != '.' && s[i] != '@' {
			buf.WriteByte(s[i])
			i++
		}
		l.country = buf.String()
		buf.Reset()
	}
	if i < len(s) && s[i] == '.' {
		i++
		for i < len(s) && s[i] != '@' {
			buf.WriteByte(s[i])
			i++
		}
		l.encoding = buf.String()
		buf.Reset()
	}
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
	if l.lang != "" {
		buf.WriteString(l.lang)
	}
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

// Variants returns a sorted slice of locale variation strings that
// should be checked for when resolving a localestring.
func (l *Locale) Variants() []string {
	variants := []string{}

	if l.lang != "" && l.country != "" && l.modifier != "" {
		variants = append(variants, fmt.Sprintf("%v_%v@%v", l.lang, l.country, l.modifier))
	}
	if l.lang != "" && l.country != "" {
		variants = append(variants, fmt.Sprintf("%v_%v", l.lang, l.country))
	}
	if l.lang != "" && l.modifier != "" {
		variants = append(variants, fmt.Sprintf("%v@%v", l.lang, l.modifier))
	}
	if l.lang != "" {
		variants = append(variants, l.lang)
	}
	return variants
}
