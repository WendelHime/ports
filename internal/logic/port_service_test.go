package logic

import (
	"context"
	"io"
	"strings"
	"testing"

	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/WendelHime/ports/internal/shared/models"
	"github.com/WendelHime/ports/internal/storage"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSyncPorts(t *testing.T) {
	ctx := context.Background()
	var tests = []struct {
		name   string
		assert func(t *testing.T, err error)
		setup  func(t *testing.T) (io.Reader, PortDomainService)
	}{
		{
			name: "add new ports with success",
			assert: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			setup: func(t *testing.T) (io.Reader, PortDomainService) {
				ctrl := gomock.NewController(t)
				portRepo := storage.NewMockPortRepository(ctrl)
				portRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(models.Port{}, localErrs.ErrNotFound).MaxTimes(3)
				portRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).MaxTimes(3)

				return strings.NewReader(threeRandomPorts()), NewPortDomainService(portRepo)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, service := tt.setup(t)
			err := service.SyncPorts(ctx, input)
			tt.assert(t, err)
		})
	}
}

func threeRandomPorts() string {
	return `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  }
}
`
}
