package memory

import (
	"context"
	"sync"

	"github.com/emzola/ridewise/riderservice/internal/repository"
	"github.com/emzola/ridewise/riderservice/pkg/model"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Rider
}

func New() *Repository {
	return &Repository{data: map[string]*model.Rider{}}
}

func (r *Repository) Create(ctx context.Context, phone string) (*model.Rider, error) {
	_, ok := r.data["uuid1"]
	if ok {
		return nil, repository.ErrDuplicatePhone
	}
	r.data["uuid1"] = &model.Rider{Phone: phone}
	return r.data["uuid1"], nil
}

func (r *Repository) Get(ctx context.Context, id string) (*model.Rider, error) {
	r.RLock()
	defer r.RUnlock()
	rider, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return rider, nil
}

func (r *Repository) Update(ctx context.Context, rider *model.Rider) error {
	r.Lock()
	defer r.Unlock()
	r.data[rider.ID] = rider
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.data[id]
	if !ok {
		return repository.ErrNotFound
	}
	delete(r.data, id)
	return nil
}
