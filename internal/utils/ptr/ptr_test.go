package ptr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	t.Run("int ptr", func(t *testing.T) {
		var i = 10

		p := Ptr(i)
		require.Equal(t, &i, p)
	})
	t.Run("float ptr", func(t *testing.T) {
		var f = 10.

		p := Ptr(f)
		require.Equal(t, &f, p)
	})
	t.Run("struct ptr", func(t *testing.T) {
		var s = struct {
			i int
			f float64
		}{
			i: 1,
			f: 2,
		}

		p := Ptr(s)
		require.Equal(t, &s, p)
	})
}

func TestVal(t *testing.T) {
	t.Run("int val", func(t *testing.T) {
		var i = 10

		v := Val(&i)
		require.Equal(t, i, v)
	})
	t.Run("float val", func(t *testing.T) {
		var f = 10.

		v := Val(&f)
		require.Equal(t, f, v)
	})
	t.Run("struct val", func(t *testing.T) {
		var s = struct {
			i int
			f float64
		}{
			i: 1,
			f: 2,
		}

		v := Val(&s)
		require.Equal(t, s, v)
	})
	t.Run("struct val", func(t *testing.T) {
		v := Val[struct {
			i int
			f float64
		}](nil)
		require.Zero(t, v)
	})
}
