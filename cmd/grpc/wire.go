//+build wireinject

package main

import (
	"github.com/FernandoCagale/calculator/grpc"
	"github.com/FernandoCagale/calculator/internal/datastore"
	"github.com/google/wire"
	"gopkg.in/mgo.v2"
)

func SetupApplication(session *mgo.Session) (*grpc.ServerGRPC, error) {
	wire.Build(grpc.Set)
	return nil, nil
}

func SetupMongoDB() (*mgo.Session, error) {
	wire.Build(datastore.Set)
	return nil, nil
}