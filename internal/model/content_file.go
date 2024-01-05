package model

import "path"

type ContentFile struct {
	Id         string
	ParentPath string
}

func (f *ContentFile) GetPath() string {
	return path.Join(f.ParentPath, f.Id)
}
