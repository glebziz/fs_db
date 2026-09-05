package wpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWithNumWorkers(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name          string
		numWorkers    int
		resNumWorkers int
	}{
		{
			name:          "success",
			numWorkers:    minNumWorkers + 1,
			resNumWorkers: minNumWorkers + 1,
		},
		{
			name:          "negative",
			numWorkers:    -1,
			resNumWorkers: minNumWorkers,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var opts options
			WithNumWorkers(tt.numWorkers)(&opts)

			require.Equal(t, tt.resNumWorkers, opts.numWorkers)
		})
	}
}

func TestWithSendDuration(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		sendDuration    time.Duration
		resSendDuration time.Duration
	}{
		{
			name:            "success",
			sendDuration:    time.Second,
			resSendDuration: time.Second,
		},
		{
			name:            "negative",
			sendDuration:    -time.Second,
			resSendDuration: minSendDuration,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var opts options
			WithSendDuration(tt.sendDuration)(&opts)

			require.Equal(t, tt.resSendDuration, opts.sendDuration)
		})
	}
}
