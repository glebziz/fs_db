package model

import "path"

type ContentFile struct {
	Id     string
	Parent string
}

func (f *ContentFile) Path() string {
	return path.Join(f.Parent, f.Id)
}
