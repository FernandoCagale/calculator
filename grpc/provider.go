package grpc

import (
	"github.com/FernandoCagale/calculator/pkg/domain/calculator"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewCalculatorGRPC, calculator.Set)
