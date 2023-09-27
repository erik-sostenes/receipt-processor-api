package receipt

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

type ReceiptTotal struct {
	value float64
}

func NewReceiptTotal(value string) (ReceiptTotal, error) {
	v, err := common.Float(value).Validate()

	return ReceiptTotal{
		value: v,
	}, err
}

func (r ReceiptTotal) Value() float64 {
	return r.value
}
