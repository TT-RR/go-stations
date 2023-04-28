package handler

import (
	"encoding/json"
	//"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := &model.HealthzResponse{}
	u.Message = "OK"

	/*
		if err != nil {
			log.Println(err)
		}
	*/

	json.NewEncoder(w).Encode(u)

}
