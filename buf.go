package ical

import (
	"bufio"
	"io"
)

// buffer wraps bufio.Writer with some ical-specific helper methods.
type buffer bufio.Writer

func newBuffer(w io.Writer) *buffer {
	return (*buffer)(bufio.NewWriter(w))
}

const crnl = "\r\n"

func (b *buffer) WriteString(ss ...string) (n int, err error) {
	for _, s := range ss {
		if n, err = (*bufio.Writer)(b).WriteString(s); err != nil {
			return n, err
		}
	}
	return n, err
}

func (b *buffer) IfWriteString(predicate bool, ss ...string) (n int, err error) {
	if predicate {
		n, err = b.WriteString(ss...)
	}
	return
}

func (b *buffer) WriteLine(ss ...string) (err error) {
	if _, err = b.WriteString(ss...); err != nil {
		return err
	}
	_, err = b.WriteString(crnl)
	return err
}

func (b *buffer) IfWriteLine(predicate bool, ss ...string) (err error) {
	if predicate {
		err = b.WriteLine(ss...)
	}
	return err
}

func (b *buffer) Flush() error {
	return (*bufio.Writer)(b).Flush()
}
