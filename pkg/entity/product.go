package entity

type Product struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PriceInCents int    `json:"price_in_cents"`
}

type Discount struct {
	Percentage   int `json:"percentage"`
	ValueInCents int `json:"value_in_cents"`
}
