package calculator

import (
	"github.com/FernandoCagale/calculator/internal/errors"
	"github.com/FernandoCagale/calculator/pkg/entity"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	PRODUCTS = "products"
	USERS    = "users"
	DATABASE = "calculator"
)

type MongodbRepository struct {
	session *mgo.Session
}

func NewMongodbRepository(session *mgo.Session) *MongodbRepository {
	repo := &MongodbRepository{session}

	productsColl := repo.session.DB(DATABASE).C(PRODUCTS)
	_, _ = productsColl.UpsertId("001", &entity.Product{
		ID:           "001",
		Description:  "Chivas",
		Title:        "Whisky Chivas",
		PriceInCents: 12000,
	})

	_, _ = productsColl.UpsertId("002", &entity.Product{
		ID:           "002",
		Description:  "Blue Label",
		Title:        "Whisky Blue Label",
		PriceInCents: 25000,
	})
	_, _ = productsColl.UpsertId("003", &entity.Product{
		ID:           "003",
		Description:  "Dalmore",
		Title:        "Whisky Dalmore",
		PriceInCents: 55000,
	})

	usersColl := repo.session.DB(DATABASE).C(USERS)
	_, _ = usersColl.UpsertId("001", &entity.User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, 12, 02, 0, 0, 0, 0, time.UTC),
	})
	_, _ = usersColl.UpsertId("002", &entity.User{
		Id:          "002",
		FirstName:   "Matt",
		LastName:    "Bridges",
		DateOfBirth: time.Date(1988, 11, 25, 0, 0, 0, 0, time.UTC),
	})
	_, _ = usersColl.UpsertId("003", &entity.User{
		Id:          "003",
		FirstName:   "Aidan",
		LastName:    "Craig",
		DateOfBirth: time.Date(1985, 05, 10, 0, 0, 0, 0, time.UTC),
	})

	return repo
}

func (repo *MongodbRepository) FindByIdProduct(ID string) (product *entity.Product, err error) {
	coll := repo.session.DB(DATABASE).C(PRODUCTS)

	err = coll.FindId(ID).One(&product)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return product, nil
}

func (repo *MongodbRepository) FindByIdUser(ID string) (user *entity.User, err error) {
	coll := repo.session.DB(DATABASE).C(USERS)

	err = coll.FindId(ID).One(&user)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return user, nil
}
