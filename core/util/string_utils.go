package util

import (
	"fmt"
	"unicode"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func remove_vietnamese_accent(s string) string {
	fmt.Println(s)
	b := make([]byte, len(s))
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, e := t.Transform(b, []byte(s), true)
	if e != nil {
		panic(e)
	}

	return string(b)
}