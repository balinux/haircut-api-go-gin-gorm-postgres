package main

import (
	"haircut-api-go-gin-gorm-postgres/config"
	"haircut-api-go-gin-gorm-postgres/controllers"
	"haircut-api-go-gin-gorm-postgres/repositories"
	"haircut-api-go-gin-gorm-postgres/routes"
	"haircut-api-go-gin-gorm-postgres/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// set the gin Release mode
	gin.SetMode(gin.ReleaseMode)
	//
	// gin.SetMode(gin.DebugMode)

	// config inisialisasi database koneksi
	config.ConnectDatabase()

	// inisialisasi repository, service, controller
	haircutRepo := repositories.NewHaircutRepository(config.DB)
	haircutService := services.NewHaircutService(haircutRepo)
	haircutController := controllers.NewHaircutController(haircutService)

	// setup router
	router := routes.SetupRouter(haircutController)

	if err := router.Run(":8888"); err != nil {
		panic(err)
	}
}
