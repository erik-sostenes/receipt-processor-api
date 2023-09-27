package ports

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
)

// ports left side -> (drives)
type (
	ReceiptCreator interface {
		CreateReceipt(context.Context, *receipt.Receipt) (receipt.ReceiptId, error)
	}
)

// ports right side -> (driven)
type (
	ReceiptSaver interface {
		SaveReceipt(context.Context, *receipt.Receipt) (receipt.ReceiptId, error)
	}
)
