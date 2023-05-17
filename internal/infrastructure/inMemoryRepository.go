package infrastructure

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type InMemoryRepository struct{}

var Ads = make([]Ad, 0)

func (r InMemoryRepository) Save(ad Ad) {
	Ads = append(Ads, ad)
}

func (r InMemoryRepository) FindById(id UUID) *Ad {
	for index, ad := range Ads {
		if ad.Id == id {
			return &Ads[index]
		}
	}
	return nil
}

func (r InMemoryRepository) FindAllAds() []Ad {
	ads := make([]Ad, 0, 5)
	for index, ad := range Ads {
		if index < 5 {
			ads = append(ads, ad)
		}
	}
	return ads
}
