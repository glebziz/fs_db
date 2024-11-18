package model

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model/sequence"
)

func TestFile_Latest(t *testing.T) {
	var (
		testKey1       = gofakeit.UUID()
		testKey2       = gofakeit.UUID()
		testContentId1 = gofakeit.UUID()
		testContentId2 = gofakeit.UUID()
		testSeq        = sequence.Seq(1)
	)
	for _, tc := range []struct {
		name   string
		file   File
		other  File
		latest File
	}{
		{
			name: "other latest",
			file: File{
				Key:       testKey1,
				ContentId: testContentId1,
				Seq:       testSeq,
			},
			other: File{
				Key:       testKey2,
				ContentId: testContentId2,
				Seq:       testSeq + 1,
			},
			latest: File{
				Key:       testKey2,
				ContentId: testContentId2,
				Seq:       testSeq + 1,
			},
		},
		{
			name: "file latest",
			file: File{
				Key:       testKey1,
				ContentId: testContentId1,
				Seq:       testSeq,
			},
			other: File{
				Key:       testKey2,
				ContentId: testContentId2,
				Seq:       testSeq - 1,
			},
			latest: File{
				Key:       testKey1,
				ContentId: testContentId1,
				Seq:       testSeq,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			latest := tc.file.Latest(tc.other)
			require.Equal(t, tc.latest, latest)
		})
	}
}
