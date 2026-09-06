package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	N = 10_000
)

func TestGen_Generate(t *testing.T) {
	g := New()
	strs := make(map[string]struct{}, N)

	for range N {
		str := g.Generate()
		strs[str] = struct{}{}
	}

	require.Len(t, strs, N)
}
