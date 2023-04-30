// Package storage contains structures that communicate directly with the storage layer
package storage

//go:generate mockgen -destination=./ports_mock.go -package=storage github.com/WendelHime/ports/internal/storage PortRepository

import (
	"context"

	"github.com/WendelHime/ports/internal/shared/models"
)

// PortRepository is an repository for creating and updating port correlated data
type PortRepository interface {
	Create(ctx context.Context, port models.Port) error
	Update(ctx context.Context, port models.Port) error
	Get(ctx context.Context, unlocs string) (models.Port, error)
}
