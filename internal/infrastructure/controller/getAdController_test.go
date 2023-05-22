package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAdController(t *testing.T) {
	service := new(mocks.FindAdService)
	id, _ := uuid.NewRandom()
	date := time.Now()
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some description",
		Price:       1.2345,
		Date:        date,
	}
	service.On("Execute", id).Return(&ad).Once()
	controller := NewGetAdController(service)
	router := gin.Default()
	router = controller.SetUpRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/ads/%s", id.String()), nil)
	router.ServeHTTP(w, req)

	var expectedBody GetAdControllerResponse
	_ = json.Unmarshal(w.Body.Bytes(), &expectedBody)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, ad.Id.String(), expectedBody.Id)
	assert.Equal(t, ad.Title, expectedBody.Title)
	assert.Equal(t, ad.Description, expectedBody.Description)
	assert.Equal(t, fmt.Sprintf("%f", ad.Price), expectedBody.Price)
	assert.Equal(t, ad.Date.String(), expectedBody.Date)
}

func TestGetAdControllerWithNoResults(t *testing.T) {
	service := new(mocks.FindAdService)
	id, _ := uuid.NewRandom()
	service.On("Execute", id).Return(nil).Once()
	controller := NewGetAdController(service)
	router := gin.Default()
	router = controller.SetUpRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/ads/%s", id.String()), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
