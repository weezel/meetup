package hex

import (
	"testing"
	"unicode/utf8"
)

func Print(t *testing.T, s string) {
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		t.Logf("Rune: %q, Hex: %x\n", r, r)
		s = s[size:] // Advance to the next rune
	}
	t.Log()
}
