package application

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/mocks"
	"math/rand"
	"testing"
	"time"
)

func TestFindAllService(t *testing.T) {
	repository := new(mocks.AdRepository)
	service := NewFindAllService(repository)
	repositoryResponse := generateRepositoryResponse()
	lenWanted := 5
	adsWanted := repositoryResponse[:5]

	repository.On("FindAllAds").Return(repositoryResponse)

	got := service.Execute()

	assert.Equal(t, lenWanted, len(got))
	assert.Equal(t, adsWanted, got)
}

func generateRepositoryResponse() []domain.Ad {
	ads := make([]domain.Ad, 0)
	for i := 0; i < 10; i++ {
		ad := domain.Ad{
			Id:          generateUUID(),
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Price:       rand.Float64(),
			Date:        time.Now(),
		}
		ads = append(ads, ad)
	}
	return ads
}

func generateUUID() uuid.UUID {
	id, _ := uuid.NewRandom()
	return id
}
