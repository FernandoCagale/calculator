package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_IsBlackFriday_OK(t *testing.T) {
	calculator := &Calculator{}

	blackFriday := calculator.isBlackFriday(time.Date(2020, 11, 25, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, true, blackFriday)
}

func Test_IsBlackFriday_NOK(t *testing.T) {
	calculator := &Calculator{}

	blackFriday := calculator.isBlackFriday(time.Date(2020, 01, 10, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, false, blackFriday)
}

func Test_DiscountCalculator_Discount_Zero(t *testing.T) {

	calculator := &Calculator{
		Product: Product{
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
	}

	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, 01, 01, 0, 0, 0, 0, time.UTC),
	}

	calculator.DiscountCalculator(time.Now(), user)

	assert.Equal(t, 0, calculator.Discount.ValueInCents)
	assert.Equal(t, 0, calculator.Discount.Percentage)
}

func Test_DiscountCalculator_Is_DateOfBirth_Discount_5(t *testing.T) {
	calculator := &Calculator{
		Product: Product{
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
	}

	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	}

	calculator.DiscountCalculator(time.Now(), user)

	assert.Equal(t, 11400, calculator.Discount.ValueInCents)
	assert.Equal(t, 5, calculator.Discount.Percentage)
}

func Test_DiscountCalculator_Is_BlackFriday_Discount_10(t *testing.T) {
	calculator := &Calculator{
		Product: Product{
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
	}

	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	calculator.DiscountCalculator(time.Date(2020, 11, 25, 0, 0, 0, 0, time.UTC), user)

	assert.Equal(t, 10800, calculator.Discount.ValueInCents)
	assert.Equal(t, 10, calculator.Discount.Percentage)
}

func Test_DiscountCalculator_Is_BlackFriday_And_DateOfBirth_Discount_10(t *testing.T) {
	calculator := &Calculator{
		Product: Product{
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
	}

	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	}

	calculator.DiscountCalculator(time.Date(2020, 11, 25, 0, 0, 0, 0, time.UTC), user)

	assert.Equal(t, 10800, calculator.Discount.ValueInCents)
	assert.Equal(t, 10, calculator.Discount.Percentage)
}

func Test_DiscountCalculator_Calculate(t *testing.T) {
	calculator := &Calculator{
		Product: Product{
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
		Discount: Discount{
			Percentage: 10,
			ValueInCents: 0,
		},
	}

	calculator.calculate()

	assert.Equal(t, 10800, calculator.Discount.ValueInCents)
}