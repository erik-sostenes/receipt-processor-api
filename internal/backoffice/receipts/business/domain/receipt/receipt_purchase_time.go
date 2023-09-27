package receipt

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

type ReceiptPurchaseTime struct {
	value int64
}

func NewReceiptPurchaseTime(value string) (ReceiptPurchaseTime, error) {
	v, err := common.Timestamp(value).Validate(TimeLayout)

	return ReceiptPurchaseTime{v}, err
}

func (r ReceiptPurchaseTime) Value() int64 {
	return r.value
}
