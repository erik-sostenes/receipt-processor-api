package item

import (
	"math"
	"strings"

	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
)

type (
	Item struct {
		ItemShortDescription ItemShortDescription
		ItemPrice            ItemPrice
	}

	Items []Item
)

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

func (i Item) CalculatePoints() uint8 {
	if len(strings.TrimSpace(i.ItemShortDescription.Value()))%3 == 0 {
		return uint8(math.Ceil(i.ItemPrice.Value() * 0.2))
	}
	return 0
}

func (i Items) CalculatePoints() uint8 {
	var counter, accumulate uint8

	for _, v := range i {
		counter++

		accumulate += v.CalculatePoints()

		if counter == 2 {
			accumulate += 5
			counter = 0
		}
	}

	return accumulate
}
