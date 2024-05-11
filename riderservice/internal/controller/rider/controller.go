package rider

import (
	"context"
	"errors"

	"github.com/emzola/ridewise/riderservice/internal/controller"
	"github.com/emzola/ridewise/riderservice/internal/repository"
	"github.com/emzola/ridewise/riderservice/pkg/model"
)

type riderRepository interface {
	Create(ctx context.Context, phone string) (*model.Rider, error)
	Get(ctx context.Context, id string) (*model.Rider, error)
	Update(ctx context.Context, rider *model.Rider) error
	Delete(ctx context.Context, id string) error
}

type Controller struct {
	repo riderRepository
}

func New(repo riderRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Create(ctx context.Context, phone string) (*model.Rider, error) {
	rider, err := c.repo.Create(ctx, phone)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicatePhone):
			return nil, controller.ErrDuplicatePhone
		default:
			return nil, err
		}
	}
	return rider, nil
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Rider, error) {
	rider, err := c.repo.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, controller.ErrNotFound
		default:
			return nil, err
		}
	}
	return rider, nil
}

func (c *Controller) Update(ctx context.Context, req UpdateRiderRequest) (*model.Rider, error) {
	rider, err := c.repo.Get(ctx, req.ID)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, controller.ErrNotFound
		default:
			return nil, err
		}
	}
	if req.FirstName != "" {
		rider.FirstName = req.FirstName
	}
	if req.LastName != "" {
		rider.LastName = req.LastName
	}
	// TODO: Implement verification for change of phone
	if req.Phone != "" {
		rider.Phone = req.Phone
	}
	// TODO: Implement verification for change of email
	if req.Email != "" {
		rider.Email = req.Email
	}
	if req.SavedLocations != nil {
		rider.SavedLocations = req.SavedLocations
	}
	err = c.repo.Update(ctx, rider)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, controller.ErrNotFound
		default:
			return nil, err
		}
	}
	return rider, nil
}

func (c *Controller) Delete(ctx context.Context, id string) error {
	err := c.repo.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return controller.ErrNotFound
		default:
			return err
		}
	}
	return nil
}

type UpdateRiderRequest struct {
	ID             string
	FirstName      string
	LastName       string
	Phone          string
	Email          string
	SavedLocations map[string]model.Location
}
