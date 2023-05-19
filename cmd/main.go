package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/repository"
	"math/rand"
)

func main() {
	fmt.Println("Hello world!")
	// create repository instance
	repository := NewInMemoryRepository()
	// create id generator instance
	idGenerator := NewUUIDGenerator()
	// create time generator instance
	timeGenerator := NewDateGenerator()

	// save an ad
	saveAdService := application.NewSaveAdService(idGenerator, timeGenerator, repository)
	title := "Some title"
	description := "Some Description"
	price := 0.1234
	savedAd := saveAdService.Execute(title, description, price)
	fmt.Println("Created Ad")
	savedAd.Print()

	//ad := CreateAd("Some title", description, price)
	//fmt.Println("Created Ad")
	//ad.Print()
	//repository.Save(ad)

	// find an ad
	findByIdService := application.NewFindByIdService(repository)
	foundAd := findByIdService.Execute(savedAd.Id)
	fmt.Println("Saved Ad")
	foundAd.Print()

	//foundAd := repository.FindById(ad.Id)
	//fmt.Println("Saved Ad")
	//(*foundAd).Print()

	for i := 0; i < 10; i++ {
		ad := Ad{
			Id:          idGenerator.Generate(),
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Price:       rand.Float64(),
			Date:        timeGenerator.Generate(),
		}
		repository.Save(ad)
	}

	findAllService := application.NewFindAllService(repository)
	foundAds := findAllService.Execute()
	//foundAds := repository.FindAllAds()
	fmt.Println("list of ads")
	for _, ad := range foundAds {
		ad.Print()
	}
}
