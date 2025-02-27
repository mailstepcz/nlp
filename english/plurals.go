package english

import "strings"

var irreg = map[string]string{
	"fish":  "fish",
	"sheep": "sheep",
}

// Pluralise returns the plural form of a lemma.
func Pluralise(word string) (string, bool) {
	plural, ok := irreg[word]
	if ok {
		return plural, true
	}

	if strings.HasSuffix(word, "fe") {
		return word[:len(word)-2] + "ves", true
	}

	if strings.HasSuffix(word, "y") {
		return word[:len(word)-1] + "ies", true
	}

	return word + "s", true // fallback
}
