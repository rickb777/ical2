// Package ics provides low-level I/O support for the ical2 api.
// Notably, it implements the iCalendar line-folding algorithm.
package ics

import (
	"bufio"
	"io"
)

// Valuer holds an iCalendar property value.
type Valuer interface {
	IsDefined() bool
	WriteTo(w StringWriter) error
}

// IsDefined tests whether a valuer is defined.
func IsDefined(v Valuer) bool {
	return v != nil && v.IsDefined()
}

//-------------------------------------------------------------------------------------------------

// StringWriter provides an iCalendar writing interface.
type StringWriter interface {
	io.Writer
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
}

// Flusher is implemented by buffers that need to be flushed.
type Flusher interface {
	Flush() error
}

//-------------------------------------------------------------------------------------------------

// MaxLineLength is the maximum length of lines emitted by the line-folding
// writer. It defaults to 75.
var MaxLineLength = 75

// foldWriter implements the max-75 character line folding.
// It also collapses any i/o errors.
type foldWriter struct {
	w          *bufio.Writer
	n          int
	err        error
	lineEnding string // usually "\r\n"
}

// NewFoldWriter returns a StringWriter wrapping an io.Writer that folds
// lines at the 75th column (determined by MaxLineLength). The line
// ending defaults to "\r\n" if blank.
//
// The data is buffered; therefore Flush must be called at the end.
func NewFoldWriter(w io.Writer, lineEnding string) StringWriter {
	if lineEnding == "" {
		lineEnding = "\r\n"
	}
	return &foldWriter{w: bufio.NewWriter(w), lineEnding: lineEnding}
}

func (fw *foldWriter) Write(s []byte) (i int, err error) {
	if fw.err != nil {
		return
	}

	remaining := MaxLineLength - fw.n

	for i = 0; i < len(s) && fw.err == nil; i++ {
		c := s[i]
		if remaining < 1 {
			fw.wrapLine()
			remaining = MaxLineLength
		}

		fw.err = fw.w.WriteByte(c)
		fw.n++
		remaining--
	}

	return i, fw.err
}

func (fw *foldWriter) WriteByte(c byte) error {
	if fw.err != nil {
		return fw.err
	}

	if fw.n >= MaxLineLength {
		fw.wrapLine()
	}

	fw.err = fw.w.WriteByte(c)
	fw.n++
	return fw.err
}

func (fw *foldWriter) wrapLine() error {
	if fw.err == nil {
		fw.newline()
		fw.err = fw.w.WriteByte(' ')
		fw.n = 1
	}

	return fw.err
}

func (fw *foldWriter) WriteString(s string) (n int, err error) {
	// treat s as a sequence of bytes, not runes
	return fw.Write([]byte(s))
}

func (fw *foldWriter) newline() error {
	if fw.err == nil {
		_, fw.err = fw.w.WriteString(fw.lineEnding)
		fw.n = 0
	}

	return fw.err
}

func (fw *foldWriter) Flush() error {
	e := fw.w.Flush()
	if e != nil {
		return e
	}
	return fw.err
}

//-------------------------------------------------------------------------------------------------

// Buffer wraps bufio.Writer with some iCalendar-specific helper methods.
// It folds long lines to meet the iCalendar max-75 characters per line limit.
// If coallesces errors so they don't have to be checked after every method;
// it is sufficient to check once at the end.
type Buffer struct {
	fw *foldWriter
}

// NewBuffer constructs a Buffer that wraps some Writer. The lineEnding can be
// "" or "\r\n" for normal iCalendar formatting, or "\n" in other cases.
func NewBuffer(w io.Writer, lineEnding string) *Buffer {
	return &Buffer{NewFoldWriter(w, lineEnding).(*foldWriter)}
}

// WriteString writes the string supplied.
func (b *Buffer) WriteString(s string) error {
	b.fw.WriteString(s)
	return b.fw.err
}

// WriteLine writes a string and a newline to the buffer. In order to comply with
// the line folding rules, you should use this method and not write strings
// containing additional newline '\n' or carriage return '\r' characters.
func (b *Buffer) WriteLine(s string) error {
	b.WriteString(s)
	return b.fw.newline()
}

// WriteValuerLine conditionally writes a valuer along with its property name. If
// the predicate is false, nothing is written.
func (b *Buffer) WriteValuerLine(predicate bool, label string, v Valuer) error {
	if !predicate {
		return b.fw.err // skip
	}

	b.WriteString(label)
	v.WriteTo(b.fw)
	return b.fw.newline()
}

// Flush flushes the buffered data through to the underlying writer. This must be
// called at least once, at the end.
func (b *Buffer) Flush() error {
	return b.fw.Flush()
}
