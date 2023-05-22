package application

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/mocks"
	"testing"
	"time"
)

func TestFindByIdService(t *testing.T) {
	repository := new(mocks.AdRepository)
	service := NewFindByIdService(repository)
	existingId, _ := uuid.NewRandom()
	noExistingId, _ := uuid.NewRandom()
	date := time.Now()
	ad := Ad{
		Id:          existingId,
		Title:       "Some title",
		Description: "Some description",
		Price:       1.2345,
		Date:        date,
	}
	dataSet := []struct {
		name  string
		input uuid.UUID
		want  *Ad
	}{
		{name: "test with some result", input: existingId, want: &ad},
		{name: "test with no result", input: noExistingId, want: nil},
	}
	repository.On("FindById", existingId).Return(&ad).Once()
	repository.On("FindById", noExistingId).Return(nil).Once()

	for _, tc := range dataSet {
		t.Run(tc.name, func(t *testing.T) {
			got := service.Execute(tc.input)
			assert.Equal(t, tc.want, got, "FindById test")
		})
	}
}
