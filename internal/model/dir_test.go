package model

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

type randSource struct{}

func (randSource) Uint64() uint64 { return 1 }

func TestParseDir(t *testing.T) {
	for _, tc := range []struct {
		name string
		path string
		dir  Dir
	}{
		{
			name: "success",
			path: "testdata/dir",
			dir: Dir{
				Name: "dir",
				Root: "testdata",
			},
		},
		{
			name: "dir without root",
			path: "dir",
			dir: Dir{
				Name: "dir",
				Root: ".",
			},
		},
		{
			name: "empty path",
			path: "",
			dir: Dir{
				Name: ".",
				Root: ".",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			dir := ParseDir(tc.path)
			require.Equal(t, tc.dir, dir)
		})
	}
}

func TestDir_Path(t *testing.T) {
	var (
		testId     = gofakeit.UUID()
		testParent = gofakeit.UUID()
	)

	for _, tc := range []struct {
		name   string
		id     string
		parent string
		path   string
	}{
		{
			name:   "with id and parent",
			id:     testId,
			parent: testParent,
			path:   fmt.Sprintf("%s/%s", testParent, testId),
		},
		{
			name: "with id",
			id:   testId,
			path: testId,
		},
		{
			name:   "with parent",
			parent: testParent,
			path:   testParent,
		},
		{
			name: "empty id and empty parent",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := Dir{
				Name: tc.id,
				Root: tc.parent,
			}

			path := dir.Path()

			require.Equal(t, tc.path, path)
		})
	}
}

func TestDirs_Iterate(t *testing.T) {
	var (
		r = rand.New(randSource{})

		dirs = Dirs{{
			Name: gofakeit.UUID(),
			Root: gofakeit.UUID(),
		}, {
			Name: gofakeit.UUID(),
			Root: gofakeit.UUID(),
		}}

		dirs2 = make(Dirs, 0, len(dirs))
	)

	for dir, ok := range dirs.Iterate(r) {
		if !ok {
			break
		}

		dirs2 = append(dirs2, dir)
	}
	for dir, ok := range dirs.Iterate(r) {
		_, _ = dir, ok
		break
	}

	require.Equal(t, dirs, dirs2)
}
