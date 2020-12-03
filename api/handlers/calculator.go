package handlers

import (
	"github.com/FernandoCagale/calculator/api/render"
	"github.com/FernandoCagale/calculator/internal/errors"
	"github.com/FernandoCagale/calculator/pkg/domain/calculator"
	"github.com/gorilla/mux"
	"net/http"
)

type CalculatorHandler struct {
	usecase calculator.UseCase
}

func NewCalculatorHandler(usecase calculator.UseCase) *CalculatorHandler {
	return &CalculatorHandler{
		usecase: usecase,
	}
}

func (handler *CalculatorHandler) Calculator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID := vars["user-id"]
	productID := vars["product-id"]

	response, err := handler.usecase.Calculator(productID, userID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, response, http.StatusOK)
}