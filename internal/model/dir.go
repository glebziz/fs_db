package model

import (
	"iter"
	"math/rand/v2"
	"path"
)

type Dir struct {
	Name  string
	Root  string
	Count uint64
	Free  uint64
}

func ParseDir(dirPath string) Dir {
	return Dir{
		Name: path.Base(dirPath),
		Root: path.Dir(dirPath),
	}
}

func (d *Dir) Path() string {
	return path.Join(d.Root, d.Name)
}

type Dirs []Dir

func (ds Dirs) Iterate(r *rand.Rand) iter.Seq2[Dir, bool] {
	r.Shuffle(len(ds), func(i, j int) {
		ds[i], ds[j] = ds[j], ds[i]
	})

	return func(yield func(dir Dir, ok bool) bool) {
		for _, dir := range ds {
			if !yield(dir, true) {
				return
			}
		}

		yield(Dir{}, false)
	}
}
