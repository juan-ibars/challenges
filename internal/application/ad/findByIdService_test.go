package ad

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
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
		want  Ad
		err   error
	}{
		{name: "test with some result", input: existingId, want: ad, err: nil},
		{name: "test with no result", input: noExistingId, want: Ad{}, err: &application.AdErrors{Msg: fmt.Sprintf("ad with id %s not found", noExistingId.String()), Id: noExistingId}},
	}
	repository.On("FindById", existingId).Return(&ad).Once()
	repository.On("FindById", noExistingId).Return(nil).Once()

	for _, tc := range dataSet {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.Execute(tc.input)
			assert.Equal(t, tc.want, got, "FindById test")
			assert.Equal(t, tc.err, err, "FindById test")
		})
	}
}
