package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/pkg/set"
	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
)

// ReceiptInMemory implements the ports.Saver interface
var _ ports.Saver = &ReceiptInMemory{}

type ReceiptInMemory struct {
	*set.Set[string, domain.Receipt]
}

func NewReciptInMemory() *ReceiptInMemory {
	return &ReceiptInMemory{
		Set: set.NewSet[string, domain.Receipt](),
	}
}

func (r *ReceiptInMemory) Save(_ context.Context, receipt *domain.Receipt) (err error) {
	_, ok := r.Get(receipt.Id)
	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("receipt with id '%s' already exists receipt", receipt.Id))
		return
	}

	r.Add(receipt.Id, *receipt)

	return
}
