package receipt

import (
	"time"

	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
)

type ReceiptPurchaseTime struct {
	value string
}

func NewReceiptPurchaseTime(value string) (ReceiptPurchaseTime, error) {
	_, err := common.Timestamp(value).Validate(TimeLayout)

	return ReceiptPurchaseTime{value}, err
}

func (r ReceiptPurchaseTime) Value() string {
	return r.value
}

func (r ReceiptPurchaseTime) CalculatePoints() uint8 {
	hour, err := time.Parse(TimeLayout, r.value)
	if err != nil {
		return 0
	}

	start := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	end := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)

	if hour.After(start) && hour.Before(end) {
		return 10
	}

	return 0
}
