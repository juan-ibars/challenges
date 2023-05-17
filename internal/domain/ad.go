package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID
	Title       string
	Description string
	Price       float64
	Date        time.Time
}

func GenerateUUID() uuid.UUID {
	result, _ := uuid.NewRandom()
	return result
}

func CreateAd(title string, description string, price float64) Ad {
	return Ad{
		Id:          GenerateUUID(),
		Title:       title,
		Description: description,
		Price:       price,
		Date:        time.Now(),
	}
}

func (ad Ad) Print() {
	fmt.Printf("ad:{\n id:%s\n title:%s\n description:%s\n price:%f\n date:%s\n}\n", ad.Id.String(), ad.Title, ad.Description, ad.Price, ad.Date.String())
}