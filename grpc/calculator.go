package grpc

import (
	"context"
	"github.com/FernandoCagale/calculator/pkg/domain/calculator"
)

type ServerGRPC struct {
	usecase calculator.UseCase
}

func NewCalculatorGRPC(usecase calculator.UseCase) *ServerGRPC {
	return &ServerGRPC{
		usecase: usecase,
	}
}

func (server *ServerGRPC) Calculator(ctx context.Context, in *Request) (*Reply, error) {
	calculator, err := server.usecase.Calculator(in.ProductId, in.UserId)
	if err != nil {
		return nil, err
	}
	return &Reply{
		Id:           calculator.ID,
		Title:        calculator.Title,
		Description:  calculator.Description,
		PriceInCents: int32(calculator.PriceInCents),
		Discount: &Discount{
			Percentage:   int32(calculator.Discount.Percentage),
			ValueInCents: int32(calculator.Discount.ValueInCents),
		},
	}, nil
}
