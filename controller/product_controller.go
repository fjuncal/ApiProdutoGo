package controller

import (
	"TestandoGo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
