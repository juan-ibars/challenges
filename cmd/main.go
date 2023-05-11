package main

import (
	"fmt"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure"
	"math/rand"
)

func main() {
	fmt.Println("Hello world!")

	ad := CreateAd("Some title", "Some Description", 0.0)
	fmt.Println("Created Ad")
	ad.Print()

	Save(ad)
	foundAd := FindById(ad.Id)
	fmt.Println("Saved Ad")
	foundAd.Print()

	for i := 0; i < 10; i++ {
		ad := CreateAd("Some title", "Some Description", rand.Float64())
		InMemoryRepository = append(InMemoryRepository, ad)
	}

	foundAds := FindAllAds()
	fmt.Println("list of Ads")
	for _, ad := range foundAds {
		ad.Print()
	}
}
