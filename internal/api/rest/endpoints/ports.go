// Package endpoints holds all the REST HTTP endpoints implementations
package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/WendelHime/ports/internal/logic"
	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/go-chi/chi/v5"
)

// PortHandlers holds the logic service being used by the port endpoints
type PortHandlers struct {
	service logic.PortDomainService
}

func NewPortHTTPHandlers(service logic.PortDomainService) *PortHandlers {
	return &PortHandlers{
		service: service,
	}
}

// SyncPorts is an upsert endpoint that insert/update ports data
func (h *PortHandlers) SyncPorts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := h.service.SyncPorts(r.Context(), r.Body)
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetPortByUnloc retrieves the port data based on unloc provided parameter
func (h *PortHandlers) GetPortByUnloc(w http.ResponseWriter, r *http.Request) {
	unloc := chi.URLParam(r, "unloc")
	port, err := h.service.GetPort(r.Context(), unloc)
	if err != nil {
		respondError(w, err)
		return
	}

	b, err := json.Marshal(port)
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func respondError(w http.ResponseWriter, err error) {
	if err != nil {
		var statusCode int
		switch err {
		case localErrs.ErrBadRequest:
			statusCode = http.StatusBadRequest
		case localErrs.ErrInternalServerError:
			statusCode = http.StatusInternalServerError
		case localErrs.ErrNotFound:
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, err.Error(), statusCode)
	}
}
