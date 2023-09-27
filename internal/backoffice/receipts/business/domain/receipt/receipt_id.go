package receipt

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

type ReceiptId struct {
	value string
}

func NewReceiptId(value string) (ReceiptId, error) {
	v, err := common.Identifier(value).Validate()

	return ReceiptId{v}, err
}

func (r ReceiptId) Value() string {
	return r.value
}
