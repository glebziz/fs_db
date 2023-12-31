package model

import "path"

type Dir struct {
	Id         string
	FileCount  uint64
	ParentPath string
}

func (d *Dir) GetPath() string {
	return path.Join(d.ParentPath, d.Id)
}
