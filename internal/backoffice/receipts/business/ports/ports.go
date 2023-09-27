package ports

import (
	"context"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain"
)

// ports right side
type (
	Saver interface {
		Save(context.Context, *domain.Receipt) error
	}
)
