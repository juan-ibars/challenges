package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/application/ad"
)

type GetAdController struct {
	service FindAdService
}

func NewGetAdController(service FindAdService) *GetAdController {
	return &GetAdController{service: service}
}

func (c *GetAdController) SetUpRouter(router *gin.Engine) *gin.Engine {
	router.GET("/ads/:id", c.handler())
	return router
}

func (c *GetAdController) handler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, _ := uuid.Parse(ctx.Param("id"))
		ad, err := c.service.Execute(id)
		if err != nil {
			if e, ok := err.(*application.AdErrors); ok {
				ctx.JSONP(404, e.Error())
			} else {
				ctx.JSONP(500, nil)
			}
		} else {
			response := GetAdControllerResponse{
				Id:          ad.Id.String(),
				Title:       ad.Title,
				Description: ad.Description,
				Price:       fmt.Sprintf("%f", ad.Price),
				Date:        ad.Date.String(),
			}
			ctx.JSONP(200, response)
		}
	}
}
