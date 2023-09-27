package dto

type ItemRequest struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptRequest struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Total        string        `json:"total"`
	ItemsRequest []ItemRequest `json:"items"`
}
