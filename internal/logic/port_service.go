// Package logic implements the services logic
package logic

//go:generate mockgen -destination=./port_service_mock.go -package=logic github.com/WendelHime/ports/internal/logic PortDomainService

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	localErrs "github.com/WendelHime/ports/internal/shared/errors"
	"github.com/WendelHime/ports/internal/shared/models"
	"github.com/WendelHime/ports/internal/storage"
)

type PortDomainService interface {
	SyncPorts(ctx context.Context, ports io.Reader) error
	GetPort(ctx context.Context, unloc string) (models.Port, error)
}

type portLogic struct {
	repository storage.PortRepository
}

func NewPortDomainService(repo storage.PortRepository) PortDomainService {
	return &portLogic{
		repository: repo,
	}
}

func (l portLogic) GetPort(ctx context.Context, unloc string) (models.Port, error) {
	if unloc == "" {
		return models.Port{}, errors.Wrap(localErrs.ErrBadRequest, "invalid unloc provided")
	}
	port, err := l.repository.Get(ctx, unloc)
	return port, err
}

func (l portLogic) SyncPorts(ctx context.Context, ports io.Reader) error {
	decoder := json.NewDecoder(ports)
	portsIsEmpty := decoder.More()
	if !portsIsEmpty {
		return errors.Wrap(localErrs.ErrBadRequest, "port input is empty")
	}

	// getting first token "{"
	_, err := decoder.Token()
	if err != nil {
		return errors.Wrap(localErrs.ErrBadRequest, fmt.Sprintf("couldn't acquire first token from input: %+v", err))
	}

	for decoder.More() {
		// retrieving unloc
		unlocToken, err := decoder.Token()
		if err != nil {
			return errors.Wrap(localErrs.ErrInternalServerError, fmt.Sprintf("failed to acquire unloc: %+v", err))
		}
		unloc := unlocToken.(string)

		// decoding object
		var port models.Port
		err = decoder.Decode(&port)
		if err != nil {
			return errors.Wrap(localErrs.ErrInternalServerError, fmt.Sprintf("failed to decode port: %+v", err))
		}

		// checking if port/unloc exists on database
		_, err = l.repository.Get(ctx, unloc)
		if err != nil && err != localErrs.ErrNotFound {
			return errors.Wrap(localErrs.ErrInternalServerError, fmt.Sprintf("unexpected error when retrieving port info from database: %+v", err))
		}

		// if port doesn't exist, let's create!
		if err == localErrs.ErrNotFound {
			err = l.repository.Create(ctx, port)
			if err != nil {
				return errors.Wrap(localErrs.ErrInternalServerError, fmt.Sprintf("failed to create port [%+v] on storage: %+v", port, err))
			}
			continue
		}

		// if port already exists, let's update
		err = l.repository.Update(ctx, port)
		if err != nil {
			return errors.Wrap(localErrs.ErrInternalServerError, fmt.Sprintf("failed to update port [%+v] on storage: %+v", port, err))
		}
	}

	return nil
}
