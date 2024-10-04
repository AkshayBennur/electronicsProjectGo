package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"electronicsProjectGo/models"
	"electronicsProjectGo/services"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsService *services.ProductsService
}

func NewProductsController(productsService *services.ProductsService) *ProductsController {
	return &ProductsController{
		productsService: productsService,
	}
}

func (rh ProductsController) CreateProduct(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create product request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var product models.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println("Error while unmarshaling create product request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.productsService.CreateProduct(&product)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh ProductsController) UpdateProduct(ctx *gin.Context) {
	productName := ctx.Param("name")
	
	responseErr := rh.productsService.UpdateProduct(productName)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh ProductsController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	responseErr := rh.productsService.DeleteProduct(productId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// func (rh ProductsController) GetProduct(ctx *gin.Context) {
// 	productId := ctx.Param("id")

// 	response, responseErr := rh.productsService.GetProduct(productId)
// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

func (rh ProductsController) GetProductsBatch(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")

	response, responseErr := rh.productsService.GetProductsBatch(country, year)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
