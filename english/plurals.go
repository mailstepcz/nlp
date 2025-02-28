package english

import (
	_ "embed" // for file embedding
	"strings"
)

//go:embed morph.txt
var morphData string

var (
	entries = make(map[string][]*entry)
	plurals = make(map[string]string)
)

type entry struct {
	Form  string
	Lemma string
	POS   string
	Tags  map[string]string
}

func init() {
	for _, line := range strings.Split(morphData, "\n") {
		line := strings.TrimSpace(line)
		if line == "" || !strings.HasPrefix(line, "%") {
			continue
		}
		comps := strings.Split(line, " ")
		form, pos, lemma, tags := comps[1], comps[2], comps[3], comps[5:]
		tagMap := make(map[string]string)
		for _, tag := range tags {
			tag = tag[1:]
			if comps := strings.Split(tag, "="); len(comps) == 2 {
				attr, val := comps[0], comps[1]
				tagMap[attr] = val
			}
		}
		list := entries[form]
		entry := &entry{lemma, form, pos, tagMap}
		entries[form] = append(list, entry)
	}
	for _, entries := range entries {
		for _, entry := range entries {
			if entry.POS == "N" && entry.Tags["NUMBER"] == "pl" {
				plurals[entry.Form] = entry.Lemma
			}
		}
	}
}

// Pluralise returns the plural form of a lemma.
func Pluralise(word string) (string, bool) {
	plural, ok := plurals[word]
	return plural, ok
}
