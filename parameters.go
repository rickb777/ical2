package ical2

// Parameters holds a set of key-value parameters.
type Parameters []Parameter

func (pp Parameters) WriteTo(w StringWriter) error {
	for _, p := range pp {
		w.WriteByte(';')
		w.WriteString(p.Key)
		w.WriteByte('=')
		w.WriteString(p.Value)
	}

	return nil
}

// Remove removes all parameters with a key (or keys) from the list.
func (pp Parameters) Remove(key ...string) Parameters {
	for i := 0; i < len(pp); i++ {
		for _, k := range key {
			if pp[i].Key == k {
				pp = append(pp[:i], pp[i+1:]...)
			}
		}
	}
	return pp
}

// Prepend appends a parameter (or parameters), ensuring that keys remain unique.
func (pp Parameters) Prepend(ps ...Parameter) Parameters {
	for _, p := range ps {
		pp = append(Parameters{p}, pp.Remove(p.Key)...)
	}
	return pp
}

// Append appends a parameter (or parameters), ensuring that keys remain unique.
func (pp Parameters) Append(ps ...Parameter) Parameters {
	for _, p := range ps {
		pp = append(pp.Remove(p.Key), p)
	}
	return pp
}

//-------------------------------------------------------------------------------------------------

