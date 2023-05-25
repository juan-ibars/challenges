package domain

import . "time"

type DateGenerator struct{}

func NewDateGenerator() *DateGenerator {
	return &DateGenerator{}
}

func (d *DateGenerator) Generate() time.Time {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return t
}
