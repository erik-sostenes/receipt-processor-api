package receipt

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/item"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
)

// DateLayout format the dates
const DateLayout = "2006-01-02"

// TimeLayout format the dates
const TimeLayout = "15:04"

// Receipt represents a domain object
type Receipt struct {
	ReceiptId           ReceiptId
	ReceiptRetailer     ReceiptRetailer
	ReceiptPurchaseDate ReceiptPurchaseDate
	ReceiptPurchaseTime ReceiptPurchaseTime
	ReceiptTotal        ReceiptTotal
	Items               item.Items
}

func NewReceipt(receiptRequest dto.ReceiptRequest) (*Receipt, error) {
	receiptRetailer, err := NewReceiptRetailer(receiptRequest.Retailer)
	if err != nil {
		return &Receipt{}, err
	}

	receiptPurchaseDate, err := NewReceiptPurchaseDate(receiptRequest.PurchaseDate)
	if err != nil {
		return &Receipt{}, err
	}

	receiptPurchaseTime, err := NewReceiptPurchaseTime(receiptRequest.PurchaseTime)
	if err != nil {
		return &Receipt{}, err
	}

	receiptTotal, err := NewReceiptTotal(receiptRequest.Total)
	if err != nil {
		return &Receipt{}, err
	}

	items, err := item.ToMap(receiptRequest.ItemsRequest)
	if err != nil {
		return &Receipt{}, err
	}
	return &Receipt{
		ReceiptRetailer:     receiptRetailer,
		ReceiptPurchaseDate: receiptPurchaseDate,
		ReceiptPurchaseTime: receiptPurchaseTime,
		ReceiptTotal:        receiptTotal,
		Items:               items,
	}, nil
}