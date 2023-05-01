package endpoints

import (
	"net/http"

	"github.com/WendelHime/ports/internal/logic"
	localErrs "github.com/WendelHime/ports/internal/shared/errors"
)

type PortHTTP struct {
	service logic.PortDomainService
}

func NewPortHTTP(service logic.PortDomainService) *PortHTTP {
	return &PortHTTP{
		service: service,
	}
}

func (h *PortHTTP) SyncPorts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	err := h.service.SyncPorts(r.Context(), r.Body)
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

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
