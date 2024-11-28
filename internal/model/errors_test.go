package model

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type closer struct {
	io.Reader
	count int
}

func (c *closer) Close() error {
	c.count++
	return nil
}

func TestNotEnoughSpaceError_Error(t *testing.T) {
	err := NotEnoughSpaceError{
		Err: assert.AnError,
	}

	assert.Equal(t, fmt.Sprintf("not enough space: %s", assert.AnError), err.Error())
}

func TestNotEnoughSpaceError_Unwrap(t *testing.T) {
	err := NotEnoughSpaceError{
		Err: assert.AnError,
	}

	require.Equal(t, assert.AnError, err.Unwrap())
}

func TestNotEnoughSpaceError_Reader(t *testing.T) {
	err := NotEnoughSpaceError{
		Start:  io.NopCloser(strings.NewReader("123")),
		Middle: strings.NewReader("456"),
		End:    strings.NewReader("789"),
	}

	r := err.Reader()
	content, e := io.ReadAll(r)
	require.NoError(t, e)
	assert.EqualValues(t, "123456789", content)
}

func TestNotEnoughSpaceError_Close(t *testing.T) {
	var (
		c   = &closer{}
		err = NotEnoughSpaceError{
			Start: c,
		}
	)

	require.NoError(t, err.Close())
	require.Equal(t, 1, c.count)
}
