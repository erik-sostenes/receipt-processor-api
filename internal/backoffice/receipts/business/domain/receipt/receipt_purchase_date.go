package receipt

import (
	"time"

	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
)

type ReceiptPurchaseDate struct {
	value string
}

func NewReceiptPurchaseDate(value string) (ReceiptPurchaseDate, error) {
	_, err := common.Timestamp(value).Validate(DateLayout)

	return ReceiptPurchaseDate{value: value}, err
}

func (r ReceiptPurchaseDate) Value() string {
	return r.value
}

func (r ReceiptPurchaseDate) CalculatePoints() uint8 {
	date, err := time.Parse(DateLayout, r.value)
	if err != nil {
		return 0
	}

	if date.Day()%2 != 0 {
		return 6
	}

	return 0
}
