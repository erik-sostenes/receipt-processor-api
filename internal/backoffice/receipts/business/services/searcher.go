package services

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
)

var _ ports.ReceiptSearcher = &ReceiptSearcher{}

type ReceiptSearcher struct{}

func NewReceiptSearcher() *ReceiptSearcher {
	return &ReceiptSearcher{}
}

func (r *ReceiptSearcher) SearchReceipt(ctx context.Context, receiptId *receipt.ReceiptId) (receipt.ReceiptPoints, error) {
	return receipt.NewReceiptPoints(40)
}
