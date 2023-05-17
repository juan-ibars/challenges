package main

import (
	"fmt"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure"
	"math/rand"
)

var repository InMemoryRepository

func main() {
	fmt.Println("Hello world!")

	// save an ad
	saveAdService := application.NewSaveAdService(repository)
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
		ad := CreateAd("Some title", "Some Description", rand.Float64())
		repository.Save(ad)
	}

	findAllService := application.NewFindAllService(repository)
	foundAds := findAllService.Execute()
	//foundAds := repository.FindAllAds()
	fmt.Println("list of Ads")
	for _, ad := range foundAds {
		ad.Print()
	}
}
