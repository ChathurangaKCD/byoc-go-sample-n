package repositories

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"example.choreo.dev/internal/models"
)

var ErrRecordNotFound = errors.New("record not found")

type petRepository struct {
	store *[]models.Pet
	lock  *sync.RWMutex
}

func NewPetRepository() models.PetRepository {
	pets := make([]models.Pet, 0)
	return &petRepository{
		store: &pets,
		lock:  &sync.RWMutex{},
	}
}

func (p petRepository) Add(ctx context.Context, pet *models.Pet) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if pet.Id == "" {
		pet.Id = uuid.NewString()
	}
	v := *pet
	*p.store = append(*p.store, v)
	return nil
}

func (p petRepository) Update(ctx context.Context, updated *models.Pet) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	idx := slices.IndexFunc(*p.store, func(e models.Pet) bool {
		return e.Id == updated.Id
	})
	if idx == -1 {
		return fmt.Errorf("petRepository:Update: %w", ErrRecordNotFound)
	}
	v := *updated
	(*p.store)[idx] = v
	return nil
}

func (p petRepository) List(ctx context.Context) ([]models.Pet, error) {
	return *p.store, nil
}

func (p petRepository) GetPetById(ctx context.Context, id string) (*models.Pet, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	idx := slices.IndexFunc(*p.store, func(e models.Pet) bool {
		return e.Id == id
	})
	if idx == -1 {
		return nil, fmt.Errorf("petRepository:Update: %w", ErrRecordNotFound)
	}
	v := (*p.store)[idx]
	return &v, nil
}