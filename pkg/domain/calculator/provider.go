package calculator

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase, NewInMemoryRepository,
	wire.Bind(new(Repository), new(*InMemoryRepository)), // in InMemoryRepository
	wire.Bind(new(UseCase), new(*UseCaseCalculator)))
