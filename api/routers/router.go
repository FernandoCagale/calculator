package routers

import (
	"github.com/FernandoCagale/calculator/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler     *handlers.HealthHandler
	calculatorHandler *handlers.CalculatorHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/calculator/user/{user-id}/product/{product-id}", routes.calculatorHandler.Calculator).Methods("GET")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, calculatorHandler *handlers.CalculatorHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler:     healthHandler,
		calculatorHandler: calculatorHandler,
	}
}
