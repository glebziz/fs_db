package core

import (
	"github.com/brianvoe/gofakeit/v6"

	"github.com/glebziz/fs_db/internal/model/sequence"
)

var (
	testTxId       = gofakeit.UUID()
	testKey        = gofakeit.UUID()
	testKey2       = gofakeit.UUID()
	testContentId  = gofakeit.UUID()
	testContentId2 = gofakeit.UUID()
	testSeq        = sequence.Seq(10)
)
