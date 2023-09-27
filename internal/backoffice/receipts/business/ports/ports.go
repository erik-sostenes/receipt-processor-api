package ports

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
)

// ports right side
type (
	Saver interface {
		Save(context.Context, *receipt.Receipt) (receipt.ReceiptId, error)
	}
)
