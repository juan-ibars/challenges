package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application/ad"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/controller"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/infrastructure/repository/inmemory"
	"math/rand"
	"sync"
	"time"
)

var connections = make([]Connection, 0)
var wg = sync.WaitGroup{}
var mu sync.Mutex

func main() {
	fmt.Println("Hello world!")
	//challenge3()

	var ads = []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "00"}

	var sites = []Site{
		{Id: "1", Name: "first"},
		{Id: "2", Name: "second"},
		{Id: "3", Name: "third"},
		{Id: "4", Name: "fourth"},
		{Id: "5", Name: "fifth"},
	}

	wg.Add(2 * len(ads))

	channel := make(chan Site)

	for _, a := range ads {
		// get the site of the ad
		go getSite(sites, channel)
		// save connection
		site := <-channel
		go saveConnection(a, site)
	}

	// wait to all coroutines to finish
	wg.Wait()

	// Print saved connections
	fmt.Println("Connections:", connections)
}

func getSite(sites []Site, ch chan<- Site) {
	defer wg.Done()
	site := sites[rand.Intn(5)]
	fmt.Printf("Name-> Id: %v, Name: %v\n", site.Id, site.Name)
	ch <- site
}

func saveConnection(ad string, site Site) {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	connection := Connection{
		AdId:   ad,
		SiteId: site.Id,
	}
	mu.Lock()
	defer mu.Unlock()
	connections = append(connections, connection)
	fmt.Println("Saved connection for ad", ad)
}

func challenge3() {
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
