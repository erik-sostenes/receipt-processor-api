package services

import (
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/business/domain"
	"github.com/erik-sostenes/receipt-processor-api/internal/backoffice/receipts/infrastructure/drives/handlers/dto"
)

// map dto data to a business object
func ToReceipt(dto *dto.ReceiptRequest) *domain.Receipt {
	var items []domain.Item
	for _, v := range dto.ItemsRequest {
		items = append(items, domain.Item(v))
	}

	return &domain.Receipt{
		Id:           dto.Id,
		Retailer:     dto.Retailer,
		PurchaseDate: dto.PurchaseDate,
		PurchaseTime: dto.PurchaseTime,
		Total:        dto.Total,
		Items:        items,
	}
}
