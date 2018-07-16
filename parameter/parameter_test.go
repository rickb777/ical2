package parameter

import (
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
		{Delegator("one"), Delegator("one", "two"), false},
		{Delegator("one", "2"), Delegator("one", "two"), false},
	}

	for i, c := range cases {
		assertTrue(t, c.a.Equals(c.b) == c.mustEqual, "%d: %v must equal %v", i, c.a, c.b)
	}
}

func assertTrue(t *testing.T, predicate bool, hint string, args ...interface{}) {
	t.Helper()
	if !predicate {
		t.Errorf(hint, args...)
	}
}
