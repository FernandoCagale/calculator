package handlers

import (
	"github.com/hellofresh/health-go"
	"net/http"
)

type HealthHandler struct {
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	health.HandlerFunc(w, r)
}
