package pkg

import (
	"github.com/FernandoCagale/calculator/api/handlers"
	"github.com/FernandoCagale/calculator/api/routers"
	"github.com/FernandoCagale/calculator/pkg/domain/calculator"
	"github.com/google/wire"
)

var Container = wire.NewSet(calculator.Set, handlers.Set, routers.Set)
