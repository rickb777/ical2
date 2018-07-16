package value

import (
	"testing"
)

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
