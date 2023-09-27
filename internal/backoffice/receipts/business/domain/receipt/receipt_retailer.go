package receipt

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

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
