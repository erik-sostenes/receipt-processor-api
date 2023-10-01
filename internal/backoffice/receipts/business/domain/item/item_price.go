package item

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

// Item represents a domain object
type ItemPrice struct {
	value float64
}

func NewItemPrice(value string) (ItemPrice, error) {
	v, err := common.Float(value).Validate()
	return ItemPrice{
		value: v,
	}, err
}

func (r ItemPrice) Value() float64 {
	return r.value
}
