package sequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeq_After(t *testing.T) {
	for _, tc := range []struct {
		name  string
		seq   Seq
		other Seq
		after bool
	}{
		{
			name:  "other after than seq",
			seq:   Next(),
			other: Next(),
			after: false,
		},
		{
			name:  "seq after than other",
			other: Next(),
			seq:   Next(),
			after: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			after := tc.seq.After(tc.other)
			require.Equal(t, tc.after, after)
		})
	}
}

func TestSeq_Before(t *testing.T) {
	for _, tc := range []struct {
		name   string
		seq    Seq
		other  Seq
		before bool
	}{
		{
			name:   "other before than seq",
			other:  Next(),
			seq:    Next(),
			before: false,
		},
		{
			name:   "seq before than other",
			seq:    Next(),
			other:  Next(),
			before: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			before := tc.seq.Before(tc.other)
			require.Equal(t, tc.before, before)
		})
	}
}

func TestSeq_Zero(t *testing.T) {
	for _, tc := range []struct {
		name string
		seq  Seq
		zero bool
	}{
		{
			name: "zero",
			seq:  0,
			zero: true,
		},
		{
			name: "non zero",
			seq:  1,
			zero: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			zero := tc.seq.Zero()
			require.Equal(t, tc.zero, zero)
		})
	}
}
