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

	ReceiptSearcher interface {
		SearchReceipt(context.Context, *receipt.ReceiptId) (receipt.ReceiptPoints, error)
	}
)

// ports right side -> (driven)
type (
	ReceiptSaver interface {
		SaveReceipt(context.Context, *receipt.Receipt) (receipt.ReceiptId, error)
	}

	ReceiptFinder interface {
		FindReceipt(context.Context, string, *receipt.ReceiptId) (receipt.ReceiptPoints, error)
	}
)
