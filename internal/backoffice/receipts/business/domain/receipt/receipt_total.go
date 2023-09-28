package receipt

import (
	"math"

	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
)

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

func (r ReceiptTotal) CalculatePoints() uint8 {
	var accumulate uint8

	if math.Mod(r.value, 1) == 0.0 {
		accumulate += 50
	}

	if math.Mod(float64(r.value), float64(0.25)) == 0 {
		accumulate += 25
	}

	return accumulate
}
