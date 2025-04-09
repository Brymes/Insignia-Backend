package api

import (
	"Insignia-Backend/controllers"

	"github.com/gin-gonic/gin"
)

func CoreRoutes(engine *gin.Engine) {
	coreRouter := engine.Group("/")
	{
		// TODO they receive invite link with organization ID In request payload, and they get signed up and assigned to organization
		coreRouter.POST("repair", controllers.CreateBoilerRepair)
		coreRouter.POST("install", controllers.CreateBoilerInstallation)
	}
}
