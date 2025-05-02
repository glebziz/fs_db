package model

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SingleContent(t *testing.T) {
	t.Parallel()

	const (
		someContent = "someContent"
	)

	for _, tc := range []struct {
		name    string
		r       io.Reader
		content []byte
	}{
		{
			name:    "success",
			r:       strings.NewReader(someContent),
			content: []byte(someContent),
		},
		{
			name:    "empty reader",
			r:       strings.NewReader(""),
			content: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				count int
				buf   bytes.Buffer
			)
			for content := range SingleContent(tc.r) {
				count++

				require.Equal(t, Seek{
					Pos: 0,
					Dir: SeekStart,
				}, content.Seek)

				_, err := io.Copy(&buf, content.Reader)
				require.NoError(t, err)
			}

			require.Equal(t, 1, count)
			require.Equal(t, tc.content, buf.Bytes())
		})
	}
}

func TestContents_InsertAtStart(t *testing.T) {
	t.Parallel()

	const (
		prefix = "Hello"
		suffix = ", World"
	)

	for _, tc := range []struct {
		name     string
		contents func(t *testing.T) (context.Context, Contents, func(contents Contents))
	}{
		{
			name: "success",
			contents: func(t *testing.T) (context.Context, Contents, func(contents Contents)) {
				contents := make(chan Content)
				go func() {
					defer close(contents)

					for _, b := range suffix {
						contents <- Content{
							Reader: strings.NewReader(string(b)),
						}
					}
				}()

				return context.Background(), contents, func(contents Contents) {
					var allContent []byte
					for content := range contents {
						buf, err := io.ReadAll(content.Reader)
						require.NoError(t, err)

						allContent = append(allContent, buf...)
					}

					require.EqualValues(t, prefix+suffix, allContent)
				}
			},
		},
		{
			name: "success with context.Done on read",
			contents: func(t *testing.T) (context.Context, Contents, func(contents Contents)) {
				ctx, cancel := context.WithCancel(context.Background())
				return ctx, nil, func(contents Contents) {
					defer cancel()

					content := <-contents
					buf, err := io.ReadAll(content.Reader)
					require.NoError(t, err)
					require.EqualValues(t, prefix, buf)
				}
			},
		},
		{
			name: "success with context.Done on write",
			contents: func(t *testing.T) (context.Context, Contents, func(contents Contents)) {
				ctx, cancel := context.WithCancel(context.Background())
				contents := make(chan Content)

				return ctx, contents, func(Contents) {
					contents <- Content{}
					cancel()
					close(contents)
				}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, contents, checkOut := tc.contents(t)
			contents = contents.InsertAtStart(ctx, Content{
				Reader: strings.NewReader(prefix),
			})

			checkOut(contents)
		})
	}
}
