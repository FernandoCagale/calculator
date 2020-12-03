package calculator

import (
	"github.com/FernandoCagale/calculator/pkg/entity"
	"time"
)

type UseCaseCalculator struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCaseCalculator {
	return &UseCaseCalculator{
		repo: repo,
	}
}

func (usecase *UseCaseCalculator) Calculator(productID string, userID string) (calculator *entity.Calculator, err error) {
	product, err := usecase.repo.FindByIdProduct(productID)
	if err != nil {
		return nil, err
	}

	calculator = entity.NewCalculator(product)

	user, err := usecase.repo.FindByIdUser(userID)
	if err != nil {
		return nil, err
	}

	calculator.DiscountCalculator(time.Now(), user)

	return calculator, nil
}
