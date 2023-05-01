package storage

import (
	"context"
	"sync"
	"testing"

	"github.com/WendelHime/ports/internal/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestInMemRepository(t *testing.T) {
	var tests = []struct {
		name   string
		assert func(t *testing.T, repo portRepo, err error)
		exec   func(repo portRepo) error
		setup  func(t *testing.T) portRepo
	}{
		{
			name: "Create port with success",
			assert: func(t *testing.T, repo portRepo, err error) {
				assert.Nil(t, err)
				assert.Contains(t, repo.ports, "UNLOC")
			},

			exec: func(repo portRepo) error {
				err := repo.Create(context.Background(), models.Port{
					Unlocs: []string{"UNLOC"},
				})
				return err
			},
			setup: func(*testing.T) portRepo {
				return portRepo{
					ports: make(map[string]models.Port),
					mutex: new(sync.Mutex),
				}
			},
		},
		{
			name: "Get port with success",
			assert: func(t *testing.T, repo portRepo, err error) {
				// TODO: get tests should be moved to a different test since we're not validating here if they're
				// returning the stored data properly
				assert.Nil(t, err)
				assert.Contains(t, repo.ports, "UNLOC")
			},

			exec: func(repo portRepo) error {
				_, err := repo.Get(context.Background(), "UNLOC")
				return err
			},
			setup: func(*testing.T) portRepo {
				repo := portRepo{
					ports: make(map[string]models.Port),
					mutex: new(sync.Mutex),
				}
				err := repo.Create(context.Background(), models.Port{
					Unlocs: []string{"UNLOC"},
				})
				assert.NoError(t, err)
				return repo
			},
		},
		{
			name: "Get port with success",
			assert: func(t *testing.T, repo portRepo, err error) {
				assert.Nil(t, err)
				assert.Contains(t, repo.ports, "UNLOC")
				port, err := repo.Get(context.Background(), "UNLOC")
				assert.NoError(t, err)
				assert.Equal(t, port.Code, "updated")
			},

			exec: func(repo portRepo) error {
				err := repo.Update(context.Background(), models.Port{
					Code:   "updated",
					Unlocs: []string{"UNLOC"},
				})
				return err
			},
			setup: func(*testing.T) portRepo {
				repo := portRepo{
					ports: make(map[string]models.Port),
					mutex: new(sync.Mutex),
				}
				err := repo.Create(context.Background(), models.Port{
					Unlocs: []string{"UNLOC"},
				})
				assert.NoError(t, err)
				return repo
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			portRepo := tt.setup(t)
			err := tt.exec(portRepo)
			tt.assert(t, portRepo, err)
		})
	}
}
