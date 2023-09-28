package receipt

import (
	"strings"
	"unicode"

	"github.com/erik-sostenes/receipt-processor-api/pkg/common"
)

type ReceiptRetailer struct {
	value string
}

func NewReceiptRetailer(value string) (ReceiptRetailer, error) {
	v, err := common.String(value).Validate("receipt -> retailer")
	if err != nil {
		return ReceiptRetailer{}, err
	}

	return ReceiptRetailer{
		value: v,
	}, nil
}

func (r ReceiptRetailer) Value() string {
	return r.value
}

func (r ReceiptRetailer) CalculatePoints() uint8 {
	var accumulate uint8

	for _, v := range strings.ReplaceAll(r.value, " ", "") {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			accumulate++
		}
	}

	return accumulate
}
