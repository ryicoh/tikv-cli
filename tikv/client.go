package tikv

import (
	"context"

	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/txnkv"
)

func NewClient(ctx context.Context, endpoints []string) (*txnkv.Client, error) {
	txn, err := txnkv.NewClient(ctx, endpoints, config.Default())
	if err != nil {
		return nil, err
	}

	return txn, nil
}
