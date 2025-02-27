package english

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlurals(t *testing.T) {
	req := require.New(t)

	for _, pair := range []struct {
		lemma  string
		plural string
	}{
		{"table", "tables"},
		{"knife", "knives"},
		{"valency", "valencies"},
	} {
		p, ok := Pluralise(pair.lemma)
		req.True(ok)
		req.Equal(pair.plural, p)
	}
}
