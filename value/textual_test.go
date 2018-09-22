package value

import (
	"bytes"
	"github.com/rickb777/ical2/ics"
	"github.com/rickb777/ical2/parameter/value"
	"testing"
)

func TestTextConstructors(t *testing.T) {
	cases := []struct {
		v   ics.Valuer
		exp string
	}{
		{Text("abc").With(value.Type(value.TEXT)), ";VALUE=TEXT:abc\n"},
		{Text("a,b,c"), ":a\\,b\\,c\n"},
		{CalAddress("j@x.org"), ":mailto:j@x.org\n"},
		{List("APPOINTMENT", "EDUCATION"), ":APPOINTMENT,EDUCATION\n"},
		{Public(), ":PUBLIC\n"},
		{Private(), ":PRIVATE\n"},
		{Confidential(), ":CONFIDENTIAL\n"},
		{Publish(), ":PUBLISH\n"},
		{Request(), ":REQUEST\n"},
		{Tentative(), ":TENTATIVE\n"},
		{Confirmed(), ":CONFIRMED\n"},
		{Cancelled(), ":CANCELLED\n"},
		{Completed(), ":COMPLETED\n"},
		{NeedsAction(), ":NEEDS-ACTION\n"},
		{InProcess(), ":IN-PROCESS\n"},
		{Draft(), ":DRAFT\n"},
		{Final(), ":FINAL\n"},
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

func TestLists(t *testing.T) {
	cases := []struct {
		vv  []ListValue
		exp []string
	}{
		{Lists("APPOINTMENT", "EDUCATION"),
			[]string{":APPOINTMENT,EDUCATION\n"}},

		{Lists("ABCDEFGHIJKLMNOPQRSTUVWXYZ1", "ABCDEFGHIJKLMNOPQRSTUVWXYZ2", "ABCDEFGHIJKLMNOPQRSTUVWXYZ3",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ4"),
			[]string{":ABCDEFGHIJKLMNOPQRSTUVWXYZ1,ABCDEFGHIJKLMNOPQRSTUVWXYZ2\n", ":ABCDEFGHIJKLMNOPQRSTUVWXYZ3,ABCDEFGHIJKLMNOPQRSTUVWXYZ4\n"}},
	}

	for i, c := range cases {
		if len(c.vv) != len(c.exp) {
			t.Errorf("%d: expected %d but got %d\n%+v", i, len(c.exp), len(c.vv), c.vv)
		}
		for j, v := range c.vv {
			b := &bytes.Buffer{}
			x := ics.NewBuffer(b, "\n")
			x.WriteValuerLine(true, "X", v)
			err := x.Flush()
			if err != nil {
				t.Errorf("%d: unexpected error %v", i, err)
			}
			s := b.String()
			// ignore X prefix
			s = s[1:]
			if s != c.exp[j] {
				t.Errorf("%d: expected %q but got %q", i, c.exp, s)
			}
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
