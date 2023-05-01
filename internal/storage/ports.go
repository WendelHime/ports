// Package storage contains structures that communicate directly with the storage layer
package storage

//go:generate mockgen -destination=./ports_mock.go -package=storage github.com/WendelHime/ports/internal/storage PortRepository

import (
	"context"

	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/WendelHime/ports/internal/shared/models"
)

// PortRepository is an repository for creating and updating port correlated data
type PortRepository interface {
	Create(ctx context.Context, port models.Port) error
	Update(ctx context.Context, port models.Port) error
	Get(ctx context.Context, unloc string) (models.Port, error)
}

type portRepo struct {
	ports map[string]models.Port
}

func NewPortRepository() PortRepository {
	return &portRepo{
		ports: make(map[string]models.Port),
	}
}

func (r *portRepo) Create(ctx context.Context, port models.Port) error {
	for _, unloc := range port.Unlocs {
		r.ports[unloc] = port
	}
	return nil
}

func (r *portRepo) Update(ctx context.Context, port models.Port) error {
	for _, unloc := range port.Unlocs {
		r.ports[unloc] = port
	}
	return nil
}

func (r *portRepo) Get(ctx context.Context, unloc string) (models.Port, error) {
	if port, exists := r.ports[unloc]; exists {
		return port, nil
	}
	return models.Port{}, localErrs.ErrNotFound
}
