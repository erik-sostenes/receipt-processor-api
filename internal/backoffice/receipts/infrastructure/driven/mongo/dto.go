package mongo

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain/receipt"
	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
)

// NewReceipt builds an instance of *Receipt based on the *receipt.Receipt
func NewReceipt(receipt *receipt.Receipt) *Receipt {
	items := make([]*Item, 0, len(receipt.Items))

	for _, item := range receipt.Items {
		items = append(items, &Item{
			ShortDescription: item.ItemShortDescription.Value(),
			Price:            item.ItemPrice.Value(),
		})
	}

	var receiptId string

	if receipt.ReceiptId.Value() == "" {
		receiptId = common.GenerateUuID()
	} else {
		receiptId = receipt.ReceiptId.Value()
	}

	return &Receipt{
		ID:           receiptId,
		Reteiler:     receipt.ReceiptRetailer.Value(),
		PurchaseDate: receipt.ReceiptPurchaseDate.Value(),
		PurchaseTime: receipt.ReceiptPurchaseTime.Value(),
		Items:        items,
		TotalPoints:  receipt.ReceiptPoints.Value(),
	}
}

type Receipt struct {
	ID           string  `bson:"_id"`
	Reteiler     string  `bson:"retailer,omitempty"`
	PurchaseDate string  `bson:"purchase_date,omitempty"`
	PurchaseTime string  `bson:"purchase_time,omitempty"`
	Items        []*Item `bson:"items,omitempty"`
	TotalPoints  uint8   `bson:"total_points, omitempty"`
}

type Item struct {
	ShortDescription string  `bson:"short_description,omitempty"`
	Price            float64 `bson:"price,omitempty"`
}
