package parameter

import (
	"github.com/rickb777/ical2/ics"
	"strings"
)

// Parameters holds a set of key-value parameters.
type Parameters []Parameter

// WriteTo serialises the parameters in iCal ics format to the writer.
func (pp Parameters) WriteTo(w ics.StringWriter) error {
	for _, p := range pp {
		w.WriteByte(';')
		p.WriteTo(w)
	}

	return nil
}

// RemoveByKey removes all parameters with a key (or keys) from the list.
func (pp Parameters) RemoveByKey(key ...string) Parameters {
	for i := 0; i < len(pp); i++ {
		for _, k := range key {
			// the length might have changed: need to check it again
			if i < len(pp) {
				k = strings.ToUpper(k)
				if pp[i].Key == k {
					pp = append(pp[:i], pp[i+1:]...)
				}
			}
		}
	}
	return pp
}

// Prepend appends a parameter (or parameters), ensuring that keys remain unique.
func (pp Parameters) Prepend(ps ...Parameter) Parameters {
	for _, p := range ps {
		p.Key = strings.ToUpper(p.Key)
		pp = append(Parameters{p}, pp.RemoveByKey(p.Key)...)
	}
	return pp
}

// Append appends a parameter (or parameters), ensuring that keys remain unique.
func (pp Parameters) Append(ps ...Parameter) Parameters {
	for _, p := range ps {
		p.Key = strings.ToUpper(p.Key)
		pp = append(pp.RemoveByKey(p.Key), p)
	}
	return pp
}
