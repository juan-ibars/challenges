package main

import (
	"fmt"
	. "github.mpi-internal.com/juan-ibars/learning-go/domain"
)

func main() {
	fmt.Println("Hello world!")
	ad := CreateAd("Some title", "Some Description", 0.0)
	ad.Print()
}
