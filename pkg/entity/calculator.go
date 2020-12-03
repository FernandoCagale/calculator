package entity

import (
	"time"
)

type Calculator struct {
	Product
	Discount Discount `json:"discount"`
}

func NewCalculator(product *Product) *Calculator {
	return &Calculator{
		Product: Product{
			ID:           product.ID,
			Title:        product.Title,
			Description:  product.Description,
			PriceInCents: product.PriceInCents,
		},
	}
}

func (e *Calculator) DiscountCalculator(date time.Time, user *User) {
	percentage := 0

	if user.IsDateOfBirth(date) {
		percentage = 5
	}

	if e.isBlackFriday(date) {
		percentage += 10
	}

	if percentage > 10 {
		percentage = 10
	}

	e.Discount = Discount{
		Percentage: percentage,
	}

	e.calculate()
}

func (e *Calculator) calculate() {
	e.Discount.ValueInCents = e.PriceInCents - (e.PriceInCents / 100 * e.Discount.Percentage)
}

func (e *Calculator) isBlackFriday(date time.Time) bool {
	if date.Month() == 11 && date.Day() == 25 {
		return true
	}
	return false
}
