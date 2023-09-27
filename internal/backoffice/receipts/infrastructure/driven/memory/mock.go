package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/ports"
	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
	"github.com/erik-sostenes/receipt-processor-api/pkg/set"
	"github.com/erik-sostenes/receipt-processor-api/pkg/wrongs"
)

// ReceiptInMemory implements the ports.Saver interface
var _ ports.Saver = &ReceiptInMemory{}

type ReceiptInMemory struct {
	*set.Set[receipt.ReceiptId, receipt.Receipt]
}

func NewReciptInMemory() *ReceiptInMemory {
	return &ReceiptInMemory{
		Set: set.NewSet[receipt.ReceiptId, receipt.Receipt](),
	}
}

func (r *ReceiptInMemory) Save(_ context.Context, rc *receipt.Receipt) (_ receipt.ReceiptId, err error) {
	_, ok := r.Get(rc.ReceiptId)
	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("receipt with id '%s' already exists receipt", rc.ReceiptId.Value()))
		return
	}

	uuid, err := receipt.NewReceiptId(common.GenerateUuID())
	if err != nil {
		return
	}

	r.Add(uuid, *rc)

	return uuid, nil
}
