package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application/ad"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/controller"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/repository"
)

func main() {
	fmt.Println("Hello world!")
	// create repository instance
	repository := NewInMemoryRepository(&Ads)
	// create id generator instance
	idGenerator := NewUUIDGenerator()
	// create time generator instance
	timeGenerator := NewDateGenerator()

	// save an ad
	// service
	saveAdService := ad.NewSaveAdService(idGenerator, timeGenerator, repository)
	// controller
	saveAdController := controller.NewPostAdController(saveAdService)

	// find an ad
	// service
	findByIdService := ad.NewFindByIdService(repository)
	// controller
	getAdController := controller.NewGetAdController(findByIdService)

	router := gin.Default()
	router = saveAdController.SetUpRouter(router)
	router = getAdController.SetUpRouter(router)

	_ = router.Run(":8000")
}
