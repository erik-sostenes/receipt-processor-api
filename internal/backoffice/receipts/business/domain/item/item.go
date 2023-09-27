package item

import "github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"

type (
	Item struct {
		ItemShortDescription ItemShortDescription
		ItemPrice            ItemPrice
	}

	Items []Item
)

func NewItem(itemRequest dto.ItemRequest) (*Item, error) {
	itemShortDescription, err := NewItemShortDescription(itemRequest.ShortDescription)
	if err != nil {
		return &Item{}, err
	}

	itemPrice, err := NewItemPrice(itemRequest.Price)
	if err != nil {
		return &Item{}, err
	}

	return &Item{
		ItemShortDescription: itemShortDescription,
		ItemPrice:            itemPrice,
	}, nil
}

func ToMap(items []dto.ItemRequest) (Items, error) {
	var i Items

	for _, v := range items {
		item, err := NewItem(v)
		if err != nil {
			return i, err
		}

		i = append(i, *item)
	}

	return i, nil
}
