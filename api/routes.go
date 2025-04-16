package api

import (
	"Insignia-Backend/controllers"

	"github.com/gin-gonic/gin"
)

func CoreRoutes(engine *gin.Engine) {
	coreRouter := engine.Group("/")
	{
		coreRouter.POST("repair", controllers.CreateBoilerRepair)
		coreRouter.POST("install", controllers.CreateBoilerInstallation)
		coreRouter.POST("help", controllers.CreateHelpRequest)
	}
}
