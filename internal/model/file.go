package model

import "time"

type File struct {
	Key       string
	ContentId string
}

type FileFilter struct {
	TxId     *string
	BeforeTs *time.Time
}
