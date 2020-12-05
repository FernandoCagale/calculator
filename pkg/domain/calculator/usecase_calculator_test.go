package calculator

import (
	"github.com/FernandoCagale/calculator/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calculator_OK(t *testing.T) {
	usecase := NewUseCase(NewInMemoryRepository())

	calculator, err := usecase.Calculator("001", "001")

	assert.Nil(t, err)
	assert.Equal(t, calculator.ID, "001")
	assert.Equal(t, calculator.Title, "Whisky Chivas")
	assert.Equal(t, calculator.Description, "Chivas")
	assert.Equal(t, calculator.PriceInCents, 12000)
	assert.Equal(t, calculator.Discount.ValueInCents, 0)
	assert.Equal(t, calculator.Discount.Percentage, 0)
}

func Test_Calculator_NotFound_User(t *testing.T) {
	usecase := NewUseCase(NewInMemoryRepository())

	calculator, err := usecase.Calculator("001", "004")

	assert.Nil(t, calculator)
	assert.Equal(t, err, errors.ErrNotFound)
	assert.Equal(t, err.Error(), "Not found")
}

func Test_Calculator_NotFound_Product(t *testing.T) {
	usecase := NewUseCase(NewInMemoryRepository())

	calculator, err := usecase.Calculator("004", "001")

	assert.Nil(t, calculator)
	assert.Equal(t, err, errors.ErrNotFound)
	assert.Equal(t, err.Error(), "Not found")
}
