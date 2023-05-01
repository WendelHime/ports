// Package endpoints holds all the REST HTTP endpoints implementations
package endpoints

import (
	"net/http"

	"github.com/WendelHime/ports/internal/logic"
	localErrs "github.com/WendelHime/ports/internal/shared/errors"
)

// PortHTTP holds the logic service being used by the port endpoints
type PortHTTP struct {
	service logic.PortDomainService
}

func NewPortHTTP(service logic.PortDomainService) *PortHTTP {
	return &PortHTTP{
		service: service,
	}
}

// SyncPorts is an upsert endpoint that insert/update ports data
func (h *PortHTTP) SyncPorts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := h.service.SyncPorts(r.Context(), r.Body)
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetPortByUnloc retrieves the port data based on unloc provided parameter
func (h *PortHTTP) GetPortByUnloc(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func respondError(w http.ResponseWriter, err error) {
	if err != nil {
		switch err {
		case localErrs.ErrBadRequest:
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
		case localErrs.ErrInternalServerError:
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
