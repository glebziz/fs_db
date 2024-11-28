package ptr

func Ptr[T any](v T) *T {
	return &v
}

func Val[T any](p *T) T {
	if p == nil {
		return *new(T)
	}

	return *p
}
