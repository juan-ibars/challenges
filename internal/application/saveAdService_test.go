package application

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/mocks"
	"testing"
	"time"
)

func TestSaveAdService(t *testing.T) {
	repository := new(mocks.AdRepository)
	idGenerator := new(mocks.IdGenerator)
	timeGenerator := new(mocks.TimeGenerator)
	service := NewSaveAdService(idGenerator, timeGenerator, repository)
	id, _ := uuid.NewRandom()
	title := "Some title"
	description := "Some description"
	price := 1.2345
	date := time.Now()
	ad := Ad{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Date:        date,
	}
	repository.On("Save", ad).Return()
	idGenerator.On("Generate").Return(id)
	timeGenerator.On("Generate").Return(date)

	got := service.Execute(title, description, price)

	assert.Equal(t, ad, got)
	repository.AssertCalled(t, "Save", ad)
}
