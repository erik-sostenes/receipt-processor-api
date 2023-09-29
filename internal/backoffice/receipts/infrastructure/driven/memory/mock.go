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

// ReceiptInMemory implements the ports.ReceiptSaver interface
var _ ports.ReceiptSaver = &ReceiptInMemory{}

// ReceiptInMemory implements the ports.ReceiptFinder
var _ ports.ReceiptFinder = &ReceiptInMemory{}

type ReceiptInMemory struct {
	*set.Set[receipt.ReceiptId, receipt.Receipt]
}

func NewReciptInMemory() *ReceiptInMemory {
	return &ReceiptInMemory{
		Set: set.NewSet[receipt.ReceiptId, receipt.Receipt](),
	}
}

func (r *ReceiptInMemory) SaveReceipt(_ context.Context, rcpt *receipt.Receipt) (_ receipt.ReceiptId, err error) {
	_, ok := r.Get(rcpt.ReceiptId)
	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("receipt with id '%s' already exists", rcpt.ReceiptId.Value()))
		return
	}

	uuid, err := receipt.NewReceiptId(common.GenerateUuID())
	if err != nil {
		return
	}

	r.Add(uuid, *rcpt)

	return uuid, nil
}

func (r *ReceiptInMemory) FindReceipt(_ context.Context, receiptId *receipt.ReceiptId) (receipt.ReceiptPoints, error) {
	receipt, ok := r.Get(*receiptId)
	if !ok {
		return receipt.ReceiptPoints, wrongs.StatusNotFound(fmt.Sprintf("receipt with id '%s' not found", receiptId.Value()))
	}

	return receipt.ReceiptPoints, nil
}
