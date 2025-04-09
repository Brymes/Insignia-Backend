package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func Server() {

	router := prepareServer()

	CoreRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome 🚀",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	log.Println("Server is running")
	err := router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
		return
	}
}
func prepareServer() *gin.Engine {

	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowOrigins = []string{"*", "http://localhost:3000"}

	router := gin.Default()
	// Registering MiddleWares
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router

}
