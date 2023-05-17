package main

import (
	"fmt"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure"
	"math/rand"
)

var adRepository InMemoryRepository

func main() {
	fmt.Println("Hello world!")

	ad := CreateAd("Some title", "Some Description", 0.0)
	fmt.Println("Created Ad")
	ad.Print()

	adRepository.Save(ad)
	foundAd := adRepository.FindById(ad.Id)
	fmt.Println("Saved Ad")
	(*foundAd).Print()

	for i := 0; i < 10; i++ {
		ad := CreateAd("Some title", "Some Description", rand.Float64())
		adRepository.Save(ad)
	}

	foundAds := adRepository.FindAllAds()
	fmt.Println("list of Ads")
	for _, ad := range foundAds {
		ad.Print()
	}
}
