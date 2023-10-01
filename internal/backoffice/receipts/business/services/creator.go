package services

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
)

// ReceiptCreator implements the ports.ReceiptCreator interface
var _ ports.ReceiptCreator = &ReceiptCreator{}

type ReceiptCreator struct {
	ports.ReceiptSaver
}

func NewReciptCreator(saver ports.ReceiptSaver) *ReceiptCreator {
	return &ReceiptCreator{
		saver,
	}
}

func (r *ReceiptCreator) CreateReceipt(ctx context.Context, receipt *receipt.Receipt) (receipt.ReceiptId, error) {
	return r.SaveReceipt(ctx, receipt)
}
