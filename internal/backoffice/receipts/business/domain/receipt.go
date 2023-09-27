package domain

// Item represents a domain object
type Item struct {
	ShortDescription string
	Price            float64
}

// Receipt represents a domain object
type Receipt struct {
	Id           string
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        float64
	Items        []Item
}
