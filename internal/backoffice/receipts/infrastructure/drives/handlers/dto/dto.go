package dto

type ItemRequest struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type ReceiptRequest struct {
	Id           string        `json:"id"`
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        float64       `json:"total"`
	ItemsRequest []ItemRequest `json:"items"`
}
