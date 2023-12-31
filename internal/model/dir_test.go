package model

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestDir_GetPath(t *testing.T) {
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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			dir := Dir{
				Id:         tc.id,
				ParentPath: tc.parent,
			}

			path := dir.GetPath()

			require.Equal(t, tc.path, path)
		})
	}
}
