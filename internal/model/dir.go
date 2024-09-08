package model

import (
	"math/rand/v2"
	"path"
)

type Dir struct {
	Name  string
	Root  string
	Count uint64
	Free  uint64
}

func (d *Dir) Path() string {
	return path.Join(d.Root, d.Name)
}

type Dirs []Dir

func (ds Dirs) Iterate(r *rand.Rand) (nextFn func() (Dir, bool)) {
	r.Shuffle(len(ds), func(i, j int) {
		ds[i], ds[j] = ds[j], ds[i]
	})

	var i int
	return func() (Dir, bool) {
		defer func() { i++ }()

		if i >= len(ds) {
			return Dir{}, false
		}

		return ds[i], true
	}
}
