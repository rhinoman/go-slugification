// Package slugification provides methods for generating 'slugged' versions
// of strings suitable for use in URLs.
package slugification

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Returns a slugified string
// Example: "Page Title" becomes "page-title"
func Slugify(inputString string) string {

	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	replaceChar := func(r rune) rune {
		switch {
		case unicode.IsLetter(r):
			return unicode.ToLower(r)
		case unicode.IsNumber(r), r == '_', r == '-', r == '+':
			return r
		case unicode.IsSpace(r):
			return '-'
		default:
			return -1
		}
	}

	slug, _, _ := transform.String(t, strings.Map(replaceChar, inputString))

	return slug
}
