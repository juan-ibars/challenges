package inmemory

import (
	"fmt"
	. "github.com/google/uuid"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/repository"
)

var Ads = make([]Ad, 0)

type InMemoryRepository struct {
	ads interface{}
}

func NewInMemoryRepository(table interface{}) *InMemoryRepository {
	return &InMemoryRepository{
		ads: table,
	}
}

func (r *InMemoryRepository) Save(ad Ad) error {
	ads, ok := r.ads.(*[]Ad)
	if ok {
		*ads = append(*ads, ad)
		return nil
	} else {
		return &repository.InfrastructureErrors{
			Msg: fmt.Sprintf("DB error"),
		}
	}
}

func (r *InMemoryRepository) FindById(id UUID) (*Ad, error) {
	ads, ok := r.ads.(*[]Ad)
	if ok {
		for index, ad := range *ads {
			if ad.Id == id {
				return &(*ads)[index], nil
			}
		}
	} else {
		return nil, &repository.InfrastructureErrors{Msg: "DB error"}
	}
	return nil, &application.AdErrors{
		Msg: "Ad not found",
		Id:  id,
	}
}

func (r *InMemoryRepository) FindAllAds() ([]Ad, error) {
	ads, ok := r.ads.(*[]Ad)
	if ok {
		return *ads, nil
	} else {
		return make([]Ad, 0), &repository.InfrastructureErrors{Msg: "DB error"}
	}
}
