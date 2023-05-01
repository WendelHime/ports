package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WendelHime/ports/internal/logic"
	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncPorts(t *testing.T) {
	var tests = []struct {
		name   string
		assert func(t *testing.T, w *httptest.ResponseRecorder)
		setup  func(t *testing.T) (*PortHTTP, *http.Request, *httptest.ResponseRecorder)
	}{
		{
			name: "Sync with success should return a ok response",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)
			},
			setup: func(t *testing.T) (*PortHTTP, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(nil).Times(1)

				portHTTP := NewPortHTTP(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with invalid body should return a bad request",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
			setup: func(t *testing.T) (*PortHTTP, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(localErrs.ErrBadRequest).Times(1)

				portHTTP := NewPortHTTP(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with internal server error should return a internal server error",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
			setup: func(t *testing.T) (*PortHTTP, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(localErrs.ErrInternalServerError).Times(1)

				portHTTP := NewPortHTTP(portService)
				return portHTTP, req, w
			},
		},
		{
			name: "Sync with unexpected error should return a internal server error",
			assert: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			},
			setup: func(t *testing.T) (*PortHTTP, *http.Request, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/ports", nil)
				w := httptest.NewRecorder()

				ctrl := gomock.NewController(t)
				portService := logic.NewMockPortDomainService(ctrl)
				portService.EXPECT().SyncPorts(req.Context(), gomock.Any()).Return(errors.New("random error")).Times(1)

				portHTTP := NewPortHTTP(portService)
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
