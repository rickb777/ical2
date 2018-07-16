package parameter

import (
	"testing"
)

func TestParametersAddIsUnique(t *testing.T) {
	pp := Parameters{}
	pp = pp.Append(Encoding(true)).Append(CommonName("Joe")).Append(Encoding(false))

	assertTrue(t, len(pp) == 2, "length should be 2: %v", pp)
	assertTrue(t, pp[0].Equals(CommonName("Joe")), "expected CommonName('Joe'): %v", pp)
	assertTrue(t, pp[1].Equals(Encoding(false)), "expected Encoding(false): %v", pp)

	pp = pp.Prepend(Encoding(true))

	assertTrue(t, len(pp) == 2, "length should be 2: %v", pp)
	assertTrue(t, pp[0].Equals(Encoding(true)), "expected Encoding(false): %v", pp)
	assertTrue(t, pp[1].Equals(CommonName("Joe")), "expected CommonName('Joe'): %v", pp)
}
