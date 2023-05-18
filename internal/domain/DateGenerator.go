package domain

import . "time"

type DateGenerator struct{}

func NewDateGenerator() *DateGenerator {
	return &DateGenerator{}
}

func (d *DateGenerator) Generate() Time {
	return Now()
}
