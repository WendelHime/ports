package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WendelHime/ports/internal/logic"
	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/WendelHime/ports/internal/shared/models"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncPorts(t *testing.T) {
	var tests = []struct {
		name   string
		assert func(t *testing.T, w *httptest.ResponseRecorder)
		setup  func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder)
	}{
		{
			name: "Sync with success should return a ok response",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(nil).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with invalid body should return a bad request",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(localErrs.ErrBadRequest).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with internal server error should return a internal server error",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(localErrs.ErrInternalServerError).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with unexpected error should return a internal server error",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(errors.New("random error")).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			portHTTP, req, w := tt.setup(t)
			portHTTP.SyncPorts(w, req)
			tt.assert(t, w)
		})
	}
}

func TestGetPortByUnloc(t *testing.T) {
	var tests = []struct {
		name   string
		assert func(t *testing.T, w *httptest.ResponseRecorder)
		setup  func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder)
	}{
		{
			name: "Get port with success should return a ok response",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
				returnedPort := models.Port{}
				err := json.Unmarshal(w.Body.Bytes(), &returnedPort)
				assert.Nil(t, err)
				assert.Equal(t, models.Port{Unlocs: []string{"aaaa"}}, returnedPort)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/ports/{unloc}", nil)
				w := httptest.NewRecorder()

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("unloc", "aaaa")
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().GetPort(req.Context(), "aaaa").Return(models.Port{
					Unlocs: []string{"aaaa"},
				}, nil).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Port not found should return a not found error",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, w.Code)
			},
			setup: func(t *testing.T) (*PortHandlers, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/ports/{unloc}", nil)
				w := httptest.NewRecorder()

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("unloc", "aaaa")
				req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().GetPort(req.Context(), "aaaa").Return(models.Port{}, localErrs.ErrNotFound).Times(1)

				portHTTP := NewPortHTTPHandlers(portService)
				return portHTTP, req, w
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			portHTTP, req, w := tt.setup(t)
			portHTTP.GetPortByUnloc(w, req)
			tt.assert(t, w)
		})
	}
}
