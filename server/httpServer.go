package server

import (
	"database/sql"
	"log"
	"runners-mysql/controllers"
	"runners-mysql/repositories"
	"runners-mysql/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	productsController *controllers.ProductsController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	productsRepository := repositories.NewProductsRepository(dbHandler)
	productsService := services.NewProductsService(productsRepository)
	productsController := controllers.NewProductsController(productsService)

	router := gin.Default()

	router.GET("/product", productsController.GetProductsBatch)
	router.POST("admin/product", productsController.CreateProduct)
	router.DELETE("admin/product/:id", productsController.DeleteProduct)
	// router.GET("/product/:id", productsController.GetProduct)
	router.PUT("user/select/:name", productsController.UpdateProduct)

	// router.POST("/result", resultsController.CreateResult)
	// router.DELETE("/result/:id", resultsController.DeleteResult)

	return HttpServer{
		config:            config,
		router:            router,
		productsController: productsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
