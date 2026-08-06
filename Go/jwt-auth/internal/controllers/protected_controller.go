package controllers

import (
	"encoding/json"
	"jwt-auth/internal/services"
	"net/http"
)

type ProtectedController struct {
	protectedService *services.ProtectedService
}

func NewProtectedController() *ProtectedController {
	return &ProtectedController{}
}

func (c *ProtectedController) Hey(w http.ResponseWriter, r *http.Request) {

	msg, err := c.protectedService.Hey()
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(msg)
}
