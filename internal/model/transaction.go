package model

import (
	"context"
	"time"
)

const (
	MainTxId = "00000000-0000-0000-0000-000000000000"

	ctxTxIdKey = "ctxTxIdKey"
)

type TxIsoLevel uint8

type Transaction struct {
	Id       string
	IsoLevel TxIsoLevel
	CreateTs time.Time
}

func StoreTxId(ctx context.Context, txId string) context.Context {
	return context.WithValue(ctx, ctxTxIdKey, txId)
}

func GetTxId(ctx context.Context) string {
	txId, _ := ctx.Value(ctxTxIdKey).(string)
	if txId == "" {
		txId = MainTxId
	}

	return txId
}
