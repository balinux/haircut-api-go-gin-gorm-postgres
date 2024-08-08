package routes

import (
	"haircut-api-go-gin-gorm-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(haircutController *controllers.HaircutController) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/haircuts", haircutController.GetAllHaircuts)
		v1.GET("/haircuts2", haircutController.GetAllHaircuts)
		v1.POST("/haircuts", haircutController.CreateHaircut)
		v1.GET("/haircuts/:id", haircutController.GetHaircut)
	}
	return router
}
