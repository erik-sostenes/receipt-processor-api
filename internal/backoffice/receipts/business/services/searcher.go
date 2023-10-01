package services

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
)

var _ ports.ReceiptSearcher = &ReceiptSearcher{}

type ReceiptSearcher struct {
	ports.ReceiptFinder
}

func NewReceiptSearcher(finder ports.ReceiptFinder) *ReceiptSearcher {
	return &ReceiptSearcher{
		finder,
	}
}

func (r *ReceiptSearcher) SearchReceipt(ctx context.Context, receiptId *receipt.ReceiptId) (receipt.ReceiptPoints, error) {
	return r.ReceiptFinder.FindReceipt(ctx, receiptId)
}
