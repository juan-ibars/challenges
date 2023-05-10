package main

import (
	"fmt"
	"github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure"
)

func main() {
	fmt.Println("Hello world!")
	ad := CreateAd("Some title", "Some Description", 0.0)
	ad.Print()

	Save(ad)
	foundAd := FindById(ad.Id)
	foundAd.Print()

	id, _ := uuid.NewRandom()
	notFoundAd := FindById(id)
	if notFoundAd == nil {
		println("ad not found")
	}
}
