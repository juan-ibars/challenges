package ad

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

	got, err := service.Execute(title, description, price)

	assert.Equal(t, ad, got)
	assert.Equal(t, err, nil)
	repository.AssertCalled(t, "Save", ad)
}

func TestSaveAdServiceWhenDescriptionLongerThanFiftyChars(t *testing.T) {
	repository := new(mocks.AdRepository)
	idGenerator := new(mocks.IdGenerator)
	timeGenerator := new(mocks.TimeGenerator)
	service := NewSaveAdService(idGenerator, timeGenerator, repository)
	title := "Some title"
	description := "123456789012345678901234567890123456789012345678901234567890"
	price := 1.2345

	ad, err := service.Execute(title, description, price)

	assert.NotEqual(t, nil, err)
	assert.Equal(t, Ad{}, ad)
}
