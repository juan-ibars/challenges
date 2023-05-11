package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID
	title       string
	description string
	price       float64
	date        time.Time
}

func GenerateUUID() uuid.UUID {
	result, _ := uuid.NewRandom()
	return result
}

func CreateAd(title string, description string, price float64) Ad {
	return Ad{
		Id:          GenerateUUID(),
		title:       title,
		description: description,
		price:       price,
		date:        time.Now(),
	}
}

func (ad Ad) Print() {
	fmt.Printf("ad:{\n id:%s\n title:%s\n description:%s\n price:%f\n date:%s\n}\n", ad.Id.String(), ad.title, ad.description, ad.price, ad.date.String())
}
