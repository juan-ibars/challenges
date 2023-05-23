package repository

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"math/rand"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	ads := make([]domain.Ad, 0)
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	ads = append(ads, ad)
	sut := NewInMemoryRepository(&ads)

	got, err := sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
	assert.Nil(t, err)
}

func TestSave(t *testing.T) {
	ads := make([]domain.Ad, 0)
	sut := NewInMemoryRepository(&ads)
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	saveErr := sut.Save(ad)

	got, findErr := sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
	assert.Nil(t, saveErr)
	assert.Nil(t, findErr)
}

func TestFindAll(t *testing.T) {
	ads := make([]domain.Ad, 0)
	for i := 0; i < 10; i++ {
		id, _ := uuid.NewRandom()
		ad := domain.Ad{
			Id:          id,
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Price:       rand.Float64(),
			Date:        time.Now(),
		}
		ads = append(ads, ad)
	}
	sut := NewInMemoryRepository(&ads)
	want := 10

	foundAds, err := sut.FindAllAds()
	got := len(foundAds)

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
	assert.Nil(t, err)
}

func TestDBError(t *testing.T) {
	ads := make([]int, 0)
	sut := NewInMemoryRepository(&ads)
	id, _ := uuid.NewRandom()

	_, err := sut.FindById(id)

	assert.NotNil(t, err)
}
