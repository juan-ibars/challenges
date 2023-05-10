package infrastructure

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type Ads []Ad

var adRepository Ads

func Save(ad Ad) {
	adRepository = append(adRepository, ad)
}

func FindById(id UUID) *Ad {
	for index, ad := range adRepository {
		if ad.Id == id {
			return &adRepository[index]
		}
	}
	return nil
}
