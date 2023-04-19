package controller

import (
	"context"
	"github.com/nishinoyama/kobuy-2/pkg/ent"
	"github.com/nishinoyama/kobuy-2/pkg/service"
)

type LedgerController struct {
	LedgerService service.LedgerService
}

type SomeLedgersResponse struct {
	Ledgers []*ent.Ledger `json:"ledgers"`
}

func (c *LedgerController) GetAll() (*SomeLedgersResponse, error) {
	ledgers, err := c.LedgerService.GetAllWithPayerAndReceiver(context.TODO())
	if err != nil {
		return nil, err
	}
	return &SomeLedgersResponse{Ledgers: ledgers}, nil
}
