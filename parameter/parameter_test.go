package parameter

import (
	"bytes"
	"testing"
)

func TestParameterEquals(t *testing.T) {
	cases := []struct {
		a, b      Parameter
		mustEqual bool
	}{
		{Encoding(true), Encoding(true), true},
		{Encoding(true), Encoding(false), false},
		{Encoding(true), Language("en"), false},
		{Delegator("one"), Delegator("another"), false},
		{Delegator("one"), Delegator("one"), true},
		{Delegator("one"), Delegatee("one"), false},
		{Delegator("one"), Delegator("one", "two"), false},
		{Delegator("one", "2"), Delegator("one", "two"), false},
	}

	for i, c := range cases {
		if c.mustEqual {
			assertTrue(t, c.a.Equals(c.b), "%d: %v must equal %v", i, c.a, c.b)
		} else {
			assertTrue(t, !c.a.Equals(c.b), "%d: %v must not equal %v", i, c.a, c.b)
		}
	}
}

func TestParameterConstructorsAndWriteTo(t *testing.T) {
	cases := []struct {
		v   Parameter
		exp string
	}{
		{AltRep("abc"), "ALTREP=abc"},
		{CommonName("abc"), "CN=abc"},
		{Dir("abc"), "DIR=abc"},
		{Email("a@b.it"), `EMAIL=a@b.it`},
		{FmtTypeOf("image", "png"), "FMTTYPE=image/png"},
		{Label("zap"), `LABEL=zap`},
		{Language("en"), "LANGUAGE=en"},
		{Member("a", "b"), "MEMBER=a,b"},
		{Member("a,z", "b", "c;u", "d:1"), `MEMBER="a,z","b","c;u","d:1"`},
		{Rsvp(true), `RSVP=TRUE`},
		{SentBy("Joe"), `SENT-BY=Joe`},
		{TZid("UTC"), `TZID=UTC`},
	}

	for i, c := range cases {
		b := &bytes.Buffer{}
		c.v.WriteTo(b)
		s := b.String()
		if s != c.exp {
			t.Errorf("%d: expected %q but got %q", i, c.exp, s)
		}
	}
}

func assertTrue(t *testing.T, predicate bool, hint string, args ...interface{}) {
	t.Helper()
	if !predicate {
		t.Errorf(hint, args...)
	}
}
