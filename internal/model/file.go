package model

import (
	"github.com/glebziz/fs_db/internal/model/sequence"
)

type File struct {
	Key       string
	TxId      string
	ContentId string
	Seq       sequence.Seq
}

func (f File) Latest(o File) File {
	if f.Seq.After(o.Seq) {
		return f
	}

	return o
}

func (f File) Deleted() bool {
	return f.ContentId == ""
}

type FileFilter struct {
	TxId      *string
	BeforeSeq *sequence.Seq
}
