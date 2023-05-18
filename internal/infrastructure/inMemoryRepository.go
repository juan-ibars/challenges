package infrastructure

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type InMemoryRepository struct {
	ads []Ad
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		ads: make([]Ad, 0),
	}
}

func (r *InMemoryRepository) Save(ad Ad) {
	r.ads = append(r.ads, ad)
}

func (r *InMemoryRepository) FindById(id UUID) *Ad {
	for index, ad := range r.ads {
		if ad.Id == id {
			return &r.ads[index]
		}
	}
	return nil
}

func (r *InMemoryRepository) FindAllAds() []Ad {
	ads := make([]Ad, 0, 5)
	for index, ad := range r.ads {
		if index < 5 {
			ads = append(ads, ad)
		}
	}
	return ads
}
