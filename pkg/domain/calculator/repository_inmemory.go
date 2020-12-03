package calculator

import (
	"github.com/FernandoCagale/calculator/internal/errors"
	"github.com/FernandoCagale/calculator/pkg/entity"
	"time"
)

type InMemoryRepository struct {
	dbProducts map[string]*entity.Product
	dbUsers    map[string]*entity.User
}

func NewInMemoryRepository() *InMemoryRepository {
	dbProducts := map[string]*entity.Product{
		"001": {
			ID:           "001",
			Description:  "Chivas",
			Title:        "Whisky Chivas",
			PriceInCents: 12000,
		},
		"002": {
			ID:           "002",
			Description:  "Blue Label",
			Title:        "Whisky Blue Label",
			PriceInCents: 25000,
		},
		"003": {
			ID:           "003",
			Description:  "Dalmore",
			Title:        "Whisky Dalmore",
			PriceInCents: 55000,
		},
	}
	dbUsers := map[string]*entity.User{
		"001": {
			Id:          "001",
			FirstName:   "Rami",
			LastName:    "Marriott",
			DateOfBirth: time.Date(1990, 12, 02, 0, 0, 0, 0, time.UTC),
		},
		"002": {
			Id:          "002",
			FirstName:   "Matt",
			LastName:    "Bridges",
			DateOfBirth: time.Date(1988, 11, 25, 0, 0, 0, 0, time.UTC),
		},
		"003": {
			Id:          "003",
			FirstName:   "Aidan",
			LastName:    "Craig",
			DateOfBirth: time.Date(1985, 05, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	return &InMemoryRepository{dbProducts, dbUsers}
}

func (repo *InMemoryRepository) FindByIdProduct(ID string) (product *entity.Product, err error) {
	if repo.dbProducts[ID] == nil {
		return nil, errors.ErrNotFound
	}
	return repo.dbProducts[ID], nil
}

func (repo *InMemoryRepository) FindByIdUser(ID string) (user *entity.User, err error) {
	if repo.dbUsers[ID] == nil {
		return nil, errors.ErrNotFound
	}
	return repo.dbUsers[ID], nil

}
