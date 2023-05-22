package controller

import (
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/application/ad"
	"strconv"
)

type PostAdController struct {
	service SaveService
}

func NewPostAdController(service SaveService) *PostAdController {
	return &PostAdController{service: service}
}

func (c *PostAdController) SetUpRouter(router *gin.Engine) *gin.Engine {
	router.POST("/ads", c.handler())
	return router
}

func (c *PostAdController) handler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request PostAdControllerBodyRequest
		_ = ctx.BindJSON(&request)
		ad := c.service.Execute(request.Title, request.Description, c.getPrice(request.Price))
		response := PostAdControllerBodyResponse{
			Id: ad.Id.String(),
		}
		ctx.JSONP(201, response)
	}
}

func (c *PostAdController) getPrice(price string) float64 {
	value, _ := strconv.ParseFloat(price, 64)
	return value
}
