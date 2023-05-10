package main

import (
	"fmt"
	. "github.mpi-internal.com/juan-ibars/learning-go/domain"
	. "github.mpi-internal.com/juan-ibars/learning-go/infrastructure"
)

func main() {
	fmt.Println("Hello world!")
	ad := CreateAd("Some title", "Some Description", 0.0)
	Save(ad)
	ad.Print()
}
