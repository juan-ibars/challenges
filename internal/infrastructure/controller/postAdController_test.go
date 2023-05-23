package controller

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestPostAdController(t *testing.T) {
	service := new(mocks.SaveService)
	controller := NewPostAdController(service)
	id, _ := uuid.NewRandom()
	router := gin.Default()
	router = controller.SetUpRouter(router)
	title := "Some title"
	description := "Some description"
	price := 123.45
	data := PostAdControllerBodyRequest{
		Title:       title,
		Description: description,
		Price:       fmt.Sprintf("%.2f", price),
	}
	body, _ := json.Marshal(data)
	ad := domain.Ad{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Date:        time.Time{},
	}
	service.On("Execute", title, description, price).Return(ad, nil).Once()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ads", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	var response PostAdControllerBodyResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, id.String(), response.Id)
}

func TestPostControllerWithBadRequest(t *testing.T) {
	service := new(mocks.SaveService)
	controller := NewPostAdController(service)
	router := gin.Default()
	router = controller.SetUpRouter(router)
	title := "Some title"
	description := "Some description"
	price := 123.45
	data := PostAdControllerBodyRequest{
		Title:       title,
		Description: description,
		Price:       fmt.Sprintf("%.2f", price),
	}
	body, _ := json.Marshal(data)
	err := errors.New("description field could not be longer than 50")
	service.On("Execute", title, description, price).Return(domain.Ad{}, err).Once()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ads", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
