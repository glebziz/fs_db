package model

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestTransaction(t *testing.T) {
	var (
		testTxId = gofakeit.UUID()
	)

	for _, tc := range []struct {
		name string
		in   *string
		exp  string
	}{
		{
			name: "tx id not empty",
			in:   ptr.Ptr(testTxId),
			exp:  testTxId,
		},
		{
			name: "empty tx id",
			in:   ptr.Ptr(""),
			exp:  MainTxId,
		},
		{
			name: "null tx id",
			exp:  MainTxId,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			if tc.in != nil {
				ctx = StoreTxId(context.Background(), ptr.Val(tc.in))
			}

			txId := GetTxId(ctx)

			require.Equal(t, tc.exp, txId)
		})
	}
}
