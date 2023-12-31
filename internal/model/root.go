package model

type Root struct {
	Path  string
	Free  uint64
	Count uint64
}

type RootMap map[string]*Root

func (rm RootMap) Select(size uint64) *Root {
	var sRoot *Root

	for _, root := range rm {
		if root.Free < size {
			continue
		}

		if sRoot == nil || root.Count < sRoot.Count {
			sRoot = root
		}
	}

	return sRoot
}
