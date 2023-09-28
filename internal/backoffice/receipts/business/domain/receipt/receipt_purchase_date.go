package receipt

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

type ReceiptPurchaseDate struct {
	value string
}

func NewReceiptPurchaseDate(value string) (ReceiptPurchaseDate, error) {
	v, err := common.Timestamp(value).Validate(DateLayout)

	return ReceiptPurchaseDate{v}, err
}

func (r ReceiptPurchaseDate) Value() string {
	return r.value
}
