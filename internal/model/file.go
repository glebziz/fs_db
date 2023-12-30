package model

import "path"

type File struct {
	Id         string
	Key        string
	ParentPath string
}

func (f *File) GetPath() string {
	return path.Join(f.ParentPath, f.Id)
}
