package calculator

import "github.com/FernandoCagale/calculator/pkg/entity"

type Repository interface {
	FindByIdProduct(ID string) (product *entity.Product, err error)
	FindByIdUser(ID string) (user *entity.User, err error)
}
