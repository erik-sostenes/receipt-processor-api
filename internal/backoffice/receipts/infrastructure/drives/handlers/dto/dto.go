package dto

type ItemRequest struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type ReceiptsRequest struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        float64       `json:"total"`
	ItemRequest  []ItemRequest `json:"items"`
}
