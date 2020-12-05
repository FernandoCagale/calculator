package calculator

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase, NewInMemoryRepository, NewMongodbRepository,
	//wire.Bind(new(Repository), new(*InMemoryRepository)), // in Memory
	wire.Bind(new(Repository), new(*MongodbRepository)),    // in Mongo
	wire.Bind(new(UseCase), new(*UseCaseCalculator)))
