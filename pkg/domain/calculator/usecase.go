package calculator

import (
	"github.com/FernandoCagale/calculator/pkg/entity"
)

type UseCase interface {
	Calculator(productID string, userID string) (calculator *entity.Calculator, err error)
}
