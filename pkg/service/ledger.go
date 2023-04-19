package service

import (
	"context"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
)

type LedgerService struct {
	LedgerClient *ent.LedgerClient
}

func (s *LedgerService) GetAllWithPayerAndReceiver(ctx context.Context) ([]*ent.Ledger, error) {
	return s.LedgerClient.Query().WithPayer().WithReceiver().All(ctx)
}
