package infrastructure

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"math/rand"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	sut := NewInMemoryRepository()
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	sut.ads = append(sut.ads, ad)

	got := sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
}

func TestSave(t *testing.T) {
	sut := NewInMemoryRepository()
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	sut.Save(ad)

	got := sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
}

func TestFindAll(t *testing.T) {
	sut := NewInMemoryRepository()
	want := 10
	for i := 0; i < want; i++ {
		id, _ := uuid.NewRandom()
		ad := domain.Ad{
			Id:          id,
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Price:       rand.Float64(),
			Date:        time.Now(),
		}
		sut.ads = append(sut.ads, ad)
	}

	foundAds := sut.FindAllAds()
	got := len(foundAds)

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
