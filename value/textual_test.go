package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"testing"
	"github.com/rickb777/ical2/parameter/valuetype"
)

func TestTextConstructors(t *testing.T) {
	cases := []struct {
		v   ics.Valuer
		exp string
	}{
		{Text("abc").With(valuetype.Type(valuetype.TEXT)), ";VALUE=TEXT:abc\n"},
		{Text("a,b,c"), ":a\\,b\\,c\n"},
		{CalAddress("j@x.org"), ":mailto:j@x.org\n"},
		{Public(), ":PUBLIC\n"},
		{Private(), ":PRIVATE\n"},
		{Confidential(), ":CONFIDENTIAL\n"},
		{Publish(), ":PUBLISH\n"},
		{Request(), ":REQUEST\n"},
		{TentativeStatus(), ":TENTATIVE\n"},
		{ConfirmedStatus(), ":CONFIRMED\n"},
		{CancelledStatus(), ":CANCELLED\n"},
		{CompletedStatus(), ":COMPLETED\n"},
		{NeedsActionStatus(), ":NEEDS-ACTION\n"},
		{InProcessStatus(), ":IN-PROCESS\n"},
		{DraftStatus(), ":DRAFT\n"},
		{FinalStatus(), ":FINAL\n"},
		{Transparent(), ":TRANSPARENT\n"},
		{Opaque(), ":OPAQUE\n"},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		x := ics.NewBuffer(b, "\n")
		x.WriteValuerLine(true, "X", c.v)
		err := x.Flush()
		if err != nil {
			t.Errorf("%d: unexpected error %v", i, err)
		}
		s := b.String()
		// ignore X prefix
		s = s[1:]
		if s != c.exp {
			t.Errorf("%d: expected %q but got %q", i, c.exp, s)
		}
	}
}

func TestEscapeText(t *testing.T) {
	cases := []struct {
		input, exp string
	}{
		// blank string stays blank
		{"", ""},
		// backslash
		{`\`, `\\`},
		// semicolon
		{`;`, `\;`},
		// comma
		{`,`, `\,`},
		// newline
		{"\n", `\n`},

		// unescaped characters are unchanged
		{"0123456789 ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz",
			"0123456789 ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz"},
		// some higher characters and most punctuation
		{"ł¶←↓→øþæßðđŋħłæ»¢“”nµ.:<>[]{}-_+=~#|$%^&*()",
			"ł¶←↓→øþæßðđŋħłæ»¢“”nµ.:<>[]{}-_+=~#|$%^&*()"},
	}

	for i, c := range cases {
		got := escapeText(c.input)
		if got != c.exp {
			t.Errorf("%d: expected %s, got %s", i, c.exp, got)
		}
	}
}
