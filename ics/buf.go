package ics

import (
	"bufio"
	"io"
)

// Valuer holds an iCal value.
type Valuer interface {
	IsDefined() bool
	WriteTo(w StringWriter) error
}

// IsDefined tests whether a valuer is defined.
func IsDefined(v Valuer) bool {
	return v != nil && v.IsDefined()
}

//-------------------------------------------------------------------------------------------------

// StringWriter provides an iCal writing interface.
type StringWriter interface {
	io.Writer
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
}

//-------------------------------------------------------------------------------------------------

// foldWriter implements the max-75 character line folding.
// It also collapses any i/o errors.
type foldWriter struct {
	w          *bufio.Writer
	n          int
	err        error
	lineEnding string // usually "\r\n"
}

func newFoldWriter(w io.Writer, lineEnding string) *foldWriter {
	if lineEnding == "" {
		lineEnding = "\r\n"
	}
	return &foldWriter{w: bufio.NewWriter(w), lineEnding: lineEnding}
}

func (fw *foldWriter) Write(s []byte) (i int, err error) {
	if fw.err != nil {
		return
	}

	remaining := 75 - fw.n

	for i = 0; i < len(s) && fw.err == nil; i++ {
		c := s[i]
		if remaining < 1 {
			fw.wrapLine()
			remaining = 75
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

	if fw.n > 74 {
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

func (fw *foldWriter) flush() error {
	e := fw.w.Flush()
	if e != nil {
		return e
	}
	return fw.err
}

//-------------------------------------------------------------------------------------------------

// Buffer wraps bufio.Writer with some iCal-specific helper methods.
// It folds long lines to meet the iCal max-75 characters per line limit.
// If coallesces errors so they don't have to be checked after every method;
// it is sufficient to check once at the end.
type Buffer struct {
	fw *foldWriter
}

// NewBuffer constructs a Buffer that wraps some Writer. The lineEnding can be "" or "\r\n" for
// normal iCal formatting, or "\n" in other cases.
func NewBuffer(w io.Writer, lineEnding string) *Buffer {
	return &Buffer{newFoldWriter(w, lineEnding)}
}

// WriteString writes the string supplied.
func (b *Buffer) WriteString(s string) error {
	b.fw.WriteString(s)
	return b.fw.err
}

func (b *Buffer) WriteLine(s string) error {
	b.WriteString(s)
	return b.fw.newline()
}

func (b *Buffer) WriteValuerLine(predicate bool, label string, v Valuer) error {
	if !predicate {
		return b.fw.err // skip
	}

	b.WriteString(label)
	v.WriteTo(b.fw)
	return b.fw.newline()
}

func (b *Buffer) Flush() error {
	return b.fw.flush()
}
