package domain

import . "time"

type DateGenerator struct{}

func NewDateGenerator() *DateGenerator {
	return &DateGenerator{}
}

func (d *DateGenerator) Generate() Time {
	time, _ := Parse(RFC3339, Now().Format(RFC3339))
	return time
}
