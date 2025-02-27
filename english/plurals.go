package english

// Pluralise returns the plural form of a word.
func Pluralise(word string) (string, bool) {
	return word + "s", true // the baseline
}
