package services

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
)

type ReceiptCreator struct {
	ports.Saver
}

func NewReciptCreator(saver ports.Saver) *ReceiptCreator {
	return &ReceiptCreator{
		saver,
	}
}

func (r ReceiptCreator) Create(ctx context.Context, receiptRequest *dto.ReceiptRequest) (receipt.ReceiptId, error) {
	receipt, err := receipt.NewReceipt(*receiptRequest)
	if err != nil {
		return receipt.ReceiptId, err
	}

	return r.Save(ctx, receipt)
}
