package ical

import (
	"io"
	"bufio"
)

// buffer wraps bufio.Writer with some ical-specific helper methods.
type buffer bufio.Writer

func newBuffer(w io.Writer) *buffer {
	return (*buffer)(bufio.NewWriter(w))
}

const crnl = "\r\n"

func (b *buffer) WriteString(s string) (int, error) {
	return (*bufio.Writer)(b).WriteString(s)
}

func (b *buffer) WriteLine(ss ...string) error {
	var err error
	for _, s := range ss {
		if _, err = b.WriteString(s); err != nil {
			return err
		}
	}
	_, err = b.WriteString(crnl)
	return err
}

func (b *buffer) Flush() error {
	return (*bufio.Writer)(b).Flush()
}
