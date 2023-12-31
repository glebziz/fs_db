package model

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestRootMap_Select(t *testing.T) {
	var (
		size = uint64(1000)

		root = &Root{
			Path:  gofakeit.UUID(),
			Free:  10000,
			Count: 1000,
		}
		root2 = &Root{
			Path:  gofakeit.UUID(),
			Free:  100,
			Count: 10,
		}
		root3 = &Root{
			Path:  gofakeit.UUID(),
			Free:  1000,
			Count: 100,
		}

		rMap = RootMap{
			root.Path:  root,
			root2.Path: root2,
			root3.Path: root3,
		}
	)

	sRoot := rMap.Select(size)

	require.Equal(t, root3, sRoot)
}
