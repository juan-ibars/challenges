package infrastructure

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

func Save(ad Ad) {
	InMemoryRepository = append(InMemoryRepository, ad)
}

func FindById(id UUID) *Ad {
	for index, ad := range InMemoryRepository {
		if ad.Id == id {
			return &InMemoryRepository[index]
		}
	}
	return nil
}

func FindAllAds() []Ad {
	var ads []Ad
	for index, ad := range InMemoryRepository {
		if index < 5 {
			ads = append(ads, ad)
		}
	}
	return ads
}
