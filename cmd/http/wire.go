//+build wireinject

package main

import (
	"github.com/FernandoCagale/calculator/api/routers"
	"github.com/FernandoCagale/calculator/pkg"
	"github.com/FernandoCagale/calculator/internal/datastore"
	"github.com/google/wire"
	"gopkg.in/mgo.v2"
)

func SetupApplication(*mgo.Session) (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}

func SetupMongoDB() (*mgo.Session, error) {
	wire.Build(datastore.Set)
	return nil, nil
}