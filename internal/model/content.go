package model

import (
	"context"
	"io"
)

type Content struct {
	Seek   Seek
	Reader io.Reader
}

type Contents <-chan Content

func SingleContent(r io.Reader) Contents {
	ch := make(chan Content, 1)
	defer close(ch)

	ch <- Content{
		Reader: r,
	}

	return ch
}

func (cs Contents) InsertAtStart(ctx context.Context, content Content) Contents {
	contents := make(chan Content, 1)
	contents <- content

	go func() {
		defer close(contents)

		var ok bool
		for {
			select {
			case <-ctx.Done():
				return
			case content, ok = <-cs:
			}

			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case contents <- content:
			}
		}
	}()

	return contents
}
