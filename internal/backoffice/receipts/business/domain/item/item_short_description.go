package item

import "github.com/erik-sostenes/receipt-processor-api/pkg/common"

type ItemShortDescription struct {
	value string
}

func NewItemShortDescription(value string) (ItemShortDescription, error) {
	v, err := common.String(value).Validate("item -> short description")
	if err != nil {
		return ItemShortDescription{}, err
	}

	return ItemShortDescription{
		value: v,
	}, nil
}

func (i ItemShortDescription) Value() string {
	return i.value
}
