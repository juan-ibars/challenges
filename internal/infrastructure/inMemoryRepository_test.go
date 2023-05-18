package infrastructure

import (
	"github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"math/rand"
	"testing"
	"time"
)

var Sut InMemoryRepository

func TestFind(t *testing.T) {
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	Sut.ads = append(Sut.ads, ad)

	got := Sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
}

func TestSave(t *testing.T) {
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some Description",
		Price:       0.12345,
		Date:        date,
	}
	Sut.Save(ad)

	got := Sut.FindById(id)

	if *got != ad {
		t.Errorf("want %v got %v", ad, *got)
	}
}

func TestFindAll(t *testing.T) {
	want := 5
	for i := 0; i < 10; i++ {
		ad := CreateAd("Some title", "Some Description", rand.Float64())
		Sut.Save(ad)
	}

	foundAds := Sut.FindAllAds()
	got := len(foundAds)

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
