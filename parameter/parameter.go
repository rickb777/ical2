// Package parameter handles iCal parameters. The key is required to be
// case-insensitive. Here, this this generally implemented by using upper-case
// keys, a pattern that must be carefully followed if Parameter structs are
// created on the fly.
//
// All the factory functions in this package that return a Parameter will use
// an upper-case key.
//
// See https://tools.ietf.org/html/rfc5545#section-3.2
package parameter

import (
	"github.com/rickb777/ical2/ics"
	"strings"
)

// Parameter holds an iCal parameter. The key must be uppercase, this being a
// pattern that simplifies the requirement for keys to be case-insensitive.
//
// For most parameters, the value is singular, i.e. there is exactly one string.
// There are several exceptions.
type Parameter struct {
	Key   string
	Value []string
}

// Equals tests whether two parameters have the same key and the same value(s).
func (p Parameter) Equals(q Parameter) bool {
	if !strings.EqualFold(p.Key, q.Key) || len(p.Value) != len(q.Value) {
		return false
	}
	for i, v := range p.Value {
		if v != q.Value[i] {
			return false
		}
	}
	return true
}

// WriteTo serialises the parameter in iCal ics format to the writer.
// Parameters with multiple values are serialised using a comma-separated list.
//
// Parameters with values containing a COLON character, a SEMICOLON character
// or a COMMA character are placed in quoted text.
func (p Parameter) WriteTo(w ics.StringWriter) error {
	w.WriteString(p.Key)
	w.WriteByte('=')

	needQuotes := false
	for _, v := range p.Value {
		needQuotes = needQuotes || strings.IndexAny(v, ":;,") >= 0
	}

	if needQuotes {
		sep := ""
		for _, v := range p.Value {
			w.WriteString(sep)
			w.WriteByte('"')
			w.WriteString(v)
			w.WriteByte('"')
			sep = ","
		}
	} else {
		w.WriteString(strings.Join(p.Value, ","))
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

// Single returns a Parameter with a single string value.
func Single(k, v string) Parameter {
	return Parameter{k, []string{v}}
}

func either(key string, predicate bool, yes, no string) Parameter {
	v := no
	if predicate {
		v = yes
	}
	return Single(key, v)
}
